// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the settings of a data store.
func (c *Client) UpdateDatastore(ctx context.Context, params *UpdateDatastoreInput, optFns ...func(*Options)) (*UpdateDatastoreOutput, error) {
	if params == nil {
		params = &UpdateDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDatastore", params, optFns, c.addOperationUpdateDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatastoreInput struct {

	// The name of the data store to be updated.
	//
	// This member is required.
	DatastoreName *string

	// Where data store data is stored. You can choose one of serviceManagedS3 or
	// customerManagedS3 storage. If not specified, the default isserviceManagedS3. You
	// cannot change this storage option after the data store is created.
	DatastoreStorage types.DatastoreStorage

	// Contains the configuration information of file formats. AWS IoT Analytics data
	// stores support JSON and Parquet (https://parquet.apache.org/). The default file
	// format is JSON. You can specify only one format. You can't change the file
	// format after you create the data store.
	FileFormatConfiguration *types.FileFormatConfiguration

	// How long, in days, message data is kept for the data store. The retention period
	// cannot be updated if the data store's S3 storage is customer-managed.
	RetentionPeriod *types.RetentionPeriod

	noSmithyDocumentSerde
}

type UpdateDatastoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDatastore{}, middleware.After)
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
	if err = addOpUpdateDatastoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDatastore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDatastore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "UpdateDatastore",
	}
}
