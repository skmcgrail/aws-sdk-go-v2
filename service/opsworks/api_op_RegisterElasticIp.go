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

// Registers an Elastic IP address with a specified stack. An address can be
// registered with only one stack at a time. If the address is already registered,
// you must first deregister it by calling DeregisterElasticIp (). For more
// information, see Resource Management
// (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html). Required
// Permissions: To use this action, an IAM user must have a Manage permissions
// level for the stack, or an attached policy that explicitly grants permissions.
// For more information on user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) RegisterElasticIp(ctx context.Context, params *RegisterElasticIpInput, optFns ...func(*Options)) (*RegisterElasticIpOutput, error) {
	stack := middleware.NewStack("RegisterElasticIp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterElasticIpMiddlewares(stack)
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
	addOpRegisterElasticIpValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterElasticIp(options.Region), middleware.Before)
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
			OperationName: "RegisterElasticIp",
			Err:           err,
		}
	}
	out := result.(*RegisterElasticIpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterElasticIpInput struct {

	// The Elastic IP address.
	//
	// This member is required.
	ElasticIp *string

	// The stack ID.
	//
	// This member is required.
	StackId *string
}

// Contains the response to a RegisterElasticIp request.
type RegisterElasticIpOutput struct {

	// The Elastic IP address.
	ElasticIp *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterElasticIpMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterElasticIp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterElasticIp{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterElasticIp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "RegisterElasticIp",
	}
}
