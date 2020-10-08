// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a TypedLinkFacet (). For more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) CreateTypedLinkFacet(ctx context.Context, params *CreateTypedLinkFacetInput, optFns ...func(*Options)) (*CreateTypedLinkFacetOutput, error) {
	stack := middleware.NewStack("CreateTypedLinkFacet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateTypedLinkFacetMiddlewares(stack)
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
	addOpCreateTypedLinkFacetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTypedLinkFacet(options.Region), middleware.Before)
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
			OperationName: "CreateTypedLinkFacet",
			Err:           err,
		}
	}
	out := result.(*CreateTypedLinkFacetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTypedLinkFacetInput struct {

	// Facet () structure that is associated with the typed link facet.
	//
	// This member is required.
	Facet *types.TypedLinkFacet

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns ().
	//
	// This member is required.
	SchemaArn *string
}

type CreateTypedLinkFacetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateTypedLinkFacetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateTypedLinkFacet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTypedLinkFacet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTypedLinkFacet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "CreateTypedLinkFacet",
	}
}
