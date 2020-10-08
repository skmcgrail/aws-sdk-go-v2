// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes the app.
func (c *Client) DescribeApp(ctx context.Context, params *DescribeAppInput, optFns ...func(*Options)) (*DescribeAppOutput, error) {
	stack := middleware.NewStack("DescribeApp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAppMiddlewares(stack)
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
	addOpDescribeAppValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApp(options.Region), middleware.Before)
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
			OperationName: "DescribeApp",
			Err:           err,
		}
	}
	out := result.(*DescribeAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppInput struct {

	// The name of the app.
	//
	// This member is required.
	AppName *string

	// The type of app.
	//
	// This member is required.
	AppType types.AppType

	// The domain ID.
	//
	// This member is required.
	DomainId *string

	// The user profile name.
	//
	// This member is required.
	UserProfileName *string
}

type DescribeAppOutput struct {

	// The app's Amazon Resource Name (ARN).
	AppArn *string

	// The name of the app.
	AppName *string

	// The type of app.
	AppType types.AppType

	// The creation time.
	CreationTime *time.Time

	// The domain ID.
	DomainId *string

	// The failure reason.
	FailureReason *string

	// The timestamp of the last health check.
	LastHealthCheckTimestamp *time.Time

	// The timestamp of the last user's activity.
	LastUserActivityTimestamp *time.Time

	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image
	// created on the instance.
	ResourceSpec *types.ResourceSpec

	// The status.
	Status types.AppStatus

	// The user profile name.
	UserProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAppMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApp{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeApp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeApp",
	}
}
