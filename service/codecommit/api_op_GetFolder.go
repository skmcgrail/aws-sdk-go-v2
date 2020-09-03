// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the contents of a specified folder in a repository.
func (c *Client) GetFolder(ctx context.Context, params *GetFolderInput, optFns ...func(*Options)) (*GetFolderOutput, error) {
	stack := middleware.NewStack("GetFolder", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetFolderMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetFolderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFolder(options.Region), middleware.Before)

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
			OperationName: "GetFolder",
			Err:           err,
		}
	}
	out := result.(*GetFolderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFolderInput struct {
	// The fully qualified path to the folder whose contents are returned, including
	// the folder name. For example, /examples is a fully-qualified path to a folder
	// named examples that was created off of the root directory (/) of a repository.
	FolderPath *string
	// The name of the repository.
	RepositoryName *string
	// A fully qualified reference used to identify a commit that contains the version
	// of the folder's content to return. A fully qualified reference can be a commit
	// ID, branch name, tag, or reference such as HEAD. If no specifier is provided,
	// the folder content is returned as it exists in the HEAD commit.
	CommitSpecifier *string
}

type GetFolderOutput struct {
	// The full SHA-1 pointer of the tree information for the commit that contains the
	// folder.
	TreeId *string
	// The fully qualified path of the folder whose contents are returned.
	FolderPath *string
	// The list of files in the specified folder, if any.
	Files []*types.File
	// The full commit ID used as a reference for the returned version of the folder
	// content.
	CommitId *string
	// The list of submodules in the specified folder, if any.
	SubModules []*types.SubModule
	// The list of symbolic links to other files and folders in the specified folder,
	// if any.
	SymbolicLinks []*types.SymbolicLink
	// The list of folders that exist under the specified folder, if any.
	SubFolders []*types.Folder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetFolderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetFolder{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFolder{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFolder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetFolder",
	}
}