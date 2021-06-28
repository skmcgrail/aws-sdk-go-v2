// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Forces a failover for a DB cluster. A failover for a DB cluster promotes one of
// the Aurora Replicas (read-only instances) in the DB cluster to be the primary
// instance (the cluster writer). Amazon Aurora will automatically fail over to an
// Aurora Replica, if one exists, when the primary instance fails. You can force a
// failover when you want to simulate a failure of a primary instance for testing.
// Because each instance in a DB cluster has its own endpoint address, you will
// need to clean up and re-establish any existing connections that use those
// endpoint addresses when the failover is complete. For more information on Amazon
// Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) FailoverDBCluster(ctx context.Context, params *FailoverDBClusterInput, optFns ...func(*Options)) (*FailoverDBClusterOutput, error) {
	if params == nil {
		params = &FailoverDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FailoverDBCluster", params, optFns, c.addOperationFailoverDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FailoverDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type FailoverDBClusterInput struct {

	// A DB cluster identifier to force a failover for. This parameter isn't
	// case-sensitive. Constraints:
	//
	// * Must match the identifier of an existing
	// DBCluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the instance to promote to the primary instance. You must specify
	// the instance identifier for an Aurora Replica in the DB cluster. For example,
	// mydbcluster-replica1.
	TargetDBInstanceIdentifier *string

	noSmithyDocumentSerde
}

type FailoverDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster. This data type is used as a
	// response element in the DescribeDBClusters, StopDBCluster, and StartDBCluster
	// actions.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFailoverDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFailoverDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpFailoverDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverDBCluster(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opFailoverDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "FailoverDBCluster",
	}
}
