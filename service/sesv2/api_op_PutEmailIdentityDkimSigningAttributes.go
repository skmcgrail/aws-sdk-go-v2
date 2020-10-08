// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used to configure or change the DKIM authentication settings for an email domain
// identity. You can use this operation to do any of the following:
//
//     * Update
// the signing attributes for an identity that uses Bring Your Own DKIM
// (BYODKIM).
//
//     * Change from using no DKIM authentication to using Easy DKIM.
//
//
// * Change from using no DKIM authentication to using BYODKIM.
//
//     * Change from
// using Easy DKIM to using BYODKIM.
//
//     * Change from using BYODKIM to using Easy
// DKIM.
func (c *Client) PutEmailIdentityDkimSigningAttributes(ctx context.Context, params *PutEmailIdentityDkimSigningAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityDkimSigningAttributesOutput, error) {
	stack := middleware.NewStack("PutEmailIdentityDkimSigningAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutEmailIdentityDkimSigningAttributesMiddlewares(stack)
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
	addOpPutEmailIdentityDkimSigningAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailIdentityDkimSigningAttributes(options.Region), middleware.Before)
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
			OperationName: "PutEmailIdentityDkimSigningAttributes",
			Err:           err,
		}
	}
	out := result.(*PutEmailIdentityDkimSigningAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change the DKIM attributes for an email identity.
type PutEmailIdentityDkimSigningAttributesInput struct {

	// The email identity that you want to configure DKIM for.
	//
	// This member is required.
	EmailIdentity *string

	// The method that you want to use to configure DKIM for the identity. There are
	// two possible values:
	//
	//     * AWS_SES – Configure DKIM for the identity by using
	// Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
	//
	//     *
	// EXTERNAL – Configure DKIM for the identity by using Bring Your Own DKIM
	// (BYODKIM).
	//
	// This member is required.
	SigningAttributesOrigin types.DkimSigningAttributesOrigin

	// An object that contains information about the private key and selector that you
	// want to use to configure DKIM for the identity. This object is only required if
	// you want to configure Bring Your Own DKIM (BYODKIM) for the identity.
	SigningAttributes *types.DkimSigningAttributes
}

// If the action is successful, the service sends back an HTTP 200 response. The
// following data is returned in JSON format by the service.
type PutEmailIdentityDkimSigningAttributesOutput struct {

	// The DKIM authentication status of the identity. Amazon SES determines the
	// authentication status by searching for specific records in the DNS configuration
	// for your domain. If you used Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) to set up
	// DKIM authentication, Amazon SES tries to find three unique CNAME records in the
	// DNS configuration for your domain. If you provided a public key to perform DKIM
	// authentication, Amazon SES tries to find a TXT record that uses the selector
	// that you specified. The value of the TXT record must be a public key that's
	// paired with the private key that you specified in the process of creating the
	// identity. The status can be one of the following:
	//
	//     * PENDING – The
	// verification process was initiated, but Amazon SES hasn't yet detected the DKIM
	// records in the DNS configuration for the domain.
	//
	//     * SUCCESS – The
	// verification process completed successfully.
	//
	//     * FAILED – The verification
	// process failed. This typically occurs when Amazon SES fails to find the DKIM
	// records in the DNS configuration of the domain.
	//
	//     * TEMPORARY_FAILURE – A
	// temporary issue is preventing Amazon SES from determining the DKIM
	// authentication status of the domain.
	//
	//     * NOT_STARTED – The DKIM verification
	// process hasn't been initiated for the domain.
	DkimStatus types.DkimStatus

	// If you used Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) to
	// configure DKIM authentication for the domain, then this object contains a set of
	// unique strings that you use to create a set of CNAME records that you add to the
	// DNS configuration for your domain. When Amazon SES detects these records in the
	// DNS configuration for your domain, the DKIM authentication process is complete.
	// If you configured DKIM authentication for the domain by providing your own
	// public-private key pair, then this object contains the selector that's
	// associated with your public key. Regardless of the DKIM authentication method
	// you use, Amazon SES searches for the appropriate records in the DNS
	// configuration of the domain for up to 72 hours.
	DkimTokens []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutEmailIdentityDkimSigningAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutEmailIdentityDkimSigningAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEmailIdentityDkimSigningAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutEmailIdentityDkimSigningAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutEmailIdentityDkimSigningAttributes",
	}
}
