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

// `Run` allows programmatically triggering runs in response to arbitrary changes in the keepers section.
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
//			thisStack, err := spacelift.NewStack(ctx, "thisStack", &spacelift.StackArgs{
//				Repository: pulumi.String("test"),
//				Branch:     pulumi.String("main"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.NewRun(ctx, "thisRun", &spacelift.RunArgs{
//				StackId: thisStack.ID(),
//				Keepers: pulumi.Map{
//					"branch": thisStack.Branch,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Run struct {
	pulumi.CustomResourceState

	// The commit SHA for which to trigger a run.
	CommitSha pulumi.StringPtrOutput `pulumi:"commitSha"`
	// Arbitrary map of values that, when changed, will trigger recreation of the resource.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// Whether the run is a proposed run. Defaults to `false`.
	Proposed pulumi.BoolPtrOutput `pulumi:"proposed"`
	// ID of the stack on which the run is to be triggered.
	StackId pulumi.StringOutput `pulumi:"stackId"`
}

// NewRun registers a new resource with the given unique name, arguments, and options.
func NewRun(ctx *pulumi.Context,
	name string, args *RunArgs, opts ...pulumi.ResourceOption) (*Run, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Run
	err := ctx.RegisterResource("spacelift:index/run:Run", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRun gets an existing Run resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RunState, opts ...pulumi.ResourceOption) (*Run, error) {
	var resource Run
	err := ctx.ReadResource("spacelift:index/run:Run", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Run resources.
type runState struct {
	// The commit SHA for which to trigger a run.
	CommitSha *string `pulumi:"commitSha"`
	// Arbitrary map of values that, when changed, will trigger recreation of the resource.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// Whether the run is a proposed run. Defaults to `false`.
	Proposed *bool `pulumi:"proposed"`
	// ID of the stack on which the run is to be triggered.
	StackId *string `pulumi:"stackId"`
}

type RunState struct {
	// The commit SHA for which to trigger a run.
	CommitSha pulumi.StringPtrInput
	// Arbitrary map of values that, when changed, will trigger recreation of the resource.
	Keepers pulumi.MapInput
	// Whether the run is a proposed run. Defaults to `false`.
	Proposed pulumi.BoolPtrInput
	// ID of the stack on which the run is to be triggered.
	StackId pulumi.StringPtrInput
}

func (RunState) ElementType() reflect.Type {
	return reflect.TypeOf((*runState)(nil)).Elem()
}

type runArgs struct {
	// The commit SHA for which to trigger a run.
	CommitSha *string `pulumi:"commitSha"`
	// Arbitrary map of values that, when changed, will trigger recreation of the resource.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// Whether the run is a proposed run. Defaults to `false`.
	Proposed *bool `pulumi:"proposed"`
	// ID of the stack on which the run is to be triggered.
	StackId string `pulumi:"stackId"`
}

// The set of arguments for constructing a Run resource.
type RunArgs struct {
	// The commit SHA for which to trigger a run.
	CommitSha pulumi.StringPtrInput
	// Arbitrary map of values that, when changed, will trigger recreation of the resource.
	Keepers pulumi.MapInput
	// Whether the run is a proposed run. Defaults to `false`.
	Proposed pulumi.BoolPtrInput
	// ID of the stack on which the run is to be triggered.
	StackId pulumi.StringInput
}

func (RunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runArgs)(nil)).Elem()
}

type RunInput interface {
	pulumi.Input

	ToRunOutput() RunOutput
	ToRunOutputWithContext(ctx context.Context) RunOutput
}

func (*Run) ElementType() reflect.Type {
	return reflect.TypeOf((**Run)(nil)).Elem()
}

func (i *Run) ToRunOutput() RunOutput {
	return i.ToRunOutputWithContext(context.Background())
}

func (i *Run) ToRunOutputWithContext(ctx context.Context) RunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunOutput)
}

func (i *Run) ToOutput(ctx context.Context) pulumix.Output[*Run] {
	return pulumix.Output[*Run]{
		OutputState: i.ToRunOutputWithContext(ctx).OutputState,
	}
}

// RunArrayInput is an input type that accepts RunArray and RunArrayOutput values.
// You can construct a concrete instance of `RunArrayInput` via:
//
//	RunArray{ RunArgs{...} }
type RunArrayInput interface {
	pulumi.Input

	ToRunArrayOutput() RunArrayOutput
	ToRunArrayOutputWithContext(context.Context) RunArrayOutput
}

