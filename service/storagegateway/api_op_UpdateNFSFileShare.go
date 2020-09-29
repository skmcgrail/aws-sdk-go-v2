// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a Network File System (NFS) file share. This operation is only supported
// in the file gateway type.  <note> <p>To leave a file share field unchanged, set
// the corresponding input field to null.</p> </note> <p>Updates the following file
// share setting:</p> <ul> <li> <p>Default storage class for your S3 bucket</p>
// </li> <li> <p>Metadata defaults for your S3 bucket</p> </li> <li> <p>Allowed NFS
// clients for your file share</p> </li> <li> <p>Squash settings</p> </li> <li>
// <p>Write status of your file share</p> </li> </ul> <note> <p>To leave a file
// share field unchanged, set the corresponding input field to null. This operation
// is only supported in file gateways.</p> </note>
func (c *Client) UpdateNFSFileShare(ctx context.Context, params *UpdateNFSFileShareInput, optFns ...func(*Options)) (*UpdateNFSFileShareOutput, error) {
	stack := middleware.NewStack("UpdateNFSFileShare", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateNFSFileShareMiddlewares(stack)
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
	addOpUpdateNFSFileShareValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNFSFileShare(options.Region), middleware.Before)
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
			OperationName: "UpdateNFSFileShare",
			Err:           err,
		}
	}
	out := result.(*UpdateNFSFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateNFSFileShareInput
type UpdateNFSFileShareInput struct {
	// A value that sets the access control list (ACL) permission for objects in the S3
	// bucket that a file gateway puts objects into. The default value is private.
	ObjectACL types.ObjectACL
	// The name of the file share. Optional.  <note> <p> <code>FileShareName</code>
	// must be set if an S3 prefix name is set in <code>LocationARN</code>.</p> </note>
	FileShareName *string
	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false. The default value is true.  <p>Valid Values: <code>true</code> |
	// <code>false</code> </p>
	GuessMIMETypeEnabled *bool
	// The list of clients that are allowed to access the file gateway. The list must
	// contain either valid IP addresses or valid CIDR blocks.
	ClientList []*string
	// The default values for the file share. Optional.
	NFSFileShareDefaults *types.NFSFileShareDefaults
	// A value that sets the write status of a file share. Set this value to true to
	// set the write status to read-only, otherwise set to false.  <p>Valid Values:
	// <code>true</code> | <code>false</code> </p>
	ReadOnly *bool
	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true, the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data.  <note> <p>
	// <code>RequesterPays</code> is a configuration for the S3 bucket that backs the
	// file share, so make sure that the configuration on the file share is the same as
	// the S3 bucket configuration.</p> </note> <p>Valid Values: <code>true</code> |
	// <code>false</code> </p>
	RequesterPays *bool
	// Refresh cache information.
	CacheAttributes *types.CacheAttributes
	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string
	// The Amazon Resource Name (ARN) of the file share to be updated.
	FileShareARN *string
	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS key,
	// or false to use a key managed by Amazon S3. Optional.  <p>Valid Values:
	// <code>true</code> | <code>false</code> </p>
	KMSEncrypted *bool
	// The user mapped to anonymous user.  <p>Valid values are the following:</p> <ul>
	// <li> <p> <code>RootSquash</code>: Only root is mapped to anonymous user.</p>
	// </li> <li> <p> <code>NoSquash</code>: No one is mapped to anonymous user.</p>
	// </li> <li> <p> <code>AllSquash</code>: Everyone is mapped to anonymous user.</p>
	// </li> </ul>
	Squash *string
	// The default storage class for objects put into an Amazon S3 bucket by the file
	// gateway. The default value is S3_INTELLIGENT_TIERING. Optional.  <p>Valid
	// Values: <code>S3_STANDARD</code> | <code>S3_INTELLIGENT_TIERING</code> |
	// <code>S3_STANDARD_IA</code> | <code>S3_ONEZONE_IA</code> </p>
	DefaultStorageClass *string
}

// UpdateNFSFileShareOutput
type UpdateNFSFileShareOutput struct {
	// The Amazon Resource Name (ARN) of the updated file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateNFSFileShareMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNFSFileShare{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNFSFileShare{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateNFSFileShare(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateNFSFileShare",
	}
}