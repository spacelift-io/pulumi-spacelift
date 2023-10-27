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

// `Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
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
//			development, err := spacelift.NewSpace(ctx, "development", &spacelift.SpaceArgs{
//				ParentSpaceId: pulumi.String("root"),
//				Description:   pulumi.String("This a child of the root space. It contains all the resources common to the development infrastructure."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.NewSpace(ctx, "development-frontend", &spacelift.SpaceArgs{
//				ParentSpaceId:   development.ID(),
//				InheritEntities: pulumi.Bool(true),
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
//	$ pulumi import spacelift:index/space:Space development $SPACE_ID
//
// ```
type Space struct {
	pulumi.CustomResourceState

	// free-form space description for users
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
	InheritEntities pulumi.BoolPtrOutput `pulumi:"inheritEntities"`
	// list of labels describing a space
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// name of the space
	Name pulumi.StringOutput `pulumi:"name"`
	// immutable ID (slug) of parent space. Defaults to `root`.
	ParentSpaceId pulumi.StringPtrOutput `pulumi:"parentSpaceId"`
}

// NewSpace registers a new resource with the given unique name, arguments, and options.
func NewSpace(ctx *pulumi.Context,
	name string, args *SpaceArgs, opts ...pulumi.ResourceOption) (*Space, error) {
	if args == nil {
		args = &SpaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Space
	err := ctx.RegisterResource("spacelift:index/space:Space", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpace gets an existing Space resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpaceState, opts ...pulumi.ResourceOption) (*Space, error) {
	var resource Space
	err := ctx.ReadResource("spacelift:index/space:Space", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Space resources.
type spaceState struct {
	// free-form space description for users
	Description *string `pulumi:"description"`
	// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
	InheritEntities *bool `pulumi:"inheritEntities"`
	// list of labels describing a space
	Labels []string `pulumi:"labels"`
	// name of the space
	Name *string `pulumi:"name"`
	// immutable ID (slug) of parent space. Defaults to `root`.
	ParentSpaceId *string `pulumi:"parentSpaceId"`
}

type SpaceState struct {
	// free-form space description for users
	Description pulumi.StringPtrInput
	// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
	InheritEntities pulumi.BoolPtrInput
	// list of labels describing a space
	Labels pulumi.StringArrayInput
	// name of the space
	Name pulumi.StringPtrInput
	// immutable ID (slug) of parent space. Defaults to `root`.
	ParentSpaceId pulumi.StringPtrInput
}

func (SpaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*spaceState)(nil)).Elem()
}

type spaceArgs struct {
	// free-form space description for users
	Description *string `pulumi:"description"`
	// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
	InheritEntities *bool `pulumi:"inheritEntities"`
	// list of labels describing a space
	Labels []string `pulumi:"labels"`
	// name of the space
	Name *string `pulumi:"name"`
	// immutable ID (slug) of parent space. Defaults to `root`.
	ParentSpaceId *string `pulumi:"parentSpaceId"`
}

// The set of arguments for constructing a Space resource.
type SpaceArgs struct {
	// free-form space description for users
	Description pulumi.StringPtrInput
	// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
	InheritEntities pulumi.BoolPtrInput
	// list of labels describing a space
	Labels pulumi.StringArrayInput
	// name of the space
	Name pulumi.StringPtrInput
	// immutable ID (slug) of parent space. Defaults to `root`.
	ParentSpaceId pulumi.StringPtrInput
}

func (SpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spaceArgs)(nil)).Elem()
}

type SpaceInput interface {
	pulumi.Input

	ToSpaceOutput() SpaceOutput
	ToSpaceOutputWithContext(ctx context.Context) SpaceOutput
}

func (*Space) ElementType() reflect.Type {
	return reflect.TypeOf((**Space)(nil)).Elem()
}

func (i *Space) ToSpaceOutput() SpaceOutput {
	return i.ToSpaceOutputWithContext(context.Background())
}

func (i *Space) ToSpaceOutputWithContext(ctx context.Context) SpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceOutput)
}

func (i *Space) ToOutput(ctx context.Context) pulumix.Output[*Space] {
	return pulumix.Output[*Space]{
		OutputState: i.ToSpaceOutputWithContext(ctx).OutputState,
	}
}

// SpaceArrayInput is an input type that accepts SpaceArray and SpaceArrayOutput values.
// You can construct a concrete instance of `SpaceArrayInput` via:
//
//	SpaceArray{ SpaceArgs{...} }
type SpaceArrayInput interface {
	pulumi.Input

	ToSpaceArrayOutput() SpaceArrayOutput
	ToSpaceArrayOutputWithContext(context.Context) SpaceArrayOutput
}

