// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of media capture pipelines.
func (c *Client) ListMediaCapturePipelines(ctx context.Context, params *ListMediaCapturePipelinesInput, optFns ...func(*Options)) (*ListMediaCapturePipelinesOutput, error) {
	if params == nil {
		params = &ListMediaCapturePipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMediaCapturePipelines", params, optFns, c.addOperationListMediaCapturePipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMediaCapturePipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMediaCapturePipelinesInput struct {

	// The maximum number of results to return in a single call. Valid Range: 1 - 99.
	MaxResults *int32

	// The token used to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMediaCapturePipelinesOutput struct {

	// The media capture pipeline objects in the list.
	MediaCapturePipelines []types.MediaCapturePipeline

	// The token used to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMediaCapturePipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMediaCapturePipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMediaCapturePipelines{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMediaCapturePipelines(options.Region), middleware.Before); err != nil {
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

// ListMediaCapturePipelinesAPIClient is a client that implements the
// ListMediaCapturePipelines operation.
type ListMediaCapturePipelinesAPIClient interface {
	ListMediaCapturePipelines(context.Context, *ListMediaCapturePipelinesInput, ...func(*Options)) (*ListMediaCapturePipelinesOutput, error)
}

var _ ListMediaCapturePipelinesAPIClient = (*Client)(nil)

// ListMediaCapturePipelinesPaginatorOptions is the paginator options for
// ListMediaCapturePipelines
type ListMediaCapturePipelinesPaginatorOptions struct {
	// The maximum number of results to return in a single call. Valid Range: 1 - 99.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMediaCapturePipelinesPaginator is a paginator for ListMediaCapturePipelines
type ListMediaCapturePipelinesPaginator struct {
	options   ListMediaCapturePipelinesPaginatorOptions
	client    ListMediaCapturePipelinesAPIClient
	params    *ListMediaCapturePipelinesInput
	nextToken *string
	firstPage bool
}

// NewListMediaCapturePipelinesPaginator returns a new
// ListMediaCapturePipelinesPaginator
func NewListMediaCapturePipelinesPaginator(client ListMediaCapturePipelinesAPIClient, params *ListMediaCapturePipelinesInput, optFns ...func(*ListMediaCapturePipelinesPaginatorOptions)) *ListMediaCapturePipelinesPaginator {
	if params == nil {
		params = &ListMediaCapturePipelinesInput{}
	}

	options := ListMediaCapturePipelinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMediaCapturePipelinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMediaCapturePipelinesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListMediaCapturePipelines page.
func (p *ListMediaCapturePipelinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMediaCapturePipelinesOutput, error) {
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

	result, err := p.client.ListMediaCapturePipelines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMediaCapturePipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListMediaCapturePipelines",
	}
}
