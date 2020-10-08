// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of vocabularies that match the specified criteria. You get the
// entire list of vocabularies if you don't enter a value in any of the request
// parameters.
func (c *Client) ListMedicalVocabularies(ctx context.Context, params *ListMedicalVocabulariesInput, optFns ...func(*Options)) (*ListMedicalVocabulariesOutput, error) {
	stack := middleware.NewStack("ListMedicalVocabularies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListMedicalVocabulariesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMedicalVocabularies(options.Region), middleware.Before)
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
			OperationName: "ListMedicalVocabularies",
			Err:           err,
		}
	}
	out := result.(*ListMedicalVocabulariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMedicalVocabulariesInput struct {

	// The maximum number of vocabularies to return in the response.
	MaxResults *int32

	// Returns vocabularies in the list whose name contains the specified string. The
	// search is case-insensitive, ListMedicalVocabularies returns both
	// "vocabularyname" and "VocabularyName" in the response list.
	NameContains *string

	// If the result of your previous request to ListMedicalVocabularies was truncated,
	// include the NextToken to fetch the next set of jobs.
	NextToken *string

	// When specified, only returns vocabularies with the VocabularyState equal to the
	// specified vocabulary state.
	StateEquals types.VocabularyState
}

type ListMedicalVocabulariesOutput struct {

	// The ListMedicalVocabularies operation returns a page of vocabularies at a time.
	// The maximum size of the page is set by the MaxResults parameter. If there are
	// more jobs in the list than the page size, Amazon Transcribe Medical returns the
	// NextPage token. Include the token in the next request to the
	// ListMedicalVocabularies operation to return the next page of jobs.
	NextToken *string

	// The requested vocabulary state.
	Status types.VocabularyState

	// A list of objects that describe the vocabularies that match the search criteria
	// in the request.
	Vocabularies []*types.VocabularyInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListMedicalVocabulariesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListMedicalVocabularies{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMedicalVocabularies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMedicalVocabularies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListMedicalVocabularies",
	}
}
