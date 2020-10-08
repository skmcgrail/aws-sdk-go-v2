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
	"time"
)

// Creates a journal stream for a given Amazon QLDB ledger. The stream captures
// every document revision that is committed to the ledger's journal and delivers
// the data to a specified Amazon Kinesis Data Streams resource.
func (c *Client) StreamJournalToKinesis(ctx context.Context, params *StreamJournalToKinesisInput, optFns ...func(*Options)) (*StreamJournalToKinesisOutput, error) {
	stack := middleware.NewStack("StreamJournalToKinesis", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStreamJournalToKinesisMiddlewares(stack)
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
	addOpStreamJournalToKinesisValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStreamJournalToKinesis(options.Region), middleware.Before)
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
			OperationName: "StreamJournalToKinesis",
			Err:           err,
		}
	}
	out := result.(*StreamJournalToKinesisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StreamJournalToKinesisInput struct {

	// The inclusive start date and time from which to start streaming journal data.
	// This parameter must be in ISO 8601 date and time format and in Universal
	// Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z The InclusiveStartTime
	// cannot be in the future and must be before ExclusiveEndTime. If you provide an
	// InclusiveStartTime that is before the ledger's CreationDateTime, QLDB
	// effectively defaults it to the ledger's CreationDateTime.
	//
	// This member is required.
	InclusiveStartTime *time.Time

	// The configuration settings of the Kinesis Data Streams destination for your
	// stream request.
	//
	// This member is required.
	KinesisConfiguration *types.KinesisConfiguration

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for
	// a journal stream to write data records to a Kinesis Data Streams resource.
	//
	// This member is required.
	RoleArn *string

	// The name that you want to assign to the QLDB journal stream. User-defined names
	// can help identify and indicate the purpose of a stream. Your stream name must be
	// unique among other active streams for a given ledger. Stream names have the same
	// naming constraints as ledger names, as defined in Quotas in Amazon QLDB
	// (https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming)
	// in the Amazon QLDB Developer Guide.
	//
	// This member is required.
	StreamName *string

	// The exclusive date and time that specifies when the stream ends. If you don't
	// define this parameter, the stream runs indefinitely until you cancel it. The
	// ExclusiveEndTime must be in ISO 8601 date and time format and in Universal
	// Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z
	ExclusiveEndTime *time.Time

	// The key-value pairs to add as tags to the stream that you want to create. Tag
	// keys are case sensitive. Tag values are case sensitive and can be null.
	Tags map[string]*string
}

type StreamJournalToKinesisOutput struct {

	// The unique ID that QLDB assigns to each QLDB journal stream.
	StreamId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStreamJournalToKinesisMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStreamJournalToKinesis{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStreamJournalToKinesis{}, middleware.After)
}

func newServiceMetadataMiddleware_opStreamJournalToKinesis(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "StreamJournalToKinesis",
	}
}
