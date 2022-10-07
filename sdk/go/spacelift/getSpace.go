// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Space` represents a Spacelift **space** - a collection of resources such as stacks, modules, policies, etc. Allows for more granular access control. Can have a parent space.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-spacelift/sdk/go/spacelift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/spacelift-io/pulumi-spacelift/sdk/go/spacelift"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		space, err := spacelift.LookupSpace(ctx, &GetSpaceArgs{
// 			SpaceId: spacelift_space.Space.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("spaceDescription", space.Description)
// 		return nil
// 	})
// }
// ```
func LookupSpace(ctx *pulumi.Context, args *LookupSpaceArgs, opts ...pulumi.InvokeOption) (*LookupSpaceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSpaceResult
	err := ctx.Invoke("spacelift:index/getSpace:getSpace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpace.
type LookupSpaceArgs struct {
	SpaceId string `pulumi:"spaceId"`
}

// A collection of values returned by getSpace.
type LookupSpaceResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	InheritEntities bool   `pulumi:"inheritEntities"`
	Name            string `pulumi:"name"`
	ParentSpaceId   string `pulumi:"parentSpaceId"`
	SpaceId         string `pulumi:"spaceId"`
}

func LookupSpaceOutput(ctx *pulumi.Context, args LookupSpaceOutputArgs, opts ...pulumi.InvokeOption) LookupSpaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpaceResult, error) {
			args := v.(LookupSpaceArgs)
			r, err := LookupSpace(ctx, &args, opts...)
			var s LookupSpaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSpaceResultOutput)
}

// A collection of arguments for invoking getSpace.
type LookupSpaceOutputArgs struct {
	SpaceId pulumi.StringInput `pulumi:"spaceId"`
}

func (LookupSpaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpaceArgs)(nil)).Elem()
}

// A collection of values returned by getSpace.
type LookupSpaceResultOutput struct{ *pulumi.OutputState }

func (LookupSpaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpaceResult)(nil)).Elem()
}

func (o LookupSpaceResultOutput) ToLookupSpaceResultOutput() LookupSpaceResultOutput {
	return o
}

func (o LookupSpaceResultOutput) ToLookupSpaceResultOutputWithContext(ctx context.Context) LookupSpaceResultOutput {
	return o
}

func (o LookupSpaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpaceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSpaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSpaceResultOutput) InheritEntities() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSpaceResult) bool { return v.InheritEntities }).(pulumi.BoolOutput)
}

func (o LookupSpaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSpaceResultOutput) ParentSpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpaceResult) string { return v.ParentSpaceId }).(pulumi.StringOutput)
}

func (o LookupSpaceResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpaceResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpaceResultOutput{})
}
