// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
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
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-spacelift/sdk/go/spacelift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/spacelift-io/spacelift-spacelift/sdk/go/spacelift/"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "k8s-module"
// 		_, err := spacelift.LookupStackAwsRole(ctx, &spacelift.LookupStackAwsRoleArgs{
// 			ModuleId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "k8s-core"
// 		_, err = spacelift.LookupStackAwsRole(ctx, &spacelift.LookupStackAwsRoleArgs{
// 			StackId: &opt1,
// 		}, nil)
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
// ### Optional
//
// - **id** (String) The ID of this resource.
// - **module_id** (String) ID of the module which assumes the AWS IAM role
// - **stack_id** (String) ID of the stack which assumes the AWS IAM role
//
// ### Read-Only
//
// - **external_id** (String) Custom external ID (works only for private workers).
// - **generate_credentials_in_worker** (Boolean) Generate AWS credentials in the private worker
// - **role_arn** (String) ARN of the AWS IAM role to attach
func LookupStackAwsRole(ctx *pulumi.Context, args *LookupStackAwsRoleArgs, opts ...pulumi.InvokeOption) (*LookupStackAwsRoleResult, error) {
	var rv LookupStackAwsRoleResult
	err := ctx.Invoke("spacelift:index/getStackAwsRole:getStackAwsRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStackAwsRole.
type LookupStackAwsRoleArgs struct {
	ModuleId *string `pulumi:"moduleId"`
	StackId  *string `pulumi:"stackId"`
}

// A collection of values returned by getStackAwsRole.
type LookupStackAwsRoleResult struct {
	ExternalId                  string `pulumi:"externalId"`
	GenerateCredentialsInWorker bool   `pulumi:"generateCredentialsInWorker"`
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	ModuleId *string `pulumi:"moduleId"`
	RoleArn  string  `pulumi:"roleArn"`
	StackId  *string `pulumi:"stackId"`
}
