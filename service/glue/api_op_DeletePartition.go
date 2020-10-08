// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specified partition.
func (c *Client) DeletePartition(ctx context.Context, params *DeletePartitionInput, optFns ...func(*Options)) (*DeletePartitionOutput, error) {
	stack := middleware.NewStack("DeletePartition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeletePartitionMiddlewares(stack)
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
	addOpDeletePartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePartition(options.Region), middleware.Before)
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
			OperationName: "DeletePartition",
			Err:           err,
		}
	}
	out := result.(*DeletePartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePartitionInput struct {

	// The name of the catalog database in which the table in question resides.
	//
	// This member is required.
	DatabaseName *string

	// The values that define the partition.
	//
	// This member is required.
	PartitionValues []*string

	// The name of the table that contains the partition to be deleted.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partition to be deleted resides. If none is
	// provided, the AWS account ID is used by default.
	CatalogId *string
}

type DeletePartitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeletePartitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePartition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePartition{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeletePartition",
	}
}
