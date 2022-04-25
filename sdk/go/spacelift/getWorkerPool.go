// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `WorkerPool` represents a worker pool assigned to the Spacelift account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-spacelift/sdk/go/spacelift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/spacelift-io/spacelift-spacelift/sdk/go/spacelift/"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spacelift.LookupWorkerPool(ctx, &spacelift.LookupWorkerPoolArgs{
// 			WorkerPoolId: "k8s-core",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// <!-- schema generated by tfplugindocs -->
// ## Schema
//
// ### Required
//
// - **worker_pool_id** (String) ID of the worker pool
//
// ### Optional
//
// - **id** (String) The ID of this resource.
//
// ### Read-Only
//
// - **config** (String, Sensitive) credentials necessary to connect WorkerPool's workers to the control plane
// - **description** (String) description of the worker pool
// - **labels** (Set of String)
// - **name** (String) name of the worker pool
func LookupWorkerPool(ctx *pulumi.Context, args *LookupWorkerPoolArgs, opts ...pulumi.InvokeOption) (*LookupWorkerPoolResult, error) {
	var rv LookupWorkerPoolResult
	err := ctx.Invoke("spacelift:index/getWorkerPool:getWorkerPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkerPool.
type LookupWorkerPoolArgs struct {
	WorkerPoolId string `pulumi:"workerPoolId"`
}

// A collection of values returned by getWorkerPool.
type LookupWorkerPoolResult struct {
	Config      string `pulumi:"config"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id           string   `pulumi:"id"`
	Labels       []string `pulumi:"labels"`
	Name         string   `pulumi:"name"`
	WorkerPoolId string   `pulumi:"workerPoolId"`
}
