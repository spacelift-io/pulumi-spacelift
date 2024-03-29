// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"errors"
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
//			_, err := spacelift.NewStack(ctx, "k8s-core", nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.NewScheduledDeleteTask(ctx, "k8s-core-delete", &spacelift.ScheduledDeleteTaskArgs{
//				StackId:         k8s_core.ID(),
//				At:              pulumi.Int(1663336895),
//				DeleteResources: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import spacelift:index/scheduledDeleteTask:ScheduledDeleteTask ireland-kubeconfig $STACK_ID/$SCHEDULED_DELETE_STACK_ID
//
// ```
type ScheduledDeleteTask struct {
	pulumi.CustomResourceState

	// Timestamp (unix timestamp) at which time the scheduling should happen.
	At pulumi.IntOutput `pulumi:"at"`
	// Indicates whether the resources of the stack should be deleted.
	DeleteResources pulumi.BoolPtrOutput `pulumi:"deleteResources"`
	// ID of the schedule
	ScheduleId pulumi.StringOutput `pulumi:"scheduleId"`
	// ID of the stack for which to set up scheduling
	StackId pulumi.StringOutput `pulumi:"stackId"`
}

// NewScheduledDeleteTask registers a new resource with the given unique name, arguments, and options.
func NewScheduledDeleteTask(ctx *pulumi.Context,
	name string, args *ScheduledDeleteTaskArgs, opts ...pulumi.ResourceOption) (*ScheduledDeleteTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.At == nil {
		return nil, errors.New("invalid value for required argument 'At'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScheduledDeleteTask
	err := ctx.RegisterResource("spacelift:index/scheduledDeleteTask:ScheduledDeleteTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledDeleteTask gets an existing ScheduledDeleteTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledDeleteTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledDeleteTaskState, opts ...pulumi.ResourceOption) (*ScheduledDeleteTask, error) {
	var resource ScheduledDeleteTask
	err := ctx.ReadResource("spacelift:index/scheduledDeleteTask:ScheduledDeleteTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledDeleteTask resources.
type scheduledDeleteTaskState struct {
	// Timestamp (unix timestamp) at which time the scheduling should happen.
	At *int `pulumi:"at"`
	// Indicates whether the resources of the stack should be deleted.
	DeleteResources *bool `pulumi:"deleteResources"`
	// ID of the schedule
	ScheduleId *string `pulumi:"scheduleId"`
	// ID of the stack for which to set up scheduling
	StackId *string `pulumi:"stackId"`
}

type ScheduledDeleteTaskState struct {
	// Timestamp (unix timestamp) at which time the scheduling should happen.
	At pulumi.IntPtrInput
	// Indicates whether the resources of the stack should be deleted.
	DeleteResources pulumi.BoolPtrInput
	// ID of the schedule
	ScheduleId pulumi.StringPtrInput
	// ID of the stack for which to set up scheduling
	StackId pulumi.StringPtrInput
}

func (ScheduledDeleteTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledDeleteTaskState)(nil)).Elem()
}

type scheduledDeleteTaskArgs struct {
	// Timestamp (unix timestamp) at which time the scheduling should happen.
	At int `pulumi:"at"`
	// Indicates whether the resources of the stack should be deleted.
	DeleteResources *bool `pulumi:"deleteResources"`
	// ID of the schedule
	ScheduleId *string `pulumi:"scheduleId"`
	// ID of the stack for which to set up scheduling
	StackId string `pulumi:"stackId"`
}

// The set of arguments for constructing a ScheduledDeleteTask resource.
type ScheduledDeleteTaskArgs struct {
	// Timestamp (unix timestamp) at which time the scheduling should happen.
	At pulumi.IntInput
	// Indicates whether the resources of the stack should be deleted.
	DeleteResources pulumi.BoolPtrInput
	// ID of the schedule
	ScheduleId pulumi.StringPtrInput
	// ID of the stack for which to set up scheduling
	StackId pulumi.StringInput
}

func (ScheduledDeleteTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledDeleteTaskArgs)(nil)).Elem()
}

type ScheduledDeleteTaskInput interface {
	pulumi.Input

	ToScheduledDeleteTaskOutput() ScheduledDeleteTaskOutput
	ToScheduledDeleteTaskOutputWithContext(ctx context.Context) ScheduledDeleteTaskOutput
}

func (*ScheduledDeleteTask) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledDeleteTask)(nil)).Elem()
}

