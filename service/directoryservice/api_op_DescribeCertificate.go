// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays information about the certificate registered for a secured LDAP
// connection.
func (c *Client) DescribeCertificate(ctx context.Context, params *DescribeCertificateInput, optFns ...func(*Options)) (*DescribeCertificateOutput, error) {
	stack := middleware.NewStack("DescribeCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeCertificateMiddlewares(stack)
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
	addOpDescribeCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificate(options.Region), middleware.Before)
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
			OperationName: "DescribeCertificate",
			Err:           err,
		}
	}
	out := result.(*DescribeCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCertificateInput struct {

	// The identifier of the certificate.
	//
	// This member is required.
	CertificateId *string

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string
}

type DescribeCertificateOutput struct {

	// Information about the certificate, including registered date time, certificate
	// state, the reason for the state, expiration date time, and certificate common
	// name.
	Certificate *types.Certificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeCertificate",
	}
}
