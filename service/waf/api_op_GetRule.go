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
// global use. Returns the Rule () that is specified by the RuleId that you
// included in the GetRule request.
func (c *Client) GetRule(ctx context.Context, params *GetRuleInput, optFns ...func(*Options)) (*GetRuleOutput, error) {
	stack := middleware.NewStack("GetRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRuleMiddlewares(stack)
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
	addOpGetRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRule(options.Region), middleware.Before)
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
			OperationName: "GetRule",
			Err:           err,
		}
	}
	out := result.(*GetRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRuleInput struct {

	// The RuleId of the Rule () that you want to get. RuleId is returned by CreateRule
	// () and by ListRules ().
	//
	// This member is required.
	RuleId *string
}

type GetRuleOutput struct {

	// Information about the Rule () that you specified in the GetRule request. For
	// more information, see the following topics:
	//
	//     * Rule (): Contains MetricName,
	// Name, an array of Predicate objects, and RuleId
	//
	//     * Predicate (): Each
	// Predicate object contains DataId, Negated, and Type
	Rule *types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "GetRule",
	}
}
