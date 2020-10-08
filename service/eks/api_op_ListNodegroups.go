// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the Amazon EKS managed node groups associated with the specified cluster
// in your AWS account in the specified Region. Self-managed node groups are not
// listed.
func (c *Client) ListNodegroups(ctx context.Context, params *ListNodegroupsInput, optFns ...func(*Options)) (*ListNodegroupsOutput, error) {
	stack := middleware.NewStack("ListNodegroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListNodegroupsMiddlewares(stack)
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
	addOpListNodegroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListNodegroups(options.Region), middleware.Before)
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
			OperationName: "ListNodegroups",
			Err:           err,
		}
	}
	out := result.(*ListNodegroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNodegroupsInput struct {

	// The name of the Amazon EKS cluster that you would like to list node groups in.
	//
	// This member is required.
	ClusterName *string

	// The maximum number of node group results returned by ListNodegroups in paginated
	// output. When you use this parameter, ListNodegroups returns only maxResults
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another ListNodegroups
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If you don't use this parameter, ListNodegroups returns up to 100 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListNodegroups request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string
}

type ListNodegroupsOutput struct {

	// The nextToken value to include in a future ListNodegroups request. When the
	// results of a ListNodegroups request exceed maxResults, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// A list of all of the node groups associated with the specified cluster.
	Nodegroups []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListNodegroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListNodegroups{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListNodegroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListNodegroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "ListNodegroups",
	}
}
