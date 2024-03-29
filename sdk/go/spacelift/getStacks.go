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

// `getStacks` represents all the stacks in the Spacelift account visible to the API user, matching predicates.
func GetStacks(ctx *pulumi.Context, args *GetStacksArgs, opts ...pulumi.InvokeOption) (*GetStacksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStacksResult
	err := ctx.Invoke("spacelift:index/getStacks:getStacks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStacks.
type GetStacksArgs struct {
	Administrative *GetStacksAdministrative `pulumi:"administrative"`
	Branch         *GetStacksBranch         `pulumi:"branch"`
	// Require stacks to be on one of the commits
	Commit *GetStacksCommit `pulumi:"commit"`
	Labels []GetStacksLabel `pulumi:"labels"`
	// Require stacks to be locked
	Locked      *GetStacksLocked      `pulumi:"locked"`
	Name        *GetStacksName        `pulumi:"name"`
	ProjectRoot *GetStacksProjectRoot `pulumi:"projectRoot"`
	Repository  *GetStacksRepository  `pulumi:"repository"`
	// Require stacks to have one of the states
	State *GetStacksState `pulumi:"state"`
	// Require stacks to use one of the IaC vendors
	Vendor *GetStacksVendor `pulumi:"vendor"`
	// Require stacks to use one of the worker pools
	WorkerPool *GetStacksWorkerPool `pulumi:"workerPool"`
}

// A collection of values returned by getStacks.
type GetStacksResult struct {
	// Require stacks to be administrative or not
	Administrative *GetStacksAdministrative `pulumi:"administrative"`
	// Require stacks to be on one of the branches
	Branch *GetStacksBranch `pulumi:"branch"`
	// Require stacks to be on one of the commits
	Commit *GetStacksCommit `pulumi:"commit"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Require stacks to have one of the labels
	Labels []GetStacksLabel `pulumi:"labels"`
	// Require stacks to be locked
	Locked *GetStacksLocked `pulumi:"locked"`
	// Require stacks to have one of the names
	Name *GetStacksName `pulumi:"name"`
	// Require stacks to be in one of the project roots
	ProjectRoot *GetStacksProjectRoot `pulumi:"projectRoot"`
	// Require stacks to be in one of the repositories
	Repository *GetStacksRepository `pulumi:"repository"`
	// List of stacks matching the predicates
	Stacks []GetStacksStack `pulumi:"stacks"`
	// Require stacks to have one of the states
	State *GetStacksState `pulumi:"state"`
	// Require stacks to use one of the IaC vendors
	Vendor *GetStacksVendor `pulumi:"vendor"`
	// Require stacks to use one of the worker pools
	WorkerPool *GetStacksWorkerPool `pulumi:"workerPool"`
}

func GetStacksOutput(ctx *pulumi.Context, args GetStacksOutputArgs, opts ...pulumi.InvokeOption) GetStacksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStacksResult, error) {
			args := v.(GetStacksArgs)
			r, err := GetStacks(ctx, &args, opts...)
			var s GetStacksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStacksResultOutput)
}

// A collection of arguments for invoking getStacks.
type GetStacksOutputArgs struct {
	Administrative GetStacksAdministrativePtrInput `pulumi:"administrative"`
	Branch         GetStacksBranchPtrInput         `pulumi:"branch"`
	// Require stacks to be on one of the commits
	Commit GetStacksCommitPtrInput  `pulumi:"commit"`
	Labels GetStacksLabelArrayInput `pulumi:"labels"`
	// Require stacks to be locked
	Locked      GetStacksLockedPtrInput      `pulumi:"locked"`
	Name        GetStacksNamePtrInput        `pulumi:"name"`
	ProjectRoot GetStacksProjectRootPtrInput `pulumi:"projectRoot"`
	Repository  GetStacksRepositoryPtrInput  `pulumi:"repository"`
	// Require stacks to have one of the states
	State GetStacksStatePtrInput `pulumi:"state"`
	// Require stacks to use one of the IaC vendors
	Vendor GetStacksVendorPtrInput `pulumi:"vendor"`
	// Require stacks to use one of the worker pools
	WorkerPool GetStacksWorkerPoolPtrInput `pulumi:"workerPool"`
}

func (GetStacksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStacksArgs)(nil)).Elem()
}

// A collection of values returned by getStacks.
type GetStacksResultOutput struct{ *pulumi.OutputState }

func (GetStacksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStacksResult)(nil)).Elem()
}

func (o GetStacksResultOutput) ToGetStacksResultOutput() GetStacksResultOutput {
	return o
}

func (o GetStacksResultOutput) ToGetStacksResultOutputWithContext(ctx context.Context) GetStacksResultOutput {
	return o
}

func (o GetStacksResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetStacksResult] {
	return pulumix.Output[GetStacksResult]{
		OutputState: o.OutputState,
	}
}

// Require stacks to be administrative or not
func (o GetStacksResultOutput) Administrative() GetStacksAdministrativePtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksAdministrative { return v.Administrative }).(GetStacksAdministrativePtrOutput)
}

// Require stacks to be on one of the branches
func (o GetStacksResultOutput) Branch() GetStacksBranchPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksBranch { return v.Branch }).(GetStacksBranchPtrOutput)
}

// Require stacks to be on one of the commits
func (o GetStacksResultOutput) Commit() GetStacksCommitPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksCommit { return v.Commit }).(GetStacksCommitPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStacksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStacksResult) string { return v.Id }).(pulumi.StringOutput)
}

// Require stacks to have one of the labels
func (o GetStacksResultOutput) Labels() GetStacksLabelArrayOutput {
	return o.ApplyT(func(v GetStacksResult) []GetStacksLabel { return v.Labels }).(GetStacksLabelArrayOutput)
}

// Require stacks to be locked
func (o GetStacksResultOutput) Locked() GetStacksLockedPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksLocked { return v.Locked }).(GetStacksLockedPtrOutput)
}

// Require stacks to have one of the names
func (o GetStacksResultOutput) Name() GetStacksNamePtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksName { return v.Name }).(GetStacksNamePtrOutput)
}

// Require stacks to be in one of the project roots
func (o GetStacksResultOutput) ProjectRoot() GetStacksProjectRootPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksProjectRoot { return v.ProjectRoot }).(GetStacksProjectRootPtrOutput)
}

// Require stacks to be in one of the repositories
func (o GetStacksResultOutput) Repository() GetStacksRepositoryPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksRepository { return v.Repository }).(GetStacksRepositoryPtrOutput)
}

// List of stacks matching the predicates
func (o GetStacksResultOutput) Stacks() GetStacksStackArrayOutput {
	return o.ApplyT(func(v GetStacksResult) []GetStacksStack { return v.Stacks }).(GetStacksStackArrayOutput)
}

// Require stacks to have one of the states
func (o GetStacksResultOutput) State() GetStacksStatePtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksState { return v.State }).(GetStacksStatePtrOutput)
}

// Require stacks to use one of the IaC vendors
func (o GetStacksResultOutput) Vendor() GetStacksVendorPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksVendor { return v.Vendor }).(GetStacksVendorPtrOutput)
}

// Require stacks to use one of the worker pools
func (o GetStacksResultOutput) WorkerPool() GetStacksWorkerPoolPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *GetStacksWorkerPool { return v.WorkerPool }).(GetStacksWorkerPoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStacksResultOutput{})
}
