// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a paginated list of all the outgoing TypedLinkSpecifier () information
// for an object. It also supports filtering by typed link facet and identity
// attributes. For more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) ListOutgoingTypedLinks(ctx context.Context, params *ListOutgoingTypedLinksInput, optFns ...func(*Options)) (*ListOutgoingTypedLinksOutput, error) {
	stack := middleware.NewStack("ListOutgoingTypedLinks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListOutgoingTypedLinksMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListOutgoingTypedLinksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOutgoingTypedLinks(options.Region), middleware.Before)
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
			OperationName: "ListOutgoingTypedLinks",
			Err:           err,
		}
	}
	out := result.(*ListOutgoingTypedLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOutgoingTypedLinksInput struct {
	// The Amazon Resource Name (ARN) of the directory where you want to list the typed
	// links.
	DirectoryArn *string
	// The consistency level to execute the request at.
	ConsistencyLevel types.ConsistencyLevel
	// The maximum number of results to retrieve.
	MaxResults *int32
	// A reference that identifies the object whose attributes will be listed.
	ObjectReference *types.ObjectReference
	// Filters are interpreted in the order of the attributes defined on the typed link
	// facet, not the order they are supplied to any API calls.
	FilterTypedLink *types.TypedLinkSchemaAndFacetName
	// The pagination token.
	NextToken *string
	// Provides range filters for multiple attributes. When providing ranges to typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	FilterAttributeRanges []*types.TypedLinkAttributeRange
}

type ListOutgoingTypedLinksOutput struct {
	// Returns a typed link specifier as output.
	TypedLinkSpecifiers []*types.TypedLinkSpecifier
	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListOutgoingTypedLinksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListOutgoingTypedLinks{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListOutgoingTypedLinks{}, middleware.After)
}

func newServiceMetadataMiddleware_opListOutgoingTypedLinks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListOutgoingTypedLinks",
	}
}