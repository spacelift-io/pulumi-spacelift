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

// `getGitlabWebhookEndpoint` returns details about Gitlab webhook endpoint
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
//			_, err := spacelift.GetGitlabWebhookEndpoint(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetGitlabWebhookEndpoint(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetGitlabWebhookEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGitlabWebhookEndpointResult
	err := ctx.Invoke("spacelift:index/getGitlabWebhookEndpoint:getGitlabWebhookEndpoint", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getGitlabWebhookEndpoint.
type GetGitlabWebhookEndpointResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Gitlab webhook endpoint address
	WebhookEndpoint string `pulumi:"webhookEndpoint"`
}

func GetGitlabWebhookEndpointOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetGitlabWebhookEndpointResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetGitlabWebhookEndpointResult, error) {
		r, err := GetGitlabWebhookEndpoint(ctx, opts...)
		var s GetGitlabWebhookEndpointResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetGitlabWebhookEndpointResultOutput)
}

// A collection of values returned by getGitlabWebhookEndpoint.
type GetGitlabWebhookEndpointResultOutput struct{ *pulumi.OutputState }

func (GetGitlabWebhookEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGitlabWebhookEndpointResult)(nil)).Elem()
}

func (o GetGitlabWebhookEndpointResultOutput) ToGetGitlabWebhookEndpointResultOutput() GetGitlabWebhookEndpointResultOutput {
	return o
}

func (o GetGitlabWebhookEndpointResultOutput) ToGetGitlabWebhookEndpointResultOutputWithContext(ctx context.Context) GetGitlabWebhookEndpointResultOutput {
	return o
}

func (o GetGitlabWebhookEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetGitlabWebhookEndpointResult] {
	return pulumix.Output[GetGitlabWebhookEndpointResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetGitlabWebhookEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitlabWebhookEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gitlab webhook endpoint address
func (o GetGitlabWebhookEndpointResultOutput) WebhookEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitlabWebhookEndpointResult) string { return v.WebhookEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGitlabWebhookEndpointResultOutput{})
}
