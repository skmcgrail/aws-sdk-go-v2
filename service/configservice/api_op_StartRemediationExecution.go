// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs an on-demand remediation for the specified AWS Config rules against the
// last known remediation configuration. It runs an execution against the current
// state of your resources. Remediation execution is asynchronous. You can specify
// up to 100 resource keys per request. An existing StartRemediationExecution call
// for the specified resource keys must complete before you can call the API again.
func (c *Client) StartRemediationExecution(ctx context.Context, params *StartRemediationExecutionInput, optFns ...func(*Options)) (*StartRemediationExecutionOutput, error) {
	stack := middleware.NewStack("StartRemediationExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartRemediationExecutionMiddlewares(stack)
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
	addOpStartRemediationExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartRemediationExecution(options.Region), middleware.Before)
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
			OperationName: "StartRemediationExecution",
			Err:           err,
		}
	}
	out := result.(*StartRemediationExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRemediationExecutionInput struct {

	// The list of names of AWS Config rules that you want to run remediation execution
	// for.
	//
	// This member is required.
	ConfigRuleName *string

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	//
	// This member is required.
	ResourceKeys []*types.ResourceKey
}

type StartRemediationExecutionOutput struct {

	// For resources that have failed to start execution, the API returns a resource
	// key object.
	FailedItems []*types.ResourceKey

	// Returns a failure message. For example, the resource is already compliant.
	FailureMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartRemediationExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartRemediationExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartRemediationExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartRemediationExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "StartRemediationExecution",
	}
}
