// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `AwsIntegrationAttachment` represents the attachment between a reusable AWS integration and a single stack or module.
func LookupAwsIntegrationAttachment(ctx *pulumi.Context, args *LookupAwsIntegrationAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupAwsIntegrationAttachmentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAwsIntegrationAttachmentResult
	err := ctx.Invoke("spacelift:index/getAwsIntegrationAttachment:getAwsIntegrationAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAwsIntegrationAttachment.
type LookupAwsIntegrationAttachmentArgs struct {
	IntegrationId string  `pulumi:"integrationId"`
	ModuleId      *string `pulumi:"moduleId"`
	StackId       *string `pulumi:"stackId"`
}

// A collection of values returned by getAwsIntegrationAttachment.
type LookupAwsIntegrationAttachmentResult struct {
	AttachmentId string `pulumi:"attachmentId"`
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	IntegrationId string  `pulumi:"integrationId"`
	ModuleId      *string `pulumi:"moduleId"`
	Read          bool    `pulumi:"read"`
	StackId       *string `pulumi:"stackId"`
	Write         bool    `pulumi:"write"`
}

func LookupAwsIntegrationAttachmentOutput(ctx *pulumi.Context, args LookupAwsIntegrationAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupAwsIntegrationAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAwsIntegrationAttachmentResult, error) {
			args := v.(LookupAwsIntegrationAttachmentArgs)
			r, err := LookupAwsIntegrationAttachment(ctx, &args, opts...)
			var s LookupAwsIntegrationAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAwsIntegrationAttachmentResultOutput)
}

// A collection of arguments for invoking getAwsIntegrationAttachment.
type LookupAwsIntegrationAttachmentOutputArgs struct {
	IntegrationId pulumi.StringInput    `pulumi:"integrationId"`
	ModuleId      pulumi.StringPtrInput `pulumi:"moduleId"`
	StackId       pulumi.StringPtrInput `pulumi:"stackId"`
}

func (LookupAwsIntegrationAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsIntegrationAttachmentArgs)(nil)).Elem()
}

// A collection of values returned by getAwsIntegrationAttachment.
type LookupAwsIntegrationAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupAwsIntegrationAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsIntegrationAttachmentResult)(nil)).Elem()
}

func (o LookupAwsIntegrationAttachmentResultOutput) ToLookupAwsIntegrationAttachmentResultOutput() LookupAwsIntegrationAttachmentResultOutput {
	return o
}

func (o LookupAwsIntegrationAttachmentResultOutput) ToLookupAwsIntegrationAttachmentResultOutputWithContext(ctx context.Context) LookupAwsIntegrationAttachmentResultOutput {
	return o
}

func (o LookupAwsIntegrationAttachmentResultOutput) AttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationAttachmentResult) string { return v.AttachmentId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAwsIntegrationAttachmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationAttachmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAwsIntegrationAttachmentResultOutput) IntegrationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationAttachmentResult) string { return v.IntegrationId }).(pulumi.StringOutput)
}

func (o LookupAwsIntegrationAttachmentResultOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsIntegrationAttachmentResult) *string { return v.ModuleId }).(pulumi.StringPtrOutput)
}

func (o LookupAwsIntegrationAttachmentResultOutput) Read() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAwsIntegrationAttachmentResult) bool { return v.Read }).(pulumi.BoolOutput)
}

func (o LookupAwsIntegrationAttachmentResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsIntegrationAttachmentResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

func (o LookupAwsIntegrationAttachmentResultOutput) Write() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAwsIntegrationAttachmentResult) bool { return v.Write }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAwsIntegrationAttachmentResultOutput{})
}