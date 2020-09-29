// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new DB parameter group.  <p>A DB parameter group is initially created
// with the default parameters for the database engine used by the DB instance. To
// provide custom values for any of the parameters, you must modify the group after
// creating it using <i>ModifyDBParameterGroup</i>. Once you've created a DB
// parameter group, you need to associate it with your DB instance using
// <i>ModifyDBInstance</i>. When you associate a new DB parameter group with a
// running DB instance, you need to reboot the DB instance without failover for the
// new DB parameter group and associated settings to take effect.</p> <important>
// <p>After you create a DB parameter group, you should wait at least 5 minutes
// before creating your first DB instance that uses that DB parameter group as the
// default parameter group. This allows Amazon Neptune to fully complete the create
// action before the parameter group is used as the default for a new DB instance.
// This is especially important for parameters that are critical when creating the
// default database for a DB instance, such as the character set for the default
// database defined by the <code>character_set_database</code> parameter. You can
// use the <i>Parameter Groups</i> option of the Amazon Neptune console or the
// <i>DescribeDBParameters</i> command to verify that your DB parameter group has
// been created or modified.</p> </important>
func (c *Client) CreateDBParameterGroup(ctx context.Context, params *CreateDBParameterGroupInput, optFns ...func(*Options)) (*CreateDBParameterGroupOutput, error) {
	stack := middleware.NewStack("CreateDBParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBParameterGroupMiddlewares(stack)
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
	addOpCreateDBParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBParameterGroup(options.Region), middleware.Before)
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
			OperationName: "CreateDBParameterGroup",
			Err:           err,
		}
	}
	out := result.(*CreateDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBParameterGroupInput struct {
	// The name of the DB parameter group. Constraints:
	//
	//     * Must be 1 to 255
	// letters, numbers, or hyphens.
	//
	//     * First character must be a letter
	//
	//     *
	// Cannot end with a hyphen or contain two consecutive hyphens
	//
	// This value is
	// stored as a lowercase string.
	DBParameterGroupName *string
	// The tags to be assigned to the new DB parameter group.
	Tags []*types.Tag
	// The description for the DB parameter group.
	Description *string
	// The DB parameter group family name. A DB parameter group can be associated with
	// one and only one DB parameter group family, and can be applied only to a DB
	// instance running a database engine and engine version compatible with that DB
	// parameter group family.
	DBParameterGroupFamily *string
}

type CreateDBParameterGroupOutput struct {
	// Contains the details of an Amazon Neptune DB parameter group. This data type is
	// used as a response element in the DescribeDBParameterGroups () action.
	DBParameterGroup *types.DBParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBParameterGroup",
	}
}