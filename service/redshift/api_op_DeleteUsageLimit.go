// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a usage limit from a cluster.
func (c *Client) DeleteUsageLimit(ctx context.Context, params *DeleteUsageLimitInput, optFns ...func(*Options)) (*DeleteUsageLimitOutput, error) {
	stack := middleware.NewStack("DeleteUsageLimit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteUsageLimitMiddlewares(stack)
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
	addOpDeleteUsageLimitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUsageLimit(options.Region), middleware.Before)
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
			OperationName: "DeleteUsageLimit",
			Err:           err,
		}
	}
	out := result.(*DeleteUsageLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUsageLimitInput struct {

	// The identifier of the usage limit to delete.
	//
	// This member is required.
	UsageLimitId *string
}

type DeleteUsageLimitOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteUsageLimitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteUsageLimit{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteUsageLimit{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUsageLimit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeleteUsageLimit",
	}
}
