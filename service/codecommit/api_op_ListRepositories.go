// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more repositories.
func (c *Client) ListRepositories(ctx context.Context, params *ListRepositoriesInput, optFns ...func(*Options)) (*ListRepositoriesOutput, error) {
	stack := middleware.NewStack("ListRepositories", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRepositoriesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositories(options.Region), middleware.Before)
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
			OperationName: "ListRepositories",
			Err:           err,
		}
	}
	out := result.(*ListRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a list repositories operation.
type ListRepositoriesInput struct {

	// An enumeration token that allows the operation to batch the results of the
	// operation. Batch sizes are 1,000 for list repository operations. When the client
	// sends the token back to AWS CodeCommit, another page of 1,000 records is
	// retrieved.
	NextToken *string

	// The order in which to sort the results of a list repositories operation.
	Order types.OrderEnum

	// The criteria used to sort the results of a list repositories operation.
	SortBy types.SortByEnum
}

// Represents the output of a list repositories operation.
type ListRepositoriesOutput struct {

	// An enumeration token that allows the operation to batch the results of the
	// operation. Batch sizes are 1,000 for list repository operations. When the client
	// sends the token back to AWS CodeCommit, another page of 1,000 records is
	// retrieved.
	NextToken *string

	// Lists the repositories called by the list repositories operation.
	Repositories []*types.RepositoryNameIdPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRepositoriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRepositories{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRepositories{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRepositories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "ListRepositories",
	}
}
