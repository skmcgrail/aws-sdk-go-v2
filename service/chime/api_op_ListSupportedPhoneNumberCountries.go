// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists supported phone number countries.
func (c *Client) ListSupportedPhoneNumberCountries(ctx context.Context, params *ListSupportedPhoneNumberCountriesInput, optFns ...func(*Options)) (*ListSupportedPhoneNumberCountriesOutput, error) {
	if params == nil {
		params = &ListSupportedPhoneNumberCountriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSupportedPhoneNumberCountries", params, optFns, c.addOperationListSupportedPhoneNumberCountriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSupportedPhoneNumberCountriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSupportedPhoneNumberCountriesInput struct {

	// The phone number product type.
	//
	// This member is required.
	ProductType types.PhoneNumberProductType

	noSmithyDocumentSerde
}

type ListSupportedPhoneNumberCountriesOutput struct {

	// The supported phone number countries.
	PhoneNumberCountries []types.PhoneNumberCountry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSupportedPhoneNumberCountriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSupportedPhoneNumberCountries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSupportedPhoneNumberCountries{}, middleware.After)
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
	if err = addOpListSupportedPhoneNumberCountriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSupportedPhoneNumberCountries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSupportedPhoneNumberCountries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListSupportedPhoneNumberCountries",
	}
}
