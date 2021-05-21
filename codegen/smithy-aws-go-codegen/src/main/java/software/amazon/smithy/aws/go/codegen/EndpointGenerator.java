/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */
package software.amazon.smithy.aws.go.codegen;

import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.TreeMap;
import java.util.function.Consumer;
import java.util.stream.Collectors;
import software.amazon.smithy.aws.go.codegen.customization.S3ModelUtils;
import software.amazon.smithy.aws.traits.ServiceTrait;
import software.amazon.smithy.codegen.core.CodegenException;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.go.codegen.GoSettings;
import software.amazon.smithy.go.codegen.GoStackStepMiddlewareGenerator;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.MiddlewareIdentifier;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.SymbolUtils;
import software.amazon.smithy.go.codegen.TriConsumer;
import software.amazon.smithy.go.codegen.integration.ConfigField;
import software.amazon.smithy.go.codegen.integration.ProtocolUtils;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.node.Node;
import software.amazon.smithy.model.node.ObjectNode;
import software.amazon.smithy.model.node.StringNode;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.utils.IoUtils;
import software.amazon.smithy.utils.ListUtils;
import software.amazon.smithy.utils.SmithyBuilder;

/**
 * Writes out a file that resolves endpoints using endpoints.json, but the
 * created resolver resolves endpoints for a single service.
 */
public final class EndpointGenerator implements Runnable {
    public static final String MIDDLEWARE_NAME = "ResolveEndpoint";
    public static final String ADD_MIDDLEWARE_HELPER_NAME = String.format("add%sMiddleware", MIDDLEWARE_NAME);
    public static final String RESOLVER_INTERFACE_NAME = "EndpointResolver";
    public static final String RESOLVER_FUNC_NAME = "EndpointResolverFunc";
    public static final String RESOLVER_OPTIONS = "EndpointResolverOptions";
    public static final String CLIENT_CONFIG_RESOLVER = "resolveDefaultEndpointConfiguration";
    public static final String RESOLVER_CONSTRUCTOR_NAME = "NewDefaultEndpointResolver";
    public static final String AWS_ENDPOINT_RESOLVER_HELPER = "withEndpointResolver";
    public static final String DUAL_STACK_ENDPOINT_OPTION = "DualStackEndpoint";

    private static final String EndpointResolverFromURL = "EndpointResolverFromURL";
    private static final String ENDPOINT_SOURCE_CUSTOM = "EndpointSourceCustom";
    private static final Symbol AWS_ENDPOINT = SymbolUtils.createPointableSymbolBuilder(
            "Endpoint", AwsGoDependency.AWS_CORE).build();

    private static final int ENDPOINT_MODEL_VERSION = 3;
    private static final String INTERNAL_ENDPOINT_PACKAGE = "internal/endpoints";
    private static final String INTERNAL_RESOLVER_NAME = "Resolver";
    private static final String INTERNAL_RESOLVER_OPTIONS_NAME = "Options";
    private static final String INTERNAL_ENDPOINTS_DATA_NAME = "defaultPartitions";
    private static final String DISABLE_HTTPS = "DisableHTTPS";

    // dual-stack related constants
    private static final String USE_DUAL_STACK_SHARED_OPTION = "UseDualStack";
    private static final String DUAL_STACK_ENDPOINT_TYPE_NAME = DUAL_STACK_ENDPOINT_OPTION;
    private static final String USE_DUAL_STACK_SHARED_CONFIG_RESOLVER = "isDualStackEndpointEnabled";

    private static final List<EndpointOption> ENDPOINT_OPTIONS = ListUtils.of(
            EndpointOption.builder()
                    .name(DISABLE_HTTPS)
                    .type(SymbolUtils.createValueSymbolBuilder("bool").build())
                    .shared(true)
                    .build(),
            EndpointOption.builder()
                    .name(DUAL_STACK_ENDPOINT_OPTION)
                    .type(SymbolUtils.createValueSymbolBuilder(DUAL_STACK_ENDPOINT_TYPE_NAME,
                            AwsGoDependency.AWS_CORE).build())
                    .shared(true)
                    .sharedOptionName(USE_DUAL_STACK_SHARED_OPTION)
                    .sharedResolver(SymbolUtils.createValueSymbolBuilder(USE_DUAL_STACK_SHARED_CONFIG_RESOLVER).build())
                    .build()
    );

    private final GoSettings settings;
    private final Model model;
    private final TriConsumer<String, String, Consumer<GoWriter>> writerFactory;
    private final ServiceShape serviceShape;
    private final ObjectNode endpointData;
    private final String endpointPrefix;
    private final Map<String, Partition> partitions = new TreeMap<>();
    private final boolean isInternalOnly;
    private final boolean isGenerateModelQueryHelpers;
    private final String resolvedSdkID;

    private EndpointGenerator(Builder builder) {
        settings = SmithyBuilder.requiredState("settings", builder.settings);
        model = SmithyBuilder.requiredState("model", builder.model);
        writerFactory = SmithyBuilder.requiredState("writerFactory", builder.writerFactory);
        isInternalOnly = builder.internalOnly;
        serviceShape = settings.getService(model);
        isGenerateModelQueryHelpers = builder.modelQueryHelpers;

        ServiceTrait serviceTrait = serviceShape.expectTrait(ServiceTrait.class);

        if (builder.sdkID != null) {
            resolvedSdkID = builder.sdkID;
        } else {
            resolvedSdkID = serviceTrait.getSdkId();
        }

        String arnNamespace;
        if (builder.arnNamespace != null) {
            arnNamespace = builder.arnNamespace;
        } else {
            arnNamespace = serviceTrait.getArnNamespace();
        }

        endpointPrefix = getEndpointPrefix(resolvedSdkID, arnNamespace);
        endpointData = Node.parse(IoUtils.readUtf8Resource(getClass(), "endpoints.json")).expectObjectNode();

        validateVersion();
        loadPartitions();
    }

