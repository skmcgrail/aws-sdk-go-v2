// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Puts an Api resource.
func (c *Client) ReimportApi(ctx context.Context, params *ReimportApiInput, optFns ...func(*Options)) (*ReimportApiOutput, error) {
	stack := middleware.NewStack("ReimportApi", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpReimportApiMiddlewares(stack)
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
	addOpReimportApiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReimportApi(options.Region), middleware.Before)
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
			OperationName: "ReimportApi",
			Err:           err,
		}
	}
	out := result.(*ReimportApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ReimportApiInput struct {
	// The OpenAPI definition. Supported only for HTTP APIs.
	Body *string
	// Specifies whether to rollback the API creation when a warning is encountered. By
	// default, API creation continues if a warning is encountered.
	FailOnWarnings *bool
	// Specifies how to interpret the base path of the API during import. Valid values
	// are ignore, prepend, and split. The default value is ignore. To learn more, see
	// Set the OpenAPI basePath Property
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html).
	// Supported only for HTTP APIs.
	Basepath *string
	// The API identifier.
	ApiId *string
}

type ReimportApiOutput struct {
	// The validation information during API import. This may include particular
	// properties of your OpenAPI definition which are ignored during import. Supported
	// only for HTTP APIs.
	ImportInfo []*string
	// Avoid validating models when creating a deployment. Supported only for WebSocket
	// APIs.
	DisableSchemaValidation *bool
	// A collection of tags associated with the API.
	Tags map[string]*string
	// An API key selection expression. Supported only for WebSocket APIs. See API Key
	// Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression *string
	// The API ID.
	ApiId *string
	// The description of the API.
	Description *string
	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	RouteSelectionExpression *string
	// The API protocol.
	ProtocolType types.ProtocolType
	// The timestamp when the API was created.
	CreatedDate *time.Time
	// The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The
	// stage name is typically appended to this URI to form a complete path to a
	// deployed API stage.
	ApiEndpoint *string
	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []*string
	// The name of the API.
	Name *string
	// A CORS configuration. Supported only for HTTP APIs.
	CorsConfiguration *types.Cors
	// A version identifier for the API.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpReimportApiMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpReimportApi{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpReimportApi{}, middleware.After)
}

func newServiceMetadataMiddleware_opReimportApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ReimportApi",
	}
}