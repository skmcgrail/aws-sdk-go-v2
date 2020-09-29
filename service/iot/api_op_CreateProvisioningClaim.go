// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a provisioning claim.
func (c *Client) CreateProvisioningClaim(ctx context.Context, params *CreateProvisioningClaimInput, optFns ...func(*Options)) (*CreateProvisioningClaimOutput, error) {
	stack := middleware.NewStack("CreateProvisioningClaim", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateProvisioningClaimMiddlewares(stack)
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
	addOpCreateProvisioningClaimValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisioningClaim(options.Region), middleware.Before)
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
			OperationName: "CreateProvisioningClaim",
			Err:           err,
		}
	}
	out := result.(*CreateProvisioningClaimOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisioningClaimInput struct {
	// The name of the provisioning template to use.
	TemplateName *string
}

type CreateProvisioningClaimOutput struct {
	// The provisioning claim certificate.
	CertificatePem *string
	// The provisioning claim expiration time.
	Expiration *time.Time
	// The ID of the certificate.
	CertificateId *string
	// The provisioning claim key pair.
	KeyPair *types.KeyPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateProvisioningClaimMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateProvisioningClaim{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProvisioningClaim{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateProvisioningClaim(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateProvisioningClaim",
	}
}