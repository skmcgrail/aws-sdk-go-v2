// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an instance fleet to a running cluster. The instance fleet configuration is
// available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x.
func (c *Client) AddInstanceFleet(ctx context.Context, params *AddInstanceFleetInput, optFns ...func(*Options)) (*AddInstanceFleetOutput, error) {
	stack := middleware.NewStack("AddInstanceFleet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddInstanceFleetMiddlewares(stack)
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
	addOpAddInstanceFleetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddInstanceFleet(options.Region), middleware.Before)
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
			OperationName: "AddInstanceFleet",
			Err:           err,
		}
	}
	out := result.(*AddInstanceFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddInstanceFleetInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// Specifies the configuration of the instance fleet.
	//
	// This member is required.
	InstanceFleet *types.InstanceFleetConfig
}

type AddInstanceFleetOutput struct {

	// The Amazon Resource Name of the cluster.
	ClusterArn *string

	// The unique identifier of the cluster.
	ClusterId *string

	// The unique identifier of the instance fleet.
	InstanceFleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddInstanceFleetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddInstanceFleet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddInstanceFleet{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddInstanceFleet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "AddInstanceFleet",
	}
}
