// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) GetDocumentationPart(ctx context.Context, params *GetDocumentationPartInput, optFns ...func(*Options)) (*GetDocumentationPartOutput, error) {
	stack := middleware.NewStack("GetDocumentationPart", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDocumentationPartMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetDocumentationPartValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentationPart(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "GetDocumentationPart",
			Err:           err,
		}
	}
	out := result.(*GetDocumentationPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets a specified documentation part of a given API.
type GetDocumentationPartInput struct {
	// [Required] The string identifier of the associated RestApi ().
	DocumentationPartId *string
	TemplateSkipList    []*string
	Template            *bool
	// [Required] The string identifier of the associated RestApi ().
	RestApiId *string
	Name      *string
	Title     *string
}

// A documentation part for a targeted API entity. A documentation part consists of
// a content map (properties) and a target (location). The target specifies an API
// entity to which the documentation content applies. The supported API entity
// types are API, AUTHORIZER, MODEL, RESOURCE, METHOD, PATH_PARAMETER,
// QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and
// RESPONSE_BODY. Valid location fields depend on the API entity type. All valid
// fields are not required. The content map is a JSON string of API-specific
// key-value pairs. Although an API can use any shape for the content map, only the
// OpenAPI-compliant documentation fields will be injected into the associated API
// entity definition in the exported OpenAPI definition file. Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationParts ()
type GetDocumentationPartOutput struct {
	// The DocumentationPart () identifier, generated by API Gateway when the
	// DocumentationPart is created.
	Id *string
	// The location of the API entity to which the documentation applies. Valid fields
	// depend on the targeted API entity type. All the valid location fields are not
	// required. If not explicitly specified, a valid location field is treated as a
	// wildcard and associated documentation content may be inherited by matching
	// entities, unless overridden.
	Location *types.DocumentationPartLocation
	// A content map of API-specific key-value pairs describing the targeted API
	// entity. The map must be encoded as a JSON string, e.g., "{ \"description\":
	// \"The API does ...\" }". Only OpenAPI-compliant documentation-related fields
	// from the properties map are exported and, hence, published as part of the API
	// entity definitions, while the original documentation parts are exported in a
	// OpenAPI extension of x-amazon-apigateway-documentation.
	Properties *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDocumentationPartMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentationPart{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentationPart{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDocumentationPart(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDocumentationPart",
	}
}