// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Starts the replication task. For more information about AWS DMS tasks, see
// Working with Migration Tasks
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html) in the AWS
// Database Migration Service User Guide.
func (c *Client) StartReplicationTask(ctx context.Context, params *StartReplicationTaskInput, optFns ...func(*Options)) (*StartReplicationTaskOutput, error) {
	stack := middleware.NewStack("StartReplicationTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartReplicationTaskMiddlewares(stack)
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
	addOpStartReplicationTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartReplicationTask(options.Region), middleware.Before)
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
			OperationName: "StartReplicationTask",
			Err:           err,
		}
	}
	out := result.(*StartReplicationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type StartReplicationTaskInput struct {

	// The Amazon Resource Name (ARN) of the replication task to be started.
	//
	// This member is required.
	ReplicationTaskArn *string

	// The type of replication task.
	//
	// This member is required.
	StartReplicationTaskType types.StartReplicationTaskTypeValue

	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error. The value can be in date,
	// checkpoint, or LSN/SCN format. Date Example: --cdc-start-position
	// “2018-03-08T12:12:12” Checkpoint Example: --cdc-start-position
	// "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373” When you use
	// this task setting with a source PostgreSQL database, a logical replication slot
	// should already be created and associated with the source endpoint. You can
	// verify this by setting the slotName extra connection attribute to the name of
	// this logical replication slot. For more information, see Extra Connection
	// Attributes When Using PostgreSQL as a Source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib).
	CdcStartPosition *string

	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation to
	// start. Specifying both values results in an error. Timestamp Example:
	// --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time

	// Indicates when you want a change data capture (CDC) operation to stop. The value
	// can be either server time or commit time. Server time example:
	// --cdc-stop-position “server_time:3018-02-09T12:12:12” Commit time example:
	// --cdc-stop-position “commit_time: 3018-02-09T12:12:12 “
	CdcStopPosition *string
}

//
type StartReplicationTaskOutput struct {

	// The replication task started.
	ReplicationTask *types.ReplicationTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartReplicationTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartReplicationTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReplicationTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartReplicationTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "StartReplicationTask",
	}
}
