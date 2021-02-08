// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupMountedFile(ctx *pulumi.Context, args *LookupMountedFileArgs, opts ...pulumi.InvokeOption) (*LookupMountedFileResult, error) {
	var rv LookupMountedFileResult
	err := ctx.Invoke("spacelift:index/getMountedFile:getMountedFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMountedFile.
type LookupMountedFileArgs struct {
	ContextId    *string `pulumi:"contextId"`
	ModuleId     *string `pulumi:"moduleId"`
	RelativePath string  `pulumi:"relativePath"`
	StackId      *string `pulumi:"stackId"`
}

// A collection of values returned by getMountedFile.
type LookupMountedFileResult struct {
	Checksum  string  `pulumi:"checksum"`
	Content   string  `pulumi:"content"`
	ContextId *string `pulumi:"contextId"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	ModuleId     *string `pulumi:"moduleId"`
	RelativePath string  `pulumi:"relativePath"`
	StackId      *string `pulumi:"stackId"`
	WriteOnly    bool    `pulumi:"writeOnly"`
}
