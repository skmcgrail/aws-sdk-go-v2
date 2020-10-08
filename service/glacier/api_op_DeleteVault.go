// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation deletes a vault. Amazon S3 Glacier will delete a vault only if
// there are no archives in the vault as of the last inventory and there have been
// no writes to the vault since the last inventory. If either of these conditions
// is not satisfied, the vault deletion fails (that is, the vault is not removed)
// and Amazon S3 Glacier returns an error. You can use DescribeVault () to return
// the number of archives in a vault, and you can use Initiate a Job (POST jobs)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-initiate-job-post.html)
// to initiate a new inventory retrieval for a vault. The inventory contains the
// archive IDs you use to delete archives using Delete Archive (DELETE archive)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-delete.html).
// <p>This operation is idempotent.</p> <p>An AWS account has full permission to
// perform all operations (actions). However, AWS Identity and Access Management
// (IAM) users don't have any permissions by default. You must grant them explicit
// permission to perform specific actions. For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p> For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/deleting-vaults.html">Deleting
// a Vault in Amazon Glacier</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-delete.html">Delete
// Vault </a> in the <i>Amazon S3 Glacier Developer Guide</i>. </p>
func (c *Client) DeleteVault(ctx context.Context, params *DeleteVaultInput, optFns ...func(*Options)) (*DeleteVaultOutput, error) {
	stack := middleware.NewStack("DeleteVault", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteVaultMiddlewares(stack)
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
	addOpDeleteVaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVault(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)

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
			OperationName: "DeleteVault",
			Err:           err,
		}
	}
	out := result.(*DeleteVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for deleting a vault from Amazon S3 Glacier.
type DeleteVaultInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

type DeleteVaultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteVaultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVault{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVault{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "DeleteVault",
	}
}
