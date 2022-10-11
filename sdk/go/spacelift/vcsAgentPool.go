// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `VcsAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/spacelift-io/pulumi-spacelift/sdk/go/spacelift"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := spacelift.NewVcsAgentPool(ctx, "ghe", &spacelift.VcsAgentPoolArgs{
//				Description: pulumi.String("VCS agent pool for our internal GitHub Enterprise"),
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
//	$ pulumi import spacelift:index/vcsAgentPool:VcsAgentPool ghe $VCS_AGENT_POOL_ID
//
// ```
type VcsAgentPool struct {
	pulumi.CustomResourceState

	// VCS agent pool configuration, encoded using base64
	Config pulumi.StringOutput `pulumi:"config"`
	// Free-form VCS agent pool description for users
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the VCS agent pool, must be unique within an account
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewVcsAgentPool registers a new resource with the given unique name, arguments, and options.
func NewVcsAgentPool(ctx *pulumi.Context,
	name string, args *VcsAgentPoolArgs, opts ...pulumi.ResourceOption) (*VcsAgentPool, error) {
	if args == nil {
		args = &VcsAgentPoolArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VcsAgentPool
	err := ctx.RegisterResource("spacelift:index/vcsAgentPool:VcsAgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVcsAgentPool gets an existing VcsAgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVcsAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VcsAgentPoolState, opts ...pulumi.ResourceOption) (*VcsAgentPool, error) {
	var resource VcsAgentPool
	err := ctx.ReadResource("spacelift:index/vcsAgentPool:VcsAgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VcsAgentPool resources.
type vcsAgentPoolState struct {
	// VCS agent pool configuration, encoded using base64
	Config *string `pulumi:"config"`
	// Free-form VCS agent pool description for users
	Description *string `pulumi:"description"`
	// Name of the VCS agent pool, must be unique within an account
	Name *string `pulumi:"name"`
}

type VcsAgentPoolState struct {
	// VCS agent pool configuration, encoded using base64
	Config pulumi.StringPtrInput
	// Free-form VCS agent pool description for users
	Description pulumi.StringPtrInput
	// Name of the VCS agent pool, must be unique within an account
	Name pulumi.StringPtrInput
}

func (VcsAgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*vcsAgentPoolState)(nil)).Elem()
}

type vcsAgentPoolArgs struct {
	// Free-form VCS agent pool description for users
	Description *string `pulumi:"description"`
	// Name of the VCS agent pool, must be unique within an account
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a VcsAgentPool resource.
type VcsAgentPoolArgs struct {
	// Free-form VCS agent pool description for users
	Description pulumi.StringPtrInput
	// Name of the VCS agent pool, must be unique within an account
	Name pulumi.StringPtrInput
}

func (VcsAgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vcsAgentPoolArgs)(nil)).Elem()
}

type VcsAgentPoolInput interface {
	pulumi.Input

	ToVcsAgentPoolOutput() VcsAgentPoolOutput
	ToVcsAgentPoolOutputWithContext(ctx context.Context) VcsAgentPoolOutput
}

func (*VcsAgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**VcsAgentPool)(nil)).Elem()
}

func (i *VcsAgentPool) ToVcsAgentPoolOutput() VcsAgentPoolOutput {
	return i.ToVcsAgentPoolOutputWithContext(context.Background())
}

func (i *VcsAgentPool) ToVcsAgentPoolOutputWithContext(ctx context.Context) VcsAgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VcsAgentPoolOutput)
}

// VcsAgentPoolArrayInput is an input type that accepts VcsAgentPoolArray and VcsAgentPoolArrayOutput values.
// You can construct a concrete instance of `VcsAgentPoolArrayInput` via:
//
//	VcsAgentPoolArray{ VcsAgentPoolArgs{...} }
type VcsAgentPoolArrayInput interface {
	pulumi.Input

	ToVcsAgentPoolArrayOutput() VcsAgentPoolArrayOutput
	ToVcsAgentPoolArrayOutputWithContext(context.Context) VcsAgentPoolArrayOutput
}

type VcsAgentPoolArray []VcsAgentPoolInput

func (VcsAgentPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VcsAgentPool)(nil)).Elem()
}

