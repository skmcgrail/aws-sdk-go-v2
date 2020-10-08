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

// Returns information about a repository branch, including its name and the last
// commit ID.
func (c *Client) GetBranch(ctx context.Context, params *GetBranchInput, optFns ...func(*Options)) (*GetBranchOutput, error) {
	stack := middleware.NewStack("GetBranch", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetBranchMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBranch(options.Region), middleware.Before)
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
			OperationName: "GetBranch",
			Err:           err,
		}
	}
	out := result.(*GetBranchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a get branch operation.
type GetBranchInput struct {

	// The name of the branch for which you want to retrieve information.
	BranchName *string

	// The name of the repository that contains the branch for which you want to
	// retrieve information.
	RepositoryName *string
}

// Represents the output of a get branch operation.
type GetBranchOutput struct {

	// The name of the branch.
	Branch *types.BranchInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetBranchMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetBranch{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBranch{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBranch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetBranch",
	}
}
