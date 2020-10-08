// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables enhanced monitoring.
func (c *Client) DisableEnhancedMonitoring(ctx context.Context, params *DisableEnhancedMonitoringInput, optFns ...func(*Options)) (*DisableEnhancedMonitoringOutput, error) {
	stack := middleware.NewStack("DisableEnhancedMonitoring", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableEnhancedMonitoringMiddlewares(stack)
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
	addOpDisableEnhancedMonitoringValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableEnhancedMonitoring(options.Region), middleware.Before)
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
			OperationName: "DisableEnhancedMonitoring",
			Err:           err,
		}
	}
	out := result.(*DisableEnhancedMonitoringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for DisableEnhancedMonitoring ().
type DisableEnhancedMonitoringInput struct {

	// List of shard-level metrics to disable. The following are the valid shard-level
	// metrics. The value "ALL" disables every metric.
	//
	//     * IncomingBytes
	//
	//     *
	// IncomingRecords
	//
	//     * OutgoingBytes
	//
	//     * OutgoingRecords
	//
	//     *
	// WriteProvisionedThroughputExceeded
	//
	//     * ReadProvisionedThroughputExceeded
	//
	//
	// * IteratorAgeMilliseconds
	//
	//     * ALL
	//
	// For more information, see Monitoring the
	// Amazon Kinesis Data Streams Service with Amazon CloudWatch
	// (https://docs.aws.amazon.com/kinesis/latest/dev/monitoring-with-cloudwatch.html)
	// in the Amazon Kinesis Data Streams Developer Guide.
	//
	// This member is required.
	ShardLevelMetrics []types.MetricsName

	// The name of the Kinesis data stream for which to disable enhanced monitoring.
	//
	// This member is required.
	StreamName *string
}

// Represents the output for EnableEnhancedMonitoring () and
// DisableEnhancedMonitoring ().
type DisableEnhancedMonitoringOutput struct {

	// Represents the current state of the metrics that are in the enhanced state
	// before the operation.
	CurrentShardLevelMetrics []types.MetricsName

	// Represents the list of all the metrics that would be in the enhanced state after
	// the operation.
	DesiredShardLevelMetrics []types.MetricsName

	// The name of the Kinesis data stream.
	StreamName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableEnhancedMonitoringMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableEnhancedMonitoring{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableEnhancedMonitoring{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableEnhancedMonitoring(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DisableEnhancedMonitoring",
	}
}
