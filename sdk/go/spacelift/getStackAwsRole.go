// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift/internal"
)

// > **Note:** `StackAwsRole` is deprecated. Please use `AwsRole` instead. The functionality is identical.
//
// `StackAwsRole` represents [cross-account IAM role delegation](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html) between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.
//
// If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).
//
// Note: when assuming credentials for **shared worker**, Spacelift will use `$accountName@$stackID` or `$accountName@$moduleID` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole).
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
//			_, err := spacelift.LookupStackAwsRole(ctx, &spacelift.LookupStackAwsRoleArgs{
//				ModuleId: pulumi.StringRef("k8s-module"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.LookupStackAwsRole(ctx, &spacelift.LookupStackAwsRoleArgs{
//				StackId: pulumi.StringRef("k8s-core"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupStackAwsRole(ctx *pulumi.Context, args *LookupStackAwsRoleArgs, opts ...pulumi.InvokeOption) (*LookupStackAwsRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStackAwsRoleResult
	err := ctx.Invoke("spacelift:index/getStackAwsRole:getStackAwsRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStackAwsRole.
type LookupStackAwsRoleArgs struct {
	// ID of the module which assumes the AWS IAM role
	ModuleId *string `pulumi:"moduleId"`
	// ID of the stack which assumes the AWS IAM role
	StackId *string `pulumi:"stackId"`
}

// A collection of values returned by getStackAwsRole.
type LookupStackAwsRoleResult struct {
	// AWS IAM role session duration in seconds
	DurationSeconds int `pulumi:"durationSeconds"`
	// Custom external ID (works only for private workers).
	ExternalId string `pulumi:"externalId"`
	// Generate AWS credentials in the private worker
	GenerateCredentialsInWorker bool `pulumi:"generateCredentialsInWorker"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the module which assumes the AWS IAM role
	ModuleId *string `pulumi:"moduleId"`
	// ARN of the AWS IAM role to attach
	RoleArn string `pulumi:"roleArn"`
	// ID of the stack which assumes the AWS IAM role
	StackId *string `pulumi:"stackId"`
}

func LookupStackAwsRoleOutput(ctx *pulumi.Context, args LookupStackAwsRoleOutputArgs, opts ...pulumi.InvokeOption) LookupStackAwsRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackAwsRoleResult, error) {
			args := v.(LookupStackAwsRoleArgs)
			r, err := LookupStackAwsRole(ctx, &args, opts...)
			var s LookupStackAwsRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackAwsRoleResultOutput)
}

// A collection of arguments for invoking getStackAwsRole.
type LookupStackAwsRoleOutputArgs struct {
	// ID of the module which assumes the AWS IAM role
	ModuleId pulumi.StringPtrInput `pulumi:"moduleId"`
	// ID of the stack which assumes the AWS IAM role
	StackId pulumi.StringPtrInput `pulumi:"stackId"`
}

func (LookupStackAwsRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackAwsRoleArgs)(nil)).Elem()
}

// A collection of values returned by getStackAwsRole.
type LookupStackAwsRoleResultOutput struct{ *pulumi.OutputState }

func (LookupStackAwsRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackAwsRoleResult)(nil)).Elem()
}

func (o LookupStackAwsRoleResultOutput) ToLookupStackAwsRoleResultOutput() LookupStackAwsRoleResultOutput {
	return o
}

func (o LookupStackAwsRoleResultOutput) ToLookupStackAwsRoleResultOutputWithContext(ctx context.Context) LookupStackAwsRoleResultOutput {
	return o
}

// AWS IAM role session duration in seconds
func (o LookupStackAwsRoleResultOutput) DurationSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStackAwsRoleResult) int { return v.DurationSeconds }).(pulumi.IntOutput)
}

// Custom external ID (works only for private workers).
func (o LookupStackAwsRoleResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackAwsRoleResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// Generate AWS credentials in the private worker
func (o LookupStackAwsRoleResultOutput) GenerateCredentialsInWorker() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStackAwsRoleResult) bool { return v.GenerateCredentialsInWorker }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStackAwsRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackAwsRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the module which assumes the AWS IAM role
func (o LookupStackAwsRoleResultOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackAwsRoleResult) *string { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// ARN of the AWS IAM role to attach
func (o LookupStackAwsRoleResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackAwsRoleResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

// ID of the stack which assumes the AWS IAM role
func (o LookupStackAwsRoleResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackAwsRoleResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackAwsRoleResultOutput{})
}
