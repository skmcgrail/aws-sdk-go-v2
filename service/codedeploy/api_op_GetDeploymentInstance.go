// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about an instance as part of a deployment.
func (c *Client) GetDeploymentInstance(ctx context.Context, params *GetDeploymentInstanceInput, optFns ...func(*Options)) (*GetDeploymentInstanceOutput, error) {
	stack := middleware.NewStack("GetDeploymentInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDeploymentInstanceMiddlewares(stack)
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
	addOpGetDeploymentInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeploymentInstance(options.Region), middleware.Before)
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
			OperationName: "GetDeploymentInstance",
			Err:           err,
		}
	}
	out := result.(*GetDeploymentInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetDeploymentInstance operation.
type GetDeploymentInstanceInput struct {

	// The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	// The unique ID of an instance in the deployment group.
	//
	// This member is required.
	InstanceId *string
}

// Represents the output of a GetDeploymentInstance operation.
type GetDeploymentInstanceOutput struct {

	// Information about the instance.
	InstanceSummary *types.InstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDeploymentInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeploymentInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeploymentInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDeploymentInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetDeploymentInstance",
	}
}
