// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified WebACL. This operation completely replaces the mutable
// specifications that you already have for the web ACL with the ones that you
// provide to this call. To modify the web ACL, retrieve it by calling GetWebACL,
// update the settings as needed, and then provide the complete web ACL
// specification to this call. A web ACL defines a collection of rules to use to
// inspect and control web requests. Each rule has an action defined (allow, block,
// or count) for requests that match the statement of the rule. In the web ACL, you
// assign a default action to take (allow, block) for any request that does not
// match any of the rules. The rules in a web ACL can be a combination of the types
// Rule, RuleGroup, and managed rule group. You can associate a web ACL with one or
// more Amazon Web Services resources to protect. The resources can be an Amazon
// CloudFront distribution, an Amazon API Gateway REST API, an Application Load
// Balancer, or an AppSync GraphQL API.
func (c *Client) UpdateWebACL(ctx context.Context, params *UpdateWebACLInput, optFns ...func(*Options)) (*UpdateWebACLOutput, error) {
	if params == nil {
		params = &UpdateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWebACL", params, optFns, c.addOperationUpdateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWebACLInput struct {

	// The action to perform if none of the Rules contained in the WebACL match.
	//
	// This member is required.
	DefaultAction *types.DefaultAction

	// The unique identifier for the web ACL. This ID is returned in the responses to
	// create and list commands. You provide it to operations like update and delete.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. WAF returns a token to your get and list
	// requests, to mark the state of the entity at the time of the request. To make
	// changes to the entity associated with the token, you provide the token to
	// operations like update and delete. WAF uses the token to ensure that no changes
	// have been made to the entity since you last retrieved it. If a change has been
	// made, the update fails with a WAFOptimisticLockException. If this happens,
	// perform another get, and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the web ACL. You cannot change the name of a web ACL after you
	// create it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, or an AppSync GraphQL API. To work with
	// CloudFront, you must also specify the Region US East (N. Virginia) as
	// follows:
	//
	// * CLI - Specify the Region when you use the CloudFront scope:
	// --scope=CLOUDFRONT --region=us-east-1.
	//
	// * API and SDKs - For all calls, use the
	// Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// This member is required.
	VisibilityConfig *types.VisibilityConfig

	// A map of custom response keys and content bodies. When you create a rule with a
	// block action, you can send a custom response to the web request. You define
	// these for the web ACL, and then use them in the rules and default actions that
	// you define in the web ACL. For information about customizing web requests and
	// responses, see Customizing web requests and responses in WAF
	// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html)
	// in the WAF Developer Guide
	// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). For
	// information about the limits on count and size for custom request and response
	// settings, see WAF quotas
	// (https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the WAF
	// Developer Guide
	// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
	CustomResponseBodies map[string]types.CustomResponseBody

	// A description of the web ACL that helps with identification.
	Description *string

	// The Rule statements used to identify the web requests that you want to allow,
	// block, or count. Each rule includes one top-level statement that WAF uses to
	// identify matching web requests, and parameters that govern how WAF handles them.
	Rules []types.Rule

	noSmithyDocumentSerde
}

type UpdateWebACLOutput struct {

	// A token used for optimistic locking. WAF returns this token to your update
	// requests. You use NextLockToken in the same manner as you use LockToken.
	NextLockToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWebACL{}, middleware.After)
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
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = addOpUpdateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "UpdateWebACL",
	}
}
