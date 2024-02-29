// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift/internal"
)

// `EnvironmentVariable` defines an environment variable on the context (`Context`), stack (`Stack`) or a module (`Module`), thereby allowing to pass and share various secrets and configuration with Spacelift stacks.
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
//			_, err := spacelift.LookupEnvironmentVariable(ctx, &spacelift.LookupEnvironmentVariableArgs{
//				ContextId: pulumi.StringRef("prod-k8s-ie"),
//				Name:      "KUBECONFIG",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.LookupEnvironmentVariable(ctx, &spacelift.LookupEnvironmentVariableArgs{
//				ModuleId: pulumi.StringRef("k8s-module"),
//				Name:     "KUBECONFIG",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.LookupEnvironmentVariable(ctx, &spacelift.LookupEnvironmentVariableArgs{
//				Name:    "KUBECONFIG",
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
func LookupEnvironmentVariable(ctx *pulumi.Context, args *LookupEnvironmentVariableArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentVariableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentVariableResult
	err := ctx.Invoke("spacelift:index/getEnvironmentVariable:getEnvironmentVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnvironmentVariable.
type LookupEnvironmentVariableArgs struct {
	// ID of the context on which the environment variable is defined
	ContextId *string `pulumi:"contextId"`
	// ID of the module on which the environment variable is defined
	ModuleId *string `pulumi:"moduleId"`
	// name of the environment variable
	Name string `pulumi:"name"`
	// ID of the stack on which the environment variable is defined
	StackId *string `pulumi:"stackId"`
}

// A collection of values returned by getEnvironmentVariable.
type LookupEnvironmentVariableResult struct {
	// SHA-256 checksum of the value
	Checksum string `pulumi:"checksum"`
	// ID of the context on which the environment variable is defined
	ContextId *string `pulumi:"contextId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the module on which the environment variable is defined
	ModuleId *string `pulumi:"moduleId"`
	// name of the environment variable
	Name string `pulumi:"name"`
	// ID of the stack on which the environment variable is defined
	StackId *string `pulumi:"stackId"`
	// value of the environment variable
	Value string `pulumi:"value"`
	// indicates whether the value can be read back outside a Run
	WriteOnly bool `pulumi:"writeOnly"`
}

func LookupEnvironmentVariableOutput(ctx *pulumi.Context, args LookupEnvironmentVariableOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentVariableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentVariableResult, error) {
			args := v.(LookupEnvironmentVariableArgs)
			r, err := LookupEnvironmentVariable(ctx, &args, opts...)
			var s LookupEnvironmentVariableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentVariableResultOutput)
}

// A collection of arguments for invoking getEnvironmentVariable.
type LookupEnvironmentVariableOutputArgs struct {
	// ID of the context on which the environment variable is defined
	ContextId pulumi.StringPtrInput `pulumi:"contextId"`
	// ID of the module on which the environment variable is defined
	ModuleId pulumi.StringPtrInput `pulumi:"moduleId"`
	// name of the environment variable
	Name pulumi.StringInput `pulumi:"name"`
	// ID of the stack on which the environment variable is defined
	StackId pulumi.StringPtrInput `pulumi:"stackId"`
}

func (LookupEnvironmentVariableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVariableArgs)(nil)).Elem()
}

// A collection of values returned by getEnvironmentVariable.
type LookupEnvironmentVariableResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVariableResult)(nil)).Elem()
}

func (o LookupEnvironmentVariableResultOutput) ToLookupEnvironmentVariableResultOutput() LookupEnvironmentVariableResultOutput {
	return o
}

func (o LookupEnvironmentVariableResultOutput) ToLookupEnvironmentVariableResultOutputWithContext(ctx context.Context) LookupEnvironmentVariableResultOutput {
	return o
}

// SHA-256 checksum of the value
func (o LookupEnvironmentVariableResultOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) string { return v.Checksum }).(pulumi.StringOutput)
}

// ID of the context on which the environment variable is defined
func (o LookupEnvironmentVariableResultOutput) ContextId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) *string { return v.ContextId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEnvironmentVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the module on which the environment variable is defined
func (o LookupEnvironmentVariableResultOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) *string { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// name of the environment variable
func (o LookupEnvironmentVariableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the stack on which the environment variable is defined
func (o LookupEnvironmentVariableResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

// value of the environment variable
func (o LookupEnvironmentVariableResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) string { return v.Value }).(pulumi.StringOutput)
}

// indicates whether the value can be read back outside a Run
func (o LookupEnvironmentVariableResultOutput) WriteOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEnvironmentVariableResult) bool { return v.WriteOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentVariableResultOutput{})
}
