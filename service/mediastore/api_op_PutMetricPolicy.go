// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The metric policy that you want to add to the container. A metric policy allows
// AWS Elemental MediaStore to send metrics to Amazon CloudWatch. It takes up to 20
// minutes for the new policy to take effect.
func (c *Client) PutMetricPolicy(ctx context.Context, params *PutMetricPolicyInput, optFns ...func(*Options)) (*PutMetricPolicyOutput, error) {
	stack := middleware.NewStack("PutMetricPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutMetricPolicyMiddlewares(stack)
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
	addOpPutMetricPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricPolicy(options.Region), middleware.Before)
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
			OperationName: "PutMetricPolicy",
			Err:           err,
		}
	}
	out := result.(*PutMetricPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricPolicyInput struct {

	// The name of the container that you want to add the metric policy to.
	//
	// This member is required.
	ContainerName *string

	// The metric policy that you want to associate with the container. In the policy,
	// you must indicate whether you want MediaStore to send container-level metrics.
	// You can also include up to five rules to define groups of objects that you want
	// MediaStore to send object-level metrics for. If you include rules in the policy,
	// construct each rule with both of the following:
	//
	//     * An object group that
	// defines which objects to include in the group. The definition can be a path or a
	// file name, but it can't have more than 900 characters. Valid characters are:
	// a-z, A-Z, 0-9, _ (underscore), = (equal), : (colon), . (period), - (hyphen), ~
	// (tilde), / (forward slash), and * (asterisk). Wildcards (*) are acceptable.
	//
	//
	// * An object group name that allows you to refer to the object group. The name
	// can't have more than 30 characters. Valid characters are: a-z, A-Z, 0-9, and _
	// (underscore).
	//
	// This member is required.
	MetricPolicy *types.MetricPolicy
}

type PutMetricPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutMetricPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutMetricPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutMetricPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutMetricPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "PutMetricPolicy",
	}
}
