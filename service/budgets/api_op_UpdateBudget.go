// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a budget. You can change every part of a budget except for the
// budgetName and the calculatedSpend. When you modify a budget, the
// calculatedSpend drops to zero until AWS has new usage data to use for
// forecasting. Only one of BudgetLimit or PlannedBudgetLimits can be present in
// the syntax at one time. Use the syntax that matches your case. The Request
// Syntax section shows the BudgetLimit syntax. For PlannedBudgetLimits, see the
// Examples
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_UpdateBudget.html#API_UpdateBudget_Examples)
// section.
func (c *Client) UpdateBudget(ctx context.Context, params *UpdateBudgetInput, optFns ...func(*Options)) (*UpdateBudgetOutput, error) {
	stack := middleware.NewStack("UpdateBudget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateBudgetMiddlewares(stack)
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
	addOpUpdateBudgetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBudget(options.Region), middleware.Before)
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
			OperationName: "UpdateBudget",
			Err:           err,
		}
	}
	out := result.(*UpdateBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of UpdateBudget
type UpdateBudgetInput struct {

	// The accountId that is associated with the budget that you want to update.
	//
	// This member is required.
	AccountId *string

	// The budget that you want to update your budget to.
	//
	// This member is required.
	NewBudget *types.Budget
}

// Response of UpdateBudget
type UpdateBudgetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateBudgetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBudget{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBudget{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateBudget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "UpdateBudget",
	}
}
