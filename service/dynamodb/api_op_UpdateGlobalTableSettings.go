// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates settings for a global table.
func (c *Client) UpdateGlobalTableSettings(ctx context.Context, params *UpdateGlobalTableSettingsInput, optFns ...func(*Options)) (*UpdateGlobalTableSettingsOutput, error) {
	stack := middleware.NewStack("UpdateGlobalTableSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpUpdateGlobalTableSettingsMiddlewares(stack)
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
	addOpUpdateGlobalTableSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlobalTableSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)

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
			OperationName: "UpdateGlobalTableSettings",
			Err:           err,
		}
	}
	out := result.(*UpdateGlobalTableSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlobalTableSettingsInput struct {

	// The name of the global table
	//
	// This member is required.
	GlobalTableName *string

	// The billing mode of the global table. If GlobalTableBillingMode is not
	// specified, the global table defaults to PROVISIONED capacity billing mode.
	//
	//
	// * PROVISIONED - We recommend using PROVISIONED for predictable workloads.
	// PROVISIONED sets the billing mode to Provisioned Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
	//
	//
	// * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable
	// workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
	GlobalTableBillingMode types.BillingMode

	// Represents the settings of a global secondary index for a global table that will
	// be modified.
	GlobalTableGlobalSecondaryIndexSettingsUpdate []*types.GlobalTableGlobalSecondaryIndexSettingsUpdate

	// Auto scaling settings for managing provisioned write capacity for the global
	// table.
	GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate *types.AutoScalingSettingsUpdate

	// The maximum number of writes consumed per second before DynamoDB returns a
	// ThrottlingException.
	GlobalTableProvisionedWriteCapacityUnits *int64

	// Represents the settings for a global table in a Region that will be modified.
	ReplicaSettingsUpdate []*types.ReplicaSettingsUpdate
}

type UpdateGlobalTableSettingsOutput struct {

	// The name of the global table.
	GlobalTableName *string

	// The Region-specific settings for the global table.
	ReplicaSettings []*types.ReplicaSettingsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpUpdateGlobalTableSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateGlobalTableSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateGlobalTableSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGlobalTableSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateGlobalTableSettings",
	}
}
