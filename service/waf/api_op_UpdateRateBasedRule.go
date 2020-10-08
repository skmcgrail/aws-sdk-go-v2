// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
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
// global use. Inserts or deletes Predicate () objects in a rule and updates the
// RateLimit in the rule. Each Predicate object identifies a predicate, such as a
// ByteMatchSet () or an IPSet (), that specifies the web requests that you want to
// block or count. The RateLimit specifies the number of requests every five
// minutes that triggers the rule. If you add more than one predicate to a
// RateBasedRule, a request must match all the predicates and exceed the RateLimit
// to be counted or blocked. For example, suppose you add the following to a
// RateBasedRule:
//
//     * An IPSet that matches the IP address 192.0.2.44/32
//
//     *
// A ByteMatchSet that matches BadBot in the User-Agent header
//
// Further, you
// specify a RateLimit of 1,000. You then add the RateBasedRule to a WebACL and
// specify that you want to block requests that satisfy the rule. For a request to
// be blocked, it must come from the IP address 192.0.2.44 and the User-Agent
// header in the request must contain the value BadBot. Further, requests that
// match these two conditions much be received at a rate of more than 1,000 every
// five minutes. If the rate drops below this limit, AWS WAF no longer blocks the
// requests.  <p>As a second example, suppose you want to limit requests to a
// particular page on your site. To do this, you could add the following to a
// <code>RateBasedRule</code>:</p> <ul> <li> <p>A <code>ByteMatchSet</code> with
// <code>FieldToMatch</code> of <code>URI</code> </p> </li> <li> <p>A
// <code>PositionalConstraint</code> of <code>STARTS_WITH</code> </p> </li> <li>
// <p>A <code>TargetString</code> of <code>login</code> </p> </li> </ul>
// <p>Further, you specify a <code>RateLimit</code> of 1,000.</p> <p>By adding this
// <code>RateBasedRule</code> to a <code>WebACL</code>, you could limit requests to
// your login page without affecting the rest of your site.</p>
func (c *Client) UpdateRateBasedRule(ctx context.Context, params *UpdateRateBasedRuleInput, optFns ...func(*Options)) (*UpdateRateBasedRuleOutput, error) {
	stack := middleware.NewStack("UpdateRateBasedRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateRateBasedRuleMiddlewares(stack)
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
	addOpUpdateRateBasedRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRateBasedRule(options.Region), middleware.Before)
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
			OperationName: "UpdateRateBasedRule",
			Err:           err,
		}
	}
	out := result.(*UpdateRateBasedRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRateBasedRuleInput struct {

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string

	// The maximum number of requests, which have an identical value in the field
	// specified by the RateKey, allowed in a five-minute period. If the number of
	// requests exceeds the RateLimit and the other predicates specified in the rule
	// are also met, AWS WAF triggers the action that is specified for this rule.
	//
	// This member is required.
	RateLimit *int64

	// The RuleId of the RateBasedRule that you want to update. RuleId is returned by
	// CreateRateBasedRule and by ListRateBasedRules ().
	//
	// This member is required.
	RuleId *string

	// An array of RuleUpdate objects that you want to insert into or delete from a
	// RateBasedRule ().
	//
	// This member is required.
	Updates []*types.RuleUpdate
}

type UpdateRateBasedRuleOutput struct {

	// The ChangeToken that you used to submit the UpdateRateBasedRule request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateRateBasedRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRateBasedRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRateBasedRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRateBasedRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateRateBasedRule",
	}
}
