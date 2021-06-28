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

// Provides summary information about the users for the specified Amazon Connect
// instance.
func (c *Client) ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if params == nil {
		params = &ListUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsers", params, optFns, c.addOperationListUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsersInput struct {

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

type ListUsersOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the users.
	UserSummaryList []types.UserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUsers{}, middleware.After)
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
	if err = addOpListUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsers(options.Region), middleware.Before); err != nil {
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

// ListUsersAPIClient is a client that implements the ListUsers operation.
type ListUsersAPIClient interface {
	ListUsers(context.Context, *ListUsersInput, ...func(*Options)) (*ListUsersOutput, error)
}

var _ ListUsersAPIClient = (*Client)(nil)

// ListUsersPaginatorOptions is the paginator options for ListUsers
type ListUsersPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUsersPaginator is a paginator for ListUsers
type ListUsersPaginator struct {
	options   ListUsersPaginatorOptions
	client    ListUsersAPIClient
	params    *ListUsersInput
	nextToken *string
	firstPage bool
}

// NewListUsersPaginator returns a new ListUsersPaginator
func NewListUsersPaginator(client ListUsersAPIClient, params *ListUsersInput, optFns ...func(*ListUsersPaginatorOptions)) *ListUsersPaginator {
	if params == nil {
		params = &ListUsersInput{}
	}

	options := ListUsersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUsersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUsersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListUsers page.
func (p *ListUsersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListUsers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListUsers",
	}
}
