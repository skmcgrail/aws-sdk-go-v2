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

// Starts a new run of the specified workflow.
func (c *Client) StartWorkflowRun(ctx context.Context, params *StartWorkflowRunInput, optFns ...func(*Options)) (*StartWorkflowRunOutput, error) {
	stack := middleware.NewStack("StartWorkflowRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartWorkflowRunMiddlewares(stack)
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
	addOpStartWorkflowRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartWorkflowRun(options.Region), middleware.Before)
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
			OperationName: "StartWorkflowRun",
			Err:           err,
		}
	}
	out := result.(*StartWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartWorkflowRunInput struct {

	// The name of the workflow to start.
	//
	// This member is required.
	Name *string
}

type StartWorkflowRunOutput struct {

	// An Id for the new run.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartWorkflowRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartWorkflowRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartWorkflowRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartWorkflowRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartWorkflowRun",
	}
}