func (i VcsAgentPoolArray) ToVcsAgentPoolArrayOutput() VcsAgentPoolArrayOutput {
	return i.ToVcsAgentPoolArrayOutputWithContext(context.Background())
}

func (i VcsAgentPoolArray) ToVcsAgentPoolArrayOutputWithContext(ctx context.Context) VcsAgentPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VcsAgentPoolArrayOutput)
}

// VcsAgentPoolMapInput is an input type that accepts VcsAgentPoolMap and VcsAgentPoolMapOutput values.
// You can construct a concrete instance of `VcsAgentPoolMapInput` via:
//
//	VcsAgentPoolMap{ "key": VcsAgentPoolArgs{...} }
type VcsAgentPoolMapInput interface {
	pulumi.Input

	ToVcsAgentPoolMapOutput() VcsAgentPoolMapOutput
	ToVcsAgentPoolMapOutputWithContext(context.Context) VcsAgentPoolMapOutput
}

type VcsAgentPoolMap map[string]VcsAgentPoolInput

func (VcsAgentPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VcsAgentPool)(nil)).Elem()
}

func (i VcsAgentPoolMap) ToVcsAgentPoolMapOutput() VcsAgentPoolMapOutput {
	return i.ToVcsAgentPoolMapOutputWithContext(context.Background())
}

func (i VcsAgentPoolMap) ToVcsAgentPoolMapOutputWithContext(ctx context.Context) VcsAgentPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VcsAgentPoolMapOutput)
}

type VcsAgentPoolOutput struct{ *pulumi.OutputState }

func (VcsAgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VcsAgentPool)(nil)).Elem()
}

func (o VcsAgentPoolOutput) ToVcsAgentPoolOutput() VcsAgentPoolOutput {
	return o
}

func (o VcsAgentPoolOutput) ToVcsAgentPoolOutputWithContext(ctx context.Context) VcsAgentPoolOutput {
	return o
}

// VCS agent pool configuration, encoded using base64
func (o VcsAgentPoolOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *VcsAgentPool) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// Free-form VCS agent pool description for users
func (o VcsAgentPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VcsAgentPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the VCS agent pool, must be unique within an account
func (o VcsAgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VcsAgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type VcsAgentPoolArrayOutput struct{ *pulumi.OutputState }

func (VcsAgentPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VcsAgentPool)(nil)).Elem()
}

func (o VcsAgentPoolArrayOutput) ToVcsAgentPoolArrayOutput() VcsAgentPoolArrayOutput {
	return o
}

func (o VcsAgentPoolArrayOutput) ToVcsAgentPoolArrayOutputWithContext(ctx context.Context) VcsAgentPoolArrayOutput {
	return o
}

func (o VcsAgentPoolArrayOutput) Index(i pulumi.IntInput) VcsAgentPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VcsAgentPool {
		return vs[0].([]*VcsAgentPool)[vs[1].(int)]
	}).(VcsAgentPoolOutput)
}

type VcsAgentPoolMapOutput struct{ *pulumi.OutputState }

func (VcsAgentPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VcsAgentPool)(nil)).Elem()
}

func (o VcsAgentPoolMapOutput) ToVcsAgentPoolMapOutput() VcsAgentPoolMapOutput {
	return o
}

func (o VcsAgentPoolMapOutput) ToVcsAgentPoolMapOutputWithContext(ctx context.Context) VcsAgentPoolMapOutput {
	return o
}

func (o VcsAgentPoolMapOutput) MapIndex(k pulumi.StringInput) VcsAgentPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VcsAgentPool {
		return vs[0].(map[string]*VcsAgentPool)[vs[1].(string)]
	}).(VcsAgentPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VcsAgentPoolInput)(nil)).Elem(), &VcsAgentPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*VcsAgentPoolArrayInput)(nil)).Elem(), VcsAgentPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VcsAgentPoolMapInput)(nil)).Elem(), VcsAgentPoolMap{})
	pulumi.RegisterOutputType(VcsAgentPoolOutput{})
	pulumi.RegisterOutputType(VcsAgentPoolArrayOutput{})
	pulumi.RegisterOutputType(VcsAgentPoolMapOutput{})
}
