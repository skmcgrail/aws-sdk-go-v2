// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the contents of the specified folder, including its documents and
// subfolders. By default, Amazon WorkDocs returns the first 100 active document
// and folder metadata items. If there are more results, the response includes a
// marker that you can use to request the next set of results. You can also request
// initialized documents.
func (c *Client) DescribeFolderContents(ctx context.Context, params *DescribeFolderContentsInput, optFns ...func(*Options)) (*DescribeFolderContentsOutput, error) {
	stack := middleware.NewStack("DescribeFolderContents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeFolderContentsMiddlewares(stack)
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
	addOpDescribeFolderContentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFolderContents(options.Region), middleware.Before)
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
			OperationName: "DescribeFolderContents",
			Err:           err,
		}
	}
	out := result.(*DescribeFolderContentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFolderContentsInput struct {

	// The ID of the folder.
	//
	// This member is required.
	FolderId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The contents to include. Specify "INITIALIZED" to include initialized documents.
	Include *string

	// The maximum number of items to return with this call.
	Limit *int32

	// The marker for the next set of results. This marker was received from a previous
	// call.
	Marker *string

	// The order for the contents of the folder.
	Order types.OrderType

	// The sorting criteria.
	Sort types.ResourceSortType

	// The type of items.
	Type types.FolderContentType
}

type DescribeFolderContentsOutput struct {

	// The documents in the specified folder.
	Documents []*types.DocumentMetadata

	// The subfolders in the specified folder.
	Folders []*types.FolderMetadata

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeFolderContentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFolderContents{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFolderContents{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFolderContents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeFolderContents",
	}
}
