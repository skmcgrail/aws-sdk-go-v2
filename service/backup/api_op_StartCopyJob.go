// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a job to create a one-time copy of the specified resource. Does not
// support continuous backups.
func (c *Client) StartCopyJob(ctx context.Context, params *StartCopyJobInput, optFns ...func(*Options)) (*StartCopyJobOutput, error) {
	if params == nil {
		params = &StartCopyJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartCopyJob", params, optFns, c.addOperationStartCopyJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartCopyJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCopyJobInput struct {

	// An Amazon Resource Name (ARN) that uniquely identifies a destination backup
	// vault to copy to; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	//
	// This member is required.
	DestinationBackupVaultArn *string

	// Specifies the IAM role ARN used to copy the target recovery point; for example,
	// arn:aws:iam::123456789012:role/S3Access.
	//
	// This member is required.
	IamRoleArn *string

	// An ARN that uniquely identifies a recovery point to use for the copy job; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	//
	// This member is required.
	RecoveryPointArn *string

	// The name of a logical source container where backups are stored. Backup vaults
	// are identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// This member is required.
	SourceBackupVaultName *string

	// A customer chosen string that can be used to distinguish between calls to
	// StartCopyJob.
	IdempotencyToken *string

	// Contains an array of Transition objects specifying how long in days before a
	// recovery point transitions to cold storage or is deleted. Backups transitioned
	// to cold storage must be stored in cold storage for a minimum of 90 days.
	// Therefore, on the console, the “expire after days” setting must be 90 days
	// greater than the “transition to cold after days” setting. The “transition to
	// cold after days” setting cannot be changed after a backup has been transitioned
	// to cold. Only Amazon EFS file system backups can be transitioned to cold
	// storage.
	Lifecycle *types.Lifecycle

	noSmithyDocumentSerde
}

type StartCopyJobOutput struct {

	// Uniquely identifies a copy job.
	CopyJobId *string

	// The date and time that a copy job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartCopyJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartCopyJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartCopyJob{}, middleware.After)
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
	if err = addOpStartCopyJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartCopyJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartCopyJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "StartCopyJob",
	}
}
