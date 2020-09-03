// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a logical container where backups are stored. A CreateBackupVault
// request includes a name, optionally one or more resource tags, an encryption
// key, and a request ID. Sensitive data, such as passport numbers, should not be
// included the name of a backup vault.
func (c *Client) CreateBackupVault(ctx context.Context, params *CreateBackupVaultInput, optFns ...func(*Options)) (*CreateBackupVaultOutput, error) {
	stack := middleware.NewStack("CreateBackupVault", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateBackupVaultMiddlewares(stack)
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
	addOpCreateBackupVaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackupVault(options.Region), middleware.Before)

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
			OperationName: "CreateBackupVault",
			Err:           err,
		}
	}
	out := result.(*CreateBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupVaultInput struct {
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string
	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string
	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string
	// Metadata that you can assign to help organize the resources that you create.
	// Each tag is a key-value pair.
	BackupVaultTags map[string]*string
}

type CreateBackupVaultOutput struct {
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created. They consist of lowercase letters, numbers, and
	// hyphens.
	BackupVaultName *string
	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string
	// The date and time a backup vault is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateBackupVaultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateBackupVault{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBackupVault{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBackupVault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "CreateBackupVault",
	}
}