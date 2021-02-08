// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupContext(ctx *pulumi.Context, args *LookupContextArgs, opts ...pulumi.InvokeOption) (*LookupContextResult, error) {
	var rv LookupContextResult
	err := ctx.Invoke("spacelift:index/getContext:getContext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContext.
type LookupContextArgs struct {
	ContextId string `pulumi:"contextId"`
}

// A collection of values returned by getContext.
type LookupContextResult struct {
	ContextId   string `pulumi:"contextId"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
