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

// Returns information about the differences in a valid commit specifier (such as a
// branch, tag, HEAD, commit ID, or other fully qualified reference). Results can
// be limited to a specified path.
func (c *Client) GetDifferences(ctx context.Context, params *GetDifferencesInput, optFns ...func(*Options)) (*GetDifferencesOutput, error) {
	stack := middleware.NewStack("GetDifferences", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDifferencesMiddlewares(stack)
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
	addOpGetDifferencesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDifferences(options.Region), middleware.Before)
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
			OperationName: "GetDifferences",
			Err:           err,
		}
	}
	out := result.(*GetDifferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDifferencesInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit.
	//
	// This member is required.
	AfterCommitSpecifier *string

	// The name of the repository where you want to get differences.
	//
	// This member is required.
	RepositoryName *string

	// The file path in which to check differences. Limits the results to this path.
	// Can also be used to specify the changed name of a directory or folder, if it has
	// changed. If not specified, differences are shown for all paths.
	AfterPath *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, the full commit ID). Optional. If not specified, all
	// changes before the afterCommitSpecifier value are shown. If you do not use
	// beforeCommitSpecifier in your request, consider limiting the results with
	// maxResults.
	BeforeCommitSpecifier *string

	// The file path in which to check for differences. Limits the results to this
	// path. Can also be used to specify the previous name of a directory or folder. If
	// beforePath and afterPath are not specified, differences are shown for all paths.
	BeforePath *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type GetDifferencesOutput struct {

	// A data type object that contains information about the differences, including
	// whether the difference is added, modified, or deleted (A, D, M).
	Differences []*types.Difference

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDifferencesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDifferences{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDifferences{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDifferences(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetDifferences",
	}
}
