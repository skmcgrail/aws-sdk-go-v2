// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies the specified Amazon FPGA Image (AFI) to the current Region.
func (c *Client) CopyFpgaImage(ctx context.Context, params *CopyFpgaImageInput, optFns ...func(*Options)) (*CopyFpgaImageOutput, error) {
	stack := middleware.NewStack("CopyFpgaImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCopyFpgaImageMiddlewares(stack)
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
	addOpCopyFpgaImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyFpgaImage(options.Region), middleware.Before)
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
			OperationName: "CopyFpgaImage",
			Err:           err,
		}
	}
	out := result.(*CopyFpgaImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyFpgaImageInput struct {

	// The ID of the source AFI.
	//
	// This member is required.
	SourceFpgaImageId *string

	// The Region that contains the source AFI.
	//
	// This member is required.
	SourceRegion *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string

	// The description for the new AFI.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The name for the new AFI. The default is the name of the source AFI.
	Name *string
}

type CopyFpgaImageOutput struct {

	// The ID of the new AFI.
	FpgaImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCopyFpgaImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCopyFpgaImage{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCopyFpgaImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyFpgaImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CopyFpgaImage",
	}
}
