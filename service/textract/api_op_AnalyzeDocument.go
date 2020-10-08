// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Analyzes an input document for relationships between detected items. The types
// of information returned are as follows:
//
//     * Form data (key-value pairs). The
// related information is returned in two Block () objects, each of type
// KEY_VALUE_SET: a KEY Block object and a VALUE Block object. For example, Name:
// Ana Silva Carolina contains a key and value. Name: is the key. Ana Silva
// Carolina is the value.
//
//     * Table and table cell data. A TABLE Block object
// contains information about a detected table. A CELL Block object is returned for
// each cell in a table.
//
//     * Lines and words of text. A LINE Block object
// contains one or more WORD Block objects. All lines and words that are detected
// in the document are returned (including text that doesn't have a relationship
// with the value of FeatureTypes).
//
//     <p>Selection elements such as check boxes
// and option buttons (radio buttons) can be detected in form data and in tables. A
// SELECTION_ELEMENT <code>Block</code> object contains information about a
// selection element, including the selection status.</p> <p>You can choose which
// type of analysis to perform by specifying the <code>FeatureTypes</code> list.
// </p> <p>The output is returned in a list of <code>Block</code> objects.</p> <p>
// <code>AnalyzeDocument</code> is a synchronous operation. To analyze documents
// asynchronously, use <a>StartDocumentAnalysis</a>.</p> <p>For more information,
// see <a
// href="https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html">Document
// Text Analysis</a>.</p>
func (c *Client) AnalyzeDocument(ctx context.Context, params *AnalyzeDocumentInput, optFns ...func(*Options)) (*AnalyzeDocumentOutput, error) {
	stack := middleware.NewStack("AnalyzeDocument", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAnalyzeDocumentMiddlewares(stack)
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
	addOpAnalyzeDocumentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAnalyzeDocument(options.Region), middleware.Before)
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
			OperationName: "AnalyzeDocument",
			Err:           err,
		}
	}
	out := result.(*AnalyzeDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AnalyzeDocumentInput struct {

	// The input document as base64-encoded bytes or an Amazon S3 object. If you use
	// the AWS CLI to call Amazon Textract operations, you can't pass image bytes. The
	// document must be an image in JPEG or PNG format. If you're using an AWS SDK to
	// call Amazon Textract, you might not need to base64-encode image bytes that are
	// passed using the Bytes field.
	//
	// This member is required.
	Document *types.Document

	// A list of the types of analysis to perform. Add TABLES to the list to return
	// information about the tables that are detected in the input document. Add FORMS
	// to return detected form data. To perform both types of analysis, add TABLES and
	// FORMS to FeatureTypes. All lines and words detected in the document are included
	// in the response (including text that isn't related to the value of
	// FeatureTypes).
	//
	// This member is required.
	FeatureTypes []types.FeatureType

	// Sets the configuration for the human in the loop workflow for analyzing
	// documents.
	HumanLoopConfig *types.HumanLoopConfig
}

type AnalyzeDocumentOutput struct {

	// The version of the model used to analyze the document.
	AnalyzeDocumentModelVersion *string

	// The items that are detected and analyzed by AnalyzeDocument.
	Blocks []*types.Block

	// Metadata about the analyzed document. An example is the number of pages.
	DocumentMetadata *types.DocumentMetadata

	// Shows the results of the human in the loop evaluation.
	HumanLoopActivationOutput *types.HumanLoopActivationOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAnalyzeDocumentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAnalyzeDocument{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAnalyzeDocument{}, middleware.After)
}

func newServiceMetadataMiddleware_opAnalyzeDocument(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "AnalyzeDocument",
	}
}
