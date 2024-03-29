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

// `getContexts` represents all the contexts in the Spacelift account visible to the API user.
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
//			_, err := spacelift.GetContexts(ctx, &spacelift.GetContextsArgs{
//				Labels: []spacelift.GetContextsLabel{
//					{
//						AnyOfs: []string{
//							"foo",
//							"bar",
//						},
//					},
//					{
//						AnyOfs: []string{
//							"baz",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetContexts(ctx *pulumi.Context, args *GetContextsArgs, opts ...pulumi.InvokeOption) (*GetContextsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetContextsResult
	err := ctx.Invoke("spacelift:index/getContexts:getContexts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContexts.
type GetContextsArgs struct {
	Labels []GetContextsLabel `pulumi:"labels"`
}

// A collection of values returned by getContexts.
type GetContextsResult struct {
	Contexts []GetContextsContext `pulumi:"contexts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Require contexts to have one of the labels
	Labels []GetContextsLabel `pulumi:"labels"`
}

func GetContextsOutput(ctx *pulumi.Context, args GetContextsOutputArgs, opts ...pulumi.InvokeOption) GetContextsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContextsResult, error) {
			args := v.(GetContextsArgs)
			r, err := GetContexts(ctx, &args, opts...)
			var s GetContextsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContextsResultOutput)
}

// A collection of arguments for invoking getContexts.
type GetContextsOutputArgs struct {
	Labels GetContextsLabelArrayInput `pulumi:"labels"`
}

func (GetContextsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContextsArgs)(nil)).Elem()
}

// A collection of values returned by getContexts.
type GetContextsResultOutput struct{ *pulumi.OutputState }

func (GetContextsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContextsResult)(nil)).Elem()
}

func (o GetContextsResultOutput) ToGetContextsResultOutput() GetContextsResultOutput {
	return o
}

func (o GetContextsResultOutput) ToGetContextsResultOutputWithContext(ctx context.Context) GetContextsResultOutput {
	return o
}

func (o GetContextsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetContextsResult] {
	return pulumix.Output[GetContextsResult]{
		OutputState: o.OutputState,
	}
}

func (o GetContextsResultOutput) Contexts() GetContextsContextArrayOutput {
	return o.ApplyT(func(v GetContextsResult) []GetContextsContext { return v.Contexts }).(GetContextsContextArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetContextsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContextsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Require contexts to have one of the labels
func (o GetContextsResultOutput) Labels() GetContextsLabelArrayOutput {
	return o.ApplyT(func(v GetContextsResult) []GetContextsLabel { return v.Labels }).(GetContextsLabelArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContextsResultOutput{})
}
