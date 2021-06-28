// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides summary information about the AppIntegration associations for the
// specified Amazon Connect instance.
func (c *Client) ListIntegrationAssociations(ctx context.Context, params *ListIntegrationAssociationsInput, optFns ...func(*Options)) (*ListIntegrationAssociationsOutput, error) {
	if params == nil {
		params = &ListIntegrationAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIntegrationAssociations", params, optFns, c.addOperationListIntegrationAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIntegrationAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIntegrationAssociationsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListIntegrationAssociationsOutput struct {

	// The AppIntegration associations.
	IntegrationAssociationSummaryList []types.IntegrationAssociationSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIntegrationAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIntegrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIntegrationAssociations{}, middleware.After)
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
	if err = addOpListIntegrationAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIntegrationAssociations(options.Region), middleware.Before); err != nil {
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

// ListIntegrationAssociationsAPIClient is a client that implements the
// ListIntegrationAssociations operation.
type ListIntegrationAssociationsAPIClient interface {
	ListIntegrationAssociations(context.Context, *ListIntegrationAssociationsInput, ...func(*Options)) (*ListIntegrationAssociationsOutput, error)
}

var _ ListIntegrationAssociationsAPIClient = (*Client)(nil)

// ListIntegrationAssociationsPaginatorOptions is the paginator options for
// ListIntegrationAssociations
type ListIntegrationAssociationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIntegrationAssociationsPaginator is a paginator for
// ListIntegrationAssociations
type ListIntegrationAssociationsPaginator struct {
	options   ListIntegrationAssociationsPaginatorOptions
	client    ListIntegrationAssociationsAPIClient
	params    *ListIntegrationAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListIntegrationAssociationsPaginator returns a new
// ListIntegrationAssociationsPaginator
func NewListIntegrationAssociationsPaginator(client ListIntegrationAssociationsAPIClient, params *ListIntegrationAssociationsInput, optFns ...func(*ListIntegrationAssociationsPaginatorOptions)) *ListIntegrationAssociationsPaginator {
	if params == nil {
		params = &ListIntegrationAssociationsInput{}
	}

	options := ListIntegrationAssociationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIntegrationAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIntegrationAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListIntegrationAssociations page.
func (p *ListIntegrationAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIntegrationAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListIntegrationAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIntegrationAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListIntegrationAssociations",
	}
}
