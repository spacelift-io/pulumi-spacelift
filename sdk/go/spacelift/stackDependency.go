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

// `StackDependency` represents a Spacelift **stack dependency** - a dependency between two stacks. When one stack depends on another, the tracked runs of the stack will not start until the dependent stack is successfully finished. Additionally, changes to the dependency will trigger the dependent.
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
//			infra, err := spacelift.NewStack(ctx, "infra", &spacelift.StackArgs{
//				Branch:     pulumi.String("master"),
//				Repository: pulumi.String("core-infra"),
//			})
//			if err != nil {
//				return err
//			}
//			app, err := spacelift.NewStack(ctx, "app", &spacelift.StackArgs{
//				Branch:     pulumi.String("master"),
//				Repository: pulumi.String("app"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.NewStackDependency(ctx, "test", &spacelift.StackDependencyArgs{
//				StackId:          app.ID(),
//				DependsOnStackId: infra.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StackDependency struct {
	pulumi.CustomResourceState

	// immutable ID (slug) of stack to depend on.
	DependsOnStackId pulumi.StringOutput `pulumi:"dependsOnStackId"`
	// immutable ID (slug) of stack which has a dependency.
	StackId pulumi.StringOutput `pulumi:"stackId"`
}

// NewStackDependency registers a new resource with the given unique name, arguments, and options.
func NewStackDependency(ctx *pulumi.Context,
	name string, args *StackDependencyArgs, opts ...pulumi.ResourceOption) (*StackDependency, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DependsOnStackId == nil {
		return nil, errors.New("invalid value for required argument 'DependsOnStackId'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StackDependency
	err := ctx.RegisterResource("spacelift:index/stackDependency:StackDependency", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStackDependency gets an existing StackDependency resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStackDependency(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackDependencyState, opts ...pulumi.ResourceOption) (*StackDependency, error) {
	var resource StackDependency
	err := ctx.ReadResource("spacelift:index/stackDependency:StackDependency", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StackDependency resources.
type stackDependencyState struct {
	// immutable ID (slug) of stack to depend on.
	DependsOnStackId *string `pulumi:"dependsOnStackId"`
	// immutable ID (slug) of stack which has a dependency.
	StackId *string `pulumi:"stackId"`
}

type StackDependencyState struct {
	// immutable ID (slug) of stack to depend on.
	DependsOnStackId pulumi.StringPtrInput
	// immutable ID (slug) of stack which has a dependency.
	StackId pulumi.StringPtrInput
}

func (StackDependencyState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackDependencyState)(nil)).Elem()
}

type stackDependencyArgs struct {
	// immutable ID (slug) of stack to depend on.
	DependsOnStackId string `pulumi:"dependsOnStackId"`
	// immutable ID (slug) of stack which has a dependency.
	StackId string `pulumi:"stackId"`
}

// The set of arguments for constructing a StackDependency resource.
type StackDependencyArgs struct {
	// immutable ID (slug) of stack to depend on.
	DependsOnStackId pulumi.StringInput
	// immutable ID (slug) of stack which has a dependency.
	StackId pulumi.StringInput
}

func (StackDependencyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackDependencyArgs)(nil)).Elem()
}

type StackDependencyInput interface {
	pulumi.Input

	ToStackDependencyOutput() StackDependencyOutput
	ToStackDependencyOutputWithContext(ctx context.Context) StackDependencyOutput
}

func (*StackDependency) ElementType() reflect.Type {
	return reflect.TypeOf((**StackDependency)(nil)).Elem()
}

func (i *StackDependency) ToStackDependencyOutput() StackDependencyOutput {
	return i.ToStackDependencyOutputWithContext(context.Background())
}

func (i *StackDependency) ToStackDependencyOutputWithContext(ctx context.Context) StackDependencyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackDependencyOutput)
}

func (i *StackDependency) ToOutput(ctx context.Context) pulumix.Output[*StackDependency] {
	return pulumix.Output[*StackDependency]{
		OutputState: i.ToStackDependencyOutputWithContext(ctx).OutputState,
	}
}

// StackDependencyArrayInput is an input type that accepts StackDependencyArray and StackDependencyArrayOutput values.
// You can construct a concrete instance of `StackDependencyArrayInput` via:
//
//	StackDependencyArray{ StackDependencyArgs{...} }
type StackDependencyArrayInput interface {
	pulumi.Input

	ToStackDependencyArrayOutput() StackDependencyArrayOutput
	ToStackDependencyArrayOutputWithContext(context.Context) StackDependencyArrayOutput
}

type StackDependencyArray []StackDependencyInput

func (StackDependencyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackDependency)(nil)).Elem()
}

