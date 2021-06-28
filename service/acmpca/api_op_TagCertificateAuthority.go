// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more tags to your private CA. Tags are labels that you can use to
// identify and organize your AWS resources. Each tag consists of a key and an
// optional value. You specify the private CA on input by its Amazon Resource Name
// (ARN). You specify the tag by using a key-value pair. You can apply a tag to
// just one private CA if you want to identify a specific characteristic of that
// CA, or you can apply the same tag to multiple private CAs if you want to filter
// for a common relationship among those CAs. To remove one or more tags, use the
// UntagCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_UntagCertificateAuthority.html)
// action. Call the ListTags
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ListTags.html)
// action to see what tags are associated with your CA.
func (c *Client) TagCertificateAuthority(ctx context.Context, params *TagCertificateAuthorityInput, optFns ...func(*Options)) (*TagCertificateAuthorityOutput, error) {
	if params == nil {
		params = &TagCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagCertificateAuthority", params, optFns, c.addOperationTagCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagCertificateAuthorityInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html).
	// This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// List of tags to be associated with the CA.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type TagCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagCertificateAuthority{}, middleware.After)
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
	if err = addOpTagCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagCertificateAuthority(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "TagCertificateAuthority",
	}
}
