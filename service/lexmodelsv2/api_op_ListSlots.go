// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of slots that match the specified criteria.
func (c *Client) ListSlots(ctx context.Context, params *ListSlotsInput, optFns ...func(*Options)) (*ListSlotsOutput, error) {
	if params == nil {
		params = &ListSlotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSlots", params, optFns, c.addOperationListSlotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSlotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSlotsInput struct {

	// The identifier of the bot that contains the slot.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the slot.
	//
	// This member is required.
	BotVersion *string

	// The unique identifier of the intent that contains the slot.
	//
	// This member is required.
	IntentId *string

	// The identifier of the language and locale of the slots to list. The string must
	// match one of the supported locales. For more information, see Supported
	// languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// Provides the specification of a filter used to limit the slots in the response
	// to only those that match the filter specification. You can only specify one
	// filter and only one string to filter on.
	Filters []types.SlotFilter

	// The maximum number of slots to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListSlots operation contains more results than
	// specified in the maxResults parameter, a token is returned in the response. Use
	// that token in the nextToken parameter to return the next page of results.
	NextToken *string

	// Determines the sort order for the response from the ListSlots operation. You can
	// choose to sort by the slot name or last updated date in either ascending or
	// descending order.
	SortBy *types.SlotSortBy

	noSmithyDocumentSerde
}

type ListSlotsOutput struct {

	// The identifier of the bot that contains the slots.
	BotId *string

	// The version of the bot that contains the slots.
	BotVersion *string

	// The identifier of the intent that contains the slots.
	IntentId *string

	// The language and locale of the slots in the list.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response to
	// the ListSlots operation. If the nextToken field is present, you send the
	// contents as the nextToken parameter of a ListSlots operation request to get the
	// next page of results.
	NextToken *string

	// Summary information for the slots that meet the filter criteria specified in the
	// request. The length of the list is specified in the maxResults parameter of the
	// request. If there are more slots available, the nextToken field contains a token
	// to get the next page of results.
	SlotSummaries []types.SlotSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSlotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSlots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSlots{}, middleware.After)
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
	if err = addOpListSlotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSlots(options.Region), middleware.Before); err != nil {
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

// ListSlotsAPIClient is a client that implements the ListSlots operation.
type ListSlotsAPIClient interface {
	ListSlots(context.Context, *ListSlotsInput, ...func(*Options)) (*ListSlotsOutput, error)
}

var _ ListSlotsAPIClient = (*Client)(nil)

// ListSlotsPaginatorOptions is the paginator options for ListSlots
type ListSlotsPaginatorOptions struct {
	// The maximum number of slots to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSlotsPaginator is a paginator for ListSlots
type ListSlotsPaginator struct {
	options   ListSlotsPaginatorOptions
	client    ListSlotsAPIClient
	params    *ListSlotsInput
	nextToken *string
	firstPage bool
}

// NewListSlotsPaginator returns a new ListSlotsPaginator
func NewListSlotsPaginator(client ListSlotsAPIClient, params *ListSlotsInput, optFns ...func(*ListSlotsPaginatorOptions)) *ListSlotsPaginator {
	if params == nil {
		params = &ListSlotsInput{}
	}

	options := ListSlotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSlotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSlotsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSlots page.
func (p *ListSlotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSlotsOutput, error) {
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

	result, err := p.client.ListSlots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSlots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListSlots",
	}
}
