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
//			_, err = spacelift.NewGcpServiceAccount(ctx, "k8s-coreGcpServiceAccount", &spacelift.GcpServiceAccountArgs{
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
//				Member:  pulumi.String(fmt.Sprintf("serviceAccount:%v", k8s_coreGcpServiceAccount.ServiceAccountEmail)),
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
//	$ pulumi import spacelift:index/gcpServiceAccount:GcpServiceAccount k8s-core stack/$STACK_ID
//
// ```
//
// ```sh
//
//	$ pulumi import spacelift:index/gcpServiceAccount:GcpServiceAccount k8s-core module/$MODULE_ID
//
// ```
type GcpServiceAccount struct {
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

// NewGcpServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewGcpServiceAccount(ctx *pulumi.Context,
	name string, args *GcpServiceAccountArgs, opts ...pulumi.ResourceOption) (*GcpServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TokenScopes == nil {
		return nil, errors.New("invalid value for required argument 'TokenScopes'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GcpServiceAccount
	err := ctx.RegisterResource("spacelift:index/gcpServiceAccount:GcpServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGcpServiceAccount gets an existing GcpServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGcpServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GcpServiceAccountState, opts ...pulumi.ResourceOption) (*GcpServiceAccount, error) {
	var resource GcpServiceAccount
	err := ctx.ReadResource("spacelift:index/gcpServiceAccount:GcpServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GcpServiceAccount resources.
type gcpServiceAccountState struct {
	// ID of the module which uses GCP service account credentials
	ModuleId *string `pulumi:"moduleId"`
	// Email address of the GCP service account dedicated for this stack
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// ID of the stack which uses GCP service account credentials
	StackId *string `pulumi:"stackId"`
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes []string `pulumi:"tokenScopes"`
}

type GcpServiceAccountState struct {
	// ID of the module which uses GCP service account credentials
	ModuleId pulumi.StringPtrInput
	// Email address of the GCP service account dedicated for this stack
	ServiceAccountEmail pulumi.StringPtrInput
	// ID of the stack which uses GCP service account credentials
	StackId pulumi.StringPtrInput
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes pulumi.StringArrayInput
}

func (GcpServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpServiceAccountState)(nil)).Elem()
}

type gcpServiceAccountArgs struct {
	// ID of the module which uses GCP service account credentials
	ModuleId *string `pulumi:"moduleId"`
	// ID of the stack which uses GCP service account credentials
	StackId *string `pulumi:"stackId"`
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes []string `pulumi:"tokenScopes"`
}

// The set of arguments for constructing a GcpServiceAccount resource.
type GcpServiceAccountArgs struct {
	// ID of the module which uses GCP service account credentials
	ModuleId pulumi.StringPtrInput
	// ID of the stack which uses GCP service account credentials
	StackId pulumi.StringPtrInput
	// List of scopes that will be requested when generating temporary GCP service account credentials
	TokenScopes pulumi.StringArrayInput
}

func (GcpServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpServiceAccountArgs)(nil)).Elem()
}

type GcpServiceAccountInput interface {
	pulumi.Input

	ToGcpServiceAccountOutput() GcpServiceAccountOutput
	ToGcpServiceAccountOutputWithContext(ctx context.Context) GcpServiceAccountOutput
}

func (*GcpServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**GcpServiceAccount)(nil)).Elem()
}

func (i *GcpServiceAccount) ToGcpServiceAccountOutput() GcpServiceAccountOutput {
	return i.ToGcpServiceAccountOutputWithContext(context.Background())
}

func (i *GcpServiceAccount) ToGcpServiceAccountOutputWithContext(ctx context.Context) GcpServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpServiceAccountOutput)
}

func (i *GcpServiceAccount) ToOutput(ctx context.Context) pulumix.Output[*GcpServiceAccount] {
	return pulumix.Output[*GcpServiceAccount]{
		OutputState: i.ToGcpServiceAccountOutputWithContext(ctx).OutputState,
	}
}

// GcpServiceAccountArrayInput is an input type that accepts GcpServiceAccountArray and GcpServiceAccountArrayOutput values.
// You can construct a concrete instance of `GcpServiceAccountArrayInput` via:
//
//	GcpServiceAccountArray{ GcpServiceAccountArgs{...} }
type GcpServiceAccountArrayInput interface {
	pulumi.Input

	ToGcpServiceAccountArrayOutput() GcpServiceAccountArrayOutput
	ToGcpServiceAccountArrayOutputWithContext(context.Context) GcpServiceAccountArrayOutput
}

type GcpServiceAccountArray []GcpServiceAccountInput

func (GcpServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GcpServiceAccount)(nil)).Elem()
}

func (i GcpServiceAccountArray) ToGcpServiceAccountArrayOutput() GcpServiceAccountArrayOutput {
	return i.ToGcpServiceAccountArrayOutputWithContext(context.Background())
}

func (i GcpServiceAccountArray) ToGcpServiceAccountArrayOutputWithContext(ctx context.Context) GcpServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpServiceAccountArrayOutput)
}

