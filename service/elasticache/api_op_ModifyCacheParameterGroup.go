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

// Modifies the parameters of a cache parameter group. You can modify up to 20
// parameters in a single request by submitting a list parameter name and value
// pairs.
func (c *Client) ModifyCacheParameterGroup(ctx context.Context, params *ModifyCacheParameterGroupInput, optFns ...func(*Options)) (*ModifyCacheParameterGroupOutput, error) {
	stack := middleware.NewStack("ModifyCacheParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyCacheParameterGroupMiddlewares(stack)
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
	addOpModifyCacheParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCacheParameterGroup(options.Region), middleware.Before)
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
			OperationName: "ModifyCacheParameterGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyCacheParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ModifyCacheParameterGroup operation.
type ModifyCacheParameterGroupInput struct {
	// An array of parameter names and values for the parameter update. You must supply
	// at least one parameter name and value; subsequent arguments are optional. A
	// maximum of 20 parameters may be modified per request.
	ParameterNameValues []*types.ParameterNameValue
	// The name of the cache parameter group to modify.
	CacheParameterGroupName *string
}

// Represents the output of one of the following operations:
//
//     *
// ModifyCacheParameterGroup
//
//     * ResetCacheParameterGroup
type ModifyCacheParameterGroupOutput struct {
	// The name of the cache parameter group.
	CacheParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyCacheParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyCacheParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCacheParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyCacheParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyCacheParameterGroup",
	}
}