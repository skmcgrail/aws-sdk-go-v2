// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a detailed view of a Remediation Execution for a set of resources
// including state, timestamps for when steps for the remediation execution occur,
// and any error messages for steps that have failed. When you specify the limit
// and the next token, you receive a paginated response.
func (c *Client) DescribeRemediationExecutionStatus(ctx context.Context, params *DescribeRemediationExecutionStatusInput, optFns ...func(*Options)) (*DescribeRemediationExecutionStatusOutput, error) {
	if params == nil {
		params = &DescribeRemediationExecutionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRemediationExecutionStatus", params, optFns, c.addOperationDescribeRemediationExecutionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRemediationExecutionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRemediationExecutionStatusInput struct {

	// A list of AWS Config rule names.
	//
	// This member is required.
	ConfigRuleName *string

	// The maximum number of RemediationExecutionStatuses returned on each page. The
	// default is maximum. If you specify 0, AWS Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	ResourceKeys []types.ResourceKey

	noSmithyDocumentSerde
}

type DescribeRemediationExecutionStatusOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a list of remediation execution statuses objects.
	RemediationExecutionStatuses []types.RemediationExecutionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRemediationExecutionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRemediationExecutionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRemediationExecutionStatus{}, middleware.After)
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
	if err = addOpDescribeRemediationExecutionStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRemediationExecutionStatus(options.Region), middleware.Before); err != nil {
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

// DescribeRemediationExecutionStatusAPIClient is a client that implements the
// DescribeRemediationExecutionStatus operation.
type DescribeRemediationExecutionStatusAPIClient interface {
	DescribeRemediationExecutionStatus(context.Context, *DescribeRemediationExecutionStatusInput, ...func(*Options)) (*DescribeRemediationExecutionStatusOutput, error)
}

var _ DescribeRemediationExecutionStatusAPIClient = (*Client)(nil)

// DescribeRemediationExecutionStatusPaginatorOptions is the paginator options for
// DescribeRemediationExecutionStatus
type DescribeRemediationExecutionStatusPaginatorOptions struct {
	// The maximum number of RemediationExecutionStatuses returned on each page. The
	// default is maximum. If you specify 0, AWS Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRemediationExecutionStatusPaginator is a paginator for
// DescribeRemediationExecutionStatus
type DescribeRemediationExecutionStatusPaginator struct {
	options   DescribeRemediationExecutionStatusPaginatorOptions
	client    DescribeRemediationExecutionStatusAPIClient
	params    *DescribeRemediationExecutionStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeRemediationExecutionStatusPaginator returns a new
// DescribeRemediationExecutionStatusPaginator
func NewDescribeRemediationExecutionStatusPaginator(client DescribeRemediationExecutionStatusAPIClient, params *DescribeRemediationExecutionStatusInput, optFns ...func(*DescribeRemediationExecutionStatusPaginatorOptions)) *DescribeRemediationExecutionStatusPaginator {
	if params == nil {
		params = &DescribeRemediationExecutionStatusInput{}
	}

	options := DescribeRemediationExecutionStatusPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRemediationExecutionStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRemediationExecutionStatusPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeRemediationExecutionStatus page.
func (p *DescribeRemediationExecutionStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRemediationExecutionStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeRemediationExecutionStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRemediationExecutionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeRemediationExecutionStatus",
	}
}
