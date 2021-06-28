// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation configures Amazon Route 53 to automatically renew the specified
// domain before the domain registration expires. The cost of renewing your domain
// registration is billed to your AWS account. The period during which you can
// renew a domain name varies by TLD. For a list of TLDs and their renewal
// policies, see Domains That You Can Register with Amazon Route 53
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
// in the Amazon Route 53 Developer Guide. Route 53 requires that you renew before
// the end of the renewal period so we can complete processing before the deadline.
func (c *Client) EnableDomainAutoRenew(ctx context.Context, params *EnableDomainAutoRenewInput, optFns ...func(*Options)) (*EnableDomainAutoRenewOutput, error) {
	if params == nil {
		params = &EnableDomainAutoRenewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableDomainAutoRenew", params, optFns, c.addOperationEnableDomainAutoRenewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableDomainAutoRenewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableDomainAutoRenewInput struct {

	// The name of the domain that you want to enable automatic renewal for.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

type EnableDomainAutoRenewOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableDomainAutoRenewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableDomainAutoRenew{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableDomainAutoRenew{}, middleware.After)
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
	if err = addOpEnableDomainAutoRenewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableDomainAutoRenew(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableDomainAutoRenew(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "EnableDomainAutoRenew",
	}
}
