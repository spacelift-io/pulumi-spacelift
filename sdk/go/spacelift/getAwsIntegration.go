// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `AwsIntegration` represents an integration with an AWS account. This integration is account-level and needs to be explicitly attached to individual stacks in order to take effect.
//
// Note: when assuming credentials for **shared workers**, Spacelift will use `$accountName-$integrationID@$stackID-suffix` or `$accountName-$integrationID@$moduleID-$suffix` as [external ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html) and `$runID@$stackID@$accountName` truncated to 64 characters as [session ID](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole),$suffix will be `read` or `write`.
func LookupAwsIntegration(ctx *pulumi.Context, args *LookupAwsIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupAwsIntegrationResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAwsIntegrationResult
	err := ctx.Invoke("spacelift:index/getAwsIntegration:getAwsIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAwsIntegration.
type LookupAwsIntegrationArgs struct {
	IntegrationId string `pulumi:"integrationId"`
}

// A collection of values returned by getAwsIntegration.
type LookupAwsIntegrationResult struct {
	DurationSeconds             int    `pulumi:"durationSeconds"`
	ExternalId                  string `pulumi:"externalId"`
	GenerateCredentialsInWorker bool   `pulumi:"generateCredentialsInWorker"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	IntegrationId string   `pulumi:"integrationId"`
	Labels        []string `pulumi:"labels"`
	Name          string   `pulumi:"name"`
	RoleArn       string   `pulumi:"roleArn"`
	SpaceId       string   `pulumi:"spaceId"`
}

func LookupAwsIntegrationOutput(ctx *pulumi.Context, args LookupAwsIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupAwsIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAwsIntegrationResult, error) {
			args := v.(LookupAwsIntegrationArgs)
			r, err := LookupAwsIntegration(ctx, &args, opts...)
			var s LookupAwsIntegrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAwsIntegrationResultOutput)
}

// A collection of arguments for invoking getAwsIntegration.
type LookupAwsIntegrationOutputArgs struct {
	IntegrationId pulumi.StringInput `pulumi:"integrationId"`
}

func (LookupAwsIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsIntegrationArgs)(nil)).Elem()
}

// A collection of values returned by getAwsIntegration.
type LookupAwsIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupAwsIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsIntegrationResult)(nil)).Elem()
}

func (o LookupAwsIntegrationResultOutput) ToLookupAwsIntegrationResultOutput() LookupAwsIntegrationResultOutput {
	return o
}

func (o LookupAwsIntegrationResultOutput) ToLookupAwsIntegrationResultOutputWithContext(ctx context.Context) LookupAwsIntegrationResultOutput {
	return o
}

func (o LookupAwsIntegrationResultOutput) DurationSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) int { return v.DurationSeconds }).(pulumi.IntOutput)
}

func (o LookupAwsIntegrationResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

func (o LookupAwsIntegrationResultOutput) GenerateCredentialsInWorker() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) bool { return v.GenerateCredentialsInWorker }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAwsIntegrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAwsIntegrationResultOutput) IntegrationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) string { return v.IntegrationId }).(pulumi.StringOutput)
}

func (o LookupAwsIntegrationResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupAwsIntegrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAwsIntegrationResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

func (o LookupAwsIntegrationResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsIntegrationResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAwsIntegrationResultOutput{})
}