// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes an output from an existing flow. This request can be made only on an
// output that does not have an entitlement associated with it. If the output has
// an entitlement, you must revoke the entitlement instead. When an entitlement is
// revoked from a flow, the service automatically removes the associated output.
func (c *Client) RemoveFlowOutput(ctx context.Context, params *RemoveFlowOutputInput, optFns ...func(*Options)) (*RemoveFlowOutputOutput, error) {
	stack := middleware.NewStack("RemoveFlowOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveFlowOutputMiddlewares(stack)
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
	addOpRemoveFlowOutputValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveFlowOutput(options.Region), middleware.Before)
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
			OperationName: "RemoveFlowOutput",
			Err:           err,
		}
	}
	out := result.(*RemoveFlowOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveFlowOutputInput struct {

	// The flow that you want to remove an output from.
	//
	// This member is required.
	FlowArn *string

	// The ARN of the output that you want to remove.
	//
	// This member is required.
	OutputArn *string
}

type RemoveFlowOutputOutput struct {

	// The ARN of the flow that is associated with the output you removed.
	FlowArn *string

	// The ARN of the output that was removed.
	OutputArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveFlowOutputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveFlowOutput{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveFlowOutput{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveFlowOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "RemoveFlowOutput",
	}
}
