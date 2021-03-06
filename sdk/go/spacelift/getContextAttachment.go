// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupContextAttachment(ctx *pulumi.Context, args *LookupContextAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupContextAttachmentResult, error) {
	var rv LookupContextAttachmentResult
	err := ctx.Invoke("spacelift:index/getContextAttachment:getContextAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContextAttachment.
type LookupContextAttachmentArgs struct {
	ContextId string  `pulumi:"contextId"`
	ModuleId  *string `pulumi:"moduleId"`
	StackId   *string `pulumi:"stackId"`
}

// A collection of values returned by getContextAttachment.
type LookupContextAttachmentResult struct {
	ContextId string `pulumi:"contextId"`
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	ModuleId *string `pulumi:"moduleId"`
	Priority int     `pulumi:"priority"`
	StackId  *string `pulumi:"stackId"`
}
