// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists versions for the specified application.
func (c *Client) ListApplicationVersions(ctx context.Context, params *ListApplicationVersionsInput, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) {
	if params == nil {
		params = &ListApplicationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationVersions", params, optFns, c.addOperationListApplicationVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationVersionsInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// The total number of items to return.
	MaxItems int32

	// A token to specify where to start paginating.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationVersionsOutput struct {

	// The token to request the next page of results.
	NextToken *string

	// An array of version summaries for the application.
	Versions []types.VersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApplicationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApplicationVersions{}, middleware.After)
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
	if err = addOpListApplicationVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationVersions(options.Region), middleware.Before); err != nil {
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

// ListApplicationVersionsAPIClient is a client that implements the
// ListApplicationVersions operation.
type ListApplicationVersionsAPIClient interface {
	ListApplicationVersions(context.Context, *ListApplicationVersionsInput, ...func(*Options)) (*ListApplicationVersionsOutput, error)
}

var _ ListApplicationVersionsAPIClient = (*Client)(nil)

// ListApplicationVersionsPaginatorOptions is the paginator options for
// ListApplicationVersions
type ListApplicationVersionsPaginatorOptions struct {
	// The total number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationVersionsPaginator is a paginator for ListApplicationVersions
type ListApplicationVersionsPaginator struct {
	options   ListApplicationVersionsPaginatorOptions
	client    ListApplicationVersionsAPIClient
	params    *ListApplicationVersionsInput
	nextToken *string
	firstPage bool
}

// NewListApplicationVersionsPaginator returns a new
// ListApplicationVersionsPaginator
func NewListApplicationVersionsPaginator(client ListApplicationVersionsAPIClient, params *ListApplicationVersionsInput, optFns ...func(*ListApplicationVersionsPaginatorOptions)) *ListApplicationVersionsPaginator {
	if params == nil {
		params = &ListApplicationVersionsInput{}
	}

	options := ListApplicationVersionsPaginatorOptions{}
	if params.MaxItems != 0 {
		options.Limit = params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListApplicationVersions page.
func (p *ListApplicationVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxItems = p.options.Limit

	result, err := p.client.ListApplicationVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "ListApplicationVersions",
	}
}
