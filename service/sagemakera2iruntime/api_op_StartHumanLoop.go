// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a human loop, provided that at least one activation condition is met.
func (c *Client) StartHumanLoop(ctx context.Context, params *StartHumanLoopInput, optFns ...func(*Options)) (*StartHumanLoopOutput, error) {
	stack := middleware.NewStack("StartHumanLoop", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartHumanLoopMiddlewares(stack)
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
	addOpStartHumanLoopValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartHumanLoop(options.Region), middleware.Before)
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
			OperationName: "StartHumanLoop",
			Err:           err,
		}
	}
	out := result.(*StartHumanLoopOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartHumanLoopInput struct {

	// The Amazon Resource Name (ARN) of the flow definition associated with this human
	// loop.
	//
	// This member is required.
	FlowDefinitionArn *string

	// An object that contains information about the human loop.
	//
	// This member is required.
	HumanLoopInput *types.HumanLoopInput

	// The name of the human loop.
	//
	// This member is required.
	HumanLoopName *string

	// Attributes of the specified data. Use DataAttributes to specify if your data is
	// free of personally identifiable information and/or free of adult content.
	DataAttributes *types.HumanLoopDataAttributes
}

type StartHumanLoopOutput struct {

	// The Amazon Resource Name (ARN) of the human loop.
	HumanLoopArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartHumanLoopMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartHumanLoop{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartHumanLoop{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartHumanLoop(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StartHumanLoop",
	}
}
