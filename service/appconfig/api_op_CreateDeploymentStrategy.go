// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// A deployment strategy defines important criteria for rolling out your
// configuration to the designated targets. A deployment strategy includes: the
// overall duration required, a percentage of targets to receive the deployment
// during each interval, an algorithm that defines how percentage grows, and bake
// time.
func (c *Client) CreateDeploymentStrategy(ctx context.Context, params *CreateDeploymentStrategyInput, optFns ...func(*Options)) (*CreateDeploymentStrategyOutput, error) {
	stack := middleware.NewStack("CreateDeploymentStrategy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDeploymentStrategyMiddlewares(stack)
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
	addOpCreateDeploymentStrategyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeploymentStrategy(options.Region), middleware.Before)
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
			OperationName: "CreateDeploymentStrategy",
			Err:           err,
		}
	}
	out := result.(*CreateDeploymentStrategyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeploymentStrategyInput struct {

	// Total amount of time for a deployment to last.
	//
	// This member is required.
	DeploymentDurationInMinutes *int32

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	//
	// This member is required.
	GrowthFactor *float32

	// A name for the deployment strategy.
	//
	// This member is required.
	Name *string

	// Save the deployment strategy to a Systems Manager (SSM) document.
	//
	// This member is required.
	ReplicateTo types.ReplicateTo

	// A description of the deployment strategy.
	Description *string

	// The amount of time AppConfig monitors for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int32

	// The algorithm used to define how percentage grows over time. AWS AppConfig
	// supports the following growth types: Linear: For this type, AppConfig processes
	// the deployment by dividing the total number of targets by the value specified
	// for Step percentage. For example, a linear deployment that uses a Step
	// percentage of 10 deploys the configuration to 10 percent of the hosts. After
	// those deployments are complete, the system deploys the configuration to the next
	// 10 percent. This continues until 100% of the targets have successfully received
	// the configuration.  <p> <b>Exponential</b>: For this type, AppConfig processes
	// the deployment exponentially using the following formula: <code>G*(2^N)</code>.
	// In this formula, <code>G</code> is the growth factor specified by the user and
	// <code>N</code> is the number of steps until the configuration is deployed to all
	// targets. For example, if you specify a growth factor of 2, then the system rolls
	// out the configuration as follows:</p> <p> <code>2*(2^0)</code> </p> <p>
	// <code>2*(2^1)</code> </p> <p> <code>2*(2^2)</code> </p> <p>Expressed
	// numerically, the deployment rolls out as follows: 2% of the targets, 4% of the
	// targets, 8% of the targets, and continues until the configuration has been
	// deployed to all targets.</p>
	GrowthType types.GrowthType

	// Metadata to assign to the deployment strategy. Tags help organize and categorize
	// your AppConfig resources. Each tag consists of a key and an optional value, both
	// of which you define.
	Tags map[string]*string
}

type CreateDeploymentStrategyOutput struct {

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int32

	// The description of the deployment strategy.
	Description *string

	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int32

	// The percentage of targets that received a deployed configuration during each
	// interval.
	GrowthFactor *float32

	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType

	// The deployment strategy ID.
	Id *string

	// The name of the deployment strategy.
	Name *string

	// Save the deployment strategy to a Systems Manager (SSM) document.
	ReplicateTo types.ReplicateTo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDeploymentStrategyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeploymentStrategy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeploymentStrategy{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDeploymentStrategy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateDeploymentStrategy",
	}
}
