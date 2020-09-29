// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the usage data of a usage plan in a specified time interval.
func (c *Client) GetUsage(ctx context.Context, params *GetUsageInput, optFns ...func(*Options)) (*GetUsageOutput, error) {
	stack := middleware.NewStack("GetUsage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUsageMiddlewares(stack)
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
	addOpGetUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsage(options.Region), middleware.Before)
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
			OperationName: "GetUsage",
			Err:           err,
		}
	}
	out := result.(*GetUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get the usage data of a usage plan in a specified time
// interval.
type GetUsageInput struct {
	// The current pagination position in the paged result set.
	Position *string
	Template *bool
	// [Required] The ending date (e.g., 2016-12-31) of the usage data.
	EndDate *string
	// [Required] The starting date (e.g., 2016-01-01) of the usage data.
	StartDate *string
	// [Required] The Id of the usage plan associated with the usage data.
	UsagePlanId      *string
	TemplateSkipList []*string
	Title            *string
	// The Id of the API key associated with the resultant usage data.
	KeyId *string
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32
	Name  *string
}

// Represents the usage data of a usage plan. Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html),
// Manage Usage in a Usage Plan
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-usage-plans-with-console.html#api-gateway-usage-plan-manage-usage)
type GetUsageOutput struct {
	// The ending date of the usage data.
	EndDate *string
	// The plan Id associated with this usage data.
	UsagePlanId *string
	// The starting date of the usage data.
	StartDate *string
	// The usage data, as daily logs of used and remaining quotas, over the specified
	// time interval indexed over the API keys in a usage plan. For example, {...,
	// "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]}, where {api_key}
	// stands for an API key value and the daily log entry is of the format [used
	// quota, remaining quota].
	Items map[string][][]*int64
	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUsageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUsage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsage{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsage",
	}
}