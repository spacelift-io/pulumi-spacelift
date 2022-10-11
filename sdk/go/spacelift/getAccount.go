// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getAccount` represents the currently used Spacelift **account**
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/spacelift-io/pulumi-spacelift/sdk/go/spacelift"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := spacelift.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAccount(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetAccountResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAccountResult
	err := ctx.Invoke("spacelift:index/getAccount:getAccount", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAccount.
type GetAccountResult struct {
	// the ID of the AWS account used by Spacelift for role assumption
	AwsAccountId string `pulumi:"awsAccountId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// name of the account
	Name string `pulumi:"name"`
	// account billing tier
	Tier string `pulumi:"tier"`
}
