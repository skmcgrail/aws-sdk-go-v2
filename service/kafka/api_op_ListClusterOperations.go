// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the operations that have been performed on the specified
// MSK cluster.
func (c *Client) ListClusterOperations(ctx context.Context, params *ListClusterOperationsInput, optFns ...func(*Options)) (*ListClusterOperationsOutput, error) {
	if params == nil {
		params = &ListClusterOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusterOperations", params, optFns, c.addOperationListClusterOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClusterOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClusterOperationsInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	//
	// This member is required.
	ClusterArn *string

	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults int32

	// The paginated results marker. When the result of the operation is truncated, the
	// call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListClusterOperationsOutput struct {

	// An array of cluster operation information objects.
	ClusterOperationInfoList []types.ClusterOperationInfo

	// If the response of ListClusterOperations is truncated, it returns a NextToken in
	// the response. This Nexttoken should be sent in the subsequent request to
	// ListClusterOperations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClusterOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListClusterOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListClusterOperations{}, middleware.After)
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
	if err = addOpListClusterOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusterOperations(options.Region), middleware.Before); err != nil {
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

// ListClusterOperationsAPIClient is a client that implements the
// ListClusterOperations operation.
type ListClusterOperationsAPIClient interface {
	ListClusterOperations(context.Context, *ListClusterOperationsInput, ...func(*Options)) (*ListClusterOperationsOutput, error)
}

var _ ListClusterOperationsAPIClient = (*Client)(nil)

// ListClusterOperationsPaginatorOptions is the paginator options for
// ListClusterOperations
type ListClusterOperationsPaginatorOptions struct {
	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClusterOperationsPaginator is a paginator for ListClusterOperations
type ListClusterOperationsPaginator struct {
	options   ListClusterOperationsPaginatorOptions
	client    ListClusterOperationsAPIClient
	params    *ListClusterOperationsInput
	nextToken *string
	firstPage bool
}

// NewListClusterOperationsPaginator returns a new ListClusterOperationsPaginator
func NewListClusterOperationsPaginator(client ListClusterOperationsAPIClient, params *ListClusterOperationsInput, optFns ...func(*ListClusterOperationsPaginatorOptions)) *ListClusterOperationsPaginator {
	if params == nil {
		params = &ListClusterOperationsInput{}
	}

	options := ListClusterOperationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClusterOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClusterOperationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListClusterOperations page.
func (p *ListClusterOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClusterOperationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListClusterOperations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClusterOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "ListClusterOperations",
	}
}
