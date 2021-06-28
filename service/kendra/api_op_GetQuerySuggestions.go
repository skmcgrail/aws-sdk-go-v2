// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Fetches the queries that are suggested to your users.
func (c *Client) GetQuerySuggestions(ctx context.Context, params *GetQuerySuggestionsInput, optFns ...func(*Options)) (*GetQuerySuggestionsOutput, error) {
	if params == nil {
		params = &GetQuerySuggestionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQuerySuggestions", params, optFns, c.addOperationGetQuerySuggestionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQuerySuggestionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQuerySuggestionsInput struct {

	// The identifier of the index you want to get query suggestions from.
	//
	// This member is required.
	IndexId *string

	// The text of a user's query to generate query suggestions. A query is suggested
	// if the query prefix matches what a user starts to type as their query. Amazon
	// Kendra does not show any suggestions if a user types fewer than two characters
	// or more than 60 characters. A query must also have at least one search result
	// and contain at least one word of more than four characters.
	//
	// This member is required.
	QueryText *string

	// The maximum number of query suggestions you want to show to your users.
	MaxSuggestionsCount *int32

	noSmithyDocumentSerde
}

type GetQuerySuggestionsOutput struct {

	// The unique identifier for a list of query suggestions for an index.
	QuerySuggestionsId *string

	// A list of query suggestions for an index.
	Suggestions []types.Suggestion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQuerySuggestionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetQuerySuggestions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQuerySuggestions{}, middleware.After)
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
	if err = addOpGetQuerySuggestionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQuerySuggestions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetQuerySuggestions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "GetQuerySuggestions",
	}
}