type SpaceArray []SpaceInput

func (SpaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Space)(nil)).Elem()
}

func (i SpaceArray) ToSpaceArrayOutput() SpaceArrayOutput {
	return i.ToSpaceArrayOutputWithContext(context.Background())
}

func (i SpaceArray) ToSpaceArrayOutputWithContext(ctx context.Context) SpaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceArrayOutput)
}

func (i SpaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*Space] {
	return pulumix.Output[[]*Space]{
		OutputState: i.ToSpaceArrayOutputWithContext(ctx).OutputState,
	}
}

// SpaceMapInput is an input type that accepts SpaceMap and SpaceMapOutput values.
// You can construct a concrete instance of `SpaceMapInput` via:
//
//	SpaceMap{ "key": SpaceArgs{...} }
type SpaceMapInput interface {
	pulumi.Input

	ToSpaceMapOutput() SpaceMapOutput
	ToSpaceMapOutputWithContext(context.Context) SpaceMapOutput
}

type SpaceMap map[string]SpaceInput

func (SpaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Space)(nil)).Elem()
}

func (i SpaceMap) ToSpaceMapOutput() SpaceMapOutput {
	return i.ToSpaceMapOutputWithContext(context.Background())
}

func (i SpaceMap) ToSpaceMapOutputWithContext(ctx context.Context) SpaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceMapOutput)
}

func (i SpaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Space] {
	return pulumix.Output[map[string]*Space]{
		OutputState: i.ToSpaceMapOutputWithContext(ctx).OutputState,
	}
}

type SpaceOutput struct{ *pulumi.OutputState }

func (SpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Space)(nil)).Elem()
}

func (o SpaceOutput) ToSpaceOutput() SpaceOutput {
	return o
}

func (o SpaceOutput) ToSpaceOutputWithContext(ctx context.Context) SpaceOutput {
	return o
}

func (o SpaceOutput) ToOutput(ctx context.Context) pulumix.Output[*Space] {
	return pulumix.Output[*Space]{
		OutputState: o.OutputState,
	}
}

// free-form space description for users
func (o SpaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Space) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// indication whether access to this space inherits read access to entities from the parent space. Defaults to `false`.
func (o SpaceOutput) InheritEntities() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Space) pulumi.BoolPtrOutput { return v.InheritEntities }).(pulumi.BoolPtrOutput)
}

// list of labels describing a space
func (o SpaceOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Space) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// name of the space
func (o SpaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Space) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// immutable ID (slug) of parent space. Defaults to `root`.
func (o SpaceOutput) ParentSpaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Space) pulumi.StringPtrOutput { return v.ParentSpaceId }).(pulumi.StringPtrOutput)
}

type SpaceArrayOutput struct{ *pulumi.OutputState }

func (SpaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Space)(nil)).Elem()
}

func (o SpaceArrayOutput) ToSpaceArrayOutput() SpaceArrayOutput {
	return o
}

func (o SpaceArrayOutput) ToSpaceArrayOutputWithContext(ctx context.Context) SpaceArrayOutput {
	return o
}

func (o SpaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Space] {
	return pulumix.Output[[]*Space]{
		OutputState: o.OutputState,
	}
}

func (o SpaceArrayOutput) Index(i pulumi.IntInput) SpaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Space {
		return vs[0].([]*Space)[vs[1].(int)]
	}).(SpaceOutput)
}

type SpaceMapOutput struct{ *pulumi.OutputState }

func (SpaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Space)(nil)).Elem()
}

func (o SpaceMapOutput) ToSpaceMapOutput() SpaceMapOutput {
	return o
}

func (o SpaceMapOutput) ToSpaceMapOutputWithContext(ctx context.Context) SpaceMapOutput {
	return o
}

func (o SpaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Space] {
	return pulumix.Output[map[string]*Space]{
		OutputState: o.OutputState,
	}
}

func (o SpaceMapOutput) MapIndex(k pulumi.StringInput) SpaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Space {
		return vs[0].(map[string]*Space)[vs[1].(string)]
	}).(SpaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceInput)(nil)).Elem(), &Space{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceArrayInput)(nil)).Elem(), SpaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceMapInput)(nil)).Elem(), SpaceMap{})
	pulumi.RegisterOutputType(SpaceOutput{})
	pulumi.RegisterOutputType(SpaceArrayOutput{})
	pulumi.RegisterOutputType(SpaceMapOutput{})
}
