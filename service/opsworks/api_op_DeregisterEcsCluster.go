// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters a specified Amazon ECS cluster from a stack. For more information,
// see  Resource Management
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-ecscluster.html#workinglayers-ecscluster-delete).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information on user permissions, see
// https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DeregisterEcsCluster(ctx context.Context, params *DeregisterEcsClusterInput, optFns ...func(*Options)) (*DeregisterEcsClusterOutput, error) {
	stack := middleware.NewStack("DeregisterEcsCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterEcsClusterMiddlewares(stack)
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
	addOpDeregisterEcsClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterEcsCluster(options.Region), middleware.Before)
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
			OperationName: "DeregisterEcsCluster",
			Err:           err,
		}
	}
	out := result.(*DeregisterEcsClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterEcsClusterInput struct {

	// The cluster's Amazon Resource Number (ARN).
	//
	// This member is required.
	EcsClusterArn *string
}

type DeregisterEcsClusterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterEcsClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterEcsCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterEcsCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterEcsCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DeregisterEcsCluster",
	}
}
