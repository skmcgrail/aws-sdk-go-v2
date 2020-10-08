// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Translates input text from the source language to the target language. For a
// list of available languages and language codes, see what-is-languages ().
func (c *Client) TranslateText(ctx context.Context, params *TranslateTextInput, optFns ...func(*Options)) (*TranslateTextOutput, error) {
	stack := middleware.NewStack("TranslateText", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTranslateTextMiddlewares(stack)
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
	addOpTranslateTextValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTranslateText(options.Region), middleware.Before)
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
			OperationName: "TranslateText",
			Err:           err,
		}
	}
	out := result.(*TranslateTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TranslateTextInput struct {

	// The language code for the language of the source text. The language must be a
	// language supported by Amazon Translate. For a list of language codes, see
	// what-is-languages (). To have Amazon Translate determine the source language of
	// your text, you can specify auto in the SourceLanguageCode field. If you specify
	// auto, Amazon Translate will call Amazon Comprehend
	// (https://docs.aws.amazon.com/comprehend/latest/dg/comprehend-general.html) to
	// determine the source language.
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code requested for the language of the target text. The language
	// must be a language supported by Amazon Translate.
	//
	// This member is required.
	TargetLanguageCode *string

	// The text to translate. The text string can be a maximum of 5,000 bytes long.
	// Depending on your character set, this may be fewer than 5,000 characters.
	//
	// This member is required.
	Text *string

	// The name of the terminology list file to be used in the TranslateText request.
	// You can use 1 terminology list at most in a TranslateText request. Terminology
	// lists can contain a maximum of 256 terms.
	TerminologyNames []*string
}

type TranslateTextOutput struct {

	// The language code for the language of the source text.
	//
	// This member is required.
	SourceLanguageCode *string

	// The language code for the language of the target text.
	//
	// This member is required.
	TargetLanguageCode *string

	// The translated text.
	//
	// This member is required.
	TranslatedText *string

	// The names of the custom terminologies applied to the input text by Amazon
	// Translate for the translated text response.
	AppliedTerminologies []*types.AppliedTerminology

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTranslateTextMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTranslateText{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTranslateText{}, middleware.After)
}

func newServiceMetadataMiddleware_opTranslateText(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "TranslateText",
	}
}
