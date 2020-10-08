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
	"time"
)

// Stops a deployment. This API action works only on deployments that have a status
// of DEPLOYING. This action moves the deployment to a status of ROLLED_BACK.
func (c *Client) StopDeployment(ctx context.Context, params *StopDeploymentInput, optFns ...func(*Options)) (*StopDeploymentOutput, error) {
	stack := middleware.NewStack("StopDeployment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStopDeploymentMiddlewares(stack)
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
	addOpStopDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDeployment(options.Region), middleware.Before)
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
			OperationName: "StopDeployment",
			Err:           err,
		}
	}
	out := result.(*StopDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDeploymentInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The sequence number of the deployment.
	//
	// This member is required.
	DeploymentNumber *int32

	// The environment ID.
	//
	// This member is required.
	EnvironmentId *string
}

type StopDeploymentOutput struct {

	// The ID of the application that was deployed.
	ApplicationId *string

	// The time the deployment completed.
	CompletedAt *time.Time

	// Information about the source location of the configuration.
	ConfigurationLocationUri *string

	// The name of the configuration.
	ConfigurationName *string

	// The ID of the configuration profile that was deployed.
	ConfigurationProfileId *string

	// The configuration version that was deployed.
	ConfigurationVersion *string

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int32

	// The sequence number of the deployment.
	DeploymentNumber *int32

	// The ID of the deployment strategy that was deployed.
	DeploymentStrategyId *string

	// The description of the deployment.
	Description *string

	// The ID of the environment that was deployed.
	EnvironmentId *string

	// A list containing all events related to a deployment. The most recent events are
	// displayed first.
	EventLog []*types.DeploymentEvent

	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int32

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor *float32

	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType

	// The percentage of targets for which the deployment is available.
	PercentageComplete *float32

	// The time the deployment started.
	StartedAt *time.Time

	// The state of the deployment.
	State types.DeploymentState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStopDeploymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStopDeployment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStopDeployment{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "StopDeployment",
	}
}