func (i GcpServiceAccountArray) ToOutput(ctx context.Context) pulumix.Output[[]*GcpServiceAccount] {
	return pulumix.Output[[]*GcpServiceAccount]{
		OutputState: i.ToGcpServiceAccountArrayOutputWithContext(ctx).OutputState,
	}
}

// GcpServiceAccountMapInput is an input type that accepts GcpServiceAccountMap and GcpServiceAccountMapOutput values.
// You can construct a concrete instance of `GcpServiceAccountMapInput` via:
//
//	GcpServiceAccountMap{ "key": GcpServiceAccountArgs{...} }
type GcpServiceAccountMapInput interface {
	pulumi.Input

	ToGcpServiceAccountMapOutput() GcpServiceAccountMapOutput
	ToGcpServiceAccountMapOutputWithContext(context.Context) GcpServiceAccountMapOutput
}

type GcpServiceAccountMap map[string]GcpServiceAccountInput

func (GcpServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GcpServiceAccount)(nil)).Elem()
}

func (i GcpServiceAccountMap) ToGcpServiceAccountMapOutput() GcpServiceAccountMapOutput {
	return i.ToGcpServiceAccountMapOutputWithContext(context.Background())
}

func (i GcpServiceAccountMap) ToGcpServiceAccountMapOutputWithContext(ctx context.Context) GcpServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpServiceAccountMapOutput)
}

func (i GcpServiceAccountMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GcpServiceAccount] {
	return pulumix.Output[map[string]*GcpServiceAccount]{
		OutputState: i.ToGcpServiceAccountMapOutputWithContext(ctx).OutputState,
	}
}

type GcpServiceAccountOutput struct{ *pulumi.OutputState }

func (GcpServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcpServiceAccount)(nil)).Elem()
}

func (o GcpServiceAccountOutput) ToGcpServiceAccountOutput() GcpServiceAccountOutput {
	return o
}

func (o GcpServiceAccountOutput) ToGcpServiceAccountOutputWithContext(ctx context.Context) GcpServiceAccountOutput {
	return o
}

func (o GcpServiceAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*GcpServiceAccount] {
	return pulumix.Output[*GcpServiceAccount]{
		OutputState: o.OutputState,
	}
}

// ID of the module which uses GCP service account credentials
func (o GcpServiceAccountOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcpServiceAccount) pulumi.StringPtrOutput { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// Email address of the GCP service account dedicated for this stack
func (o GcpServiceAccountOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *GcpServiceAccount) pulumi.StringOutput { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

// ID of the stack which uses GCP service account credentials
func (o GcpServiceAccountOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcpServiceAccount) pulumi.StringPtrOutput { return v.StackId }).(pulumi.StringPtrOutput)
}

// List of scopes that will be requested when generating temporary GCP service account credentials
func (o GcpServiceAccountOutput) TokenScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GcpServiceAccount) pulumi.StringArrayOutput { return v.TokenScopes }).(pulumi.StringArrayOutput)
}

type GcpServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (GcpServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GcpServiceAccount)(nil)).Elem()
}

func (o GcpServiceAccountArrayOutput) ToGcpServiceAccountArrayOutput() GcpServiceAccountArrayOutput {
	return o
}

func (o GcpServiceAccountArrayOutput) ToGcpServiceAccountArrayOutputWithContext(ctx context.Context) GcpServiceAccountArrayOutput {
	return o
}

func (o GcpServiceAccountArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GcpServiceAccount] {
	return pulumix.Output[[]*GcpServiceAccount]{
		OutputState: o.OutputState,
	}
}

func (o GcpServiceAccountArrayOutput) Index(i pulumi.IntInput) GcpServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GcpServiceAccount {
		return vs[0].([]*GcpServiceAccount)[vs[1].(int)]
	}).(GcpServiceAccountOutput)
}

type GcpServiceAccountMapOutput struct{ *pulumi.OutputState }

func (GcpServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GcpServiceAccount)(nil)).Elem()
}

func (o GcpServiceAccountMapOutput) ToGcpServiceAccountMapOutput() GcpServiceAccountMapOutput {
	return o
}

func (o GcpServiceAccountMapOutput) ToGcpServiceAccountMapOutputWithContext(ctx context.Context) GcpServiceAccountMapOutput {
	return o
}

func (o GcpServiceAccountMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GcpServiceAccount] {
	return pulumix.Output[map[string]*GcpServiceAccount]{
		OutputState: o.OutputState,
	}
}

func (o GcpServiceAccountMapOutput) MapIndex(k pulumi.StringInput) GcpServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GcpServiceAccount {
		return vs[0].(map[string]*GcpServiceAccount)[vs[1].(string)]
	}).(GcpServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GcpServiceAccountInput)(nil)).Elem(), &GcpServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcpServiceAccountArrayInput)(nil)).Elem(), GcpServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcpServiceAccountMapInput)(nil)).Elem(), GcpServiceAccountMap{})
	pulumi.RegisterOutputType(GcpServiceAccountOutput{})
	pulumi.RegisterOutputType(GcpServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(GcpServiceAccountMapOutput{})
}
