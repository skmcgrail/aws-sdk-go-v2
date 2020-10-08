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

// Removes a table definition from the Data Catalog. After completing this
// operation, you no longer have access to the table versions and partitions that
// belong to the deleted table. AWS Glue deletes these "orphaned" resources
// asynchronously in a timely manner, at the discretion of the service. To ensure
// the immediate deletion of all related resources, before calling DeleteTable, use
// DeleteTableVersion or BatchDeleteTableVersion, and DeletePartition or
// BatchDeletePartition, to delete any resources that belong to the table.
func (c *Client) DeleteTable(ctx context.Context, params *DeleteTableInput, optFns ...func(*Options)) (*DeleteTableOutput, error) {
	stack := middleware.NewStack("DeleteTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteTableMiddlewares(stack)
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
	addOpDeleteTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTable(options.Region), middleware.Before)
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
			OperationName: "DeleteTable",
			Err:           err,
		}
	}
	out := result.(*DeleteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTableInput struct {

	// The name of the catalog database in which the table resides. For Hive
	// compatibility, this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table to be deleted. For Hive compatibility, this name is
	// entirely lowercase.
	//
	// This member is required.
	Name *string

	// The ID of the Data Catalog where the table resides. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string
}

type DeleteTableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteTable",
	}
}
