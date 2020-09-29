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

// Creates a copy of an entire cluster or replication group at a specific moment in
// time.  <note> <p>This operation is valid for Redis only.</p> </note>
func (c *Client) CreateSnapshot(ctx context.Context, params *CreateSnapshotInput, optFns ...func(*Options)) (*CreateSnapshotOutput, error) {
	stack := middleware.NewStack("CreateSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateSnapshotMiddlewares(stack)
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
	addOpCreateSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshot(options.Region), middleware.Before)
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
			OperationName: "CreateSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateSnapshot operation.
type CreateSnapshotInput struct {
	// The identifier of an existing cluster. The snapshot is created from this
	// cluster.
	CacheClusterId *string
	// A name for the snapshot being created.
	SnapshotName *string
	// The ID of the KMS key used to encrypt the snapshot.
	KmsKeyId *string
	// The identifier of an existing replication group. The snapshot is created from
	// this replication group.
	ReplicationGroupId *string
}

type CreateSnapshotOutput struct {
	// Represents a copy of an entire Redis cluster as of the time when the snapshot
	// was taken.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateSnapshot",
	}
}