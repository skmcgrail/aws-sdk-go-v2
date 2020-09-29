// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the attributes of your VPC endpoint service configuration. You can
// change the Network Load Balancers for your service, and you can specify whether
// acceptance is required for requests to connect to your endpoint service through
// an interface VPC endpoint. If you set or modify the private DNS name, you must
// prove that you own the private DNS domain name. For more information, see VPC
// Endpoint Service Private DNS Name Verification
// (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-services-dns-validation.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) ModifyVpcEndpointServiceConfiguration(ctx context.Context, params *ModifyVpcEndpointServiceConfigurationInput, optFns ...func(*Options)) (*ModifyVpcEndpointServiceConfigurationOutput, error) {
	stack := middleware.NewStack("ModifyVpcEndpointServiceConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyVpcEndpointServiceConfigurationMiddlewares(stack)
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
	addOpModifyVpcEndpointServiceConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcEndpointServiceConfiguration(options.Region), middleware.Before)
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
			OperationName: "ModifyVpcEndpointServiceConfiguration",
			Err:           err,
		}
	}
	out := result.(*ModifyVpcEndpointServiceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcEndpointServiceConfigurationInput struct {
	// The Amazon Resource Names (ARNs) of Network Load Balancers to add to your
	// service configuration.
	AddNetworkLoadBalancerArns []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The Amazon Resource Names (ARNs) of Network Load Balancers to remove from your
	// service configuration.
	RemoveNetworkLoadBalancerArns []*string
	// The ID of the service.
	ServiceId *string
	// The private DNS name to assign to the endpoint service.
	PrivateDnsName *string
	// Indicates whether requests to create an endpoint to your service must be
	// accepted.
	AcceptanceRequired *bool
	// Removes the private DNS name of the endpoint service.
	RemovePrivateDnsName *bool
}

type ModifyVpcEndpointServiceConfigurationOutput struct {
	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyVpcEndpointServiceConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcEndpointServiceConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcEndpointServiceConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyVpcEndpointServiceConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcEndpointServiceConfiguration",
	}
}