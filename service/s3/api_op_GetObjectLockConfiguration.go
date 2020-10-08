// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the Object Lock configuration for a bucket. The rule specified in the
// Object Lock configuration will be applied by default to every new object placed
// in the specified bucket. For more information, see Locking Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
func (c *Client) GetObjectLockConfiguration(ctx context.Context, params *GetObjectLockConfigurationInput, optFns ...func(*Options)) (*GetObjectLockConfigurationOutput, error) {
	stack := middleware.NewStack("GetObjectLockConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetObjectLockConfigurationMiddlewares(stack)
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
	addOpGetObjectLockConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectLockConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

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
			OperationName: "GetObjectLockConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetObjectLockConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectLockConfigurationInput struct {

	// The bucket whose Object Lock configuration you want to retrieve.
	//
	// This member is required.
	Bucket *string
}

type GetObjectLockConfigurationOutput struct {

	// The specified bucket's Object Lock configuration.
	ObjectLockConfiguration *types.ObjectLockConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetObjectLockConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetObjectLockConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetObjectLockConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetObjectLockConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetObjectLockConfiguration",
	}
}