    private void validateVersion() {
        int version = endpointData.expectNumberMember("version").getValue().intValue();
        if (version != ENDPOINT_MODEL_VERSION) {
            throw new CodegenException("Invalid endpoints.json version. Expected version 3, found " + version);
        }
    }

    // Get service's endpoint prefix from a known list. If not found, fallback to ArnNamespace
    private String getEndpointPrefix(ServiceShape service) {
        ObjectNode endpointPrefixData = Node.parse(IoUtils.readUtf8Resource(getClass(), "endpoint-prefix.json"))
                .expectObjectNode();
        ServiceTrait serviceTrait = service.getTrait(ServiceTrait.class)
                .orElseThrow(() -> new CodegenException("No service trait found on " + service.getId()));
        return endpointPrefixData.getStringMemberOrDefault(serviceTrait.getSdkId(), serviceTrait.getArnNamespace());
    }

    private String getEndpointPrefix(String sdkId, String arnNamespace) {
        ObjectNode endpointPrefixData = Node.parse(IoUtils.readUtf8Resource(getClass(), "endpoint-prefix.json"))
                .expectObjectNode();
        return endpointPrefixData.getStringMemberOrDefault(sdkId, arnNamespace);
    }

    private void loadPartitions() {
        List<ObjectNode> partitionObjects = endpointData
                .expectArrayMember("partitions")
                .getElementsAs(Node::expectObjectNode);

        for (ObjectNode partition : partitionObjects) {
            String partitionName = partition.expectStringMember("partition").getValue();
            partitions.put(partitionName, new Partition(partition, partitionName));
        }
    }

    @Override
    public void run() {
        if (!this.isInternalOnly) {
            writerFactory.accept("endpoints.go", settings.getModuleName(), writer -> {
                generatePublicResolverTypes(writer);
                generateMiddleware(writer);
                generateAwsEndpointResolverWrapper(writer);
            });
        }

        String pkgName = isInternalOnly ? INTERNAL_ENDPOINT_PACKAGE + "/" + this.endpointPrefix : INTERNAL_ENDPOINT_PACKAGE;
        writerFactory.accept(pkgName + "/endpoints.go", getInternalEndpointImportPath(), (writer) -> {
            generateInternalResolverImplementation(writer);
            generateInternalEndpointsModel(writer);
            if (isGenerateModelQueryHelpers) {
                generateInternalModelHelpers(writer);
            }
        });

        if (!this.isInternalOnly) {
            writerFactory.accept(INTERNAL_ENDPOINT_PACKAGE + "/endpoints_test.go",
                    getInternalEndpointImportPath(), (writer) -> {
                        writer.addUseImports(SmithyGoDependency.TESTING);
                        writer.openBlock("func TestRegexCompile(t *testing.T) {", "}", () -> {
                            writer.write("_ = $T",
                                    getInternalEndpointsSymbol(INTERNAL_ENDPOINTS_DATA_NAME, false).build());
                        });
                    });
        }
    }

    private void generateInternalModelHelpers(GoWriter writer) {
        generateDNSSuffixFunction(writer);
        generateDNSSuffixFromRegionFunction(writer);
    }

    private void generateDNSSuffixFunction(GoWriter writer) {
        Symbol optionsSymbol = getInternalEndpointsSymbol(INTERNAL_RESOLVER_OPTIONS_NAME, false).build();

        writer.addUseImports(SmithyGoDependency.FMT);
        writer.writeDocs("GetDNSSuffix returns the dnsSuffix URL component for the given partition id");
        writer.openBlock("func GetDNSSuffix(id string, options $T) (string, error) {", "}", optionsSymbol, () -> {
            Symbol equalFold = SymbolUtils.createValueSymbolBuilder("EqualFold", SmithyGoDependency.STRINGS)
                    .build();

            Symbol dualStackEnable = DualStackEndpointConstant.ENABLE.getSymbol();

            writer.openBlock("switch {", "}", () -> {
                partitions.forEach((s, partition) -> {
                    writer.openBlock("case $T(id, $S):", "", equalFold, partition.id, () -> {
                        writer.openBlock("if options.$L == $T {", "}", DUAL_STACK_ENDPOINT_OPTION, dualStackEnable,
                                () -> {
                                    Optional<String> suffix = partition.getDualStackDnsSuffix();
                                    if (suffix.isPresent()) {
                                        writer.write("return $S, nil", suffix.get());
                                    } else {
                                        writer.write("return \"\", fmt.Errorf(\"partition $L does not have a "
                                                + "dual-stack dns suffix\")", partition.id);
                                    }
                                });
                        writer.write("return $S, nil", partition.dnsSuffix);
                    });
                });
                writer.openBlock("default:", "", () -> writer.write("return \"\", fmt.Errorf(\"unknown partition\")"));
            });
        });
        writer.write("");
    }

