// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Infers a schema for an SQL-based Amazon Kinesis Data Analytics application by
// evaluating sample records on the specified streaming source (Kinesis data stream
// or Kinesis Data Firehose delivery stream) or Amazon S3 object. In the response,
// the operation returns the inferred schema and also the sample records that the
// operation used to infer the schema. You can use the inferred schema when
// configuring a streaming source for your application. When you create an
// application using the Kinesis Data Analytics console, the console uses this
// operation to infer a schema and show it in the console user interface.
func (c *Client) DiscoverInputSchema(ctx context.Context, params *DiscoverInputSchemaInput, optFns ...func(*Options)) (*DiscoverInputSchemaOutput, error) {
	stack := middleware.NewStack("DiscoverInputSchema", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDiscoverInputSchemaMiddlewares(stack)
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
	addOpDiscoverInputSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDiscoverInputSchema(options.Region), middleware.Before)
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
			OperationName: "DiscoverInputSchema",
			Err:           err,
		}
	}
	out := result.(*DiscoverInputSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DiscoverInputSchemaInput struct {

	// The ARN of the role that is used to access the streaming source.
	//
	// This member is required.
	ServiceExecutionRole *string

	// The InputProcessingConfiguration () to use to preprocess the records before
	// discovering the schema of the records.
	InputProcessingConfiguration *types.InputProcessingConfiguration

	// The point at which you want Kinesis Data Analytics to start reading records from
	// the specified streaming source discovery purposes.
	InputStartingPositionConfiguration *types.InputStartingPositionConfiguration

	// The Amazon Resource Name (ARN) of the streaming source.
	ResourceARN *string

	// Specify this parameter to discover a schema from data in an Amazon S3 object.
	S3Configuration *types.S3Configuration
}

type DiscoverInputSchemaOutput struct {

	// The schema inferred from the streaming source. It identifies the format of the
	// data in the streaming source and how each data element maps to corresponding
	// columns in the in-application stream that you can create.
	InputSchema *types.SourceSchema

	// An array of elements, where each element corresponds to a row in a stream record
	// (a stream record can have more than one row).
	ParsedInputRecords [][]*string

	// The stream data that was modified by the processor specified in the
	// InputProcessingConfiguration parameter.
	ProcessedInputRecords []*string

	// The raw stream data that was sampled to infer the schema.
	RawInputRecords []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDiscoverInputSchemaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDiscoverInputSchema{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDiscoverInputSchema{}, middleware.After)
}

func newServiceMetadataMiddleware_opDiscoverInputSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DiscoverInputSchema",
	}
}
