// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a snapshot of a volume.  <p>You can take snapshots of your gateway
// volumes on a scheduled or ad hoc basis. This API action enables you to delete a
// snapshot schedule for a volume. For more information, see <a
// href="https://docs.aws.amazon.com/storagegatewaylatest/userguide/backing-up-volumes.html">Backing
// up your volumes</a>. In the <code>DeleteSnapshotSchedule</code> request, you
// identify the volume by providing its Amazon Resource Name (ARN). This operation
// is only supported in stored and cached volume gateway types.</p> <note> <p>To
// list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, go to <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html">DescribeSnapshots</a>
// in the <i>Amazon Elastic Compute Cloud API Reference</i>.</p> </note>
func (c *Client) DeleteSnapshotSchedule(ctx context.Context, params *DeleteSnapshotScheduleInput, optFns ...func(*Options)) (*DeleteSnapshotScheduleOutput, error) {
	stack := middleware.NewStack("DeleteSnapshotSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteSnapshotScheduleMiddlewares(stack)
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
	addOpDeleteSnapshotScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSnapshotSchedule(options.Region), middleware.Before)
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
			OperationName: "DeleteSnapshotSchedule",
			Err:           err,
		}
	}
	out := result.(*DeleteSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSnapshotScheduleInput struct {
	// The volume which snapshot schedule to delete.
	VolumeARN *string
}

type DeleteSnapshotScheduleOutput struct {
	// The volume which snapshot schedule was deleted.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteSnapshotScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSnapshotSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSnapshotSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSnapshotSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteSnapshotSchedule",
	}
}