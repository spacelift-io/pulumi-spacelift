// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Mountedfile` represents a file mounted in each Run's workspace that is part of a configuration of a context (`Context`), stack (`Stack`) or a module (`Module`). In principle, it's very similar to an environment variable (`EnvironmentVariable`) except that the value is written to the filesystem rather than passed to the environment.
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
// 		_, err := spacelift.LookupMountedfile(ctx, &GetMountedfileArgs{
// 			ContextId:    pulumi.StringRef("prod-k8s-ie"),
// 			RelativePath: "kubeconfig",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = spacelift.LookupMountedfile(ctx, &GetMountedfileArgs{
// 			ModuleId:     pulumi.StringRef("k8s-module"),
// 			RelativePath: "kubeconfig",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = spacelift.LookupMountedfile(ctx, &GetMountedfileArgs{
// 			RelativePath: "kubeconfig",
// 			StackId:      pulumi.StringRef("k8s-core"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMountedfile(ctx *pulumi.Context, args *LookupMountedfileArgs, opts ...pulumi.InvokeOption) (*LookupMountedfileResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMountedfileResult
	err := ctx.Invoke("spacelift:index/getMountedfile:getMountedfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMountedfile.
type LookupMountedfileArgs struct {
	ContextId    *string `pulumi:"contextId"`
	ModuleId     *string `pulumi:"moduleId"`
	RelativePath string  `pulumi:"relativePath"`
	StackId      *string `pulumi:"stackId"`
}

// A collection of values returned by getMountedfile.
type LookupMountedfileResult struct {
	Checksum  string  `pulumi:"checksum"`
	Content   string  `pulumi:"content"`
	ContextId *string `pulumi:"contextId"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	ModuleId     *string `pulumi:"moduleId"`
	RelativePath string  `pulumi:"relativePath"`
	StackId      *string `pulumi:"stackId"`
	WriteOnly    bool    `pulumi:"writeOnly"`
}

func LookupMountedfileOutput(ctx *pulumi.Context, args LookupMountedfileOutputArgs, opts ...pulumi.InvokeOption) LookupMountedfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMountedfileResult, error) {
			args := v.(LookupMountedfileArgs)
			r, err := LookupMountedfile(ctx, &args, opts...)
			var s LookupMountedfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMountedfileResultOutput)
}

// A collection of arguments for invoking getMountedfile.
type LookupMountedfileOutputArgs struct {
	ContextId    pulumi.StringPtrInput `pulumi:"contextId"`
	ModuleId     pulumi.StringPtrInput `pulumi:"moduleId"`
	RelativePath pulumi.StringInput    `pulumi:"relativePath"`
	StackId      pulumi.StringPtrInput `pulumi:"stackId"`
}

func (LookupMountedfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMountedfileArgs)(nil)).Elem()
}

// A collection of values returned by getMountedfile.
type LookupMountedfileResultOutput struct{ *pulumi.OutputState }

func (LookupMountedfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMountedfileResult)(nil)).Elem()
}

func (o LookupMountedfileResultOutput) ToLookupMountedfileResultOutput() LookupMountedfileResultOutput {
	return o
}

func (o LookupMountedfileResultOutput) ToLookupMountedfileResultOutputWithContext(ctx context.Context) LookupMountedfileResultOutput {
	return o
}

func (o LookupMountedfileResultOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.Checksum }).(pulumi.StringOutput)
}

func (o LookupMountedfileResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.Content }).(pulumi.StringOutput)
}

func (o LookupMountedfileResultOutput) ContextId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMountedfileResult) *string { return v.ContextId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMountedfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMountedfileResultOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMountedfileResult) *string { return v.ModuleId }).(pulumi.StringPtrOutput)
}

func (o LookupMountedfileResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o LookupMountedfileResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMountedfileResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

func (o LookupMountedfileResultOutput) WriteOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMountedfileResult) bool { return v.WriteOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMountedfileResultOutput{})
}