// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the bandwidth rate limits of a gateway. You can update both the upload
// and download bandwidth rate limit or specify only one of the two. If you don't
// set a bandwidth rate limit, the existing rate limit remains. This operation is
// supported for the stored volume, cached volume, and tape gateway types.  <p>By
// default, a gateway's bandwidth rate limits are not set. If you don't set any
// limit, the gateway does not have any limitations on its bandwidth usage and
// could potentially use the maximum available bandwidth.</p> <p>To specify which
// gateway to update, use the Amazon Resource Name (ARN) of the gateway in your
// request.</p>
func (c *Client) UpdateBandwidthRateLimit(ctx context.Context, params *UpdateBandwidthRateLimitInput, optFns ...func(*Options)) (*UpdateBandwidthRateLimitOutput, error) {
	stack := middleware.NewStack("UpdateBandwidthRateLimit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateBandwidthRateLimitMiddlewares(stack)
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
	addOpUpdateBandwidthRateLimitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBandwidthRateLimit(options.Region), middleware.Before)
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
			OperationName: "UpdateBandwidthRateLimit",
			Err:           err,
		}
	}
	out := result.(*UpdateBandwidthRateLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:  <ul> <li> <p>
// <a>UpdateBandwidthRateLimitInput$AverageDownloadRateLimitInBitsPerSec</a> </p>
// </li> <li> <p>
// <a>UpdateBandwidthRateLimitInput$AverageUploadRateLimitInBitsPerSec</a> </p>
// </li> </ul>
type UpdateBandwidthRateLimitInput struct {
	// The average upload bandwidth rate limit in bits per second.
	AverageUploadRateLimitInBitsPerSec *int64
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
	// The average download bandwidth rate limit in bits per second.
	AverageDownloadRateLimitInBitsPerSec *int64
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// throttle information was updated.
type UpdateBandwidthRateLimitOutput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateBandwidthRateLimitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBandwidthRateLimit{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBandwidthRateLimit{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateBandwidthRateLimit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateBandwidthRateLimit",
	}
}