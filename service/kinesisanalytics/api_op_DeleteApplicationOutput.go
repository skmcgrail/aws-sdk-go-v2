// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation (). Deletes output destination configuration from
// your application configuration. Amazon Kinesis Analytics will no longer write
// data from the corresponding in-application stream to the external output
// destination. This operation requires permissions to perform the
// kinesisanalytics:DeleteApplicationOutput action.
func (c *Client) DeleteApplicationOutput(ctx context.Context, params *DeleteApplicationOutputInput, optFns ...func(*Options)) (*DeleteApplicationOutputOutput, error) {
	stack := middleware.NewStack("DeleteApplicationOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteApplicationOutputMiddlewares(stack)
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
	addOpDeleteApplicationOutputValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationOutput(options.Region), middleware.Before)

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
			OperationName: "DeleteApplicationOutput",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteApplicationOutputInput struct {
	// Amazon Kinesis Analytics application version. You can use the
	// DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified is
	// not the current version, the ConcurrentModificationException is returned.
	CurrentApplicationVersionId *int64
	// Amazon Kinesis Analytics application name.
	ApplicationName *string
	// The ID of the configuration to delete. Each output configuration that is added
	// to the application, either when the application is created or later using the
	// AddApplicationOutput
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationOutput.html)
	// operation, has a unique ID. You need to provide the ID to uniquely identify the
	// output configuration that you want to delete from the application configuration.
	// You can use the DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the specific OutputId.
	OutputId *string
}

//
type DeleteApplicationOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteApplicationOutputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationOutput{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationOutput{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationOutput",
	}
}