type RunArray []RunInput

func (RunArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Run)(nil)).Elem()
}

func (i RunArray) ToRunArrayOutput() RunArrayOutput {
	return i.ToRunArrayOutputWithContext(context.Background())
}

func (i RunArray) ToRunArrayOutputWithContext(ctx context.Context) RunArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunArrayOutput)
}

func (i RunArray) ToOutput(ctx context.Context) pulumix.Output[[]*Run] {
	return pulumix.Output[[]*Run]{
		OutputState: i.ToRunArrayOutputWithContext(ctx).OutputState,
	}
}

// RunMapInput is an input type that accepts RunMap and RunMapOutput values.
// You can construct a concrete instance of `RunMapInput` via:
//
//	RunMap{ "key": RunArgs{...} }
type RunMapInput interface {
	pulumi.Input

	ToRunMapOutput() RunMapOutput
	ToRunMapOutputWithContext(context.Context) RunMapOutput
}

type RunMap map[string]RunInput

func (RunMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Run)(nil)).Elem()
}

func (i RunMap) ToRunMapOutput() RunMapOutput {
	return i.ToRunMapOutputWithContext(context.Background())
}

func (i RunMap) ToRunMapOutputWithContext(ctx context.Context) RunMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunMapOutput)
}

func (i RunMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Run] {
	return pulumix.Output[map[string]*Run]{
		OutputState: i.ToRunMapOutputWithContext(ctx).OutputState,
	}
}

type RunOutput struct{ *pulumi.OutputState }

func (RunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Run)(nil)).Elem()
}

func (o RunOutput) ToRunOutput() RunOutput {
	return o
}

func (o RunOutput) ToRunOutputWithContext(ctx context.Context) RunOutput {
	return o
}

func (o RunOutput) ToOutput(ctx context.Context) pulumix.Output[*Run] {
	return pulumix.Output[*Run]{
		OutputState: o.OutputState,
	}
}

// The commit SHA for which to trigger a run.
func (o RunOutput) CommitSha() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Run) pulumi.StringPtrOutput { return v.CommitSha }).(pulumi.StringPtrOutput)
}

// Arbitrary map of values that, when changed, will trigger recreation of the resource.
func (o RunOutput) Keepers() pulumi.MapOutput {
	return o.ApplyT(func(v *Run) pulumi.MapOutput { return v.Keepers }).(pulumi.MapOutput)
}

// Whether the run is a proposed run. Defaults to `false`.
func (o RunOutput) Proposed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Run) pulumi.BoolPtrOutput { return v.Proposed }).(pulumi.BoolPtrOutput)
}

// ID of the stack on which the run is to be triggered.
func (o RunOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *Run) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

type RunArrayOutput struct{ *pulumi.OutputState }

func (RunArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Run)(nil)).Elem()
}

func (o RunArrayOutput) ToRunArrayOutput() RunArrayOutput {
	return o
}

func (o RunArrayOutput) ToRunArrayOutputWithContext(ctx context.Context) RunArrayOutput {
	return o
}

func (o RunArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Run] {
	return pulumix.Output[[]*Run]{
		OutputState: o.OutputState,
	}
}

func (o RunArrayOutput) Index(i pulumi.IntInput) RunOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Run {
		return vs[0].([]*Run)[vs[1].(int)]
	}).(RunOutput)
}

type RunMapOutput struct{ *pulumi.OutputState }

func (RunMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Run)(nil)).Elem()
}

func (o RunMapOutput) ToRunMapOutput() RunMapOutput {
	return o
}

func (o RunMapOutput) ToRunMapOutputWithContext(ctx context.Context) RunMapOutput {
	return o
}

func (o RunMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Run] {
	return pulumix.Output[map[string]*Run]{
		OutputState: o.OutputState,
	}
}

func (o RunMapOutput) MapIndex(k pulumi.StringInput) RunOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Run {
		return vs[0].(map[string]*Run)[vs[1].(string)]
	}).(RunOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RunInput)(nil)).Elem(), &Run{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunArrayInput)(nil)).Elem(), RunArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunMapInput)(nil)).Elem(), RunMap{})
	pulumi.RegisterOutputType(RunOutput{})
	pulumi.RegisterOutputType(RunArrayOutput{})
	pulumi.RegisterOutputType(RunMapOutput{})
}
