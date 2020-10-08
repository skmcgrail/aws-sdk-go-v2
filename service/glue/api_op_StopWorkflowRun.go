// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops the execution of the specified workflow run.
func (c *Client) StopWorkflowRun(ctx context.Context, params *StopWorkflowRunInput, optFns ...func(*Options)) (*StopWorkflowRunOutput, error) {
	stack := middleware.NewStack("StopWorkflowRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopWorkflowRunMiddlewares(stack)
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
	addOpStopWorkflowRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopWorkflowRun(options.Region), middleware.Before)
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
			OperationName: "StopWorkflowRun",
			Err:           err,
		}
	}
	out := result.(*StopWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopWorkflowRunInput struct {

	// The name of the workflow to stop.
	//
	// This member is required.
	Name *string

	// The ID of the workflow run to stop.
	//
	// This member is required.
	RunId *string
}

type StopWorkflowRunOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopWorkflowRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopWorkflowRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopWorkflowRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopWorkflowRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StopWorkflowRun",
	}
}
