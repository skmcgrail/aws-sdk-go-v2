// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the properties associated with an asycnhronous batch translation job
// including name, ID, status, source and target languages, input/output S3
// buckets, and so on.
func (c *Client) DescribeTextTranslationJob(ctx context.Context, params *DescribeTextTranslationJobInput, optFns ...func(*Options)) (*DescribeTextTranslationJobOutput, error) {
	stack := middleware.NewStack("DescribeTextTranslationJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTextTranslationJobMiddlewares(stack)
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
	addOpDescribeTextTranslationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTextTranslationJob(options.Region), middleware.Before)
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
			OperationName: "DescribeTextTranslationJob",
			Err:           err,
		}
	}
	out := result.(*DescribeTextTranslationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTextTranslationJobInput struct {

	// The identifier that Amazon Translate generated for the job. The
	// StartTextTranslationJob () operation returns this identifier in its response.
	//
	// This member is required.
	JobId *string
}

type DescribeTextTranslationJobOutput struct {

	// An object that contains the properties associated with an asynchronous batch
	// translation job.
	TextTranslationJobProperties *types.TextTranslationJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTextTranslationJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTextTranslationJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTextTranslationJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTextTranslationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "DescribeTextTranslationJob",
	}
}
