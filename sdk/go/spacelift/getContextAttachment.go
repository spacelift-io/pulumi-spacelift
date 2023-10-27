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

// `ContextAttachment` represents a Spacelift attachment of a single context to a single stack or module, with a predefined priority.
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
//			_, err := spacelift.LookupContextAttachment(ctx, &spacelift.LookupContextAttachmentArgs{
//				ContextId: "prod-k8s-ie",
//				StackId:   pulumi.StringRef("apps-cluster"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = spacelift.LookupContextAttachment(ctx, &spacelift.LookupContextAttachmentArgs{
//				ContextId: "prod-k8s-ie",
//				ModuleId:  pulumi.StringRef("terraform-aws-kafka"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupContextAttachment(ctx *pulumi.Context, args *LookupContextAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupContextAttachmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContextAttachmentResult
	err := ctx.Invoke("spacelift:index/getContextAttachment:getContextAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContextAttachment.
type LookupContextAttachmentArgs struct {
	// ID of the attached context
	ContextId string `pulumi:"contextId"`
	// ID of the attached module
	ModuleId *string `pulumi:"moduleId"`
	// ID of the attached stack
	StackId *string `pulumi:"stackId"`
}

// A collection of values returned by getContextAttachment.
type LookupContextAttachmentResult struct {
	// ID of the attached context
	ContextId string `pulumi:"contextId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the attached module
	ModuleId *string `pulumi:"moduleId"`
	// Priority of the context attachment. All the contexts attached to a stack are sorted by priority (lowest first), though values don't need to be unique. This ordering establishes precedence rules between contexts should there be a conflict and multiple contexts define the same value.
	Priority int `pulumi:"priority"`
	// ID of the attached stack
	StackId *string `pulumi:"stackId"`
}

func LookupContextAttachmentOutput(ctx *pulumi.Context, args LookupContextAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupContextAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContextAttachmentResult, error) {
			args := v.(LookupContextAttachmentArgs)
			r, err := LookupContextAttachment(ctx, &args, opts...)
			var s LookupContextAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContextAttachmentResultOutput)
}

// A collection of arguments for invoking getContextAttachment.
type LookupContextAttachmentOutputArgs struct {
	// ID of the attached context
	ContextId pulumi.StringInput `pulumi:"contextId"`
	// ID of the attached module
	ModuleId pulumi.StringPtrInput `pulumi:"moduleId"`
	// ID of the attached stack
	StackId pulumi.StringPtrInput `pulumi:"stackId"`
}

func (LookupContextAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContextAttachmentArgs)(nil)).Elem()
}

// A collection of values returned by getContextAttachment.
type LookupContextAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupContextAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContextAttachmentResult)(nil)).Elem()
}

func (o LookupContextAttachmentResultOutput) ToLookupContextAttachmentResultOutput() LookupContextAttachmentResultOutput {
	return o
}

func (o LookupContextAttachmentResultOutput) ToLookupContextAttachmentResultOutputWithContext(ctx context.Context) LookupContextAttachmentResultOutput {
	return o
}

func (o LookupContextAttachmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupContextAttachmentResult] {
	return pulumix.Output[LookupContextAttachmentResult]{
		OutputState: o.OutputState,
	}
}

// ID of the attached context
func (o LookupContextAttachmentResultOutput) ContextId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextAttachmentResult) string { return v.ContextId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContextAttachmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextAttachmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the attached module
func (o LookupContextAttachmentResultOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContextAttachmentResult) *string { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// Priority of the context attachment. All the contexts attached to a stack are sorted by priority (lowest first), though values don't need to be unique. This ordering establishes precedence rules between contexts should there be a conflict and multiple contexts define the same value.
func (o LookupContextAttachmentResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContextAttachmentResult) int { return v.Priority }).(pulumi.IntOutput)
}

// ID of the attached stack
func (o LookupContextAttachmentResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContextAttachmentResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContextAttachmentResultOutput{})
}
