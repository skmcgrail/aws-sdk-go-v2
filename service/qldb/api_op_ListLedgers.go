// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of ledger summaries that are associated with the current AWS
// account and Region. This action returns a maximum of 100 items and is paginated
// so that you can retrieve all the items by calling ListLedgers multiple times.
func (c *Client) ListLedgers(ctx context.Context, params *ListLedgersInput, optFns ...func(*Options)) (*ListLedgersOutput, error) {
	stack := middleware.NewStack("ListLedgers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListLedgersMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLedgers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListLedgers",
			Err:           err,
		}
	}
	out := result.(*ListLedgersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLedgersInput struct {

	// The maximum number of results to return in a single ListLedgers request. (The
	// actual number of results returned might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListLedgers call, then you should use that value as input here.
	NextToken *string
}

type ListLedgersOutput struct {

	// The array of ledger summaries that are associated with the current AWS account
	// and Region.
	Ledgers []*types.LedgerSummary

	// A pagination token, indicating whether there are more results available:
	//
	//     *
	// If NextToken is empty, then the last page of results has been processed and
	// there are no more results to be retrieved.
	//
	//     * If NextToken is not empty,
	// then there are more results available. To retrieve the next page of results, use
	// the value of NextToken in a subsequent ListLedgers call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListLedgersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListLedgers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListLedgers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListLedgers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListLedgers",
	}
}
