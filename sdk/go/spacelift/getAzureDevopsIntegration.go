// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift/internal"
)

// `getAzureDevopsIntegration` returns details about Azure DevOps integration
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
//			_, err := spacelift.GetAzureDevopsIntegration(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAzureDevopsIntegration(ctx *pulumi.Context, args *GetAzureDevopsIntegrationArgs, opts ...pulumi.InvokeOption) (*GetAzureDevopsIntegrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAzureDevopsIntegrationResult
	err := ctx.Invoke("spacelift:index/getAzureDevopsIntegration:getAzureDevopsIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAzureDevopsIntegration.
type GetAzureDevopsIntegrationArgs struct {
	// Azure DevOps integration id. If not provided, the default integration will be returned
	Id *string `pulumi:"id"`
}

// A collection of values returned by getAzureDevopsIntegration.
type GetAzureDevopsIntegrationResult struct {
	// Azure DevOps integration description
	Description string `pulumi:"description"`
	// Azure DevOps integration id. If not provided, the default integration will be returned
	Id *string `pulumi:"id"`
	// Azure DevOps integration is default
	IsDefault bool `pulumi:"isDefault"`
	// Azure DevOps integration labels
	Labels []string `pulumi:"labels"`
	// Azure DevOps integration name
	Name string `pulumi:"name"`
	// Azure DevOps integration organization url
	OrganizationUrl string `pulumi:"organizationUrl"`
	// Azure DevOps integration space id
	SpaceId string `pulumi:"spaceId"`
	// Azure DevOps integration webhook password
	WebhookPassword string `pulumi:"webhookPassword"`
	// Azure DevOps integration webhook url
	WebhookUrl string `pulumi:"webhookUrl"`
}

func GetAzureDevopsIntegrationOutput(ctx *pulumi.Context, args GetAzureDevopsIntegrationOutputArgs, opts ...pulumi.InvokeOption) GetAzureDevopsIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAzureDevopsIntegrationResult, error) {
			args := v.(GetAzureDevopsIntegrationArgs)
			r, err := GetAzureDevopsIntegration(ctx, &args, opts...)
			var s GetAzureDevopsIntegrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAzureDevopsIntegrationResultOutput)
}

// A collection of arguments for invoking getAzureDevopsIntegration.
type GetAzureDevopsIntegrationOutputArgs struct {
	// Azure DevOps integration id. If not provided, the default integration will be returned
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (GetAzureDevopsIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAzureDevopsIntegrationArgs)(nil)).Elem()
}

// A collection of values returned by getAzureDevopsIntegration.
type GetAzureDevopsIntegrationResultOutput struct{ *pulumi.OutputState }

func (GetAzureDevopsIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAzureDevopsIntegrationResult)(nil)).Elem()
}

func (o GetAzureDevopsIntegrationResultOutput) ToGetAzureDevopsIntegrationResultOutput() GetAzureDevopsIntegrationResultOutput {
	return o
}

func (o GetAzureDevopsIntegrationResultOutput) ToGetAzureDevopsIntegrationResultOutputWithContext(ctx context.Context) GetAzureDevopsIntegrationResultOutput {
	return o
}

// Azure DevOps integration description
func (o GetAzureDevopsIntegrationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) string { return v.Description }).(pulumi.StringOutput)
}

// Azure DevOps integration id. If not provided, the default integration will be returned
func (o GetAzureDevopsIntegrationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Azure DevOps integration is default
func (o GetAzureDevopsIntegrationResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// Azure DevOps integration labels
func (o GetAzureDevopsIntegrationResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// Azure DevOps integration name
func (o GetAzureDevopsIntegrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure DevOps integration organization url
func (o GetAzureDevopsIntegrationResultOutput) OrganizationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) string { return v.OrganizationUrl }).(pulumi.StringOutput)
}

// Azure DevOps integration space id
func (o GetAzureDevopsIntegrationResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// Azure DevOps integration webhook password
func (o GetAzureDevopsIntegrationResultOutput) WebhookPassword() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) string { return v.WebhookPassword }).(pulumi.StringOutput)
}

// Azure DevOps integration webhook url
func (o GetAzureDevopsIntegrationResultOutput) WebhookUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureDevopsIntegrationResult) string { return v.WebhookUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAzureDevopsIntegrationResultOutput{})
}
