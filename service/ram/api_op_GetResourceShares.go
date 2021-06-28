// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the resource shares that you own or the resource shares that are shared
// with you.
func (c *Client) GetResourceShares(ctx context.Context, params *GetResourceSharesInput, optFns ...func(*Options)) (*GetResourceSharesOutput, error) {
	if params == nil {
		params = &GetResourceSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceShares", params, optFns, c.addOperationGetResourceSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceSharesInput struct {

	// The type of owner.
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The name of the resource share.
	Name *string

	// The token for the next page of results.
	NextToken *string

	// The Amazon Resource Name (ARN) of the AWS RAM permission that is associated with
	// the resource share.
	PermissionArn *string

	// The ARNs of the resource shares.
	ResourceShareArns []string

	// The status of the resource share.
	ResourceShareStatus types.ResourceShareStatus

	// One or more tag filters.
	TagFilters []types.TagFilter

	noSmithyDocumentSerde
}

type GetResourceSharesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the resource shares.
	ResourceShares []types.ResourceShare

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceShares{}, middleware.After)
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
	if err = addOpGetResourceSharesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceShares(options.Region), middleware.Before); err != nil {
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

// GetResourceSharesAPIClient is a client that implements the GetResourceShares
// operation.
type GetResourceSharesAPIClient interface {
	GetResourceShares(context.Context, *GetResourceSharesInput, ...func(*Options)) (*GetResourceSharesOutput, error)
}

var _ GetResourceSharesAPIClient = (*Client)(nil)

// GetResourceSharesPaginatorOptions is the paginator options for GetResourceShares
type GetResourceSharesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourceSharesPaginator is a paginator for GetResourceShares
type GetResourceSharesPaginator struct {
	options   GetResourceSharesPaginatorOptions
	client    GetResourceSharesAPIClient
	params    *GetResourceSharesInput
	nextToken *string
	firstPage bool
}

// NewGetResourceSharesPaginator returns a new GetResourceSharesPaginator
func NewGetResourceSharesPaginator(client GetResourceSharesAPIClient, params *GetResourceSharesInput, optFns ...func(*GetResourceSharesPaginatorOptions)) *GetResourceSharesPaginator {
	if params == nil {
		params = &GetResourceSharesInput{}
	}

	options := GetResourceSharesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetResourceSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourceSharesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetResourceShares page.
func (p *GetResourceSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourceSharesOutput, error) {
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

	result, err := p.client.GetResourceShares(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetResourceShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "GetResourceShares",
	}
}
