// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift/internal"
)

// `ScheduledDeleteTask` represents a scheduling configuration for a Stack. It will trigger a stack deletion task at the given timestamp.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := spacelift.GetScheduledDeleteStack(ctx, &spacelift.GetScheduledDeleteStackArgs{
//				ScheduledDeleteStackId: "$STACK_ID/$SCHEDULED_DELETE_STACK_ID",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetScheduledDeleteStack(ctx *pulumi.Context, args *GetScheduledDeleteStackArgs, opts ...pulumi.InvokeOption) (*GetScheduledDeleteStackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetScheduledDeleteStackResult
	err := ctx.Invoke("spacelift:index/getScheduledDeleteStack:getScheduledDeleteStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScheduledDeleteStack.
type GetScheduledDeleteStackArgs struct {
	// ID of the scheduled delete*stack (stack*id/schedule_id)
	ScheduledDeleteStackId string `pulumi:"scheduledDeleteStackId"`
}

// A collection of values returned by getScheduledDeleteStack.
type GetScheduledDeleteStackResult struct {
	// Timestamp (unix timestamp) at which time the scheduling should happen.
	At int `pulumi:"at"`
	// Indicates whether the resources of the stack should be deleted.
	DeleteResources bool `pulumi:"deleteResources"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the schedule
	ScheduleId string `pulumi:"scheduleId"`
	// ID of the scheduled delete*stack (stack*id/schedule_id)
	ScheduledDeleteStackId string `pulumi:"scheduledDeleteStackId"`
	// Stack ID of the scheduling config
	StackId string `pulumi:"stackId"`
}

func GetScheduledDeleteStackOutput(ctx *pulumi.Context, args GetScheduledDeleteStackOutputArgs, opts ...pulumi.InvokeOption) GetScheduledDeleteStackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScheduledDeleteStackResult, error) {
			args := v.(GetScheduledDeleteStackArgs)
			r, err := GetScheduledDeleteStack(ctx, &args, opts...)
			var s GetScheduledDeleteStackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScheduledDeleteStackResultOutput)
}

// A collection of arguments for invoking getScheduledDeleteStack.
type GetScheduledDeleteStackOutputArgs struct {
	// ID of the scheduled delete*stack (stack*id/schedule_id)
	ScheduledDeleteStackId pulumi.StringInput `pulumi:"scheduledDeleteStackId"`
}

func (GetScheduledDeleteStackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScheduledDeleteStackArgs)(nil)).Elem()
}

// A collection of values returned by getScheduledDeleteStack.
type GetScheduledDeleteStackResultOutput struct{ *pulumi.OutputState }

func (GetScheduledDeleteStackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScheduledDeleteStackResult)(nil)).Elem()
}

func (o GetScheduledDeleteStackResultOutput) ToGetScheduledDeleteStackResultOutput() GetScheduledDeleteStackResultOutput {
	return o
}

func (o GetScheduledDeleteStackResultOutput) ToGetScheduledDeleteStackResultOutputWithContext(ctx context.Context) GetScheduledDeleteStackResultOutput {
	return o
}

func (o GetScheduledDeleteStackResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetScheduledDeleteStackResult] {
	return pulumix.Output[GetScheduledDeleteStackResult]{
		OutputState: o.OutputState,
	}
}

// Timestamp (unix timestamp) at which time the scheduling should happen.
func (o GetScheduledDeleteStackResultOutput) At() pulumi.IntOutput {
	return o.ApplyT(func(v GetScheduledDeleteStackResult) int { return v.At }).(pulumi.IntOutput)
}

// Indicates whether the resources of the stack should be deleted.
func (o GetScheduledDeleteStackResultOutput) DeleteResources() pulumi.BoolOutput {
	return o.ApplyT(func(v GetScheduledDeleteStackResult) bool { return v.DeleteResources }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetScheduledDeleteStackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScheduledDeleteStackResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the schedule
func (o GetScheduledDeleteStackResultOutput) ScheduleId() pulumi.StringOutput {
	return o.ApplyT(func(v GetScheduledDeleteStackResult) string { return v.ScheduleId }).(pulumi.StringOutput)
}

// ID of the scheduled delete*stack (stack*id/schedule_id)
func (o GetScheduledDeleteStackResultOutput) ScheduledDeleteStackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetScheduledDeleteStackResult) string { return v.ScheduledDeleteStackId }).(pulumi.StringOutput)
}

// Stack ID of the scheduling config
func (o GetScheduledDeleteStackResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetScheduledDeleteStackResult) string { return v.StackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScheduledDeleteStackResultOutput{})
}