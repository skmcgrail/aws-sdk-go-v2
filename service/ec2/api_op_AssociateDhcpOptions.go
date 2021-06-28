// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a set of DHCP options (that you've previously created) with the
// specified VPC, or associates no DHCP options with the VPC. After you associate
// the options with the VPC, any existing instances and all new instances that you
// launch in that VPC use the options. You don't need to restart or relaunch the
// instances. They automatically pick up the changes within a few hours, depending
// on how frequently the instance renews its DHCP lease. You can explicitly renew
// the lease using the operating system on the instance. For more information, see
// DHCP Options Sets
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) AssociateDhcpOptions(ctx context.Context, params *AssociateDhcpOptionsInput, optFns ...func(*Options)) (*AssociateDhcpOptionsOutput, error) {
	if params == nil {
		params = &AssociateDhcpOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateDhcpOptions", params, optFns, c.addOperationAssociateDhcpOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateDhcpOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDhcpOptionsInput struct {

	// The ID of the DHCP options set, or default to associate no DHCP options with the
	// VPC.
	//
	// This member is required.
	DhcpOptionsId *string

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type AssociateDhcpOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateDhcpOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateDhcpOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateDhcpOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateDhcpOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDhcpOptions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateDhcpOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateDhcpOptions",
	}
}
