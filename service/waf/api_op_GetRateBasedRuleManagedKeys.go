// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
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
// global use. Returns an array of IP addresses currently being blocked by the
// RateBasedRule () that is specified by the RuleId. The maximum number of managed
// keys that will be blocked is 10,000. If more than 10,000 addresses exceed the
// rate limit, the 10,000 addresses with the highest rates will be blocked.
func (c *Client) GetRateBasedRuleManagedKeys(ctx context.Context, params *GetRateBasedRuleManagedKeysInput, optFns ...func(*Options)) (*GetRateBasedRuleManagedKeysOutput, error) {
	stack := middleware.NewStack("GetRateBasedRuleManagedKeys", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRateBasedRuleManagedKeysMiddlewares(stack)
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
	addOpGetRateBasedRuleManagedKeysValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRateBasedRuleManagedKeys(options.Region), middleware.Before)
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
			OperationName: "GetRateBasedRuleManagedKeys",
			Err:           err,
		}
	}
	out := result.(*GetRateBasedRuleManagedKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRateBasedRuleManagedKeysInput struct {

	// The RuleId of the RateBasedRule () for which you want to get a list of
	// ManagedKeys. RuleId is returned by CreateRateBasedRule () and by
	// ListRateBasedRules ().
	//
	// This member is required.
	RuleId *string

	// A null value and not currently used. Do not include this in your request.
	NextMarker *string
}

type GetRateBasedRuleManagedKeysOutput struct {

	// An array of IP addresses that currently are blocked by the specified
	// RateBasedRule ().
	ManagedKeys []*string

	// A null value and not currently used.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRateBasedRuleManagedKeysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRateBasedRuleManagedKeys{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRateBasedRuleManagedKeys{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRateBasedRuleManagedKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "GetRateBasedRuleManagedKeys",
	}
}
