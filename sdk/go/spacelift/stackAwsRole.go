// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **Note:** `StackAwsRole` is deprecated. Please use `AwsRole` instead. The functionality is identical.
//
// `StackAwsRole` represents [cross-account IAM role delegation](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.
//
// If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).
//
// Note: when assuming credentials for **shared worker**, Spacelift will use `$accountName@$stackID` or `$accountName@$moduleID` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and Run ID as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole).
//
// ## Schema
//
// ### Required
//
// - **role_arn** (String) ARN of the AWS IAM role to attach
//
// ### Optional
//
// - **external_id** (String) Custom external ID (works only for private workers).
// - **generate_credentials_in_worker** (Boolean) Generate AWS credentials in the private worker
// - **id** (String) The ID of this resource.
// - **module_id** (String) ID of the module which assumes the AWS IAM role
// - **stack_id** (String) ID of the stack which assumes the AWS IAM role
type StackAwsRole struct {
	pulumi.CustomResourceState

	// Custom external ID (works only for private workers).
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// Generate AWS credentials in the private worker
	GenerateCredentialsInWorker pulumi.BoolPtrOutput `pulumi:"generateCredentialsInWorker"`
	// ID of the module which assumes the AWS IAM role
	ModuleId pulumi.StringPtrOutput `pulumi:"moduleId"`
	// ARN of the AWS IAM role to attach
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// ID of the stack which assumes the AWS IAM role
	StackId pulumi.StringPtrOutput `pulumi:"stackId"`
}

// NewStackAwsRole registers a new resource with the given unique name, arguments, and options.
func NewStackAwsRole(ctx *pulumi.Context,
	name string, args *StackAwsRoleArgs, opts ...pulumi.ResourceOption) (*StackAwsRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource StackAwsRole
	err := ctx.RegisterResource("spacelift:index/stackAwsRole:StackAwsRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStackAwsRole gets an existing StackAwsRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStackAwsRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackAwsRoleState, opts ...pulumi.ResourceOption) (*StackAwsRole, error) {
	var resource StackAwsRole
	err := ctx.ReadResource("spacelift:index/stackAwsRole:StackAwsRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StackAwsRole resources.
type stackAwsRoleState struct {
	// Custom external ID (works only for private workers).
	ExternalId *string `pulumi:"externalId"`
	// Generate AWS credentials in the private worker
	GenerateCredentialsInWorker *bool `pulumi:"generateCredentialsInWorker"`
	// ID of the module which assumes the AWS IAM role
	ModuleId *string `pulumi:"moduleId"`
	// ARN of the AWS IAM role to attach
	RoleArn *string `pulumi:"roleArn"`
	// ID of the stack which assumes the AWS IAM role
	StackId *string `pulumi:"stackId"`
}

type StackAwsRoleState struct {
	// Custom external ID (works only for private workers).
	ExternalId pulumi.StringPtrInput
	// Generate AWS credentials in the private worker
	GenerateCredentialsInWorker pulumi.BoolPtrInput
	// ID of the module which assumes the AWS IAM role
	ModuleId pulumi.StringPtrInput
	// ARN of the AWS IAM role to attach
	RoleArn pulumi.StringPtrInput
	// ID of the stack which assumes the AWS IAM role
	StackId pulumi.StringPtrInput
}

func (StackAwsRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackAwsRoleState)(nil)).Elem()
}

type stackAwsRoleArgs struct {
	// Custom external ID (works only for private workers).
	ExternalId *string `pulumi:"externalId"`
	// Generate AWS credentials in the private worker
	GenerateCredentialsInWorker *bool `pulumi:"generateCredentialsInWorker"`
	// ID of the module which assumes the AWS IAM role
	ModuleId *string `pulumi:"moduleId"`
	// ARN of the AWS IAM role to attach
	RoleArn string `pulumi:"roleArn"`
	// ID of the stack which assumes the AWS IAM role
	StackId *string `pulumi:"stackId"`
}

// The set of arguments for constructing a StackAwsRole resource.
type StackAwsRoleArgs struct {
	// Custom external ID (works only for private workers).
	ExternalId pulumi.StringPtrInput
	// Generate AWS credentials in the private worker
	GenerateCredentialsInWorker pulumi.BoolPtrInput
	// ID of the module which assumes the AWS IAM role
	ModuleId pulumi.StringPtrInput
	// ARN of the AWS IAM role to attach
	RoleArn pulumi.StringInput
	// ID of the stack which assumes the AWS IAM role
	StackId pulumi.StringPtrInput
}

func (StackAwsRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackAwsRoleArgs)(nil)).Elem()
}

type StackAwsRoleInput interface {
	pulumi.Input

	ToStackAwsRoleOutput() StackAwsRoleOutput
	ToStackAwsRoleOutputWithContext(ctx context.Context) StackAwsRoleOutput
}

func (*StackAwsRole) ElementType() reflect.Type {
	return reflect.TypeOf((*StackAwsRole)(nil))
}

func (i *StackAwsRole) ToStackAwsRoleOutput() StackAwsRoleOutput {
	return i.ToStackAwsRoleOutputWithContext(context.Background())
}

func (i *StackAwsRole) ToStackAwsRoleOutputWithContext(ctx context.Context) StackAwsRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackAwsRoleOutput)
}

type StackAwsRoleOutput struct {
	*pulumi.OutputState
}

func (StackAwsRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackAwsRole)(nil))
}

func (o StackAwsRoleOutput) ToStackAwsRoleOutput() StackAwsRoleOutput {
	return o
}

func (o StackAwsRoleOutput) ToStackAwsRoleOutputWithContext(ctx context.Context) StackAwsRoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StackAwsRoleOutput{})
}
