// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Trains or retrains an active solution. A solution is created using the
// CreateSolution () operation and must be in the ACTIVE state before calling
// CreateSolutionVersion. A new version of the solution is created every time you
// call this operation. Status A solution version can be in one of the following
// states:
//
//     * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE
// FAILED
//
// To get the status of the version, call DescribeSolutionVersion (). Wait
// until the status shows as ACTIVE before calling CreateCampaign. If the status
// shows as CREATE FAILED, the response includes a failureReason key, which
// describes why the job failed. Related APIs
//
//     * ListSolutionVersions ()
//
//     *
// DescribeSolutionVersion ()
//
//     * ListSolutions ()
//
//     * CreateSolution ()
//
//
// * DescribeSolution ()
//
//     * DeleteSolution ()
func (c *Client) CreateSolutionVersion(ctx context.Context, params *CreateSolutionVersionInput, optFns ...func(*Options)) (*CreateSolutionVersionOutput, error) {
	stack := middleware.NewStack("CreateSolutionVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateSolutionVersionMiddlewares(stack)
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
	addOpCreateSolutionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolutionVersion(options.Region), middleware.Before)
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
			OperationName: "CreateSolutionVersion",
			Err:           err,
		}
	}
	out := result.(*CreateSolutionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolutionVersionInput struct {

	// The Amazon Resource Name (ARN) of the solution containing the training
	// configuration information.
	//
	// This member is required.
	SolutionArn *string

	// The scope of training to be performed when creating the solution version. The
	// FULL option trains the solution version based on the entirety of the input
	// solution's training data, while the UPDATE option processes only the data that
	// has changed in comparison to the input solution. Choose UPDATE when you want to
	// incrementally update your solution version instead of creating an entirely new
	// one. The UPDATE option can only be used when you already have an active solution
	// version created from the input solution using the FULL option and the input
	// solution was trained with the native-recipe-hrnn-coldstart () recipe.
	TrainingMode types.TrainingMode
}

type CreateSolutionVersionOutput struct {

	// The ARN of the new solution version.
	SolutionVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateSolutionVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSolutionVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSolutionVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSolutionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateSolutionVersion",
	}
}
