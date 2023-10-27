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

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-google/sdk/v1/go/google"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := spacelift.NewStack(ctx, "k8s-coreStack", &spacelift.StackArgs{
//				Branch:     pulumi.String("master"),
//				Repository: pulumi.String("core-infra"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.NewStackGcpServiceAccount(ctx, "k8s-coreStackGcpServiceAccount", &spacelift.StackGcpServiceAccountArgs{
//				StackId: k8s_coreStack.ID(),
//				TokenScopes: pulumi.StringArray{
//					pulumi.String("https://www.googleapis.com/auth/compute"),
//					pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
//					pulumi.String("https://www.googleapis.com/auth/devstorage.full_control"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = index.NewGoogle_project(ctx, "k8s-coregoogle_project", &index.Google_projectArgs{
//				Name:      "Kubernetes code",
//				ProjectId: "unicorn-k8s-core",
//				OrgId:     _var.Gcp_organization_id,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = index.NewGoogle_project_iam_member(ctx, "k8s-coregoogle_project_iam_member", &index.Google_project_iam_memberArgs{
//				Project: k8s_coregoogle_project.Id,
//				Role:    "roles/owner",
//				Member:  pulumi.String(fmt.Sprintf("serviceAccount:%v", k8s_coreStackGcpServiceAccount.ServiceAccountEmail)),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StackGcpServiceAccount struct {
	pulumi.CustomResourceState

	// ID of the module which uses GCP service account credentials
	ModuleId pulumi.StringPtrOutput `pulumi:"moduleId"`
	// Email address of the GCP service account dedicated for this stack
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// ID of the stack which uses GCP service account credentials
	StackId pulumi.StringPtrOutput `pulumi:"stackId"`
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes pulumi.StringArrayOutput `pulumi:"tokenScopes"`
}

// NewStackGcpServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewStackGcpServiceAccount(ctx *pulumi.Context,
	name string, args *StackGcpServiceAccountArgs, opts ...pulumi.ResourceOption) (*StackGcpServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TokenScopes == nil {
		return nil, errors.New("invalid value for required argument 'TokenScopes'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StackGcpServiceAccount
	err := ctx.RegisterResource("spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStackGcpServiceAccount gets an existing StackGcpServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStackGcpServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackGcpServiceAccountState, opts ...pulumi.ResourceOption) (*StackGcpServiceAccount, error) {
	var resource StackGcpServiceAccount
	err := ctx.ReadResource("spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StackGcpServiceAccount resources.
type stackGcpServiceAccountState struct {
	// ID of the module which uses GCP service account credentials
	ModuleId *string `pulumi:"moduleId"`
	// Email address of the GCP service account dedicated for this stack
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// ID of the stack which uses GCP service account credentials
	StackId *string `pulumi:"stackId"`
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes []string `pulumi:"tokenScopes"`
}

type StackGcpServiceAccountState struct {
	// ID of the module which uses GCP service account credentials
	ModuleId pulumi.StringPtrInput
	// Email address of the GCP service account dedicated for this stack
	ServiceAccountEmail pulumi.StringPtrInput
	// ID of the stack which uses GCP service account credentials
	StackId pulumi.StringPtrInput
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes pulumi.StringArrayInput
}

func (StackGcpServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackGcpServiceAccountState)(nil)).Elem()
}

type stackGcpServiceAccountArgs struct {
	// ID of the module which uses GCP service account credentials
	ModuleId *string `pulumi:"moduleId"`
	// ID of the stack which uses GCP service account credentials
	StackId *string `pulumi:"stackId"`
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes []string `pulumi:"tokenScopes"`
}

// The set of arguments for constructing a StackGcpServiceAccount resource.
type StackGcpServiceAccountArgs struct {
	// ID of the module which uses GCP service account credentials
	ModuleId pulumi.StringPtrInput
	// ID of the stack which uses GCP service account credentials
	StackId pulumi.StringPtrInput
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes pulumi.StringArrayInput
}

func (StackGcpServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackGcpServiceAccountArgs)(nil)).Elem()
}

type StackGcpServiceAccountInput interface {
	pulumi.Input

	ToStackGcpServiceAccountOutput() StackGcpServiceAccountOutput
	ToStackGcpServiceAccountOutputWithContext(ctx context.Context) StackGcpServiceAccountOutput
}

func (*StackGcpServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**StackGcpServiceAccount)(nil)).Elem()
}

