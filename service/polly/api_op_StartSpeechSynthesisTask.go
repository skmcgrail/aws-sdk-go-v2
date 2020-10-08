// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows the creation of an asynchronous synthesis task, by starting a new
// SpeechSynthesisTask. This operation requires all the standard information needed
// for speech synthesis, plus the name of an Amazon S3 bucket for the service to
// store the output of the synthesis task and two optional parameters
// (OutputS3KeyPrefix and SnsTopicArn). Once the synthesis task is created, this
// operation will return a SpeechSynthesisTask object, which will include an
// identifier of this task as well as the current status.
func (c *Client) StartSpeechSynthesisTask(ctx context.Context, params *StartSpeechSynthesisTaskInput, optFns ...func(*Options)) (*StartSpeechSynthesisTaskOutput, error) {
	stack := middleware.NewStack("StartSpeechSynthesisTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartSpeechSynthesisTaskMiddlewares(stack)
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
	addOpStartSpeechSynthesisTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSpeechSynthesisTask(options.Region), middleware.Before)
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
			OperationName: "StartSpeechSynthesisTask",
			Err:           err,
		}
	}
	out := result.(*StartSpeechSynthesisTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSpeechSynthesisTaskInput struct {

	// The format in which the returned output will be encoded. For audio stream, this
	// will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	//
	// This member is required.
	OutputFormat types.OutputFormat

	// Amazon S3 bucket name to which the output file will be saved.
	//
	// This member is required.
	OutputS3BucketName *string

	// The input text to synthesize. If you specify ssml as the TextType, follow the
	// SSML format for the input text.
	//
	// This member is required.
	Text *string

	// Voice ID to use for the synthesis.
	//
	// This member is required.
	VoiceId types.VoiceId

	// Specifies the engine (standard or neural) for Amazon Polly to use when
	// processing input text for speech synthesis. Using a voice that is not supported
	// for the engine selected will result in an error.
	Engine types.Engine

	// Optional language code for the Speech Synthesis request. This is only necessary
	// if using a bilingual voice, such as Aditi, which can be used for either Indian
	// English (en-IN) or Hindi (hi-IN). If a bilingual voice is used and no language
	// code is specified, Amazon Polly will use the default language of the bilingual
	// voice. The default language for any voice is the one returned by the
	// DescribeVoices
	// (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html) operation
	// for the LanguageCode parameter. For example, if no language code is specified,
	// Aditi will use Indian English rather than Hindi.
	LanguageCode types.LanguageCode

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon is
	// the same as the language of the voice.
	LexiconNames []*string

	// The Amazon S3 key prefix for the output speech file.
	OutputS3KeyPrefix *string

	// The audio frequency specified in Hz. The valid values for mp3 and ogg_vorbis are
	// "8000", "16000", "22050", and "24000". The default value for standard voices is
	// "22050". The default value for neural voices is "24000". Valid values for pcm
	// are "8000" and "16000" The default value is "16000".
	SampleRate *string

	// ARN for the SNS topic optionally used for providing status notification for a
	// speech synthesis task.
	SnsTopicArn *string

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []types.SpeechMarkType

	// Specifies whether the input text is plain text or SSML. The default value is
	// plain text.
	TextType types.TextType
}

type StartSpeechSynthesisTaskOutput struct {

	// SynthesisTask object that provides information and attributes about a newly
	// submitted speech synthesis task.
	SynthesisTask *types.SynthesisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartSpeechSynthesisTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartSpeechSynthesisTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSpeechSynthesisTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartSpeechSynthesisTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "StartSpeechSynthesisTask",
	}
}
