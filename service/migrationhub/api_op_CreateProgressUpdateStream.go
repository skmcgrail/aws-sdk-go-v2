// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a progress update stream which is an AWS resource used for access
// control as well as a namespace for migration task names that is implicitly
// linked to your AWS account. It must uniquely identify the migration tool as it
// is used for all updates made by the tool; however, it does not need to be unique
// for each AWS account because it is scoped to the AWS account.
func (c *Client) CreateProgressUpdateStream(ctx context.Context, params *CreateProgressUpdateStreamInput, optFns ...func(*Options)) (*CreateProgressUpdateStreamOutput, error) {
	if params == nil {
		params = &CreateProgressUpdateStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProgressUpdateStream", params, optFns, c.addOperationCreateProgressUpdateStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProgressUpdateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProgressUpdateStreamInput struct {

	// The name of the ProgressUpdateStream. Do not store personal data in this field.
	//
	// This member is required.
	ProgressUpdateStreamName *string

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun bool

	noSmithyDocumentSerde
}

type CreateProgressUpdateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProgressUpdateStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProgressUpdateStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProgressUpdateStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateProgressUpdateStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProgressUpdateStream(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateProgressUpdateStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "CreateProgressUpdateStream",
	}
}
