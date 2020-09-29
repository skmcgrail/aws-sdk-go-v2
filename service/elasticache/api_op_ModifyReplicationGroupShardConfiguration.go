// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a replication group's shards (node groups) by allowing you to add
// shards, remove shards, or rebalance the keyspaces among exisiting shards.
func (c *Client) ModifyReplicationGroupShardConfiguration(ctx context.Context, params *ModifyReplicationGroupShardConfigurationInput, optFns ...func(*Options)) (*ModifyReplicationGroupShardConfigurationOutput, error) {
	stack := middleware.NewStack("ModifyReplicationGroupShardConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyReplicationGroupShardConfigurationMiddlewares(stack)
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
	addOpModifyReplicationGroupShardConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationGroupShardConfiguration(options.Region), middleware.Before)
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
			OperationName: "ModifyReplicationGroupShardConfiguration",
			Err:           err,
		}
	}
	out := result.(*ModifyReplicationGroupShardConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a ModifyReplicationGroupShardConfiguration operation.
type ModifyReplicationGroupShardConfigurationInput struct {
	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRetain is a list of NodeGroupIds to retain in the cluster.
	// ElastiCache for Redis will attempt to remove all node groups except those listed
	// by NodeGroupsToRetain from the cluster.
	NodeGroupsToRetain []*string
	// Indicates that the shard reconfiguration process begins immediately. At present,
	// the only permitted value for this parameter is true. Value: true
	ApplyImmediately *bool
	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
	// NodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
	// ElastiCache for Redis will attempt to remove all node groups listed by
	// NodeGroupsToRemove from the cluster.
	NodeGroupsToRemove []*string
	// The number of node groups (shards) that results from the modification of the
	// shard configuration.
	NodeGroupCount *int32
	// The name of the Redis (cluster mode enabled) cluster (replication group) on
	// which the shards are to be configured.
	ReplicationGroupId *string
	// Specifies the preferred availability zones for each node group in the cluster.
	// If the value of NodeGroupCount is greater than the current number of node groups
	// (shards), you can use this parameter to specify the preferred availability zones
	// of the cluster's shards. If you omit this parameter ElastiCache selects
	// availability zones for you. You can specify this parameter only if the value of
	// NodeGroupCount is greater than the current number of node groups (shards).
	ReshardingConfiguration []*types.ReshardingConfiguration
}

type ModifyReplicationGroupShardConfigurationOutput struct {
	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyReplicationGroupShardConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyReplicationGroupShardConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyReplicationGroupShardConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyReplicationGroupShardConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyReplicationGroupShardConfiguration",
	}
}