    private void generateDNSSuffixFromRegionFunction(GoWriter writer) {
        Symbol optionsSymbol = getInternalEndpointsSymbol(INTERNAL_RESOLVER_OPTIONS_NAME, false).build();

        writer.addUseImports(SmithyGoDependency.FMT);
        writer.writeDocs("GetDNSSuffixFromRegion returns the dnsSuffix URL component for the given partition id");
        writer.openBlock("func GetDNSSuffixFromRegion(region string, options $T) (string, error) {", "}", optionsSymbol,
                () -> {
                    Symbol dualStackEnable = DualStackEndpointConstant.ENABLE.getSymbol();

                    writer.openBlock("switch {", "}", () -> {
                        getSortedPartitions().forEach(partition -> {
                            writer.openBlock("case partitionRegexp.$L.MatchString(region):", "",
                                    getPartitionIDFieldName(partition.getId()), () -> {
                                        writer.openBlock("if options.$L == $T {", "}", DUAL_STACK_ENDPOINT_OPTION,
                                                dualStackEnable, () -> {
                                                    Optional<String> suffix = partition.getDualStackDnsSuffix();
                                                    if (suffix.isPresent()) {
                                                        writer.write("return $S, nil", suffix.get());
                                                    } else {
                                                        writer.write("return \"\", fmt.Errorf(\"partition $L does not "
                                                                + "have a dual-stack dns suffix\")", partition.id);
                                                    }
                                                });
                                        writer.write("return $S, nil", partition.dnsSuffix);
                                    });
                        });
                        writer.openBlock("default:", "", () ->
                                writer.write("return \"\", fmt.Errorf(\"unknown region partition\")"));
                    });
                });
        writer.write("");
    }

    private void generateAwsEndpointResolverWrapper(GoWriter writer) {
        Symbol awsEndpointResolver = SymbolUtils.createValueSymbolBuilder("EndpointResolver", AwsGoDependency.AWS_CORE)
                .build();
        Symbol resolverInterface = SymbolUtils.createValueSymbolBuilder(RESOLVER_INTERFACE_NAME).build();

        Symbol wrappedResolverSymbol = SymbolUtils.createPointableSymbolBuilder("wrappedEndpointResolver").build();
        writer.openBlock("type $T struct {", "}", wrappedResolverSymbol, () -> {
            writer.write("awsResolver $T", awsEndpointResolver);
            writer.write("resolver $T", resolverInterface);
        });
        writeExternalResolveEndpointImplementation(writer, wrappedResolverSymbol, "w", () -> {
            writer.openBlock("if w.awsResolver == nil {", "}", () -> writer.write("goto fallback"));

            writer.write("endpoint, err = w.awsResolver.ResolveEndpoint(ServiceID, region)");
            writer.openBlock("if err == nil {", "}", () -> writer.write("return endpoint, nil"));
            writer.write("");

            writer.addUseImports(SmithyGoDependency.ERRORS);
            writer.openBlock("if nf := (&$T{}); !errors.As(err, &nf) {", "}",
                    SymbolUtils.createValueSymbolBuilder("EndpointNotFoundError", AwsGoDependency.AWS_CORE).build(),
                    () -> writer.write("return endpoint, err"));
            writer.write("");

            writer.write("fallback:");
            writer.openBlock("if w.resolver == nil {", "}", () -> {
                writer.addUseImports(SmithyGoDependency.FMT);
                writer.write("return endpoint, fmt.Errorf(\"default endpoint resolver provided was nil\")");
            });
            writer.write("return w.resolver.ResolveEndpoint(region, options)");
        });

        // Generate exported helper for constructing a wrapper around the AWS EndpointResolver type that is compatible
        // with the clients EndpointResolver interface.
        writer.writeDocs(String.format("%s returns an EndpointResolver that first delegates endpoint resolution "
                        + "to the awsResolver. If awsResolver returns `aws.EndpointNotFoundError` error, the resolver "
                        + "will use the the provided fallbackResolver for resolution. awsResolver and fallbackResolver "
                        + "must not be nil",
                AWS_ENDPOINT_RESOLVER_HELPER));
        writer.openBlock("func $L(awsResolver $T, fallbackResolver $T) $T {", "}", AWS_ENDPOINT_RESOLVER_HELPER,
                awsEndpointResolver, resolverInterface, resolverInterface, () -> {
                    writer.openBlock("return &$T{", "}", wrappedResolverSymbol, () -> {
                        writer.write("awsResolver: awsResolver,");
                        writer.write("resolver: fallbackResolver,");
                    });
                });
    }

    private void generateMiddleware(GoWriter writer) {
        // Generate middleware definition
        GoStackStepMiddlewareGenerator middleware = GoStackStepMiddlewareGenerator.createSerializeStepMiddleware(
                MIDDLEWARE_NAME, MiddlewareIdentifier.string(MIDDLEWARE_NAME));
        middleware.writeMiddleware(writer, this::generateMiddlewareResolverBody,
                this::generateMiddlewareStructureMembers);

        Symbol stackSymbol = SymbolUtils.createPointableSymbolBuilder("Stack", SmithyGoDependency.SMITHY_MIDDLEWARE)
                .build();

        // Generate Middleware Adder Helper
        writer.openBlock("func $L(stack $P, o Options) error {", "}", ADD_MIDDLEWARE_HELPER_NAME, stackSymbol, () -> {
            writer.addUseImports(SmithyGoDependency.SMITHY_MIDDLEWARE);
            String closeBlock = String.format("}, \"%s\", middleware.Before)",
                    ProtocolUtils.OPERATION_SERIALIZER_MIDDLEWARE_ID);
            writer.write("endpointOptions := o.EndpointOptions");
            if (S3ModelUtils.isServiceS3(model, serviceShape) || S3ModelUtils.isServiceS3Control(model, serviceShape)) {
                writer.openBlock("if o.EndpointOptions.$L == $T {", "}", DUAL_STACK_ENDPOINT_OPTION,
                        DualStackEndpointConstant.UNSET.getSymbol(), () -> {
                            writer.openBlock("if o.UseDualstack {", "", () -> {
                                writer.write("endpointOptions.$L = $T", DUAL_STACK_ENDPOINT_OPTION,
                                        DualStackEndpointConstant.ENABLE.getSymbol());
                                writer.openBlock("} else {", "}", () -> {
                                    writer.write("endpointOptions.$L = $T", DUAL_STACK_ENDPOINT_OPTION,
                                            DualStackEndpointConstant.DISABLE.getSymbol());
                                });
                            });
                        });
            }
            writer.openBlock("return stack.Serialize.Insert(&$T{", closeBlock,
                    middleware.getMiddlewareSymbol(),
                    () -> {
                        writer.write("Resolver: o.EndpointResolver,");
                        writer.write("Options: endpointOptions,");
                    });
        });
        writer.write("");
        // Generate Middleware Remover Helper
        writer.openBlock("func remove$LMiddleware(stack $P) error {", "}", middleware.getMiddlewareSymbol(),
                stackSymbol, () -> {
                    writer.write("_, err := stack.Serialize.Remove((&$T{}).ID())", middleware.getMiddlewareSymbol());
                    writer.write("return err");
                });
    }

