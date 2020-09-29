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
	"time"
)

// Gets information about a Stage () resource.
func (c *Client) GetStage(ctx context.Context, params *GetStageInput, optFns ...func(*Options)) (*GetStageOutput, error) {
	stack := middleware.NewStack("GetStage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetStageMiddlewares(stack)
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
	addOpGetStageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetStage(options.Region), middleware.Before)
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
			OperationName: "GetStage",
			Err:           err,
		}
	}
	out := result.(*GetStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to get information about a Stage () resource.
type GetStageInput struct {
	// [Required] The name of the Stage () resource to get information about.
	StageName        *string
	Template         *bool
	TemplateSkipList []*string
	Title            *string
	Name             *string
	// [Required] The string identifier of the associated RestApi ().
	RestApiId *string
}

// Represents a unique identifier for a version of a deployed RestApi () that is
// callable by users. Deploy an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html)
type GetStageOutput struct {
	// The identifier of the Deployment () that the stage points to.
	DeploymentId *string
	// The stage's description.
	Description *string
	// The name of the stage is the first path segment in the Uniform Resource
	// Identifier (URI) of a call to API Gateway. Stage names can only contain
	// alphanumeric characters, hyphens, and underscores. Maximum length is 128
	// characters.
	StageName *string
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string
	// The status of the cache cluster for the stage, if enabled.
	CacheClusterStatus types.CacheClusterStatus
	// The identifier of a client certificate for an API stage.
	ClientCertificateId *string
	// Specifies whether a cache cluster is enabled for the stage.
	CacheClusterEnabled *bool
	// Settings for the canary deployment in this stage.
	CanarySettings *types.CanarySettings
	// A map that defines the stage variables for a Stage () resource. Variable names
	// can have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]*string
	// The size of the cache cluster for the stage, if enabled.
	CacheClusterSize types.CacheClusterSize
	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings
	// The version of the associated API documentation.
	DocumentationVersion *string
	// The timestamp when the stage last updated.
	LastUpdatedDate *time.Time
	// The timestamp when the stage was created.
	CreatedDate *time.Time
	// Specifies whether active tracing with X-ray is enabled for the Stage ().
	TracingEnabled *bool
	// The ARN of the WebAcl associated with the Stage ().
	WebAclArn *string
	// A map that defines the method settings for a Stage () resource. Keys (designated
	// as /{method_setting_key below) are method paths defined as
	// {resource_path}/{http_method} for an individual method override, or /\*/\* for
	// overriding all methods in the stage.
	MethodSettings map[string]*types.MethodSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetStageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetStage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStage{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetStage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetStage",
	}
}