func (i *StackGcpServiceAccount) ToStackGcpServiceAccountOutput() StackGcpServiceAccountOutput {
	return i.ToStackGcpServiceAccountOutputWithContext(context.Background())
}

func (i *StackGcpServiceAccount) ToStackGcpServiceAccountOutputWithContext(ctx context.Context) StackGcpServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackGcpServiceAccountOutput)
}

func (i *StackGcpServiceAccount) ToOutput(ctx context.Context) pulumix.Output[*StackGcpServiceAccount] {
	return pulumix.Output[*StackGcpServiceAccount]{
		OutputState: i.ToStackGcpServiceAccountOutputWithContext(ctx).OutputState,
	}
}

// StackGcpServiceAccountArrayInput is an input type that accepts StackGcpServiceAccountArray and StackGcpServiceAccountArrayOutput values.
// You can construct a concrete instance of `StackGcpServiceAccountArrayInput` via:
//
//	StackGcpServiceAccountArray{ StackGcpServiceAccountArgs{...} }
type StackGcpServiceAccountArrayInput interface {
	pulumi.Input

	ToStackGcpServiceAccountArrayOutput() StackGcpServiceAccountArrayOutput
	ToStackGcpServiceAccountArrayOutputWithContext(context.Context) StackGcpServiceAccountArrayOutput
}

type StackGcpServiceAccountArray []StackGcpServiceAccountInput

func (StackGcpServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackGcpServiceAccount)(nil)).Elem()
}

func (i StackGcpServiceAccountArray) ToStackGcpServiceAccountArrayOutput() StackGcpServiceAccountArrayOutput {
	return i.ToStackGcpServiceAccountArrayOutputWithContext(context.Background())
}

func (i StackGcpServiceAccountArray) ToStackGcpServiceAccountArrayOutputWithContext(ctx context.Context) StackGcpServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackGcpServiceAccountArrayOutput)
}

func (i StackGcpServiceAccountArray) ToOutput(ctx context.Context) pulumix.Output[[]*StackGcpServiceAccount] {
	return pulumix.Output[[]*StackGcpServiceAccount]{
		OutputState: i.ToStackGcpServiceAccountArrayOutputWithContext(ctx).OutputState,
	}
}

// StackGcpServiceAccountMapInput is an input type that accepts StackGcpServiceAccountMap and StackGcpServiceAccountMapOutput values.
// You can construct a concrete instance of `StackGcpServiceAccountMapInput` via:
//
//	StackGcpServiceAccountMap{ "key": StackGcpServiceAccountArgs{...} }
type StackGcpServiceAccountMapInput interface {
	pulumi.Input

	ToStackGcpServiceAccountMapOutput() StackGcpServiceAccountMapOutput
	ToStackGcpServiceAccountMapOutputWithContext(context.Context) StackGcpServiceAccountMapOutput
}

type StackGcpServiceAccountMap map[string]StackGcpServiceAccountInput

func (StackGcpServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackGcpServiceAccount)(nil)).Elem()
}

func (i StackGcpServiceAccountMap) ToStackGcpServiceAccountMapOutput() StackGcpServiceAccountMapOutput {
	return i.ToStackGcpServiceAccountMapOutputWithContext(context.Background())
}

func (i StackGcpServiceAccountMap) ToStackGcpServiceAccountMapOutputWithContext(ctx context.Context) StackGcpServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackGcpServiceAccountMapOutput)
}

