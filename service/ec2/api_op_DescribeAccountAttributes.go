// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes attributes of your AWS account. The following are the supported
// account attributes:
//
//     * supported-platforms: Indicates whether your account
// can launch instances into EC2-Classic and EC2-VPC, or only into EC2-VPC.
//
//     *
// default-vpc: The ID of the default VPC for your account, or none.
//
//     *
// max-instances: This attribute is no longer supported. The returned value does
// not reflect your actual vCPU limit for running On-Demand Instances. For more
// information, see On-Demand Instance Limits
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-on-demand-instances.html#ec2-on-demand-instances-limits)
// in the Amazon Elastic Compute Cloud User Guide.
//
//     *
// vpc-max-security-groups-per-interface: The maximum number of security groups
// that you can assign to a network interface.
//
//     * max-elastic-ips: The maximum
// number of Elastic IP addresses that you can allocate for use with EC2-Classic.
//
//
// * vpc-max-elastic-ips: The maximum number of Elastic IP addresses that you can
// allocate for use with EC2-VPC.
func (c *Client) DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) {
	stack := middleware.NewStack("DescribeAccountAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeAccountAttributesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAttributes(options.Region), middleware.Before)
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
			OperationName: "DescribeAccountAttributes",
			Err:           err,
		}
	}
	out := result.(*DescribeAccountAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAttributesInput struct {

	// The account attribute names.
	AttributeNames []types.AccountAttributeName

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeAccountAttributesOutput struct {

	// Information about the account attributes.
	AccountAttributes []*types.AccountAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeAccountAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeAccountAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAccountAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccountAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeAccountAttributes",
	}
}
