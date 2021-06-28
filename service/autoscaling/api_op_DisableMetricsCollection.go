// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables group metrics for the specified Auto Scaling group.
func (c *Client) DisableMetricsCollection(ctx context.Context, params *DisableMetricsCollectionInput, optFns ...func(*Options)) (*DisableMetricsCollectionOutput, error) {
	if params == nil {
		params = &DisableMetricsCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableMetricsCollection", params, optFns, c.addOperationDisableMetricsCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableMetricsCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableMetricsCollectionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Specifies one or more of the following metrics:
	//
	// * GroupMinSize
	//
	// *
	// GroupMaxSize
	//
	// * GroupDesiredCapacity
	//
	// * GroupInServiceInstances
	//
	// *
	// GroupPendingInstances
	//
	// * GroupStandbyInstances
	//
	// * GroupTerminatingInstances
	//
	// *
	// GroupTotalInstances
	//
	// * GroupInServiceCapacity
	//
	// * GroupPendingCapacity
	//
	// *
	// GroupStandbyCapacity
	//
	// * GroupTerminatingCapacity
	//
	// * GroupTotalCapacity
	//
	// *
	// WarmPoolDesiredCapacity
	//
	// * WarmPoolWarmedCapacity
	//
	// * WarmPoolPendingCapacity
	//
	// *
	// WarmPoolTerminatingCapacity
	//
	// * WarmPoolTotalCapacity
	//
	// *
	// GroupAndWarmPoolDesiredCapacity
	//
	// * GroupAndWarmPoolTotalCapacity
	//
	// If you omit
	// this parameter, all metrics are disabled.
	Metrics []string

	noSmithyDocumentSerde
}

type DisableMetricsCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableMetricsCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableMetricsCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableMetricsCollection{}, middleware.After)
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
	if err = addOpDisableMetricsCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableMetricsCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableMetricsCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DisableMetricsCollection",
	}
}
