// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Schedules a run.
func (c *Client) ScheduleRun(ctx context.Context, params *ScheduleRunInput, optFns ...func(*Options)) (*ScheduleRunOutput, error) {
	stack := middleware.NewStack("ScheduleRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpScheduleRunMiddlewares(stack)
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
	addOpScheduleRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opScheduleRun(options.Region), middleware.Before)
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
			OperationName: "ScheduleRun",
			Err:           err,
		}
	}
	out := result.(*ScheduleRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the schedule run operation.
type ScheduleRunInput struct {
	// The ARN of the project for the run to be scheduled.
	ProjectArn *string
	// Information about the settings for the run to be scheduled.
	Configuration *types.ScheduleRunConfiguration
	// The ARN of the device pool for the run to be scheduled.
	DevicePoolArn *string
	// The name for the run to be scheduled.
	Name *string
	// Specifies configuration information about a test run, such as the execution
	// timeout (in minutes).
	ExecutionConfiguration *types.ExecutionConfiguration
	// The filter criteria used to dynamically select a set of devices for a test run
	// and the maximum number of devices to be included in the run. Either
	// devicePoolArn or deviceSelectionConfiguration is required in a request.
	DeviceSelectionConfiguration *types.DeviceSelectionConfiguration
	// The ARN of an application package to run tests against, created with
	// CreateUpload (). See ListUploads ().
	AppArn *string
	// Information about the test for the run to be scheduled.
	Test *types.ScheduleRunTest
}

// Represents the result of a schedule run request.
type ScheduleRunOutput struct {
	// Information about the scheduled run.
	Run *types.Run

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpScheduleRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpScheduleRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpScheduleRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opScheduleRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ScheduleRun",
	}
}