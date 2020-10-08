// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds VPC interfaces to flow
func (c *Client) AddFlowVpcInterfaces(ctx context.Context, params *AddFlowVpcInterfacesInput, optFns ...func(*Options)) (*AddFlowVpcInterfacesOutput, error) {
	stack := middleware.NewStack("AddFlowVpcInterfaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAddFlowVpcInterfacesMiddlewares(stack)
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
	addOpAddFlowVpcInterfacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddFlowVpcInterfaces(options.Region), middleware.Before)
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
			OperationName: "AddFlowVpcInterfaces",
			Err:           err,
		}
	}
	out := result.(*AddFlowVpcInterfacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to add VPC interfaces to the flow.
type AddFlowVpcInterfacesInput struct {

	// The flow that you want to mutate.
	//
	// This member is required.
	FlowArn *string

	// A list of VPC interfaces that you want to add.
	//
	// This member is required.
	VpcInterfaces []*types.VpcInterfaceRequest
}

type AddFlowVpcInterfacesOutput struct {

	// The ARN of the flow that these VPC interfaces were added to.
	FlowArn *string

	// The details of the newly added VPC interfaces.
	VpcInterfaces []*types.VpcInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAddFlowVpcInterfacesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAddFlowVpcInterfaces{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAddFlowVpcInterfaces{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddFlowVpcInterfaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "AddFlowVpcInterfaces",
	}
}
