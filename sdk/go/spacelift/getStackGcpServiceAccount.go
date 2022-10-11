// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := spacelift.LookupStackGcpServiceAccount(ctx, &GetStackGcpServiceAccountArgs{
//				ModuleId: pulumi.StringRef("k8s-module"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.LookupStackGcpServiceAccount(ctx, &GetStackGcpServiceAccountArgs{
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
func LookupStackGcpServiceAccount(ctx *pulumi.Context, args *LookupStackGcpServiceAccountArgs, opts ...pulumi.InvokeOption) (*LookupStackGcpServiceAccountResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupStackGcpServiceAccountResult
	err := ctx.Invoke("spacelift:index/getStackGcpServiceAccount:getStackGcpServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStackGcpServiceAccount.
type LookupStackGcpServiceAccountArgs struct {
	// ID of the stack which uses GCP service account credentials
	ModuleId *string `pulumi:"moduleId"`
	// ID of the stack which uses GCP service account credentials
	StackId *string `pulumi:"stackId"`
}

// A collection of values returned by getStackGcpServiceAccount.
type LookupStackGcpServiceAccountResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the stack which uses GCP service account credentials
	ModuleId *string `pulumi:"moduleId"`
	// email address of the GCP service account dedicated for this stack
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// ID of the stack which uses GCP service account credentials
	StackId *string `pulumi:"stackId"`
	// list of Google API scopes
	TokenScopes []string `pulumi:"tokenScopes"`
}

func LookupStackGcpServiceAccountOutput(ctx *pulumi.Context, args LookupStackGcpServiceAccountOutputArgs, opts ...pulumi.InvokeOption) LookupStackGcpServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackGcpServiceAccountResult, error) {
			args := v.(LookupStackGcpServiceAccountArgs)
			r, err := LookupStackGcpServiceAccount(ctx, &args, opts...)
			var s LookupStackGcpServiceAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackGcpServiceAccountResultOutput)
}

// A collection of arguments for invoking getStackGcpServiceAccount.
type LookupStackGcpServiceAccountOutputArgs struct {
	// ID of the stack which uses GCP service account credentials
	ModuleId pulumi.StringPtrInput `pulumi:"moduleId"`
	// ID of the stack which uses GCP service account credentials
	StackId pulumi.StringPtrInput `pulumi:"stackId"`
}

func (LookupStackGcpServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackGcpServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getStackGcpServiceAccount.
type LookupStackGcpServiceAccountResultOutput struct{ *pulumi.OutputState }

func (LookupStackGcpServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackGcpServiceAccountResult)(nil)).Elem()
}

func (o LookupStackGcpServiceAccountResultOutput) ToLookupStackGcpServiceAccountResultOutput() LookupStackGcpServiceAccountResultOutput {
	return o
}

func (o LookupStackGcpServiceAccountResultOutput) ToLookupStackGcpServiceAccountResultOutputWithContext(ctx context.Context) LookupStackGcpServiceAccountResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStackGcpServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackGcpServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the stack which uses GCP service account credentials
func (o LookupStackGcpServiceAccountResultOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackGcpServiceAccountResult) *string { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// email address of the GCP service account dedicated for this stack
func (o LookupStackGcpServiceAccountResultOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStackGcpServiceAccountResult) string { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

// ID of the stack which uses GCP service account credentials
func (o LookupStackGcpServiceAccountResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackGcpServiceAccountResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

// list of Google API scopes
func (o LookupStackGcpServiceAccountResultOutput) TokenScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStackGcpServiceAccountResult) []string { return v.TokenScopes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackGcpServiceAccountResultOutput{})
}