    private void generateMiddlewareResolverBody(GoStackStepMiddlewareGenerator g, GoWriter w) {
        w.addUseImports(SmithyGoDependency.FMT);
        w.addUseImports(SmithyGoDependency.NET_URL);
        w.addUseImports(AwsGoDependency.AWS_MIDDLEWARE);
        w.addUseImports(SmithyGoDependency.SMITHY_MIDDLEWARE);
        w.addUseImports(SmithyGoDependency.SMITHY_HTTP_TRANSPORT);

        w.write("req, ok := in.Request.(*smithyhttp.Request)");
        w.openBlock("if !ok {", "}", () -> {
            w.write("return out, metadata, fmt.Errorf(\"unknown transport type %T\", in.Request)");
        });
        w.write("");

        w.openBlock("if m.Resolver == nil {", "}", () -> {
            w.write("return out, metadata, fmt.Errorf(\"expected endpoint resolver to not be nil\")");
        });
        w.write("");

        w.write("var endpoint $T", SymbolUtils.createValueSymbolBuilder("Endpoint", AwsGoDependency.AWS_CORE)
                .build());
        w.write("endpoint, err = m.Resolver.ResolveEndpoint(awsmiddleware.GetRegion(ctx), m.Options)");
        w.openBlock("if err != nil {", "}", () -> {
            w.write("return out, metadata, fmt.Errorf(\"failed to resolve service endpoint, %w\", err)");
        });
        w.write("");

        w.write("req.URL, err = url.Parse(endpoint.URL)");
        w.openBlock("if err != nil {", "}", () -> {
            w.write("return out, metadata, fmt.Errorf(\"failed to parse endpoint URL: %w\", err)");
        });
        w.write("");

        w.openBlock("if len(awsmiddleware.GetSigningName(ctx)) == 0 {", "}", () -> {
            w.write("signingName := endpoint.SigningName");
            w.openBlock("if len(signingName) == 0 {", "}", () -> {
                w.write("signingName = $S", serviceShape.expectTrait(ServiceTrait.class).getArnNamespace());
            });
            w.write("ctx = awsmiddleware.SetSigningName(ctx, signingName)");
        });

        // set endoint source on context
        w.write("ctx = awsmiddleware.SetEndpointSource(ctx, endpoint.Source)");
        // set host-name immutable on context
        w.write("ctx = smithyhttp.SetHostnameImmutable(ctx, endpoint.HostnameImmutable)");
        // set signing region on context
        w.write("ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)");
        // set partition id on context
        w.write("ctx = awsmiddleware.SetPartitionID(ctx, endpoint.PartitionID)");

        w.insertTrailingNewline();
        w.write("return next.HandleSerialize(ctx, in)");
    }

    private void generateMiddlewareStructureMembers(GoStackStepMiddlewareGenerator g, GoWriter w) {
        w.write("Resolver $L", RESOLVER_INTERFACE_NAME);
        w.write("Options $L", RESOLVER_OPTIONS);
    }

    private Symbol.Builder getInternalEndpointsSymbol(String symbolName, boolean pointable) {
        Symbol.Builder builder;
        if (pointable) {
            builder = SymbolUtils.createPointableSymbolBuilder(symbolName);
        } else {
            builder = SymbolUtils.createValueSymbolBuilder(symbolName);
        }
        return builder.namespace(getInternalEndpointImportPath(), "/")
                .putProperty(SymbolUtils.NAMESPACE_ALIAS, "internalendpoints");
    }

    private String getInternalEndpointImportPath() {
        return settings.getModuleName() + "/" + INTERNAL_ENDPOINT_PACKAGE;
    }

