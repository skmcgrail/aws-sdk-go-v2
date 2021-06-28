// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The following example serializes a payload that uses an XML namespace.
func (c *Client) HttpPayloadWithXmlNamespace(ctx context.Context, params *HttpPayloadWithXmlNamespaceInput, optFns ...func(*Options)) (*HttpPayloadWithXmlNamespaceOutput, error) {
	if params == nil {
		params = &HttpPayloadWithXmlNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPayloadWithXmlNamespace", params, optFns, c.addOperationHttpPayloadWithXmlNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPayloadWithXmlNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithXmlNamespaceInput struct {
	Nested *types.PayloadWithXmlNamespace

	noSmithyDocumentSerde
}

type HttpPayloadWithXmlNamespaceOutput struct {
	Nested *types.PayloadWithXmlNamespace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHttpPayloadWithXmlNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadWithXmlNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadWithXmlNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespace(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadWithXmlNamespace",
	}
}
