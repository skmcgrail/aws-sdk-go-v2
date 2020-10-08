// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specified cluster parameter group. The cluster parameter group to be
// deleted can't be associated with any clusters.
func (c *Client) DeleteDBClusterParameterGroup(ctx context.Context, params *DeleteDBClusterParameterGroupInput, optFns ...func(*Options)) (*DeleteDBClusterParameterGroupOutput, error) {
	stack := middleware.NewStack("DeleteDBClusterParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteDBClusterParameterGroupMiddlewares(stack)
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
	addOpDeleteDBClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBClusterParameterGroup(options.Region), middleware.Before)
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
			OperationName: "DeleteDBClusterParameterGroup",
			Err:           err,
		}
	}
	out := result.(*DeleteDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DeleteDBClusterParameterGroup ().
type DeleteDBClusterParameterGroupInput struct {

	// The name of the cluster parameter group. Constraints:
	//
	//     * Must be the name of
	// an existing cluster parameter group.
	//
	//     * You can't delete a default cluster
	// parameter group.
	//
	//     * Cannot be associated with any clusters.
	//
	// This member is required.
	DBClusterParameterGroupName *string
}

type DeleteDBClusterParameterGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteDBClusterParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBClusterParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBClusterParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDBClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBClusterParameterGroup",
	}
}