func (i *ScheduledDeleteTask) ToScheduledDeleteTaskOutput() ScheduledDeleteTaskOutput {
	return i.ToScheduledDeleteTaskOutputWithContext(context.Background())
}

func (i *ScheduledDeleteTask) ToScheduledDeleteTaskOutputWithContext(ctx context.Context) ScheduledDeleteTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledDeleteTaskOutput)
}

func (i *ScheduledDeleteTask) ToOutput(ctx context.Context) pulumix.Output[*ScheduledDeleteTask] {
	return pulumix.Output[*ScheduledDeleteTask]{
		OutputState: i.ToScheduledDeleteTaskOutputWithContext(ctx).OutputState,
	}
}

// ScheduledDeleteTaskArrayInput is an input type that accepts ScheduledDeleteTaskArray and ScheduledDeleteTaskArrayOutput values.
// You can construct a concrete instance of `ScheduledDeleteTaskArrayInput` via:
//
//	ScheduledDeleteTaskArray{ ScheduledDeleteTaskArgs{...} }
type ScheduledDeleteTaskArrayInput interface {
	pulumi.Input

	ToScheduledDeleteTaskArrayOutput() ScheduledDeleteTaskArrayOutput
	ToScheduledDeleteTaskArrayOutputWithContext(context.Context) ScheduledDeleteTaskArrayOutput
}

type ScheduledDeleteTaskArray []ScheduledDeleteTaskInput

func (ScheduledDeleteTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledDeleteTask)(nil)).Elem()
}

func (i ScheduledDeleteTaskArray) ToScheduledDeleteTaskArrayOutput() ScheduledDeleteTaskArrayOutput {
	return i.ToScheduledDeleteTaskArrayOutputWithContext(context.Background())
}

func (i ScheduledDeleteTaskArray) ToScheduledDeleteTaskArrayOutputWithContext(ctx context.Context) ScheduledDeleteTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledDeleteTaskArrayOutput)
}

func (i ScheduledDeleteTaskArray) ToOutput(ctx context.Context) pulumix.Output[[]*ScheduledDeleteTask] {
	return pulumix.Output[[]*ScheduledDeleteTask]{
		OutputState: i.ToScheduledDeleteTaskArrayOutputWithContext(ctx).OutputState,
	}
}

// ScheduledDeleteTaskMapInput is an input type that accepts ScheduledDeleteTaskMap and ScheduledDeleteTaskMapOutput values.
// You can construct a concrete instance of `ScheduledDeleteTaskMapInput` via:
//
//	ScheduledDeleteTaskMap{ "key": ScheduledDeleteTaskArgs{...} }
type ScheduledDeleteTaskMapInput interface {
	pulumi.Input

	ToScheduledDeleteTaskMapOutput() ScheduledDeleteTaskMapOutput
	ToScheduledDeleteTaskMapOutputWithContext(context.Context) ScheduledDeleteTaskMapOutput
}

type ScheduledDeleteTaskMap map[string]ScheduledDeleteTaskInput

func (ScheduledDeleteTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledDeleteTask)(nil)).Elem()
}

func (i ScheduledDeleteTaskMap) ToScheduledDeleteTaskMapOutput() ScheduledDeleteTaskMapOutput {
	return i.ToScheduledDeleteTaskMapOutputWithContext(context.Background())
}

func (i ScheduledDeleteTaskMap) ToScheduledDeleteTaskMapOutputWithContext(ctx context.Context) ScheduledDeleteTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledDeleteTaskMapOutput)
}

func (i ScheduledDeleteTaskMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ScheduledDeleteTask] {
	return pulumix.Output[map[string]*ScheduledDeleteTask]{
		OutputState: i.ToScheduledDeleteTaskMapOutputWithContext(ctx).OutputState,
	}
}

type ScheduledDeleteTaskOutput struct{ *pulumi.OutputState }

