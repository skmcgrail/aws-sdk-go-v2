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

// Enables a VPC for ClassicLink. You can then link EC2-Classic instances to your
// ClassicLink-enabled VPC to allow communication over private IP addresses. You
// cannot enable your VPC for ClassicLink if any of your VPC route tables have
// existing routes for address ranges within the 10.0.0.0/8 IP address range,
// excluding local routes for VPCs in the 10.0.0.0/16 and 10.1.0.0/16 IP address
// ranges. For more information, see ClassicLink
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) EnableVpcClassicLink(ctx context.Context, params *EnableVpcClassicLinkInput, optFns ...func(*Options)) (*EnableVpcClassicLinkOutput, error) {
	stack := middleware.NewStack("EnableVpcClassicLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpEnableVpcClassicLinkMiddlewares(stack)
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
	addOpEnableVpcClassicLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableVpcClassicLink(options.Region), middleware.Before)
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
			OperationName: "EnableVpcClassicLink",
			Err:           err,
		}
	}
	out := result.(*EnableVpcClassicLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableVpcClassicLinkInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the VPC.
	VpcId *string
}

type EnableVpcClassicLinkOutput struct {
	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpEnableVpcClassicLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpEnableVpcClassicLink{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpEnableVpcClassicLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableVpcClassicLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableVpcClassicLink",
	}
}