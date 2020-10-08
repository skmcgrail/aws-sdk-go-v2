// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests and responses are serialized when there is no
// input or output payload but there are HTTP header bindings.
func (c *Client) InputAndOutputWithHeaders(ctx context.Context, params *InputAndOutputWithHeadersInput, optFns ...func(*Options)) (*InputAndOutputWithHeadersOutput, error) {
	stack := middleware.NewStack("InputAndOutputWithHeaders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInputAndOutputWithHeadersMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInputAndOutputWithHeaders(options.Region), middleware.Before)
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
			OperationName: "InputAndOutputWithHeaders",
			Err:           err,
		}
	}
	out := result.(*InputAndOutputWithHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InputAndOutputWithHeadersInput struct {
	HeaderBooleanList []*bool

	HeaderByte *int8

	HeaderDouble *float64

	HeaderEnum types.FooEnum

	HeaderEnumList []types.FooEnum

	HeaderFalseBool *bool

	HeaderFloat *float32

	HeaderInteger *int32

	HeaderIntegerList []*int32

	HeaderLong *int64

	HeaderShort *int16

	HeaderString *string

	HeaderStringList []*string

	HeaderStringSet []*string

	HeaderTimestampList []*time.Time

	HeaderTrueBool *bool
}

type InputAndOutputWithHeadersOutput struct {
	HeaderBooleanList []*bool

	HeaderByte *int8

	HeaderDouble *float64

	HeaderEnum types.FooEnum

	HeaderEnumList []types.FooEnum

	HeaderFalseBool *bool

	HeaderFloat *float32

	HeaderInteger *int32

	HeaderIntegerList []*int32

	HeaderLong *int64

	HeaderShort *int16

	HeaderString *string

	HeaderStringList []*string

	HeaderStringSet []*string

	HeaderTimestampList []*time.Time

	HeaderTrueBool *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInputAndOutputWithHeadersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInputAndOutputWithHeaders{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInputAndOutputWithHeaders{}, middleware.After)
}

func newServiceMetadataMiddleware_opInputAndOutputWithHeaders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InputAndOutputWithHeaders",
	}
}
