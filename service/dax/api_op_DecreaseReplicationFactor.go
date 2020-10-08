// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes one or more nodes from a DAX cluster. You cannot use
// DecreaseReplicationFactor to remove the last node in a DAX cluster. If you need
// to do this, use DeleteCluster instead.
func (c *Client) DecreaseReplicationFactor(ctx context.Context, params *DecreaseReplicationFactorInput, optFns ...func(*Options)) (*DecreaseReplicationFactorOutput, error) {
	stack := middleware.NewStack("DecreaseReplicationFactor", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDecreaseReplicationFactorMiddlewares(stack)
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
	addOpDecreaseReplicationFactorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseReplicationFactor(options.Region), middleware.Before)
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
			OperationName: "DecreaseReplicationFactor",
			Err:           err,
		}
	}
	out := result.(*DecreaseReplicationFactorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecreaseReplicationFactorInput struct {

	// The name of the DAX cluster from which you want to remove nodes.
	//
	// This member is required.
	ClusterName *string

	// The new number of nodes for the DAX cluster.
	//
	// This member is required.
	NewReplicationFactor *int32

	// The Availability Zone(s) from which to remove nodes.
	AvailabilityZones []*string

	// The unique identifiers of the nodes to be removed from the cluster.
	NodeIdsToRemove []*string
}

type DecreaseReplicationFactorOutput struct {

	// A description of the DAX cluster, after you have decreased its replication
	// factor.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDecreaseReplicationFactorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDecreaseReplicationFactor{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDecreaseReplicationFactor{}, middleware.After)
}

func newServiceMetadataMiddleware_opDecreaseReplicationFactor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "DecreaseReplicationFactor",
	}
}
