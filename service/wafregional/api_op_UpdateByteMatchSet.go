// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Inserts or deletes ByteMatchTuple () objects (filters) in a
// ByteMatchSet (). For each ByteMatchTuple object, you specify the following
// values:
//
//     * Whether to insert or delete the object from the array. If you
// want to change a ByteMatchSetUpdate object, you delete the existing object and
// add a new one.
//
//     * The part of a web request that you want AWS WAF to
// inspect, such as a query string or the value of the User-Agent header.
//
//     *
// The bytes (typically a string that corresponds with ASCII characters) that you
// want AWS WAF to look for. For more information, including how you specify the
// values for the AWS WAF API and the AWS CLI or SDKs, see TargetString in the
// ByteMatchTuple () data type.
//
//     * Where to look, such as at the beginning or
// the end of a query string.
//
//     * Whether to perform any conversions on the
// request, such as converting it to lowercase, before inspecting it for the
// specified string.
//
// For example, you can add a ByteMatchSetUpdate object that
// matches web requests in which User-Agent headers contain the string BadBot. You
// can then configure AWS WAF to block those requests. To create and configure a
// ByteMatchSet, perform the following steps:
//
//     * Create a ByteMatchSet. For
// more information, see CreateByteMatchSet ().
//
//     * Use GetChangeToken () to get
// the change token that you provide in the ChangeToken parameter of an
// UpdateByteMatchSet request.
//
//     * Submit an UpdateByteMatchSet request to
// specify the part of the request that you want AWS WAF to inspect (for example,
// the header or the URI) and the value that you want AWS WAF to watch for.
//
// For
// more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateByteMatchSet(ctx context.Context, params *UpdateByteMatchSetInput, optFns ...func(*Options)) (*UpdateByteMatchSetOutput, error) {
	stack := middleware.NewStack("UpdateByteMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateByteMatchSetMiddlewares(stack)
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
	addOpUpdateByteMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateByteMatchSet(options.Region), middleware.Before)
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
			OperationName: "UpdateByteMatchSet",
			Err:           err,
		}
	}
	out := result.(*UpdateByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateByteMatchSetInput struct {

	// The ByteMatchSetId of the ByteMatchSet () that you want to update.
	// ByteMatchSetId is returned by CreateByteMatchSet () and by ListByteMatchSets ().
	//
	// This member is required.
	ByteMatchSetId *string

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string

	// An array of ByteMatchSetUpdate objects that you want to insert into or delete
	// from a ByteMatchSet (). For more information, see the applicable data types:
	//
	//
	// * ByteMatchSetUpdate (): Contains Action and ByteMatchTuple
	//
	//     *
	// ByteMatchTuple (): Contains FieldToMatch, PositionalConstraint, TargetString,
	// and TextTransformation
	//
	//     * FieldToMatch (): Contains Data and Type
	//
	// This member is required.
	Updates []*types.ByteMatchSetUpdate
}

type UpdateByteMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateByteMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateByteMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateByteMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateByteMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "UpdateByteMatchSet",
	}
}
