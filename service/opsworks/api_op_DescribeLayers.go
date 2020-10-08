// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Requests a description of one or more layers in a specified stack. This call
// accepts only one resource-identifying parameter. Required Permissions: To use
// this action, an IAM user must have a Show, Deploy, or Manage permissions level
// for the stack, or an attached policy that explicitly grants permissions. For
// more information about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeLayers(ctx context.Context, params *DescribeLayersInput, optFns ...func(*Options)) (*DescribeLayersOutput, error) {
	stack := middleware.NewStack("DescribeLayers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeLayersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLayers(options.Region), middleware.Before)
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
			OperationName: "DescribeLayers",
			Err:           err,
		}
	}
	out := result.(*DescribeLayersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLayersInput struct {

	// An array of layer IDs that specify the layers to be described. If you omit this
	// parameter, DescribeLayers returns a description of every layer in the specified
	// stack.
	LayerIds []*string

	// The stack ID.
	StackId *string
}

// Contains the response to a DescribeLayers request.
type DescribeLayersOutput struct {

	// An array of Layer objects that describe the layers.
	Layers []*types.Layer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeLayersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLayers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLayers{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLayers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeLayers",
	}
}
