// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a request to compile the specified type of information of the deployed
// environment. Setting the InfoType to tail compiles the last lines from the
// application server log files of every Amazon EC2 instance in your environment.
// Setting the InfoType to bundle compresses the application server log files for
// every Amazon EC2 instance into a .zip file. Legacy and .NET containers do not
// support bundle logs. Use RetrieveEnvironmentInfo to obtain the set of logs.
// Related Topics
//
// * RetrieveEnvironmentInfo
func (c *Client) RequestEnvironmentInfo(ctx context.Context, params *RequestEnvironmentInfoInput, optFns ...func(*Options)) (*RequestEnvironmentInfoOutput, error) {
	if params == nil {
		params = &RequestEnvironmentInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestEnvironmentInfo", params, optFns, c.addOperationRequestEnvironmentInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestEnvironmentInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to retrieve logs from an environment and store them in your Elastic
// Beanstalk storage bucket.
type RequestEnvironmentInfoInput struct {

	// The type of information to request.
	//
	// This member is required.
	InfoType types.EnvironmentInfoType

	// The ID of the environment of the requested data. If no such environment is
	// found, RequestEnvironmentInfo returns an InvalidParameterValue error. Condition:
	// You must specify either this or an EnvironmentName, or both. If you do not
	// specify either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentId *string

	// The name of the environment of the requested data. If no such environment is
	// found, RequestEnvironmentInfo returns an InvalidParameterValue error. Condition:
	// You must specify either this or an EnvironmentId, or both. If you do not specify
	// either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentName *string

	noSmithyDocumentSerde
}

type RequestEnvironmentInfoOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRequestEnvironmentInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRequestEnvironmentInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRequestEnvironmentInfo{}, middleware.After)
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
	if err = addOpRequestEnvironmentInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRequestEnvironmentInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRequestEnvironmentInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "RequestEnvironmentInfo",
	}
}
