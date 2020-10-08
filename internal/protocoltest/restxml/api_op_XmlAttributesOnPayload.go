// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes an XML attributes on a document targeted by httpPayload.
func (c *Client) XmlAttributesOnPayload(ctx context.Context, params *XmlAttributesOnPayloadInput, optFns ...func(*Options)) (*XmlAttributesOnPayloadOutput, error) {
	stack := middleware.NewStack("XmlAttributesOnPayload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpXmlAttributesOnPayloadMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlAttributesOnPayload(options.Region), middleware.Before)
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
			OperationName: "XmlAttributesOnPayload",
			Err:           err,
		}
	}
	out := result.(*XmlAttributesOnPayloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlAttributesOnPayloadInput struct {
	Payload *types.XmlAttributesInputOutput
}

type XmlAttributesOnPayloadOutput struct {
	Payload *types.XmlAttributesInputOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpXmlAttributesOnPayloadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpXmlAttributesOnPayload{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpXmlAttributesOnPayload{}, middleware.After)
}

func newServiceMetadataMiddleware_opXmlAttributesOnPayload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlAttributesOnPayload",
	}
}
