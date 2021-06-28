// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of all Amazon QLDB journal stream descriptors for a given
// ledger. The output of each stream descriptor includes the same details that are
// returned by DescribeJournalKinesisStream. This action does not return any
// expired journal streams. For more information, see Expiration for terminal
// streams
// (https://docs.aws.amazon.com/qldb/latest/developerguide/streams.create.html#streams.create.states.expiration)
// in the Amazon QLDB Developer Guide. This action returns a maximum of MaxResults
// items. It is paginated so that you can retrieve all the items by calling
// ListJournalKinesisStreamsForLedger multiple times.
func (c *Client) ListJournalKinesisStreamsForLedger(ctx context.Context, params *ListJournalKinesisStreamsForLedgerInput, optFns ...func(*Options)) (*ListJournalKinesisStreamsForLedgerOutput, error) {
	if params == nil {
		params = &ListJournalKinesisStreamsForLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJournalKinesisStreamsForLedger", params, optFns, c.addOperationListJournalKinesisStreamsForLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJournalKinesisStreamsForLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJournalKinesisStreamsForLedgerInput struct {

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The maximum number of results to return in a single
	// ListJournalKinesisStreamsForLedger request. (The actual number of results
	// returned might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListJournalKinesisStreamsForLedger call, you should use that value as input
	// here.
	NextToken *string

	noSmithyDocumentSerde
}

type ListJournalKinesisStreamsForLedgerOutput struct {

	// * If NextToken is empty, the last page of results has been processed and there
	// are no more results to be retrieved.
	//
	// * If NextToken is not empty, more results
	// are available. To retrieve the next page of results, use the value of NextToken
	// in a subsequent ListJournalKinesisStreamsForLedger call.
	NextToken *string

	// The array of QLDB journal stream descriptors that are associated with the given
	// ledger.
	Streams []types.JournalKinesisStreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJournalKinesisStreamsForLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJournalKinesisStreamsForLedger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJournalKinesisStreamsForLedger{}, middleware.After)
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
	if err = addOpListJournalKinesisStreamsForLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJournalKinesisStreamsForLedger(options.Region), middleware.Before); err != nil {
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

// ListJournalKinesisStreamsForLedgerAPIClient is a client that implements the
// ListJournalKinesisStreamsForLedger operation.
type ListJournalKinesisStreamsForLedgerAPIClient interface {
	ListJournalKinesisStreamsForLedger(context.Context, *ListJournalKinesisStreamsForLedgerInput, ...func(*Options)) (*ListJournalKinesisStreamsForLedgerOutput, error)
}

var _ ListJournalKinesisStreamsForLedgerAPIClient = (*Client)(nil)

// ListJournalKinesisStreamsForLedgerPaginatorOptions is the paginator options for
// ListJournalKinesisStreamsForLedger
type ListJournalKinesisStreamsForLedgerPaginatorOptions struct {
	// The maximum number of results to return in a single
	// ListJournalKinesisStreamsForLedger request. (The actual number of results
	// returned might be fewer.)
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJournalKinesisStreamsForLedgerPaginator is a paginator for
// ListJournalKinesisStreamsForLedger
type ListJournalKinesisStreamsForLedgerPaginator struct {
	options   ListJournalKinesisStreamsForLedgerPaginatorOptions
	client    ListJournalKinesisStreamsForLedgerAPIClient
	params    *ListJournalKinesisStreamsForLedgerInput
	nextToken *string
	firstPage bool
}

// NewListJournalKinesisStreamsForLedgerPaginator returns a new
// ListJournalKinesisStreamsForLedgerPaginator
func NewListJournalKinesisStreamsForLedgerPaginator(client ListJournalKinesisStreamsForLedgerAPIClient, params *ListJournalKinesisStreamsForLedgerInput, optFns ...func(*ListJournalKinesisStreamsForLedgerPaginatorOptions)) *ListJournalKinesisStreamsForLedgerPaginator {
	if params == nil {
		params = &ListJournalKinesisStreamsForLedgerInput{}
	}

	options := ListJournalKinesisStreamsForLedgerPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJournalKinesisStreamsForLedgerPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJournalKinesisStreamsForLedgerPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListJournalKinesisStreamsForLedger page.
func (p *ListJournalKinesisStreamsForLedgerPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJournalKinesisStreamsForLedgerOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListJournalKinesisStreamsForLedger(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListJournalKinesisStreamsForLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListJournalKinesisStreamsForLedger",
	}
}
