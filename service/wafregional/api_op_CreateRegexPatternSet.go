// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
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
// global use. Creates a RegexPatternSet. You then use UpdateRegexPatternSet () to
// specify the regular expression (regex) pattern that you want AWS WAF to search
// for, such as B[a@]dB[o0]t. You can then configure AWS WAF to reject those
// requests. To create and configure a RegexPatternSet, perform the following
// steps:
//
//     * Use GetChangeToken () to get the change token that you provide in
// the ChangeToken parameter of a CreateRegexPatternSet request.
//
//     * Submit a
// CreateRegexPatternSet request.
//
//     * Use GetChangeToken to get the change token
// that you provide in the ChangeToken parameter of an UpdateRegexPatternSet
// request.
//
//     * Submit an UpdateRegexPatternSet () request to specify the string
// that you want AWS WAF to watch for.
//
// For more information about how to use the
// AWS WAF API to allow or block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateRegexPatternSet(ctx context.Context, params *CreateRegexPatternSetInput, optFns ...func(*Options)) (*CreateRegexPatternSetOutput, error) {
	stack := middleware.NewStack("CreateRegexPatternSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateRegexPatternSetMiddlewares(stack)
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
	addOpCreateRegexPatternSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegexPatternSet(options.Region), middleware.Before)
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
			OperationName: "CreateRegexPatternSet",
			Err:           err,
		}
	}
	out := result.(*CreateRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegexPatternSetInput struct {
	// A friendly name or description of the RegexPatternSet (). You can't change Name
	// after you create a RegexPatternSet.
	Name *string
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
}

type CreateRegexPatternSetOutput struct {
	// A RegexPatternSet () that contains no objects.
	RegexPatternSet *types.RegexPatternSet
	// The ChangeToken that you used to submit the CreateRegexPatternSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateRegexPatternSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRegexPatternSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRegexPatternSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRegexPatternSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateRegexPatternSet",
	}
}