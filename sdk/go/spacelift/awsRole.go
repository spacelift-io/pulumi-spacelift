// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `AwsRole` represents [cross-account IAM role delegation](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.
//
// If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).
//
// Note: when assuming credentials for **shared worker**, Spacelift will use `$accountName@$stackID` or `$accountName@$moduleID` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole).
//
// ## Import
//
// ```sh
//  $ pulumi import spacelift:index/awsRole:AwsRole k8s-core stack/$STACK_ID
// ```
//
// ```sh
//  $ pulumi import spacelift:index/awsRole:AwsRole k8s-core module/$MODULE_ID
// ```
type AwsRole struct {
	pulumi.CustomResourceState

	// AWS IAM role session duration in seconds
	DurationSeconds pulumi.IntOutput `pulumi:"durationSeconds"`
	// Custom external ID (works only for private workers).
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// Generate AWS credentials in the private worker. Defaults to `false`.
	GenerateCredentialsInWorker pulumi.BoolPtrOutput `pulumi:"generateCredentialsInWorker"`
	// ID of the module which assumes the AWS IAM role
	ModuleId pulumi.StringPtrOutput `pulumi:"moduleId"`
	// ARN of the AWS IAM role to attach
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// ID of the stack which assumes the AWS IAM role
	StackId pulumi.StringPtrOutput `pulumi:"stackId"`
}

// NewAwsRole registers a new resource with the given unique name, arguments, and options.
func NewAwsRole(ctx *pulumi.Context,
	name string, args *AwsRoleArgs, opts ...pulumi.ResourceOption) (*AwsRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AwsRole
	err := ctx.RegisterResource("spacelift:index/awsRole:AwsRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsRole gets an existing AwsRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsRoleState, opts ...pulumi.ResourceOption) (*AwsRole, error) {
	var resource AwsRole
	err := ctx.ReadResource("spacelift:index/awsRole:AwsRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsRole resources.
type awsRoleState struct {
	// AWS IAM role session duration in seconds
	DurationSeconds *int `pulumi:"durationSeconds"`
	// Custom external ID (works only for private workers).
	ExternalId *string `pulumi:"externalId"`
	// Generate AWS credentials in the private worker. Defaults to `false`.
	GenerateCredentialsInWorker *bool `pulumi:"generateCredentialsInWorker"`
	// ID of the module which assumes the AWS IAM role
	ModuleId *string `pulumi:"moduleId"`
	// ARN of the AWS IAM role to attach
	RoleArn *string `pulumi:"roleArn"`
	// ID of the stack which assumes the AWS IAM role
	StackId *string `pulumi:"stackId"`
}

type AwsRoleState struct {
	// AWS IAM role session duration in seconds
	DurationSeconds pulumi.IntPtrInput
	// Custom external ID (works only for private workers).
	ExternalId pulumi.StringPtrInput
	// Generate AWS credentials in the private worker. Defaults to `false`.
	GenerateCredentialsInWorker pulumi.BoolPtrInput
	// ID of the module which assumes the AWS IAM role
	ModuleId pulumi.StringPtrInput
	// ARN of the AWS IAM role to attach
	RoleArn pulumi.StringPtrInput
	// ID of the stack which assumes the AWS IAM role
	StackId pulumi.StringPtrInput
}

func (AwsRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsRoleState)(nil)).Elem()
}

type awsRoleArgs struct {
	// AWS IAM role session duration in seconds
	DurationSeconds *int `pulumi:"durationSeconds"`
	// Custom external ID (works only for private workers).
	ExternalId *string `pulumi:"externalId"`
	// Generate AWS credentials in the private worker. Defaults to `false`.
	GenerateCredentialsInWorker *bool `pulumi:"generateCredentialsInWorker"`
	// ID of the module which assumes the AWS IAM role
	ModuleId *string `pulumi:"moduleId"`
	// ARN of the AWS IAM role to attach
	RoleArn string `pulumi:"roleArn"`
	// ID of the stack which assumes the AWS IAM role
	StackId *string `pulumi:"stackId"`
}

// The set of arguments for constructing a AwsRole resource.
type AwsRoleArgs struct {
	// AWS IAM role session duration in seconds
	DurationSeconds pulumi.IntPtrInput
	// Custom external ID (works only for private workers).
	ExternalId pulumi.StringPtrInput
	// Generate AWS credentials in the private worker. Defaults to `false`.
	GenerateCredentialsInWorker pulumi.BoolPtrInput
	// ID of the module which assumes the AWS IAM role
	ModuleId pulumi.StringPtrInput
	// ARN of the AWS IAM role to attach
	RoleArn pulumi.StringInput
	// ID of the stack which assumes the AWS IAM role
	StackId pulumi.StringPtrInput
}

func (AwsRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsRoleArgs)(nil)).Elem()
}

type AwsRoleInput interface {
	pulumi.Input

	ToAwsRoleOutput() AwsRoleOutput
	ToAwsRoleOutputWithContext(ctx context.Context) AwsRoleOutput
}

func (*AwsRole) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsRole)(nil)).Elem()
}