    private void generatePublicResolverTypes(GoWriter writer) {
        Symbol awsEndpointSymbol = SymbolUtils.createValueSymbolBuilder("Endpoint", AwsGoDependency.AWS_CORE).build();
        Symbol internalEndpointsSymbol = getInternalEndpointsSymbol(INTERNAL_RESOLVER_NAME, true).build();

        Symbol resolverOptionsSymbol = SymbolUtils.createPointableSymbolBuilder(RESOLVER_OPTIONS).build();
        writer.writeDocs(String.format("%s is the service endpoint resolver options",
                resolverOptionsSymbol.getName()));
        writer.write("type $T = $T", resolverOptionsSymbol, getInternalEndpointsSymbol(INTERNAL_RESOLVER_OPTIONS_NAME,
                false).build());
        writer.write("");

        // Generate Resolver Interface
        writer.writeDocs(String.format("%s interface for resolving service endpoints.", RESOLVER_INTERFACE_NAME));
        writer.openBlock("type $L interface {", "}", RESOLVER_INTERFACE_NAME, () -> {
            writer.write("ResolveEndpoint(region string, options $T) ($T, error)", resolverOptionsSymbol,
                    awsEndpointSymbol);
        });
        writer.write("var _ $L = &$T{}", RESOLVER_INTERFACE_NAME, internalEndpointsSymbol);
        writer.write("");

        // Resolver Constructor
        writer.writeDocs(String.format("%s constructs a new service endpoint resolver", RESOLVER_CONSTRUCTOR_NAME));
        writer.openBlock("func $L() $P {", "}", RESOLVER_CONSTRUCTOR_NAME, internalEndpointsSymbol, () -> {
            writer.write("return $T()", getInternalEndpointsSymbol("New", false)
                    .build());
        });

        Symbol resolverFuncSymbol = SymbolUtils.createValueSymbolBuilder(RESOLVER_FUNC_NAME).build();

        // Generate resolver function creator
        writer.writeDocs(String.format("%s is a helper utility that wraps a function so it satisfies the %s "
                + "interface. This is useful when you want to add additional endpoint resolving logic, or stub out "
                + "specific endpoints with custom values.", RESOLVER_FUNC_NAME, RESOLVER_INTERFACE_NAME));
        writer.write("type $T func(region string, options $T) ($T, error)",
                resolverFuncSymbol, resolverOptionsSymbol, awsEndpointSymbol);

        writeExternalResolveEndpointImplementation(writer, resolverFuncSymbol, "fn", () -> {
            writer.write("return fn(region, options)");
        });

        // Generate Client Options Configuration Resolver
        writer.openBlock("func $L(o $P) {", "}", CLIENT_CONFIG_RESOLVER,
                SymbolUtils.createPointableSymbolBuilder("Options").build(), () -> {
                    writer.openBlock("if o.EndpointResolver != nil {", "}", () -> writer.write("return"));
                    writer.write("o.EndpointResolver = $L()", RESOLVER_CONSTRUCTOR_NAME);
                });

        // Generate EndpointResolverFromURL helper
        writer.writeDocs(String.format("%s returns an EndpointResolver configured using the provided endpoint url. "
                        + "By default, the resolved endpoint resolver uses the client region as signing region, and  "
                        + "the endpoint source is set to EndpointSourceCustom."
                        + "You can provide functional options to configure endpoint values for the resolved endpoint.",
                EndpointResolverFromURL));
        writer.openBlock("func $L(url string, optFns ...func($P)) EndpointResolver {", "}",
                EndpointResolverFromURL, AWS_ENDPOINT, () -> {
                    Symbol customEndpointSource = SymbolUtils.createValueSymbolBuilder(
                            ENDPOINT_SOURCE_CUSTOM, AwsGoDependency.AWS_CORE
                    ).build();
                    writer.write("e := $T{ URL : url, Source : $T }", AWS_ENDPOINT, customEndpointSource);
                    writer.write("for _, fn := range optFns { fn(&e) }");
                    writer.write("");

                    writer.openBlock("return $T(", ")", resolverFuncSymbol, () -> {
                        writer.write("func(region string, options $L) ($T, error) {", RESOLVER_OPTIONS, AWS_ENDPOINT);
                        writer.write("if len(e.SigningRegion) == 0 { e.SigningRegion = region }");
                        writer.write("return e, nil },");
                    });
                });
    }

    private void writeExternalResolveEndpointImplementation(
            GoWriter writer,
            Symbol receiverType,
            String receiverIdentifier,
            Runnable body
    ) {
        Symbol resolverOptionsSymbol = SymbolUtils.createPointableSymbolBuilder(RESOLVER_OPTIONS).build();
        writeResolveEndpointImplementation(writer, receiverType, receiverIdentifier, resolverOptionsSymbol,
                body);
    }

    private void writeInternalResolveEndpointImplementation(
            GoWriter writer,
            Symbol receiverType,
            String receiverIdentifier,
            Runnable body
    ) {
        Symbol resolverOptionsSymbol = SymbolUtils.createPointableSymbolBuilder(INTERNAL_RESOLVER_OPTIONS_NAME).build();
        writeResolveEndpointImplementation(writer, receiverType, receiverIdentifier, resolverOptionsSymbol,
                body);
    }

    /**
     * Writes the ResolveEndpoint function signature to satisfy the EndpointResolver interface.
     *
     * @param writer                the code writer
     * @param receiverType          the receiver symbol type should be can be value or pointer
     * @param receiverIdentifier    the identifier to use for the receiver
     * @param resolverOptionsSymbol the symbol for the options
     * @param body                  a runnable that will populate the function implementation.
     */
    private void writeResolveEndpointImplementation(
            GoWriter writer,
            Symbol receiverType,
            String receiverIdentifier,
            Symbol resolverOptionsSymbol,
            Runnable body
    ) {
        Symbol awsEndpointSymbol = SymbolUtils.createValueSymbolBuilder("Endpoint", AwsGoDependency.AWS_CORE).build();
        writer.openBlock("func ($L $P) ResolveEndpoint(region string, options $T) (endpoint $T, err error) {", "}",
                        receiverIdentifier, receiverType, resolverOptionsSymbol, awsEndpointSymbol, body::run)
                .write("");
    }

