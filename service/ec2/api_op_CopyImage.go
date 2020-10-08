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

// Initiates the copy of an AMI from the specified source Region to the current
// Region. You specify the destination Region by using its endpoint when making the
// request. Copies of encrypted backing snapshots for the AMI are encrypted. Copies
// of unencrypted backing snapshots remain unencrypted, unless you set Encrypted
// during the copy operation. You cannot create an unencrypted copy of an encrypted
// backing snapshot. For more information about the prerequisites and limits when
// copying an AMI, see Copying an AMI
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/CopyingAMIs.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) CopyImage(ctx context.Context, params *CopyImageInput, optFns ...func(*Options)) (*CopyImageOutput, error) {
	stack := middleware.NewStack("CopyImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCopyImageMiddlewares(stack)
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
	addOpCopyImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyImage(options.Region), middleware.Before)
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
			OperationName: "CopyImage",
			Err:           err,
		}
	}
	out := result.(*CopyImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CopyImage.
type CopyImageInput struct {

	// The name of the new AMI in the destination Region.
	//
	// This member is required.
	Name *string

	// The ID of the AMI to copy.
	//
	// This member is required.
	SourceImageId *string

	// The name of the Region that contains the AMI to copy.
	//
	// This member is required.
	SourceRegion *string

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	ClientToken *string

	// A description for the new AMI in the destination Region.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Specifies whether the destination snapshots of the copied image should be
	// encrypted. You can encrypt a copy of an unencrypted snapshot, but you cannot
	// create an unencrypted copy of an encrypted snapshot. The default CMK for EBS is
	// used unless you specify a non-default AWS Key Management Service (AWS KMS) CMK
	// using KmsKeyId. For more information, see Amazon EBS Encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool

	// An identifier for the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) to use when creating the encrypted volume. This parameter is
	// only required if you want to use a non-default CMK; if this parameter is not
	// specified, the default CMK for EBS is used. If a KmsKeyId is specified, the
	// Encrypted flag must also be set. To specify a CMK, use its key ID, Amazon
	// Resource Name (ARN), alias name, or alias ARN. When using an alias name, prefix
	// it with "alias/". For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//
	// * Alias name: alias/ExampleAlias
	//
	//     * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// AWS parses KmsKeyId
	// asynchronously, meaning that the action you call may appear to complete even
	// though you provided an invalid identifier. This action will eventually report
	// failure. The specified CMK must exist in the Region that the snapshot is being
	// copied to. Amazon EBS does not support asymmetric CMKs.
	KmsKeyId *string
}

// Contains the output of CopyImage.
type CopyImageOutput struct {

	// The ID of the new AMI.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCopyImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCopyImage{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCopyImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CopyImage",
	}
}