func (i *AwsRole) ToAwsRoleOutput() AwsRoleOutput {
	return i.ToAwsRoleOutputWithContext(context.Background())
}

func (i *AwsRole) ToAwsRoleOutputWithContext(ctx context.Context) AwsRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsRoleOutput)
}

// AwsRoleArrayInput is an input type that accepts AwsRoleArray and AwsRoleArrayOutput values.
// You can construct a concrete instance of `AwsRoleArrayInput` via:
//
//          AwsRoleArray{ AwsRoleArgs{...} }
type AwsRoleArrayInput interface {
	pulumi.Input

	ToAwsRoleArrayOutput() AwsRoleArrayOutput
	ToAwsRoleArrayOutputWithContext(context.Context) AwsRoleArrayOutput
}

type AwsRoleArray []AwsRoleInput

func (AwsRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsRole)(nil)).Elem()
}

func (i AwsRoleArray) ToAwsRoleArrayOutput() AwsRoleArrayOutput {
	return i.ToAwsRoleArrayOutputWithContext(context.Background())
}

func (i AwsRoleArray) ToAwsRoleArrayOutputWithContext(ctx context.Context) AwsRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsRoleArrayOutput)
}

// AwsRoleMapInput is an input type that accepts AwsRoleMap and AwsRoleMapOutput values.
// You can construct a concrete instance of `AwsRoleMapInput` via:
//
//          AwsRoleMap{ "key": AwsRoleArgs{...} }
type AwsRoleMapInput interface {
	pulumi.Input

	ToAwsRoleMapOutput() AwsRoleMapOutput
	ToAwsRoleMapOutputWithContext(context.Context) AwsRoleMapOutput
}

type AwsRoleMap map[string]AwsRoleInput

func (AwsRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsRole)(nil)).Elem()
}

func (i AwsRoleMap) ToAwsRoleMapOutput() AwsRoleMapOutput {
	return i.ToAwsRoleMapOutputWithContext(context.Background())
}

func (i AwsRoleMap) ToAwsRoleMapOutputWithContext(ctx context.Context) AwsRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsRoleMapOutput)
}

type AwsRoleOutput struct{ *pulumi.OutputState }

func (AwsRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsRole)(nil)).Elem()
}

func (o AwsRoleOutput) ToAwsRoleOutput() AwsRoleOutput {
	return o
}

func (o AwsRoleOutput) ToAwsRoleOutputWithContext(ctx context.Context) AwsRoleOutput {
	return o
}

// AWS IAM role session duration in seconds
func (o AwsRoleOutput) DurationSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *AwsRole) pulumi.IntOutput { return v.DurationSeconds }).(pulumi.IntOutput)
}

// Custom external ID (works only for private workers).
func (o AwsRoleOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsRole) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// Generate AWS credentials in the private worker. Defaults to `false`.
func (o AwsRoleOutput) GenerateCredentialsInWorker() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AwsRole) pulumi.BoolPtrOutput { return v.GenerateCredentialsInWorker }).(pulumi.BoolPtrOutput)
}

// ID of the module which assumes the AWS IAM role
func (o AwsRoleOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsRole) pulumi.StringPtrOutput { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// ARN of the AWS IAM role to attach
func (o AwsRoleOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsRole) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// ID of the stack which assumes the AWS IAM role
func (o AwsRoleOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsRole) pulumi.StringPtrOutput { return v.StackId }).(pulumi.StringPtrOutput)
}

type AwsRoleArrayOutput struct{ *pulumi.OutputState }

func (AwsRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsRole)(nil)).Elem()
}

func (o AwsRoleArrayOutput) ToAwsRoleArrayOutput() AwsRoleArrayOutput {
	return o
}

func (o AwsRoleArrayOutput) ToAwsRoleArrayOutputWithContext(ctx context.Context) AwsRoleArrayOutput {
	return o
}

func (o AwsRoleArrayOutput) Index(i pulumi.IntInput) AwsRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsRole {
		return vs[0].([]*AwsRole)[vs[1].(int)]
	}).(AwsRoleOutput)
}

type AwsRoleMapOutput struct{ *pulumi.OutputState }

func (AwsRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsRole)(nil)).Elem()
}

func (o AwsRoleMapOutput) ToAwsRoleMapOutput() AwsRoleMapOutput {
	return o
}

func (o AwsRoleMapOutput) ToAwsRoleMapOutputWithContext(ctx context.Context) AwsRoleMapOutput {
	return o
}

func (o AwsRoleMapOutput) MapIndex(k pulumi.StringInput) AwsRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsRole {
		return vs[0].(map[string]*AwsRole)[vs[1].(string)]
	}).(AwsRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsRoleInput)(nil)).Elem(), &AwsRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsRoleArrayInput)(nil)).Elem(), AwsRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsRoleMapInput)(nil)).Elem(), AwsRoleMap{})
	pulumi.RegisterOutputType(AwsRoleOutput{})
	pulumi.RegisterOutputType(AwsRoleArrayOutput{})
	pulumi.RegisterOutputType(AwsRoleMapOutput{})
}
