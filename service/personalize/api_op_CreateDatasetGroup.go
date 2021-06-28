// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an empty dataset group. A dataset group contains related datasets that
// supply data for training a model. A dataset group can contain at most three
// datasets, one for each type of dataset:
//
// * Interactions
//
// * Items
//
// * Users
//
// To
// train a model (create a solution), a dataset group that contains an Interactions
// dataset is required. Call CreateDataset to add a dataset to the group. A dataset
// group can be in one of the following states:
//
// * CREATE PENDING > CREATE
// IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// * DELETE PENDING
//
// To get the status of
// the dataset group, call DescribeDatasetGroup. If the status shows as CREATE
// FAILED, the response includes a failureReason key, which describes why the
// creation failed. You must wait until the status of the dataset group is ACTIVE
// before adding a dataset to the group. You can specify an AWS Key Management
// Service (KMS) key to encrypt the datasets in the group. If you specify a KMS
// key, you must also include an AWS Identity and Access Management (IAM) role that
// has permission to access the key. APIs that require a dataset group ARN in the
// request
//
// * CreateDataset
//
// * CreateEventTracker
//
// * CreateSolution
//
// Related
// APIs
//
// * ListDatasetGroups
//
// * DescribeDatasetGroup
//
// * DeleteDatasetGroup
func (c *Client) CreateDatasetGroup(ctx context.Context, params *CreateDatasetGroupInput, optFns ...func(*Options)) (*CreateDatasetGroupOutput, error) {
	if params == nil {
		params = &CreateDatasetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatasetGroup", params, optFns, c.addOperationCreateDatasetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetGroupInput struct {

	// The name for the new dataset group.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of a KMS key used to encrypt the datasets.
	KmsKeyArn *string

	// The ARN of the IAM role that has permissions to access the KMS key. Supplying an
	// IAM role is only valid when also specifying a KMS key.
	RoleArn *string

	noSmithyDocumentSerde
}

type CreateDatasetGroupOutput struct {

	// The Amazon Resource Name (ARN) of the new dataset group.
	DatasetGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatasetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDatasetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDatasetGroup{}, middleware.After)
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
	if err = addOpCreateDatasetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatasetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatasetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateDatasetGroup",
	}
}