    private void generateInternalResolverImplementation(GoWriter writer) {
        // Options
        Symbol resolverOptionsSymbol = SymbolUtils.createPointableSymbolBuilder(INTERNAL_RESOLVER_OPTIONS_NAME).build();
        writer.writeDocs(String.format("%s is the endpoint resolver configuration options",
                resolverOptionsSymbol.getName()));
        writer.openBlock("type $T struct {", "}", resolverOptionsSymbol, () -> {
            ENDPOINT_OPTIONS.forEach(field -> {
                writer.write("$L $P", field.getName(), field.getType());
            });
        });
        writer.write("");

        writer.openBlock("func " + USE_DUAL_STACK_SHARED_CONFIG_RESOLVER + "(value $T) bool {", "}",
                SymbolUtils.createValueSymbolBuilder(DUAL_STACK_ENDPOINT_TYPE_NAME, AwsGoDependency.AWS_CORE)
                        .build(),
                () -> writer.write("return value == $T", DualStackEndpointConstant.ENABLE.getSymbol()));
        writer.write("");

        // Resolver
        Symbol resolverImplSymbol = SymbolUtils.createPointableSymbolBuilder(INTERNAL_RESOLVER_NAME).build();


        writer.writeDocs(String.format("%s %s endpoint resolver", resolverImplSymbol.getName(),
                this.resolvedSdkID));
        writer.openBlock("type $T struct {", "}", resolverImplSymbol, () -> {
            writer.write("partitions $T", SymbolUtils.createValueSymbolBuilder("Partitions",
                    AwsGoDependency.AWS_ENDPOINTS).build());
        });
        writer.write("");
        writer.writeDocs("ResolveEndpoint resolves the service endpoint for the given region and options");
        writeInternalResolveEndpointImplementation(writer, resolverImplSymbol, "r", () -> {
            // Currently all APIs require a region to derive the endpoint for that API. If there are ever a truly
            // region-less API then this should be gated at codegen.
            writer.addUseImports(AwsGoDependency.AWS_CORE);
            writer.write("if len(region) == 0 { return endpoint, &aws.MissingRegionError{} }");
            writer.write("");

            Symbol sharedOptions = SymbolUtils.createPointableSymbolBuilder("Options",
                    AwsGoDependency.AWS_ENDPOINTS).build();
            writer.openBlock("opt := $T{", "}", sharedOptions, () -> {
                ENDPOINT_OPTIONS.stream().filter(EndpointOption::isShared).forEach(field -> {
                    String internalName = field.getSharedOptionName().orElse(field.getName());
                    Optional<Symbol> resolver = field.getSharedResolver();
                    if (resolver.isPresent()) {
                        writer.write("$L: $T(options.$L),", internalName, resolver.get(), field.getName());
                    } else {
                        writer.write("$L: options.$L,", internalName, field.getName());
                    }
                });
            });
            writer.write("return r.partitions.ResolveEndpoint(region, opt)");
        });
        writer.write("");
        writer.writeDocs(String.format("New returns a new %s", resolverImplSymbol.getName()));
        writer.openBlock("func New() *$T {", "}", resolverImplSymbol, () -> writer.openBlock("return &$T{", "}",
                resolverImplSymbol, () -> {
                    writer.write("partitions: $L,", INTERNAL_ENDPOINTS_DATA_NAME);
                }));
    }

    private static String getPartitionIDFieldName(String id) {
        StringBuilder builder = new StringBuilder();

        char[] charArray = id.toCharArray();
        boolean capitalize = true;
        for (int i = 0; i < charArray.length; i++) {
            if (!Character.isAlphabetic(charArray[i])) {
                capitalize = true;
                continue;
            }

            if (capitalize) {
                builder.append(Character.toUpperCase(charArray[i]));
                capitalize = false;
            } else {
                builder.append(Character.toLowerCase(charArray[i]));
            }
        }

        return builder.toString();
    }

    private void generateInternalEndpointsModel(GoWriter writer) {
        writer.addUseImports(AwsGoDependency.AWS_ENDPOINTS);

        List<Partition> sortedPartitions = getSortedPartitions();

        writer.openBlock("var partitionRegexp = struct{", "}{", () -> {
            sortedPartitions.forEach(partition -> {
                writer.write("$L $P", getPartitionIDFieldName(partition.getId()),
                        SymbolUtils.createPointableSymbolBuilder("Regexp", AwsGoDependency.REGEXP).build());
            });
        }).openBlock("", "}", () -> {
            sortedPartitions.forEach(partition -> {
                writer.write("$L: regexp.MustCompile($S),", getPartitionIDFieldName(partition.getId()),
                        partition.getConfig().expectStringMember("regionRegex").getValue());
            });
        });
        writer.write("");

        Symbol partitionsSymbol = SymbolUtils.createPointableSymbolBuilder("Partitions", AwsGoDependency.AWS_ENDPOINTS)
                .build();
        writer.openBlock("var $L = $T{", "}", INTERNAL_ENDPOINTS_DATA_NAME, partitionsSymbol, () -> {
            sortedPartitions.forEach(entry -> {
                writer.openBlock("{", "},", () -> writePartition(writer, entry));
            });
        });
    }

    private List<Partition> getSortedPartitions() {
        return partitions.entrySet().stream()
                .sorted((x, y) -> {
                    // Always sort standard aws partition first
                    if (x.getKey().equals("aws")) {
                        return -1;
                    }
                    return x.getKey().compareTo(y.getKey());
                }).map(Map.Entry::getValue).collect(Collectors.toList());
    }

