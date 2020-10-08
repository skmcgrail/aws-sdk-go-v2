// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more applications. The maximum number of
// applications that can be returned is 100.
func (c *Client) BatchGetApplications(ctx context.Context, params *BatchGetApplicationsInput, optFns ...func(*Options)) (*BatchGetApplicationsOutput, error) {
	stack := middleware.NewStack("BatchGetApplications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetApplicationsMiddlewares(stack)
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
	addOpBatchGetApplicationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetApplications(options.Region), middleware.Before)
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
			OperationName: "BatchGetApplications",
			Err:           err,
		}
	}
	out := result.(*BatchGetApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetApplications operation.
type BatchGetApplicationsInput struct {

	// A list of application names separated by spaces. The maximum number of
	// application names you can specify is 100.
	//
	// This member is required.
	ApplicationNames []*string
}

// Represents the output of a BatchGetApplications operation.
type BatchGetApplicationsOutput struct {

	// Information about the applications.
	ApplicationsInfo []*types.ApplicationInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetApplicationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetApplications{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetApplications{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetApplications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "BatchGetApplications",
	}
}