func (i StackDependencyArray) ToStackDependencyArrayOutput() StackDependencyArrayOutput {
	return i.ToStackDependencyArrayOutputWithContext(context.Background())
}

func (i StackDependencyArray) ToStackDependencyArrayOutputWithContext(ctx context.Context) StackDependencyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackDependencyArrayOutput)
}

func (i StackDependencyArray) ToOutput(ctx context.Context) pulumix.Output[[]*StackDependency] {
	return pulumix.Output[[]*StackDependency]{
		OutputState: i.ToStackDependencyArrayOutputWithContext(ctx).OutputState,
	}
}

// StackDependencyMapInput is an input type that accepts StackDependencyMap and StackDependencyMapOutput values.
// You can construct a concrete instance of `StackDependencyMapInput` via:
//
//	StackDependencyMap{ "key": StackDependencyArgs{...} }
type StackDependencyMapInput interface {
	pulumi.Input

	ToStackDependencyMapOutput() StackDependencyMapOutput
	ToStackDependencyMapOutputWithContext(context.Context) StackDependencyMapOutput
}

type StackDependencyMap map[string]StackDependencyInput

func (StackDependencyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackDependency)(nil)).Elem()
}

func (i StackDependencyMap) ToStackDependencyMapOutput() StackDependencyMapOutput {
	return i.ToStackDependencyMapOutputWithContext(context.Background())
}

func (i StackDependencyMap) ToStackDependencyMapOutputWithContext(ctx context.Context) StackDependencyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackDependencyMapOutput)
}

func (i StackDependencyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*StackDependency] {
	return pulumix.Output[map[string]*StackDependency]{
		OutputState: i.ToStackDependencyMapOutputWithContext(ctx).OutputState,
	}
}

type StackDependencyOutput struct{ *pulumi.OutputState }

func (StackDependencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackDependency)(nil)).Elem()
}

func (o StackDependencyOutput) ToStackDependencyOutput() StackDependencyOutput {
	return o
}

func (o StackDependencyOutput) ToStackDependencyOutputWithContext(ctx context.Context) StackDependencyOutput {
	return o
}

func (o StackDependencyOutput) ToOutput(ctx context.Context) pulumix.Output[*StackDependency] {
	return pulumix.Output[*StackDependency]{
		OutputState: o.OutputState,
	}
}

// immutable ID (slug) of stack to depend on.
func (o StackDependencyOutput) DependsOnStackId() pulumi.StringOutput {
	return o.ApplyT(func(v *StackDependency) pulumi.StringOutput { return v.DependsOnStackId }).(pulumi.StringOutput)
}

// immutable ID (slug) of stack which has a dependency.
func (o StackDependencyOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *StackDependency) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

type StackDependencyArrayOutput struct{ *pulumi.OutputState }

func (StackDependencyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackDependency)(nil)).Elem()
}

func (o StackDependencyArrayOutput) ToStackDependencyArrayOutput() StackDependencyArrayOutput {
	return o
}

func (o StackDependencyArrayOutput) ToStackDependencyArrayOutputWithContext(ctx context.Context) StackDependencyArrayOutput {
	return o
}

func (o StackDependencyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*StackDependency] {
	return pulumix.Output[[]*StackDependency]{
		OutputState: o.OutputState,
	}
}

func (o StackDependencyArrayOutput) Index(i pulumi.IntInput) StackDependencyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StackDependency {
		return vs[0].([]*StackDependency)[vs[1].(int)]
	}).(StackDependencyOutput)
}

type StackDependencyMapOutput struct{ *pulumi.OutputState }

func (StackDependencyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackDependency)(nil)).Elem()
}

func (o StackDependencyMapOutput) ToStackDependencyMapOutput() StackDependencyMapOutput {
	return o
}

func (o StackDependencyMapOutput) ToStackDependencyMapOutputWithContext(ctx context.Context) StackDependencyMapOutput {
	return o
}

func (o StackDependencyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*StackDependency] {
	return pulumix.Output[map[string]*StackDependency]{
		OutputState: o.OutputState,
	}
}

func (o StackDependencyMapOutput) MapIndex(k pulumi.StringInput) StackDependencyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StackDependency {
		return vs[0].(map[string]*StackDependency)[vs[1].(string)]
	}).(StackDependencyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackDependencyInput)(nil)).Elem(), &StackDependency{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackDependencyArrayInput)(nil)).Elem(), StackDependencyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackDependencyMapInput)(nil)).Elem(), StackDependencyMap{})
	pulumi.RegisterOutputType(StackDependencyOutput{})
	pulumi.RegisterOutputType(StackDependencyArrayOutput{})
	pulumi.RegisterOutputType(StackDependencyMapOutput{})
}
