// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the association information for customer gateways that are associated with
// devices and links in your global network.
func (c *Client) GetCustomerGatewayAssociations(ctx context.Context, params *GetCustomerGatewayAssociationsInput, optFns ...func(*Options)) (*GetCustomerGatewayAssociationsOutput, error) {
	if params == nil {
		params = &GetCustomerGatewayAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCustomerGatewayAssociations", params, optFns, c.addOperationGetCustomerGatewayAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCustomerGatewayAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCustomerGatewayAssociationsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// One or more customer gateway Amazon Resource Names (ARNs). For more information,
	// see Resources Defined by Amazon EC2
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies).
	// The maximum is 10.
	CustomerGatewayArns []string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCustomerGatewayAssociationsOutput struct {

	// The customer gateway associations.
	CustomerGatewayAssociations []types.CustomerGatewayAssociation

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCustomerGatewayAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCustomerGatewayAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCustomerGatewayAssociations{}, middleware.After)
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
	if err = addOpGetCustomerGatewayAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCustomerGatewayAssociations(options.Region), middleware.Before); err != nil {
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

// GetCustomerGatewayAssociationsAPIClient is a client that implements the
// GetCustomerGatewayAssociations operation.
type GetCustomerGatewayAssociationsAPIClient interface {
	GetCustomerGatewayAssociations(context.Context, *GetCustomerGatewayAssociationsInput, ...func(*Options)) (*GetCustomerGatewayAssociationsOutput, error)
}

var _ GetCustomerGatewayAssociationsAPIClient = (*Client)(nil)

// GetCustomerGatewayAssociationsPaginatorOptions is the paginator options for
// GetCustomerGatewayAssociations
type GetCustomerGatewayAssociationsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCustomerGatewayAssociationsPaginator is a paginator for
// GetCustomerGatewayAssociations
type GetCustomerGatewayAssociationsPaginator struct {
	options   GetCustomerGatewayAssociationsPaginatorOptions
	client    GetCustomerGatewayAssociationsAPIClient
	params    *GetCustomerGatewayAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetCustomerGatewayAssociationsPaginator returns a new
// GetCustomerGatewayAssociationsPaginator
func NewGetCustomerGatewayAssociationsPaginator(client GetCustomerGatewayAssociationsAPIClient, params *GetCustomerGatewayAssociationsInput, optFns ...func(*GetCustomerGatewayAssociationsPaginatorOptions)) *GetCustomerGatewayAssociationsPaginator {
	if params == nil {
		params = &GetCustomerGatewayAssociationsInput{}
	}

	options := GetCustomerGatewayAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCustomerGatewayAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCustomerGatewayAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetCustomerGatewayAssociations page.
func (p *GetCustomerGatewayAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCustomerGatewayAssociationsOutput, error) {
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

	result, err := p.client.GetCustomerGatewayAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCustomerGatewayAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetCustomerGatewayAssociations",
	}
}
