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

// Modifies the parameters of a DB cluster parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName, ParameterValue, and
// ApplyMethod. A maximum of 20 parameters can be modified in a single request.
// <note> <p>Changes to dynamic parameters are applied immediately. Changes to
// static parameters require a reboot without failover to the DB cluster associated
// with the parameter group before the change can take effect.</p> </note>
// <important> <p>After you create a DB cluster parameter group, you should wait at
// least 5 minutes before creating your first DB cluster that uses that DB cluster
// parameter group as the default parameter group. This allows Amazon Neptune to
// fully complete the create action before the parameter group is used as the
// default for a new DB cluster. This is especially important for parameters that
// are critical when creating the default database for a DB cluster, such as the
// character set for the default database defined by the
// <code>character_set_database</code> parameter. You can use the <i>Parameter
// Groups</i> option of the Amazon Neptune console or the
// <a>DescribeDBClusterParameters</a> command to verify that your DB cluster
// parameter group has been created or modified.</p> </important>
func (c *Client) ModifyDBClusterParameterGroup(ctx context.Context, params *ModifyDBClusterParameterGroupInput, optFns ...func(*Options)) (*ModifyDBClusterParameterGroupOutput, error) {
	stack := middleware.NewStack("ModifyDBClusterParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBClusterParameterGroupMiddlewares(stack)
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
	addOpModifyDBClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(options.Region), middleware.Before)
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
			OperationName: "ModifyDBClusterParameterGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterParameterGroupInput struct {
	// A list of parameters in the DB cluster parameter group to modify.
	Parameters []*types.Parameter
	// The name of the DB cluster parameter group to modify.
	DBClusterParameterGroupName *string
}

type ModifyDBClusterParameterGroupOutput struct {
	// The name of the DB cluster parameter group. Constraints:
	//
	//     * Must be 1 to 255
	// letters or numbers.
	//
	//     * First character must be a letter
	//
	//     * Cannot end
	// with a hyphen or contain two consecutive hyphens
	//
	// This value is stored as a
	// lowercase string.
	DBClusterParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBClusterParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterParameterGroup",
	}
}