// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `VCSAgentPool` represents a Spacelift **VCS agent pool** - a logical group of proxies allowing Spacelift to access private VCS installations
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/spacelift-io/spacelift-spacelift/sdk/go/spacelift/"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spacelift.NewVCSAgentPool(ctx, "ghe", &spacelift.VCSAgentPoolArgs{
// 			Description: pulumi.String("VCS agent pool for our internal GitHub Enterpris"),
// 			Name:        pulumi.String("ghe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// <!-- schema generated by tfplugindocs -->
// ## Schema
//
// ### Required
//
// - **name** (String) Name of the VCS agent pool, must be unique within an account
//
// ### Optional
//
// - **description** (String) Free-form VCS agent pool description for users
// - **id** (String) The ID of this resource.
//
// ### Read-Only
//
// - **config** (String, Sensitive) VCS agent pool configuration, encoded using base64
//
// ## Import
//
// Import is supported using the following syntax
//
// ```sh
//  $ pulumi import spacelift:index/vCSAgentPool:VCSAgentPool ghe $VCS_AGENT_POOL_ID
// ```
type VCSAgentPool struct {
	pulumi.CustomResourceState

	// VCS agent pool configuration, encoded using base64
	Config pulumi.StringOutput `pulumi:"config"`
	// Free-form VCS agent pool description for users
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the VCS agent pool, must be unique within an account
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewVCSAgentPool registers a new resource with the given unique name, arguments, and options.
func NewVCSAgentPool(ctx *pulumi.Context,
	name string, args *VCSAgentPoolArgs, opts ...pulumi.ResourceOption) (*VCSAgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource VCSAgentPool
	err := ctx.RegisterResource("spacelift:index/vCSAgentPool:VCSAgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVCSAgentPool gets an existing VCSAgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVCSAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VCSAgentPoolState, opts ...pulumi.ResourceOption) (*VCSAgentPool, error) {
	var resource VCSAgentPool
	err := ctx.ReadResource("spacelift:index/vCSAgentPool:VCSAgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VCSAgentPool resources.
type vcsagentPoolState struct {
	// VCS agent pool configuration, encoded using base64
	Config *string `pulumi:"config"`
	// Free-form VCS agent pool description for users
	Description *string `pulumi:"description"`
	// Name of the VCS agent pool, must be unique within an account
	Name *string `pulumi:"name"`
}

type VCSAgentPoolState struct {
	// VCS agent pool configuration, encoded using base64
	Config pulumi.StringPtrInput
	// Free-form VCS agent pool description for users
	Description pulumi.StringPtrInput
	// Name of the VCS agent pool, must be unique within an account
	Name pulumi.StringPtrInput
}

func (VCSAgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*vcsagentPoolState)(nil)).Elem()
}

type vcsagentPoolArgs struct {
	// Free-form VCS agent pool description for users
	Description *string `pulumi:"description"`
	// Name of the VCS agent pool, must be unique within an account
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a VCSAgentPool resource.
type VCSAgentPoolArgs struct {
	// Free-form VCS agent pool description for users
	Description pulumi.StringPtrInput
	// Name of the VCS agent pool, must be unique within an account
	Name pulumi.StringInput
}

func (VCSAgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vcsagentPoolArgs)(nil)).Elem()
}

type VCSAgentPoolInput interface {
	pulumi.Input

	ToVCSAgentPoolOutput() VCSAgentPoolOutput
	ToVCSAgentPoolOutputWithContext(ctx context.Context) VCSAgentPoolOutput
}

func (*VCSAgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((*VCSAgentPool)(nil))
}

func (i *VCSAgentPool) ToVCSAgentPoolOutput() VCSAgentPoolOutput {
	return i.ToVCSAgentPoolOutputWithContext(context.Background())
}

func (i *VCSAgentPool) ToVCSAgentPoolOutputWithContext(ctx context.Context) VCSAgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VCSAgentPoolOutput)
}

type VCSAgentPoolOutput struct {
	*pulumi.OutputState
}

func (VCSAgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VCSAgentPool)(nil))
}

func (o VCSAgentPoolOutput) ToVCSAgentPoolOutput() VCSAgentPoolOutput {
	return o
}

func (o VCSAgentPoolOutput) ToVCSAgentPoolOutputWithContext(ctx context.Context) VCSAgentPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VCSAgentPoolOutput{})
}
