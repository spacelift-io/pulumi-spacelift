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

// `Policy` represents a Spacelift **policy** - a collection of customer-defined rules that are applied by Spacelift at one of the decision points within the application.
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
//			policy, err := spacelift.LookupPolicy(ctx, &spacelift.LookupPolicyArgs{
//				PolicyId: spacelift_policy.Policy.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("policyBody", policy.Body)
//			return nil
//		})
//	}
//
// ```
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("spacelift:index/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyArgs struct {
	// immutable ID (slug) of the policy
	PolicyId string `pulumi:"policyId"`
}

// A collection of values returned by getPolicy.
type LookupPolicyResult struct {
	// body of the policy
	Body string `pulumi:"body"`
	// The provider-assigned unique ID for this managed resource.
	Id     string   `pulumi:"id"`
	Labels []string `pulumi:"labels"`
	// name of the policy
	Name string `pulumi:"name"`
	// immutable ID (slug) of the policy
	PolicyId string `pulumi:"policyId"`
	// ID (slug) of the space the policy is in
	SpaceId string `pulumi:"spaceId"`
	// type of the policy
	Type string `pulumi:"type"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyOutputArgs struct {
	// immutable ID (slug) of the policy
	PolicyId pulumi.StringInput `pulumi:"policyId"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy.
type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyResult] {
	return pulumix.Output[LookupPolicyResult]{
		OutputState: o.OutputState,
	}
}

// body of the policy
func (o LookupPolicyResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Body }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// name of the policy
func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// immutable ID (slug) of the policy
func (o LookupPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

// ID (slug) of the space the policy is in
func (o LookupPolicyResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// type of the policy
func (o LookupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
