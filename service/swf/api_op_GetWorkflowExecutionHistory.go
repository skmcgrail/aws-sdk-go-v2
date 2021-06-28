// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the history of the specified workflow execution. The results may be
// split into multiple pages. To retrieve subsequent pages, make the call again
// using the nextPageToken returned by the initial call. This operation is
// eventually consistent. The results are best effort and may not exactly reflect
// recent updates and changes. Access Control You can use IAM policies to control
// this action's access to Amazon SWF resources as follows:
//
// * Use a Resource
// element with the domain name to limit the action to only specified domains.
//
// *
// Use an Action element to allow or deny permission to call this action.
//
// * You
// cannot use an IAM policy to constrain this action's parameters.
//
// If the caller
// doesn't have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails. The associated
// event attribute's cause parameter is set to OPERATION_NOT_PERMITTED. For details
// and example IAM policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) GetWorkflowExecutionHistory(ctx context.Context, params *GetWorkflowExecutionHistoryInput, optFns ...func(*Options)) (*GetWorkflowExecutionHistoryOutput, error) {
	if params == nil {
		params = &GetWorkflowExecutionHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowExecutionHistory", params, optFns, c.addOperationGetWorkflowExecutionHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowExecutionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowExecutionHistoryInput struct {

	// The name of the domain containing the workflow execution.
	//
	// This member is required.
	Domain *string

	// Specifies the workflow execution for which to return the history.
	//
	// This member is required.
	Execution *types.WorkflowExecution

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	MaximumPageSize int32

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 60 seconds. Using an expired
	// pagination token will return a 400 error: "Specified token has exceeded its
	// maximum lifetime". The configured maximumPageSize determines how many results
	// can be returned in a single call.
	NextPageToken *string

	// When set to true, returns the events in reverse order. By default the results
	// are returned in ascending order of the eventTimeStamp of the events.
	ReverseOrder bool

	noSmithyDocumentSerde
}

// Paginated representation of a workflow history for a workflow execution. This is
// the up to date, complete and authoritative record of the events related to all
// tasks and events in the life of the workflow execution.
type GetWorkflowExecutionHistoryOutput struct {

	// The list of history events.
	//
	// This member is required.
	Events []types.HistoryEvent

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken. Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowExecutionHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetWorkflowExecutionHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetWorkflowExecutionHistory{}, middleware.After)
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
	if err = addOpGetWorkflowExecutionHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowExecutionHistory(options.Region), middleware.Before); err != nil {
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

// GetWorkflowExecutionHistoryAPIClient is a client that implements the
// GetWorkflowExecutionHistory operation.
type GetWorkflowExecutionHistoryAPIClient interface {
	GetWorkflowExecutionHistory(context.Context, *GetWorkflowExecutionHistoryInput, ...func(*Options)) (*GetWorkflowExecutionHistoryOutput, error)
}

var _ GetWorkflowExecutionHistoryAPIClient = (*Client)(nil)

// GetWorkflowExecutionHistoryPaginatorOptions is the paginator options for
// GetWorkflowExecutionHistory
type GetWorkflowExecutionHistoryPaginatorOptions struct {
	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetWorkflowExecutionHistoryPaginator is a paginator for
// GetWorkflowExecutionHistory
type GetWorkflowExecutionHistoryPaginator struct {
	options   GetWorkflowExecutionHistoryPaginatorOptions
	client    GetWorkflowExecutionHistoryAPIClient
	params    *GetWorkflowExecutionHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetWorkflowExecutionHistoryPaginator returns a new
// GetWorkflowExecutionHistoryPaginator
func NewGetWorkflowExecutionHistoryPaginator(client GetWorkflowExecutionHistoryAPIClient, params *GetWorkflowExecutionHistoryInput, optFns ...func(*GetWorkflowExecutionHistoryPaginatorOptions)) *GetWorkflowExecutionHistoryPaginator {
	if params == nil {
		params = &GetWorkflowExecutionHistoryInput{}
	}

	options := GetWorkflowExecutionHistoryPaginatorOptions{}
	if params.MaximumPageSize != 0 {
		options.Limit = params.MaximumPageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetWorkflowExecutionHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetWorkflowExecutionHistoryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetWorkflowExecutionHistory page.
func (p *GetWorkflowExecutionHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetWorkflowExecutionHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextPageToken = p.nextToken

	params.MaximumPageSize = p.options.Limit

	result, err := p.client.GetWorkflowExecutionHistory(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetWorkflowExecutionHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "GetWorkflowExecutionHistory",
	}
}