    private void writePartition(GoWriter writer, Partition partition) {
        writer.write("ID: $S,", partition.getId());
        Symbol endpointSymbol = SymbolUtils.createValueSymbolBuilder("Endpoint",
                AwsGoDependency.AWS_ENDPOINTS).build();
        writer.openBlock("Defaults: $T{", "},", endpointSymbol,
                () -> writeEndpoint(writer, partition.getDefaults()));

        Optional<ObjectNode> dualStackDefaults = partition.getDualStackDefaults();
        dualStackDefaults.ifPresent(objectNode -> writer.openBlock("DualStackDefaults: $T{", "},", endpointSymbol,
                () -> writeEndpoint(writer, objectNode)));

        writer.addUseImports(AwsGoDependency.REGEXP);
        writer.write("RegionRegex: partitionRegexp.$L,", getPartitionIDFieldName(partition.getId()));

        Optional<String> optionalPartitionEndpoint = partition.getPartitionEndpoint();
        Symbol isRegionalizedValue = SymbolUtils.createValueSymbolBuilder(optionalPartitionEndpoint.isPresent()
                ? "false" : "true").build();
        writer.write("IsRegionalized: $T,", isRegionalizedValue);
        optionalPartitionEndpoint.ifPresent(s -> writer.write("PartitionEndpoint: $S,", s));

        Map<StringNode, Node> endpoints = partition.getEndpoints().getMembers();
        if (endpoints.size() > 0) {
            Symbol endpointsSymbol = SymbolUtils.createPointableSymbolBuilder("Endpoints",
                            AwsGoDependency.AWS_ENDPOINTS)
                    .build();
            writer.openBlock("Endpoints: $T{", "},", endpointsSymbol, () -> {
                endpoints.forEach((s, n) -> {
                    writer.openBlock("$S: $T{", "},", s, endpointSymbol,
                            () -> writeEndpoint(writer, n.expectObjectNode()));
                });
            });
        }

        Map<StringNode, Node> dualStackEndpoints = partition.getDualStackEndpoints().getMembers();
        if (dualStackEndpoints.size() > 0) {
            Symbol endpointsSymbol = SymbolUtils.createPointableSymbolBuilder("Endpoints",
                            AwsGoDependency.AWS_ENDPOINTS)
                    .build();
            writer.openBlock(DUAL_STACK_ENDPOINT_TYPE_NAME + "s: $T{", "},", endpointsSymbol, () -> {
                dualStackEndpoints.forEach((s, n) -> {
                    writer.openBlock("$S: $T{", "},", s, endpointSymbol,
                            () -> writeEndpoint(writer, n.expectObjectNode()));
                });
            });
        }
    }

    private void writeEndpoint(GoWriter writer, ObjectNode node) {
        node.getStringMember("hostname").ifPresent(n -> {
            writer.write("Hostname: $S,", n.getValue());
        });
        node.getArrayMember("protocols").ifPresent(nodes -> {
            writer.writeInline("Protocols: []string{");
            nodes.forEach(n -> {
                writer.writeInline("$S, ", n.expectStringNode().getValue());
            });
            writer.write("},");
        });
        node.getArrayMember("signatureVersions").ifPresent(nodes -> {
            writer.writeInline("SignatureVersions: []string{");
            nodes.forEach(n -> writer.writeInline("$S, ", n.expectStringNode().getValue()));
            writer.write("},");
        });
        node.getMember("credentialScope").ifPresent(n -> {
            ObjectNode credentialScope = n.expectObjectNode();
            Symbol credentialScopeSymbol = SymbolUtils.createValueSymbolBuilder("CredentialScope",
                            AwsGoDependency.AWS_ENDPOINTS)
                    .build();
            writer.openBlock("CredentialScope: $T{", "},", credentialScopeSymbol, () -> {
                credentialScope.getStringMember("region").ifPresent(nn -> {
                    writer.write("Region: $S,", nn.getValue());
                });
                credentialScope.getStringMember("service").ifPresent(nn -> {
                    writer.write("Service: $S,", nn.getValue());
                });
            });
        });
    }

    private static class EndpointOption extends ConfigField {
        private final boolean shared;
        private final String sharedOptionName;
        private final Symbol sharedResolver;

        public EndpointOption(Builder builder) {
            super(builder);
            this.shared = builder.shared;
            this.sharedOptionName = builder.sharedOptionName;
            this.sharedResolver = builder.sharedResolver;
        }

        public static Builder builder() {
            return new Builder();
        }

        public boolean isShared() {
            return shared;
        }

        public Optional<String> getSharedOptionName() {
            return Optional.ofNullable(sharedOptionName);
        }

        public Optional<Symbol> getSharedResolver() {
            return Optional.ofNullable(this.sharedResolver);
        }

        private static class Builder extends ConfigField.Builder {
            private boolean shared;
            private String sharedOptionName;
            private Symbol sharedResolver;

            public Builder() {
                super();
            }

            /**
             * Set the resolver config field to be shared common parameter
             *
             * @param shared whether the resolver config field is shared
             * @return the builder
             */
            public Builder shared(boolean shared) {
                this.shared = shared;
                return this;
            }

            public Builder sharedOptionName(String sharedOptionName) {
                this.sharedOptionName = sharedOptionName;
                return this;
            }

            public Builder sharedResolver(Symbol sharedResolver) {
                this.sharedResolver = sharedResolver;
                return this;
            }

            @Override
            public EndpointOption build() {
                return new EndpointOption(this);
            }

            @Override
            public Builder name(String name) {
                super.name(name);
                return this;
            }

            @Override
            public Builder type(Symbol type) {
                super.type(type);
                return this;
            }

            @Override
            public Builder documentation(String documentation) {
                super.documentation(documentation);
                return this;
            }
        }
    }

    private final class Partition {
        private static final String DEFAULTS_KEY = "defaults";
        private static final String DUAL_STACK_DEFAULTS_KEY = "dualstackDefaults";
        private static final String DNS_SUFFIX_KEY = "dnsSuffix";
        private static final String DUAL_STACK_DNS_SUFFIX_KEY = "dualstackDnsSuffix";
        public static final String HOSTNAME_KEY = "hostname";

        private final String id;
        private final ObjectNode defaults;
        private final ObjectNode dualStackDefaults;
        private final ObjectNode config;
        private final String dnsSuffix;
        private final String dualStackDnsSuffix;

