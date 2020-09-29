// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new maintenance window. The value you specify for Duration determines
// the specific end time for the maintenance window based on the time it begins. No
// maintenance window tasks are permitted to start after the resulting endtime
// minus the number of hours you specify for Cutoff. For example, if the
// maintenance window starts at 3 PM, the duration is three hours, and the value
// you specify for Cutoff is one hour, no maintenance window tasks can start after
// 5 PM.
func (c *Client) CreateMaintenanceWindow(ctx context.Context, params *CreateMaintenanceWindowInput, optFns ...func(*Options)) (*CreateMaintenanceWindowOutput, error) {
	stack := middleware.NewStack("CreateMaintenanceWindow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateMaintenanceWindowMiddlewares(stack)
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
	addIdempotencyToken_opCreateMaintenanceWindowMiddleware(stack, options)
	addOpCreateMaintenanceWindowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMaintenanceWindow(options.Region), middleware.Before)
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
			OperationName: "CreateMaintenanceWindow",
			Err:           err,
		}
	}
	out := result.(*CreateMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMaintenanceWindowInput struct {
	// An optional description for the maintenance window. We recommend specifying a
	// description to help you organize your maintenance windows.
	Description *string
	// Optional metadata that you assign to a resource. Tags enable you to categorize a
	// resource in different ways, such as by purpose, owner, or environment. For
	// example, you might want to tag a maintenance window to identify the type of
	// tasks it will run, the types of targets, and the environment it will run in. In
	// this case, you could specify the following key name/value pairs:
	//
	//     *
	// Key=TaskType,Value=AgentUpdate
	//
	//     * Key=OS,Value=Windows
	//
	//     *
	// Key=Environment,Value=Production
	//
	// To add tags to an existing maintenance window,
	// use the AddTagsToResource () action.
	Tags []*types.Tag
	// The time zone that the scheduled maintenance window executions are based on, in
	// Internet Assigned Numbers Authority (IANA) format. For example:
	// "America/Los_Angeles", "etc/UTC", or "Asia/Seoul". For more information, see the
	// Time Zone Database (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string
	// User-provided idempotency token.
	ClientToken *string
	// The date and time, in ISO-8601 Extended format, for when you want the
	// maintenance window to become active. StartDate allows you to delay activation of
	// the maintenance window until the specified future date.
	StartDate *string
	// The duration of the maintenance window in hours.
	Duration *int32
	// The name of the maintenance window.
	Name *string
	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string
	// The date and time, in ISO-8601 Extended format, for when you want the
	// maintenance window to become inactive. EndDate allows you to set a date and time
	// in the future when the maintenance window will no longer run.
	EndDate *string
	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	Cutoff *int32
	// The number of days to wait after the date and time specified by a CRON
	// expression before running the maintenance window. For example, the following
	// cron expression schedules a maintenance window to run on the third Tuesday of
	// every month at 11:30 PM. cron(0 30 23 ? * TUE#3 *) If the schedule offset is 2,
	// the maintenance window won't run until two days later.
	ScheduleOffset *int32
	// Enables a maintenance window task to run on managed instances, even if you have
	// not registered those instances as targets. If enabled, then you must specify the
	// unregistered instances (by instance ID) when you register a task with the
	// maintenance window. If you don't enable this option, then you must specify
	// previously-registered targets when you register a task with the maintenance
	// window.
	AllowUnassociatedTargets *bool
}

type CreateMaintenanceWindowOutput struct {
	// The ID of the created maintenance window.
	WindowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateMaintenanceWindowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMaintenanceWindow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMaintenanceWindow{}, middleware.After)
}

type idempotencyToken_initializeOpCreateMaintenanceWindow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMaintenanceWindow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMaintenanceWindow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMaintenanceWindowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMaintenanceWindowInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMaintenanceWindowMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateMaintenanceWindow{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMaintenanceWindow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreateMaintenanceWindow",
	}
}