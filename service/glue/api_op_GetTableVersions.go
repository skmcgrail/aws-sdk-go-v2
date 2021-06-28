// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of strings that identify available versions of a specified
// table.
func (c *Client) GetTableVersions(ctx context.Context, params *GetTableVersionsInput, optFns ...func(*Options)) (*GetTableVersionsOutput, error) {
	if params == nil {
		params = &GetTableVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTableVersions", params, optFns, c.addOperationGetTableVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTableVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTableVersionsInput struct {

	// The database in the catalog in which the table resides. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table. For Hive compatibility, this name is entirely lowercase.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the tables reside. If none is provided, the
	// Amazon Web Services account ID is used by default.
	CatalogId *string

	// The maximum number of table versions to return in one response.
	MaxResults *int32

	// A continuation token, if this is not the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetTableVersionsOutput struct {

	// A continuation token, if the list of available versions does not include the
	// last one.
	NextToken *string

	// A list of strings identifying available versions of the specified table.
	TableVersions []types.TableVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTableVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTableVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTableVersions{}, middleware.After)
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
	if err = addOpGetTableVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTableVersions(options.Region), middleware.Before); err != nil {
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

// GetTableVersionsAPIClient is a client that implements the GetTableVersions
// operation.
type GetTableVersionsAPIClient interface {
	GetTableVersions(context.Context, *GetTableVersionsInput, ...func(*Options)) (*GetTableVersionsOutput, error)
}

var _ GetTableVersionsAPIClient = (*Client)(nil)

// GetTableVersionsPaginatorOptions is the paginator options for GetTableVersions
type GetTableVersionsPaginatorOptions struct {
	// The maximum number of table versions to return in one response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTableVersionsPaginator is a paginator for GetTableVersions
type GetTableVersionsPaginator struct {
	options   GetTableVersionsPaginatorOptions
	client    GetTableVersionsAPIClient
	params    *GetTableVersionsInput
	nextToken *string
	firstPage bool
}

// NewGetTableVersionsPaginator returns a new GetTableVersionsPaginator
func NewGetTableVersionsPaginator(client GetTableVersionsAPIClient, params *GetTableVersionsInput, optFns ...func(*GetTableVersionsPaginatorOptions)) *GetTableVersionsPaginator {
	if params == nil {
		params = &GetTableVersionsInput{}
	}

	options := GetTableVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTableVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTableVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetTableVersions page.
func (p *GetTableVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTableVersionsOutput, error) {
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

	result, err := p.client.GetTableVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTableVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetTableVersions",
	}
}
