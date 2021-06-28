// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// A remediation exception is when a specific resource is no longer considered for
// auto-remediation. This API adds a new exception or updates an existing exception
// for a specific resource with a specific AWS Config rule. AWS Config generates a
// remediation exception when a problem occurs executing a remediation action to a
// specific resource. Remediation exceptions blocks auto-remediation until the
// exception is cleared.
func (c *Client) PutRemediationExceptions(ctx context.Context, params *PutRemediationExceptionsInput, optFns ...func(*Options)) (*PutRemediationExceptionsOutput, error) {
	if params == nil {
		params = &PutRemediationExceptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRemediationExceptions", params, optFns, c.addOperationPutRemediationExceptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRemediationExceptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRemediationExceptionsInput struct {

	// The name of the AWS Config rule for which you want to create remediation
	// exception.
	//
	// This member is required.
	ConfigRuleName *string

	// An exception list of resource exception keys to be processed with the current
	// request. AWS Config adds exception for each resource key. For example, AWS
	// Config adds 3 exceptions for 3 resource keys.
	//
	// This member is required.
	ResourceKeys []types.RemediationExceptionResourceKey

	// The exception is automatically deleted after the expiration date.
	ExpirationTime *time.Time

	// The message contains an explanation of the exception.
	Message *string

	noSmithyDocumentSerde
}

type PutRemediationExceptionsOutput struct {

	// Returns a list of failed remediation exceptions batch objects. Each object in
	// the batch consists of a list of failed items and failure messages.
	FailedBatches []types.FailedRemediationExceptionBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRemediationExceptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRemediationExceptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRemediationExceptions{}, middleware.After)
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
	if err = addOpPutRemediationExceptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRemediationExceptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRemediationExceptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutRemediationExceptions",
	}
}
