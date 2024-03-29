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

// `Mountedfile` represents a file mounted in each Run's workspace that is part of a configuration of a context (`Context`), stack (`Stack`) or a module (`Module`). In principle, it's very similar to an environment variable (`EnvironmentVariable`) except that the value is written to the filesystem rather than passed to the environment.
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
//			_, err := spacelift.LookupMountedfile(ctx, &spacelift.LookupMountedfileArgs{
//				ContextId:    pulumi.StringRef("prod-k8s-ie"),
//				RelativePath: "kubeconfig",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.LookupMountedfile(ctx, &spacelift.LookupMountedfileArgs{
//				ModuleId:     pulumi.StringRef("k8s-module"),
//				RelativePath: "kubeconfig",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.LookupMountedfile(ctx, &spacelift.LookupMountedfileArgs{
//				RelativePath: "kubeconfig",
//				StackId:      pulumi.StringRef("k8s-core"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupMountedfile(ctx *pulumi.Context, args *LookupMountedfileArgs, opts ...pulumi.InvokeOption) (*LookupMountedfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMountedfileResult
	err := ctx.Invoke("spacelift:index/getMountedfile:getMountedfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMountedfile.
type LookupMountedfileArgs struct {
	// ID of the context where the mounted file is stored
	ContextId *string `pulumi:"contextId"`
	// ID of the module where the mounted file is stored
	ModuleId *string `pulumi:"moduleId"`
	// relative path to the mounted file
	RelativePath string `pulumi:"relativePath"`
	// ID of the stack where the mounted file is stored
	StackId *string `pulumi:"stackId"`
}

// A collection of values returned by getMountedfile.
type LookupMountedfileResult struct {
	// SHA-256 checksum of the value
	Checksum string `pulumi:"checksum"`
	// content of the mounted file encoded using Base-64
	Content string `pulumi:"content"`
	// ID of the context where the mounted file is stored
	ContextId *string `pulumi:"contextId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the module where the mounted file is stored
	ModuleId *string `pulumi:"moduleId"`
	// relative path to the mounted file
	RelativePath string `pulumi:"relativePath"`
	// ID of the stack where the mounted file is stored
	StackId *string `pulumi:"stackId"`
	// indicates whether the value can be read back outside a Run
	WriteOnly bool `pulumi:"writeOnly"`
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
	// ID of the context where the mounted file is stored
	ContextId pulumi.StringPtrInput `pulumi:"contextId"`
	// ID of the module where the mounted file is stored
	ModuleId pulumi.StringPtrInput `pulumi:"moduleId"`
	// relative path to the mounted file
	RelativePath pulumi.StringInput `pulumi:"relativePath"`
	// ID of the stack where the mounted file is stored
	StackId pulumi.StringPtrInput `pulumi:"stackId"`
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

func (o LookupMountedfileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMountedfileResult] {
	return pulumix.Output[LookupMountedfileResult]{
		OutputState: o.OutputState,
	}
}

// SHA-256 checksum of the value
func (o LookupMountedfileResultOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.Checksum }).(pulumi.StringOutput)
}

// content of the mounted file encoded using Base-64
func (o LookupMountedfileResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.Content }).(pulumi.StringOutput)
}

// ID of the context where the mounted file is stored
func (o LookupMountedfileResultOutput) ContextId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMountedfileResult) *string { return v.ContextId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMountedfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the module where the mounted file is stored
func (o LookupMountedfileResultOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMountedfileResult) *string { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// relative path to the mounted file
func (o LookupMountedfileResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountedfileResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

// ID of the stack where the mounted file is stored
func (o LookupMountedfileResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMountedfileResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

// indicates whether the value can be read back outside a Run
func (o LookupMountedfileResultOutput) WriteOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMountedfileResult) bool { return v.WriteOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMountedfileResultOutput{})
}
