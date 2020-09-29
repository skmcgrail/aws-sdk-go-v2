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

// Returns information about cluster or replication group snapshots. By default,
// DescribeSnapshots lists all of your snapshots; it can optionally describe a
// single snapshot, or just the snapshots associated with a particular cache
// cluster.  <note> <p>This operation is valid for Redis only.</p> </note>
func (c *Client) DescribeSnapshots(ctx context.Context, params *DescribeSnapshotsInput, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	stack := middleware.NewStack("DescribeSnapshots", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeSnapshotsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshots(options.Region), middleware.Before)
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
			OperationName: "DescribeSnapshots",
			Err:           err,
		}
	}
	out := result.(*DescribeSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeSnapshotsMessage operation.
type DescribeSnapshotsInput struct {
	// A Boolean value which if true, the node group (shard) configuration is included
	// in the snapshot description.
	ShowNodeGroupConfig *bool
	// A user-supplied replication group identifier. If this parameter is specified,
	// only snapshots associated with that specific replication group are described.
	ReplicationGroupId *string
	// If set to system, the output shows snapshots that were automatically created by
	// ElastiCache. If set to user the output shows snapshots that were manually
	// created. If omitted, the output shows both automatically and manually created
	// snapshots.
	SnapshotSource *string
	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
	// A user-supplied cluster identifier. If this parameter is specified, only
	// snapshots associated with that specific cluster are described.
	CacheClusterId *string
	// A user-supplied name of the snapshot. If this parameter is specified, only this
	// snapshot are described.
	SnapshotName *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 50 Constraints: minimum
	// 20; maximum 50.
	MaxRecords *int32
}

// Represents the output of a DescribeSnapshots operation.
type DescribeSnapshotsOutput struct {
	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
	// A list of snapshots. Each item in the list contains detailed information about
	// one snapshot.
	Snapshots []*types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeSnapshotsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSnapshots{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSnapshots{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeSnapshots",
	}
}