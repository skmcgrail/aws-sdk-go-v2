// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a journal export job, including the ledger name,
// export ID, when it was created, current status, and its start and end time
// export parameters. This action does not return any expired export jobs. For more
// information, see Export Job Expiration
// (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide. If the export job with the given ExportId
// doesn't exist, then throws ResourceNotFoundException. If the ledger with the
// given Name doesn't exist, then throws ResourceNotFoundException.
func (c *Client) DescribeJournalS3Export(ctx context.Context, params *DescribeJournalS3ExportInput, optFns ...func(*Options)) (*DescribeJournalS3ExportOutput, error) {
	stack := middleware.NewStack("DescribeJournalS3Export", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeJournalS3ExportMiddlewares(stack)
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
	addOpDescribeJournalS3ExportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJournalS3Export(options.Region), middleware.Before)
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
			OperationName: "DescribeJournalS3Export",
			Err:           err,
		}
	}
	out := result.(*DescribeJournalS3ExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJournalS3ExportInput struct {

	// The unique ID of the journal export job that you want to describe.
	//
	// This member is required.
	ExportId *string

	// The name of the ledger.
	//
	// This member is required.
	Name *string
}

type DescribeJournalS3ExportOutput struct {

	// Information about the journal export job returned by a DescribeJournalS3Export
	// request.
	//
	// This member is required.
	ExportDescription *types.JournalS3ExportDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeJournalS3ExportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJournalS3Export{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJournalS3Export{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeJournalS3Export(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "DescribeJournalS3Export",
	}
}