func (ScheduledDeleteTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledDeleteTask)(nil)).Elem()
}

func (o ScheduledDeleteTaskOutput) ToScheduledDeleteTaskOutput() ScheduledDeleteTaskOutput {
	return o
}

func (o ScheduledDeleteTaskOutput) ToScheduledDeleteTaskOutputWithContext(ctx context.Context) ScheduledDeleteTaskOutput {
	return o
}

func (o ScheduledDeleteTaskOutput) ToOutput(ctx context.Context) pulumix.Output[*ScheduledDeleteTask] {
	return pulumix.Output[*ScheduledDeleteTask]{
		OutputState: o.OutputState,
	}
}

// Timestamp (unix timestamp) at which time the scheduling should happen.
func (o ScheduledDeleteTaskOutput) At() pulumi.IntOutput {
	return o.ApplyT(func(v *ScheduledDeleteTask) pulumi.IntOutput { return v.At }).(pulumi.IntOutput)
}

// Indicates whether the resources of the stack should be deleted.
func (o ScheduledDeleteTaskOutput) DeleteResources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScheduledDeleteTask) pulumi.BoolPtrOutput { return v.DeleteResources }).(pulumi.BoolPtrOutput)
}

// ID of the schedule
func (o ScheduledDeleteTaskOutput) ScheduleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledDeleteTask) pulumi.StringOutput { return v.ScheduleId }).(pulumi.StringOutput)
}

// ID of the stack for which to set up scheduling
func (o ScheduledDeleteTaskOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledDeleteTask) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

type ScheduledDeleteTaskArrayOutput struct{ *pulumi.OutputState }

func (ScheduledDeleteTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledDeleteTask)(nil)).Elem()
}

func (o ScheduledDeleteTaskArrayOutput) ToScheduledDeleteTaskArrayOutput() ScheduledDeleteTaskArrayOutput {
	return o
}

func (o ScheduledDeleteTaskArrayOutput) ToScheduledDeleteTaskArrayOutputWithContext(ctx context.Context) ScheduledDeleteTaskArrayOutput {
	return o
}

func (o ScheduledDeleteTaskArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ScheduledDeleteTask] {
	return pulumix.Output[[]*ScheduledDeleteTask]{
		OutputState: o.OutputState,
	}
}

func (o ScheduledDeleteTaskArrayOutput) Index(i pulumi.IntInput) ScheduledDeleteTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScheduledDeleteTask {
		return vs[0].([]*ScheduledDeleteTask)[vs[1].(int)]
	}).(ScheduledDeleteTaskOutput)
}

type ScheduledDeleteTaskMapOutput struct{ *pulumi.OutputState }

func (ScheduledDeleteTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledDeleteTask)(nil)).Elem()
}

func (o ScheduledDeleteTaskMapOutput) ToScheduledDeleteTaskMapOutput() ScheduledDeleteTaskMapOutput {
	return o
}

func (o ScheduledDeleteTaskMapOutput) ToScheduledDeleteTaskMapOutputWithContext(ctx context.Context) ScheduledDeleteTaskMapOutput {
	return o
}

func (o ScheduledDeleteTaskMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ScheduledDeleteTask] {
	return pulumix.Output[map[string]*ScheduledDeleteTask]{
		OutputState: o.OutputState,
	}
}

func (o ScheduledDeleteTaskMapOutput) MapIndex(k pulumi.StringInput) ScheduledDeleteTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScheduledDeleteTask {
		return vs[0].(map[string]*ScheduledDeleteTask)[vs[1].(string)]
	}).(ScheduledDeleteTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledDeleteTaskInput)(nil)).Elem(), &ScheduledDeleteTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledDeleteTaskArrayInput)(nil)).Elem(), ScheduledDeleteTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledDeleteTaskMapInput)(nil)).Elem(), ScheduledDeleteTaskMap{})
	pulumi.RegisterOutputType(ScheduledDeleteTaskOutput{})
	pulumi.RegisterOutputType(ScheduledDeleteTaskArrayOutput{})
	pulumi.RegisterOutputType(ScheduledDeleteTaskMapOutput{})
}
