// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.
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
// 		_, err := spacelift.LookupDriftDetection(ctx, &GetDriftDetectionArgs{
// 			StackId: "core-infra-production",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDriftDetection(ctx *pulumi.Context, args *LookupDriftDetectionArgs, opts ...pulumi.InvokeOption) (*LookupDriftDetectionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupDriftDetectionResult
	err := ctx.Invoke("spacelift:index/getDriftDetection:getDriftDetection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDriftDetection.
type LookupDriftDetectionArgs struct {
	StackId string `pulumi:"stackId"`
}

// A collection of values returned by getDriftDetection.
type LookupDriftDetectionResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Reconcile bool     `pulumi:"reconcile"`
	Schedules []string `pulumi:"schedules"`
	StackId   string   `pulumi:"stackId"`
	Timezone  string   `pulumi:"timezone"`
}

func LookupDriftDetectionOutput(ctx *pulumi.Context, args LookupDriftDetectionOutputArgs, opts ...pulumi.InvokeOption) LookupDriftDetectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDriftDetectionResult, error) {
			args := v.(LookupDriftDetectionArgs)
			r, err := LookupDriftDetection(ctx, &args, opts...)
			var s LookupDriftDetectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDriftDetectionResultOutput)
}

// A collection of arguments for invoking getDriftDetection.
type LookupDriftDetectionOutputArgs struct {
	StackId pulumi.StringInput `pulumi:"stackId"`
}

func (LookupDriftDetectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDriftDetectionArgs)(nil)).Elem()
}

// A collection of values returned by getDriftDetection.
type LookupDriftDetectionResultOutput struct{ *pulumi.OutputState }

func (LookupDriftDetectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDriftDetectionResult)(nil)).Elem()
}

func (o LookupDriftDetectionResultOutput) ToLookupDriftDetectionResultOutput() LookupDriftDetectionResultOutput {
	return o
}

func (o LookupDriftDetectionResultOutput) ToLookupDriftDetectionResultOutputWithContext(ctx context.Context) LookupDriftDetectionResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDriftDetectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDriftDetectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDriftDetectionResultOutput) Reconcile() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDriftDetectionResult) bool { return v.Reconcile }).(pulumi.BoolOutput)
}

func (o LookupDriftDetectionResultOutput) Schedules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDriftDetectionResult) []string { return v.Schedules }).(pulumi.StringArrayOutput)
}

func (o LookupDriftDetectionResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDriftDetectionResult) string { return v.StackId }).(pulumi.StringOutput)
}

func (o LookupDriftDetectionResultOutput) Timezone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDriftDetectionResult) string { return v.Timezone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDriftDetectionResultOutput{})
}
