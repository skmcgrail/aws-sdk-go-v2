// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a public namespace based on DNS, which will be visible on the internet.
// The namespace defines your service naming scheme. For example, if you name your
// namespace example.com and name your service backend, the resulting DNS name for
// the service will be backend.example.com. For the current limit on the number of
// namespaces that you can create using the same AWS account, see AWS Cloud Map
// Limits (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)
// in the AWS Cloud Map Developer Guide.
func (c *Client) CreatePublicDnsNamespace(ctx context.Context, params *CreatePublicDnsNamespaceInput, optFns ...func(*Options)) (*CreatePublicDnsNamespaceOutput, error) {
	stack := middleware.NewStack("CreatePublicDnsNamespace", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePublicDnsNamespaceMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addIdempotencyToken_opCreatePublicDnsNamespaceMiddleware(stack, options)
	addOpCreatePublicDnsNamespaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePublicDnsNamespace(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreatePublicDnsNamespace",
			Err:           err,
		}
	}
	out := result.(*CreatePublicDnsNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePublicDnsNamespaceInput struct {

	// The name that you want to assign to this namespace.
	//
	// This member is required.
	Name *string

	// A unique string that identifies the request and that allows failed
	// CreatePublicDnsNamespace requests to be retried without the risk of executing
	// the operation twice. CreatorRequestId can be any unique string, for example, a
	// date/time stamp.
	CreatorRequestId *string

	// A description for the namespace.
	Description *string

	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []*types.Tag
}

type CreatePublicDnsNamespaceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePublicDnsNamespaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePublicDnsNamespace{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePublicDnsNamespace{}, middleware.After)
}

type idempotencyToken_initializeOpCreatePublicDnsNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePublicDnsNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePublicDnsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePublicDnsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePublicDnsNamespaceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePublicDnsNamespaceMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreatePublicDnsNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePublicDnsNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "CreatePublicDnsNamespace",
	}
}
