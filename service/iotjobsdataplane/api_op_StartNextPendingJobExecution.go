// Code generated by smithy-go-codegen DO NOT EDIT.

package iotjobsdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets and starts the next pending (status IN_PROGRESS or QUEUED) job execution
// for a thing.
func (c *Client) StartNextPendingJobExecution(ctx context.Context, params *StartNextPendingJobExecutionInput, optFns ...func(*Options)) (*StartNextPendingJobExecutionOutput, error) {
	stack := middleware.NewStack("StartNextPendingJobExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartNextPendingJobExecutionMiddlewares(stack)
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
	addOpStartNextPendingJobExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartNextPendingJobExecution(options.Region), middleware.Before)
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
			OperationName: "StartNextPendingJobExecution",
			Err:           err,
		}
	}
	out := result.(*StartNextPendingJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartNextPendingJobExecutionInput struct {

	// The name of the thing associated with the device.
	//
	// This member is required.
	ThingName *string

	// A collection of name/value pairs that describe the status of the job execution.
	// If not specified, the statusDetails are unchanged.
	StatusDetails map[string]*string

	// Specifies the amount of time this device has to finish execution of this job. If
	// the job execution status is not set to a terminal state before this timer
	// expires, or before the timer is reset (by calling UpdateJobExecution, setting
	// the status to IN_PROGRESS and specifying a new timeout value in field
	// stepTimeoutInMinutes) the job execution status will be automatically set to
	// TIMED_OUT. Note that setting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created (CreateJob using
	// field timeoutConfig).
	StepTimeoutInMinutes *int64
}

type StartNextPendingJobExecutionOutput struct {

	// A JobExecution object.
	Execution *types.JobExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartNextPendingJobExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartNextPendingJobExecution{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartNextPendingJobExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartNextPendingJobExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot-jobs-data",
		OperationName: "StartNextPendingJobExecution",
	}
}
