// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests the validation of a policy and returns a list of findings. The findings
// help you identify issues and provide actionable recommendations to resolve the
// issue and enable you to author functional policies that meet security best
// practices.
func (c *Client) ValidatePolicy(ctx context.Context, params *ValidatePolicyInput, optFns ...func(*Options)) (*ValidatePolicyOutput, error) {
	if params == nil {
		params = &ValidatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidatePolicy", params, optFns, c.addOperationValidatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidatePolicyInput struct {

	// The JSON policy document to use as the content for the policy.
	//
	// This member is required.
	PolicyDocument *string

	// The type of policy to validate. Identity policies grant permissions to IAM
	// principals. Identity policies include managed and inline policies for IAM roles,
	// users, and groups. They also include service-control policies (SCPs) that are
	// attached to an AWS organization, organizational unit (OU), or an account.
	// Resource policies grant permissions on AWS resources. Resource policies include
	// trust policies for IAM roles and bucket policies for S3 buckets. You can provide
	// a generic input such as identity policy or resource policy or a specific input
	// such as managed policy or S3 bucket policy.
	//
	// This member is required.
	PolicyType types.PolicyType

	// The locale to use for localizing the findings.
	Locale types.Locale

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string

	noSmithyDocumentSerde
}

type ValidatePolicyOutput struct {

	// The list of findings in a policy returned by Access Analyzer based on its suite
	// of policy checks.
	//
	// This member is required.
	Findings []types.ValidatePolicyFinding

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidatePolicy{}, middleware.After)
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
	if err = addOpValidatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidatePolicy(options.Region), middleware.Before); err != nil {
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

// ValidatePolicyAPIClient is a client that implements the ValidatePolicy
// operation.
type ValidatePolicyAPIClient interface {
	ValidatePolicy(context.Context, *ValidatePolicyInput, ...func(*Options)) (*ValidatePolicyOutput, error)
}

var _ ValidatePolicyAPIClient = (*Client)(nil)

// ValidatePolicyPaginatorOptions is the paginator options for ValidatePolicy
type ValidatePolicyPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ValidatePolicyPaginator is a paginator for ValidatePolicy
type ValidatePolicyPaginator struct {
	options   ValidatePolicyPaginatorOptions
	client    ValidatePolicyAPIClient
	params    *ValidatePolicyInput
	nextToken *string
	firstPage bool
}

// NewValidatePolicyPaginator returns a new ValidatePolicyPaginator
func NewValidatePolicyPaginator(client ValidatePolicyAPIClient, params *ValidatePolicyInput, optFns ...func(*ValidatePolicyPaginatorOptions)) *ValidatePolicyPaginator {
	if params == nil {
		params = &ValidatePolicyInput{}
	}

	options := ValidatePolicyPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ValidatePolicyPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ValidatePolicyPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ValidatePolicy page.
func (p *ValidatePolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ValidatePolicyOutput, error) {
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

	result, err := p.client.ValidatePolicy(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opValidatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "ValidatePolicy",
	}
}
