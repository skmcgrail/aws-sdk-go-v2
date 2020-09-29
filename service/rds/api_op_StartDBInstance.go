// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an Amazon RDS DB instance that was stopped using the AWS console, the
// stop-db-instance AWS CLI command, or the StopDBInstance action.  <p>For more
// information, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StartInstance.html">
// Starting an Amazon RDS DB instance That Was Previously Stopped</a> in the
// <i>Amazon RDS User Guide.</i> </p> <note> <p> This command doesn't apply to
// Aurora MySQL and Aurora PostgreSQL. For Aurora DB clusters, use
// <code>StartDBCluster</code> instead. </p> </note>
func (c *Client) StartDBInstance(ctx context.Context, params *StartDBInstanceInput, optFns ...func(*Options)) (*StartDBInstanceOutput, error) {
	stack := middleware.NewStack("StartDBInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpStartDBInstanceMiddlewares(stack)
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
	addOpStartDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDBInstance(options.Region), middleware.Before)
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
			OperationName: "StartDBInstance",
			Err:           err,
		}
	}
	out := result.(*StartDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDBInstanceInput struct {
	// The user-supplied instance identifier.
	DBInstanceIdentifier *string
}

type StartDBInstanceOutput struct {
	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpStartDBInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpStartDBInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpStartDBInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StartDBInstance",
	}
}