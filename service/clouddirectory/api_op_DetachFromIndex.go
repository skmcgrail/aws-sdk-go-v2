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

// Detaches the specified object from the specified index.
func (c *Client) DetachFromIndex(ctx context.Context, params *DetachFromIndexInput, optFns ...func(*Options)) (*DetachFromIndexOutput, error) {
	stack := middleware.NewStack("DetachFromIndex", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDetachFromIndexMiddlewares(stack)
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
	addOpDetachFromIndexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachFromIndex(options.Region), middleware.Before)
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
			OperationName: "DetachFromIndex",
			Err:           err,
		}
	}
	out := result.(*DetachFromIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachFromIndexInput struct {

	// The Amazon Resource Name (ARN) of the directory the index and object exist in.
	//
	// This member is required.
	DirectoryArn *string

	// A reference to the index object.
	//
	// This member is required.
	IndexReference *types.ObjectReference

	// A reference to the object being detached from the index.
	//
	// This member is required.
	TargetReference *types.ObjectReference
}

type DetachFromIndexOutput struct {

	// The ObjectIdentifier of the object that was detached from the index.
	DetachedObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDetachFromIndexMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDetachFromIndex{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachFromIndex{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachFromIndex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "DetachFromIndex",
	}
}
