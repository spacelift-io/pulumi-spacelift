// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetWorkerPools(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetWorkerPoolsResult, error) {
	var rv GetWorkerPoolsResult
	err := ctx.Invoke("spacelift:index/getWorkerPools:getWorkerPools", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getWorkerPools.
type GetWorkerPoolsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string                     `pulumi:"id"`
	WorkerPools []GetWorkerPoolsWorkerPool `pulumi:"workerPools"`
}