        private Partition(ObjectNode config, String partition) {
            id = partition;
            this.config = config;

            // Resolve the partition defaults + the service defaults.
            ObjectNode service = getService();

            ObjectNode serviceDefaults = config.expectObjectMember(DEFAULTS_KEY).merge(service
                    .getObjectMember(DEFAULTS_KEY)
                    .orElse(Node.objectNode()));

            ObjectNode serviceDualStackDefaults = config.getObjectMember(DUAL_STACK_DEFAULTS_KEY)
                    .orElse(Node.objectNode()).merge(service
                            .getObjectMember(DUAL_STACK_DEFAULTS_KEY)
                            .orElse(Node.objectNode()));

            // Resolve the hostnameTemplate to use for this service in this partition.
            this.dnsSuffix = config.expectStringMember(DNS_SUFFIX_KEY).getValue();

            String hostnameTemplate = serviceDefaults.expectStringMember(HOSTNAME_KEY).getValue();
            hostnameTemplate = hostnameTemplate.replace("{service}", endpointPrefix);
            hostnameTemplate = hostnameTemplate.replace("{dnsSuffix}",
                    config.expectStringMember(DNS_SUFFIX_KEY).getValue());

            this.defaults = serviceDefaults.withMember("hostname", hostnameTemplate);

            if (serviceDualStackDefaults.size() > 0) {
                // Resolve the dual-stack hostnameTemplate to use for this service in this partition.
                this.dualStackDnsSuffix = service.getStringMember(DUAL_STACK_DNS_SUFFIX_KEY)
                        .orElse(config.expectStringMember(DUAL_STACK_DNS_SUFFIX_KEY)).getValue();

                String dualStackhostnameTemplate = serviceDualStackDefaults.expectStringMember("hostname").getValue();
                dualStackhostnameTemplate = dualStackhostnameTemplate.replace("{service}", endpointPrefix);
                dualStackhostnameTemplate = dualStackhostnameTemplate.replace("{dualstackDnsSuffix}",
                        this.dualStackDnsSuffix);

                this.dualStackDefaults = serviceDualStackDefaults.withMember(HOSTNAME_KEY, dualStackhostnameTemplate);
            } else {
                this.dualStackDefaults = null;
                this.dualStackDnsSuffix = null;
            }
        }

        /**
         * @return the partition defaults merged with the service defaults
         */
        ObjectNode getDefaults() {
            return defaults;
        }

        Optional<ObjectNode> getDualStackDefaults() {
            return Optional.ofNullable(dualStackDefaults);
        }

        ObjectNode getService() {
            ObjectNode services = config.getObjectMember("services").orElse(Node.objectNode());
            return services.getObjectMember(endpointPrefix).orElse(Node.objectNode());
        }

        ObjectNode getEndpoints() {
            return getService().getObjectMember("endpoints").orElse(Node.objectNode());
        }

        ObjectNode getDualStackEndpoints() {
            return getService().getObjectMember("dualstackEndpoints").orElse(Node.objectNode());
        }

        Optional<String> getPartitionEndpoint() {
            ObjectNode service = getService();
            // Note: regionalized services always use regionalized endpoints.
            return service.getBooleanMemberOrDefault("isRegionalized", true)
                    ? Optional.empty()
                    : service.getStringMember("partitionEndpoint").map(StringNode::getValue);
        }

        public String getId() {
            return id;
        }

        public ObjectNode getConfig() {
            return config;
        }

        public String getDnsSuffix() {
            return dnsSuffix;
        }

        public Optional<String> getDualStackDnsSuffix() {
            return Optional.ofNullable(dualStackDnsSuffix);
        }
    }

    public static Builder builder() {
        return new Builder();
    }

    public static final class Builder implements SmithyBuilder<EndpointGenerator> {
        private GoSettings settings;
        private Model model;
        private TriConsumer<String, String, Consumer<GoWriter>> writerFactory;
        private boolean internalOnly;
        private boolean modelQueryHelpers;
        private String sdkID;
        private String arnNamespace;

        public Builder settings(GoSettings settings) {
            this.settings = settings;
            return this;
        }

        public Builder model(Model model) {
            this.model = model;
            return this;
        }

        public Builder writerFactory(TriConsumer<String, String, Consumer<GoWriter>> writerFactory) {
            this.writerFactory = writerFactory;
            return this;
        }

        public Builder internalOnly(boolean internalOnly) {
            this.internalOnly = internalOnly;
            return this;
        }

        public Builder modelQueryHelpers(boolean modelQueryHelpers) {
            this.modelQueryHelpers = modelQueryHelpers;
            return this;
        }

        public Builder sdkID(String sdkID) {
            this.sdkID = sdkID;
            return this;
        }

        public Builder arnNamespace(String arnNamespace) {
            this.arnNamespace = arnNamespace;
            return this;
        }

        @Override
        public EndpointGenerator build() {
            return new EndpointGenerator(this);
        }
    }

    enum DualStackEndpointConstant {
        UNSET(DUAL_STACK_ENDPOINT_TYPE_NAME + "Unset"),
        ENABLE(DUAL_STACK_ENDPOINT_TYPE_NAME + "Enabled"),
        DISABLE(DUAL_STACK_ENDPOINT_TYPE_NAME + "Disabled");

        public static final String DOCUMENTATION = DUAL_STACK_ENDPOINT_TYPE_NAME + " is a constant to describe the "
                + "dual-stack endpoint resolution behavior.";

        private final String constantName;

        DualStackEndpointConstant(String name) {
            this.constantName = name;
        }

        public String getConstantName() {
            return constantName;
        }

        public Symbol getSymbol() {
            return SymbolUtils.createValueSymbolBuilder(getConstantName(), AwsGoDependency.AWS_CORE)
                    .build();
        }
    }
}
