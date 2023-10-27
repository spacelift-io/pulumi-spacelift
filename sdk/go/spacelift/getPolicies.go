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

// `getPolicies` can find all policies that have certain labels.
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
//			_, err := spacelift.GetPolicies(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.GetPolicies(ctx, &spacelift.GetPoliciesArgs{
//				Type: pulumi.StringRef("PLAN"),
//				Labels: []string{
//					"autoattach",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			var splat0 []interface{}
//			for _, val0 := range data.Spacelift_policies.This.Policies {
//				splat0 = append(splat0, val0.Id)
//			}
//			ctx.Export("policyIds", splat0)
//			return nil
//		})
//	}
//
// ```
func GetPolicies(ctx *pulumi.Context, args *GetPoliciesArgs, opts ...pulumi.InvokeOption) (*GetPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPoliciesResult
	err := ctx.Invoke("spacelift:index/getPolicies:getPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesArgs struct {
	Labels []string `pulumi:"labels"`
	Type   *string  `pulumi:"type"`
}

// A collection of values returned by getPolicies.
type GetPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// required labels to match
	Labels   []string            `pulumi:"labels"`
	Policies []GetPoliciesPolicy `pulumi:"policies"`
	// required policy type
	Type *string `pulumi:"type"`
}

func GetPoliciesOutput(ctx *pulumi.Context, args GetPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPoliciesResult, error) {
			args := v.(GetPoliciesArgs)
			r, err := GetPolicies(ctx, &args, opts...)
			var s GetPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPoliciesResultOutput)
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesOutputArgs struct {
	Labels pulumi.StringArrayInput `pulumi:"labels"`
	Type   pulumi.StringPtrInput   `pulumi:"type"`
}

func (GetPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getPolicies.
type GetPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoliciesResult)(nil)).Elem()
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutput() GetPoliciesResultOutput {
	return o
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutputWithContext(ctx context.Context) GetPoliciesResultOutput {
	return o
}

func (o GetPoliciesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetPoliciesResult] {
	return pulumix.Output[GetPoliciesResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// required labels to match
func (o GetPoliciesResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPoliciesResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o GetPoliciesResultOutput) Policies() GetPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetPoliciesResult) []GetPoliciesPolicy { return v.Policies }).(GetPoliciesPolicyArrayOutput)
}

// required policy type
func (o GetPoliciesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoliciesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPoliciesResultOutput{})
}
