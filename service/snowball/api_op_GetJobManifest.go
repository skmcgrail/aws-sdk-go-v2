// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a link to an Amazon S3 presigned URL for the manifest file associated
// with the specified JobId value. You can access the manifest file for up to 60
// minutes after this request has been made. To access the manifest file after 60
// minutes have passed, you'll have to make another call to the GetJobManifest
// action.  <p>The manifest is an encrypted file that you can download after your
// job enters the <code>WithCustomer</code> status. The manifest is decrypted by
// using the <code>UnlockCode</code> code value, when you pass both values to the
// Snowball through the Snowball client when the client is started for the first
// time.</p> <p>As a best practice, we recommend that you don't save a copy of an
// <code>UnlockCode</code> value in the same location as the manifest file for that
// job. Saving these separately helps prevent unauthorized parties from gaining
// access to the Snowball associated with that job.</p> <p>The credentials of a
// given job, including its manifest file and unlock code, expire 90 days after the
// job is created.</p>
func (c *Client) GetJobManifest(ctx context.Context, params *GetJobManifestInput, optFns ...func(*Options)) (*GetJobManifestOutput, error) {
	stack := middleware.NewStack("GetJobManifest", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetJobManifestMiddlewares(stack)
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
	addOpGetJobManifestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobManifest(options.Region), middleware.Before)
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
			OperationName: "GetJobManifest",
			Err:           err,
		}
	}
	out := result.(*GetJobManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobManifestInput struct {
	// The ID for a job that you want to get the manifest file for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string
}

type GetJobManifestOutput struct {
	// The Amazon S3 presigned URL for the manifest file associated with the specified
	// JobId value.
	ManifestURI *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetJobManifestMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetJobManifest{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetJobManifest{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetJobManifest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "GetJobManifest",
	}
}