func (i StackGcpServiceAccountMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*StackGcpServiceAccount] {
	return pulumix.Output[map[string]*StackGcpServiceAccount]{
		OutputState: i.ToStackGcpServiceAccountMapOutputWithContext(ctx).OutputState,
	}
}

type StackGcpServiceAccountOutput struct{ *pulumi.OutputState }

func (StackGcpServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackGcpServiceAccount)(nil)).Elem()
}

func (o StackGcpServiceAccountOutput) ToStackGcpServiceAccountOutput() StackGcpServiceAccountOutput {
	return o
}

func (o StackGcpServiceAccountOutput) ToStackGcpServiceAccountOutputWithContext(ctx context.Context) StackGcpServiceAccountOutput {
	return o
}

func (o StackGcpServiceAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*StackGcpServiceAccount] {
	return pulumix.Output[*StackGcpServiceAccount]{
		OutputState: o.OutputState,
	}
}

// ID of the module which uses GCP service account credentials
func (o StackGcpServiceAccountOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackGcpServiceAccount) pulumi.StringPtrOutput { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// Email address of the GCP service account dedicated for this stack
func (o StackGcpServiceAccountOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *StackGcpServiceAccount) pulumi.StringOutput { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

// ID of the stack which uses GCP service account credentials
func (o StackGcpServiceAccountOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackGcpServiceAccount) pulumi.StringPtrOutput { return v.StackId }).(pulumi.StringPtrOutput)
}

// List of scopes that will be requested when generating temporary GCP service account credentials
func (o StackGcpServiceAccountOutput) TokenScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StackGcpServiceAccount) pulumi.StringArrayOutput { return v.TokenScopes }).(pulumi.StringArrayOutput)
}

type StackGcpServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (StackGcpServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackGcpServiceAccount)(nil)).Elem()
}

func (o StackGcpServiceAccountArrayOutput) ToStackGcpServiceAccountArrayOutput() StackGcpServiceAccountArrayOutput {
	return o
}

func (o StackGcpServiceAccountArrayOutput) ToStackGcpServiceAccountArrayOutputWithContext(ctx context.Context) StackGcpServiceAccountArrayOutput {
	return o
}

func (o StackGcpServiceAccountArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*StackGcpServiceAccount] {
	return pulumix.Output[[]*StackGcpServiceAccount]{
		OutputState: o.OutputState,
	}
}

func (o StackGcpServiceAccountArrayOutput) Index(i pulumi.IntInput) StackGcpServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StackGcpServiceAccount {
		return vs[0].([]*StackGcpServiceAccount)[vs[1].(int)]
	}).(StackGcpServiceAccountOutput)
}

type StackGcpServiceAccountMapOutput struct{ *pulumi.OutputState }

func (StackGcpServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackGcpServiceAccount)(nil)).Elem()
}

func (o StackGcpServiceAccountMapOutput) ToStackGcpServiceAccountMapOutput() StackGcpServiceAccountMapOutput {
	return o
}

func (o StackGcpServiceAccountMapOutput) ToStackGcpServiceAccountMapOutputWithContext(ctx context.Context) StackGcpServiceAccountMapOutput {
	return o
}

func (o StackGcpServiceAccountMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*StackGcpServiceAccount] {
	return pulumix.Output[map[string]*StackGcpServiceAccount]{
		OutputState: o.OutputState,
	}
}

func (o StackGcpServiceAccountMapOutput) MapIndex(k pulumi.StringInput) StackGcpServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StackGcpServiceAccount {
		return vs[0].(map[string]*StackGcpServiceAccount)[vs[1].(string)]
	}).(StackGcpServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackGcpServiceAccountInput)(nil)).Elem(), &StackGcpServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackGcpServiceAccountArrayInput)(nil)).Elem(), StackGcpServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackGcpServiceAccountMapInput)(nil)).Elem(), StackGcpServiceAccountMap{})
	pulumi.RegisterOutputType(StackGcpServiceAccountOutput{})
	pulumi.RegisterOutputType(StackGcpServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(StackGcpServiceAccountMapOutput{})
}
