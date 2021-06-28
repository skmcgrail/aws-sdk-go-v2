// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists medical transcription jobs with a specified status or substring that
// matches their names.
func (c *Client) ListMedicalTranscriptionJobs(ctx context.Context, params *ListMedicalTranscriptionJobsInput, optFns ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error) {
	if params == nil {
		params = &ListMedicalTranscriptionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMedicalTranscriptionJobs", params, optFns, c.addOperationListMedicalTranscriptionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMedicalTranscriptionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMedicalTranscriptionJobsInput struct {

	// When specified, the jobs returned in the list are limited to jobs whose name
	// contains the specified string.
	JobNameContains *string

	// The maximum number of medical transcription jobs to return in the response. IF
	// there are fewer results in the list, this response contains only the actual
	// results.
	MaxResults *int32

	// If you a receive a truncated result in the previous request of
	// ListMedicalTranscriptionJobs, include NextToken to fetch the next set of jobs.
	NextToken *string

	// When specified, returns only medical transcription jobs with the specified
	// status. Jobs are ordered by creation date, with the newest jobs returned first.
	// If you don't specify a status, Amazon Transcribe Medical returns all
	// transcription jobs ordered by creation date.
	Status types.TranscriptionJobStatus

	noSmithyDocumentSerde
}

type ListMedicalTranscriptionJobsOutput struct {

	// A list of objects containing summary information for a transcription job.
	MedicalTranscriptionJobSummaries []types.MedicalTranscriptionJobSummary

	// The ListMedicalTranscriptionJobs operation returns a page of jobs at a time. The
	// maximum size of the page is set by the MaxResults parameter. If the number of
	// jobs exceeds what can fit on a page, Amazon Transcribe Medical returns the
	// NextPage token. Include the token in the next request to the
	// ListMedicalTranscriptionJobs operation to return in the next page of jobs.
	NextToken *string

	// The requested status of the medical transcription jobs returned.
	Status types.TranscriptionJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMedicalTranscriptionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMedicalTranscriptionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMedicalTranscriptionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMedicalTranscriptionJobs(options.Region), middleware.Before); err != nil {
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

// ListMedicalTranscriptionJobsAPIClient is a client that implements the
// ListMedicalTranscriptionJobs operation.
type ListMedicalTranscriptionJobsAPIClient interface {
	ListMedicalTranscriptionJobs(context.Context, *ListMedicalTranscriptionJobsInput, ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error)
}

var _ ListMedicalTranscriptionJobsAPIClient = (*Client)(nil)

// ListMedicalTranscriptionJobsPaginatorOptions is the paginator options for
// ListMedicalTranscriptionJobs
type ListMedicalTranscriptionJobsPaginatorOptions struct {
	// The maximum number of medical transcription jobs to return in the response. IF
	// there are fewer results in the list, this response contains only the actual
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMedicalTranscriptionJobsPaginator is a paginator for
// ListMedicalTranscriptionJobs
type ListMedicalTranscriptionJobsPaginator struct {
	options   ListMedicalTranscriptionJobsPaginatorOptions
	client    ListMedicalTranscriptionJobsAPIClient
	params    *ListMedicalTranscriptionJobsInput
	nextToken *string
	firstPage bool
}

// NewListMedicalTranscriptionJobsPaginator returns a new
// ListMedicalTranscriptionJobsPaginator
func NewListMedicalTranscriptionJobsPaginator(client ListMedicalTranscriptionJobsAPIClient, params *ListMedicalTranscriptionJobsInput, optFns ...func(*ListMedicalTranscriptionJobsPaginatorOptions)) *ListMedicalTranscriptionJobsPaginator {
	if params == nil {
		params = &ListMedicalTranscriptionJobsInput{}
	}

	options := ListMedicalTranscriptionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMedicalTranscriptionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMedicalTranscriptionJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListMedicalTranscriptionJobs page.
func (p *ListMedicalTranscriptionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error) {
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

	result, err := p.client.ListMedicalTranscriptionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMedicalTranscriptionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListMedicalTranscriptionJobs",
	}
}
