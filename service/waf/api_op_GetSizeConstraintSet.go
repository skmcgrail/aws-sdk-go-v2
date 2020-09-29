// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
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
// global use. Returns the SizeConstraintSet () specified by SizeConstraintSetId.
func (c *Client) GetSizeConstraintSet(ctx context.Context, params *GetSizeConstraintSetInput, optFns ...func(*Options)) (*GetSizeConstraintSetOutput, error) {
	stack := middleware.NewStack("GetSizeConstraintSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSizeConstraintSetMiddlewares(stack)
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
	addOpGetSizeConstraintSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSizeConstraintSet(options.Region), middleware.Before)
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
			OperationName: "GetSizeConstraintSet",
			Err:           err,
		}
	}
	out := result.(*GetSizeConstraintSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSizeConstraintSetInput struct {
	// The SizeConstraintSetId of the SizeConstraintSet () that you want to get.
	// SizeConstraintSetId is returned by CreateSizeConstraintSet () and by
	// ListSizeConstraintSets ().
	SizeConstraintSetId *string
}

type GetSizeConstraintSetOutput struct {
	// Information about the SizeConstraintSet () that you specified in the
	// GetSizeConstraintSet request. For more information, see the following topics:
	//
	//
	// * SizeConstraintSet (): Contains SizeConstraintSetId, SizeConstraints, and
	// Name
	//
	//     * SizeConstraints: Contains an array of SizeConstraint () objects.
	// Each SizeConstraint object contains FieldToMatch (), TextTransformation,
	// ComparisonOperator, and Size
	//
	//     * FieldToMatch (): Contains Data and Type
	SizeConstraintSet *types.SizeConstraintSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSizeConstraintSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSizeConstraintSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSizeConstraintSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSizeConstraintSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "GetSizeConstraintSet",
	}
}