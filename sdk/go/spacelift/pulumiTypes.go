// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ModuleGitlab struct {
	Namespace string `pulumi:"namespace"`
}

// ModuleGitlabInput is an input type that accepts ModuleGitlabArgs and ModuleGitlabOutput values.
// You can construct a concrete instance of `ModuleGitlabInput` via:
//
//          ModuleGitlabArgs{...}
type ModuleGitlabInput interface {
	pulumi.Input

	ToModuleGitlabOutput() ModuleGitlabOutput
	ToModuleGitlabOutputWithContext(context.Context) ModuleGitlabOutput
}

type ModuleGitlabArgs struct {
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (ModuleGitlabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleGitlab)(nil)).Elem()
}

func (i ModuleGitlabArgs) ToModuleGitlabOutput() ModuleGitlabOutput {
	return i.ToModuleGitlabOutputWithContext(context.Background())
}

func (i ModuleGitlabArgs) ToModuleGitlabOutputWithContext(ctx context.Context) ModuleGitlabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleGitlabOutput)
}

func (i ModuleGitlabArgs) ToModuleGitlabPtrOutput() ModuleGitlabPtrOutput {
	return i.ToModuleGitlabPtrOutputWithContext(context.Background())
}

func (i ModuleGitlabArgs) ToModuleGitlabPtrOutputWithContext(ctx context.Context) ModuleGitlabPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleGitlabOutput).ToModuleGitlabPtrOutputWithContext(ctx)
}

// ModuleGitlabPtrInput is an input type that accepts ModuleGitlabArgs, ModuleGitlabPtr and ModuleGitlabPtrOutput values.
// You can construct a concrete instance of `ModuleGitlabPtrInput` via:
//
//          ModuleGitlabArgs{...}
//
//  or:
//
//          nil
type ModuleGitlabPtrInput interface {
	pulumi.Input

	ToModuleGitlabPtrOutput() ModuleGitlabPtrOutput
	ToModuleGitlabPtrOutputWithContext(context.Context) ModuleGitlabPtrOutput
}

type moduleGitlabPtrType ModuleGitlabArgs

func ModuleGitlabPtr(v *ModuleGitlabArgs) ModuleGitlabPtrInput {
	return (*moduleGitlabPtrType)(v)
}

func (*moduleGitlabPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleGitlab)(nil)).Elem()
}

func (i *moduleGitlabPtrType) ToModuleGitlabPtrOutput() ModuleGitlabPtrOutput {
	return i.ToModuleGitlabPtrOutputWithContext(context.Background())
}

func (i *moduleGitlabPtrType) ToModuleGitlabPtrOutputWithContext(ctx context.Context) ModuleGitlabPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleGitlabPtrOutput)
}

type ModuleGitlabOutput struct{ *pulumi.OutputState }

func (ModuleGitlabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleGitlab)(nil)).Elem()
}

func (o ModuleGitlabOutput) ToModuleGitlabOutput() ModuleGitlabOutput {
	return o
}

func (o ModuleGitlabOutput) ToModuleGitlabOutputWithContext(ctx context.Context) ModuleGitlabOutput {
	return o
}

func (o ModuleGitlabOutput) ToModuleGitlabPtrOutput() ModuleGitlabPtrOutput {
	return o.ToModuleGitlabPtrOutputWithContext(context.Background())
}

func (o ModuleGitlabOutput) ToModuleGitlabPtrOutputWithContext(ctx context.Context) ModuleGitlabPtrOutput {
	return o.ApplyT(func(v ModuleGitlab) *ModuleGitlab {
		return &v
	}).(ModuleGitlabPtrOutput)
}
func (o ModuleGitlabOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v ModuleGitlab) string { return v.Namespace }).(pulumi.StringOutput)
}

type ModuleGitlabPtrOutput struct{ *pulumi.OutputState }

func (ModuleGitlabPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleGitlab)(nil)).Elem()
}

func (o ModuleGitlabPtrOutput) ToModuleGitlabPtrOutput() ModuleGitlabPtrOutput {
	return o
}

func (o ModuleGitlabPtrOutput) ToModuleGitlabPtrOutputWithContext(ctx context.Context) ModuleGitlabPtrOutput {
	return o
}

func (o ModuleGitlabPtrOutput) Elem() ModuleGitlabOutput {
	return o.ApplyT(func(v *ModuleGitlab) ModuleGitlab { return *v }).(ModuleGitlabOutput)
}

func (o ModuleGitlabPtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleGitlab) *string {
		if v == nil {
			return nil
		}
		return &v.Namespace
	}).(pulumi.StringPtrOutput)
}

type StackCloudformation struct {
	EntryTemplateFile string `pulumi:"entryTemplateFile"`
	Region            string `pulumi:"region"`
	StackName         string `pulumi:"stackName"`
	TemplateBucket    string `pulumi:"templateBucket"`
}

// StackCloudformationInput is an input type that accepts StackCloudformationArgs and StackCloudformationOutput values.
// You can construct a concrete instance of `StackCloudformationInput` via:
//
//          StackCloudformationArgs{...}
type StackCloudformationInput interface {
	pulumi.Input

	ToStackCloudformationOutput() StackCloudformationOutput
	ToStackCloudformationOutputWithContext(context.Context) StackCloudformationOutput
}

type StackCloudformationArgs struct {
	EntryTemplateFile pulumi.StringInput `pulumi:"entryTemplateFile"`
	Region            pulumi.StringInput `pulumi:"region"`
	StackName         pulumi.StringInput `pulumi:"stackName"`
	TemplateBucket    pulumi.StringInput `pulumi:"templateBucket"`
}

func (StackCloudformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StackCloudformation)(nil)).Elem()
}

func (i StackCloudformationArgs) ToStackCloudformationOutput() StackCloudformationOutput {
	return i.ToStackCloudformationOutputWithContext(context.Background())
}

func (i StackCloudformationArgs) ToStackCloudformationOutputWithContext(ctx context.Context) StackCloudformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackCloudformationOutput)
}

func (i StackCloudformationArgs) ToStackCloudformationPtrOutput() StackCloudformationPtrOutput {
	return i.ToStackCloudformationPtrOutputWithContext(context.Background())
}

func (i StackCloudformationArgs) ToStackCloudformationPtrOutputWithContext(ctx context.Context) StackCloudformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackCloudformationOutput).ToStackCloudformationPtrOutputWithContext(ctx)
}

// StackCloudformationPtrInput is an input type that accepts StackCloudformationArgs, StackCloudformationPtr and StackCloudformationPtrOutput values.
// You can construct a concrete instance of `StackCloudformationPtrInput` via:
//
//          StackCloudformationArgs{...}
//
//  or:
//
//          nil
type StackCloudformationPtrInput interface {
	pulumi.Input

	ToStackCloudformationPtrOutput() StackCloudformationPtrOutput
	ToStackCloudformationPtrOutputWithContext(context.Context) StackCloudformationPtrOutput
}

type stackCloudformationPtrType StackCloudformationArgs

func StackCloudformationPtr(v *StackCloudformationArgs) StackCloudformationPtrInput {
	return (*stackCloudformationPtrType)(v)
}

func (*stackCloudformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StackCloudformation)(nil)).Elem()
}

func (i *stackCloudformationPtrType) ToStackCloudformationPtrOutput() StackCloudformationPtrOutput {
	return i.ToStackCloudformationPtrOutputWithContext(context.Background())
}

func (i *stackCloudformationPtrType) ToStackCloudformationPtrOutputWithContext(ctx context.Context) StackCloudformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackCloudformationPtrOutput)
}

type StackCloudformationOutput struct{ *pulumi.OutputState }

func (StackCloudformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackCloudformation)(nil)).Elem()
}

func (o StackCloudformationOutput) ToStackCloudformationOutput() StackCloudformationOutput {
	return o
}

func (o StackCloudformationOutput) ToStackCloudformationOutputWithContext(ctx context.Context) StackCloudformationOutput {
	return o
}

func (o StackCloudformationOutput) ToStackCloudformationPtrOutput() StackCloudformationPtrOutput {
	return o.ToStackCloudformationPtrOutputWithContext(context.Background())
}

func (o StackCloudformationOutput) ToStackCloudformationPtrOutputWithContext(ctx context.Context) StackCloudformationPtrOutput {
	return o.ApplyT(func(v StackCloudformation) *StackCloudformation {
		return &v
	}).(StackCloudformationPtrOutput)
}
func (o StackCloudformationOutput) EntryTemplateFile() pulumi.StringOutput {
	return o.ApplyT(func(v StackCloudformation) string { return v.EntryTemplateFile }).(pulumi.StringOutput)
}

func (o StackCloudformationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v StackCloudformation) string { return v.Region }).(pulumi.StringOutput)
}

func (o StackCloudformationOutput) StackName() pulumi.StringOutput {
	return o.ApplyT(func(v StackCloudformation) string { return v.StackName }).(pulumi.StringOutput)
}

func (o StackCloudformationOutput) TemplateBucket() pulumi.StringOutput {
	return o.ApplyT(func(v StackCloudformation) string { return v.TemplateBucket }).(pulumi.StringOutput)
}

type StackCloudformationPtrOutput struct{ *pulumi.OutputState }

func (StackCloudformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackCloudformation)(nil)).Elem()
}

func (o StackCloudformationPtrOutput) ToStackCloudformationPtrOutput() StackCloudformationPtrOutput {
	return o
}

func (o StackCloudformationPtrOutput) ToStackCloudformationPtrOutputWithContext(ctx context.Context) StackCloudformationPtrOutput {
	return o
}

func (o StackCloudformationPtrOutput) Elem() StackCloudformationOutput {
	return o.ApplyT(func(v *StackCloudformation) StackCloudformation { return *v }).(StackCloudformationOutput)
}

func (o StackCloudformationPtrOutput) EntryTemplateFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackCloudformation) *string {
		if v == nil {
			return nil
		}
		return &v.EntryTemplateFile
	}).(pulumi.StringPtrOutput)
}

func (o StackCloudformationPtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackCloudformation) *string {
		if v == nil {
			return nil
		}
		return &v.Region
	}).(pulumi.StringPtrOutput)
}

func (o StackCloudformationPtrOutput) StackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackCloudformation) *string {
		if v == nil {
			return nil
		}
		return &v.StackName
	}).(pulumi.StringPtrOutput)
}

func (o StackCloudformationPtrOutput) TemplateBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackCloudformation) *string {
		if v == nil {
			return nil
		}
		return &v.TemplateBucket
	}).(pulumi.StringPtrOutput)
}

type StackGitlab struct {
	Namespace string `pulumi:"namespace"`
}

// StackGitlabInput is an input type that accepts StackGitlabArgs and StackGitlabOutput values.
// You can construct a concrete instance of `StackGitlabInput` via:
//
//          StackGitlabArgs{...}
type StackGitlabInput interface {
	pulumi.Input

	ToStackGitlabOutput() StackGitlabOutput
	ToStackGitlabOutputWithContext(context.Context) StackGitlabOutput
}

type StackGitlabArgs struct {
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (StackGitlabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StackGitlab)(nil)).Elem()
}

func (i StackGitlabArgs) ToStackGitlabOutput() StackGitlabOutput {
	return i.ToStackGitlabOutputWithContext(context.Background())
}

func (i StackGitlabArgs) ToStackGitlabOutputWithContext(ctx context.Context) StackGitlabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackGitlabOutput)
}

func (i StackGitlabArgs) ToStackGitlabPtrOutput() StackGitlabPtrOutput {
	return i.ToStackGitlabPtrOutputWithContext(context.Background())
}

func (i StackGitlabArgs) ToStackGitlabPtrOutputWithContext(ctx context.Context) StackGitlabPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackGitlabOutput).ToStackGitlabPtrOutputWithContext(ctx)
}

// StackGitlabPtrInput is an input type that accepts StackGitlabArgs, StackGitlabPtr and StackGitlabPtrOutput values.
// You can construct a concrete instance of `StackGitlabPtrInput` via:
//
//          StackGitlabArgs{...}
//
//  or:
//
//          nil
type StackGitlabPtrInput interface {
	pulumi.Input

	ToStackGitlabPtrOutput() StackGitlabPtrOutput
	ToStackGitlabPtrOutputWithContext(context.Context) StackGitlabPtrOutput
}

type stackGitlabPtrType StackGitlabArgs

func StackGitlabPtr(v *StackGitlabArgs) StackGitlabPtrInput {
	return (*stackGitlabPtrType)(v)
}

func (*stackGitlabPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StackGitlab)(nil)).Elem()
}

func (i *stackGitlabPtrType) ToStackGitlabPtrOutput() StackGitlabPtrOutput {
	return i.ToStackGitlabPtrOutputWithContext(context.Background())
}

func (i *stackGitlabPtrType) ToStackGitlabPtrOutputWithContext(ctx context.Context) StackGitlabPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackGitlabPtrOutput)
}

type StackGitlabOutput struct{ *pulumi.OutputState }

func (StackGitlabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackGitlab)(nil)).Elem()
}

func (o StackGitlabOutput) ToStackGitlabOutput() StackGitlabOutput {
	return o
}

func (o StackGitlabOutput) ToStackGitlabOutputWithContext(ctx context.Context) StackGitlabOutput {
	return o
}

func (o StackGitlabOutput) ToStackGitlabPtrOutput() StackGitlabPtrOutput {
	return o.ToStackGitlabPtrOutputWithContext(context.Background())
}

func (o StackGitlabOutput) ToStackGitlabPtrOutputWithContext(ctx context.Context) StackGitlabPtrOutput {
	return o.ApplyT(func(v StackGitlab) *StackGitlab {
		return &v
	}).(StackGitlabPtrOutput)
}
func (o StackGitlabOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v StackGitlab) string { return v.Namespace }).(pulumi.StringOutput)
}

type StackGitlabPtrOutput struct{ *pulumi.OutputState }

func (StackGitlabPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackGitlab)(nil)).Elem()
}

func (o StackGitlabPtrOutput) ToStackGitlabPtrOutput() StackGitlabPtrOutput {
	return o
}

func (o StackGitlabPtrOutput) ToStackGitlabPtrOutputWithContext(ctx context.Context) StackGitlabPtrOutput {
	return o
}

func (o StackGitlabPtrOutput) Elem() StackGitlabOutput {
	return o.ApplyT(func(v *StackGitlab) StackGitlab { return *v }).(StackGitlabOutput)
}

func (o StackGitlabPtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackGitlab) *string {
		if v == nil {
			return nil
		}
		return &v.Namespace
	}).(pulumi.StringPtrOutput)
}

type StackPulumi struct {
	LoginUrl  string `pulumi:"loginUrl"`
	StackName string `pulumi:"stackName"`
}

// StackPulumiInput is an input type that accepts StackPulumiArgs and StackPulumiOutput values.
// You can construct a concrete instance of `StackPulumiInput` via:
//
//          StackPulumiArgs{...}
type StackPulumiInput interface {
	pulumi.Input

	ToStackPulumiOutput() StackPulumiOutput
	ToStackPulumiOutputWithContext(context.Context) StackPulumiOutput
}

type StackPulumiArgs struct {
	LoginUrl  pulumi.StringInput `pulumi:"loginUrl"`
	StackName pulumi.StringInput `pulumi:"stackName"`
}

func (StackPulumiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StackPulumi)(nil)).Elem()
}

func (i StackPulumiArgs) ToStackPulumiOutput() StackPulumiOutput {
	return i.ToStackPulumiOutputWithContext(context.Background())
}

func (i StackPulumiArgs) ToStackPulumiOutputWithContext(ctx context.Context) StackPulumiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPulumiOutput)
}

func (i StackPulumiArgs) ToStackPulumiPtrOutput() StackPulumiPtrOutput {
	return i.ToStackPulumiPtrOutputWithContext(context.Background())
}

func (i StackPulumiArgs) ToStackPulumiPtrOutputWithContext(ctx context.Context) StackPulumiPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPulumiOutput).ToStackPulumiPtrOutputWithContext(ctx)
}

// StackPulumiPtrInput is an input type that accepts StackPulumiArgs, StackPulumiPtr and StackPulumiPtrOutput values.
// You can construct a concrete instance of `StackPulumiPtrInput` via:
//
//          StackPulumiArgs{...}
//
//  or:
//
//          nil
type StackPulumiPtrInput interface {
	pulumi.Input

	ToStackPulumiPtrOutput() StackPulumiPtrOutput
	ToStackPulumiPtrOutputWithContext(context.Context) StackPulumiPtrOutput
}

type stackPulumiPtrType StackPulumiArgs

func StackPulumiPtr(v *StackPulumiArgs) StackPulumiPtrInput {
	return (*stackPulumiPtrType)(v)
}

func (*stackPulumiPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StackPulumi)(nil)).Elem()
}

func (i *stackPulumiPtrType) ToStackPulumiPtrOutput() StackPulumiPtrOutput {
	return i.ToStackPulumiPtrOutputWithContext(context.Background())
}

func (i *stackPulumiPtrType) ToStackPulumiPtrOutputWithContext(ctx context.Context) StackPulumiPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPulumiPtrOutput)
}

type StackPulumiOutput struct{ *pulumi.OutputState }

func (StackPulumiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackPulumi)(nil)).Elem()
}

func (o StackPulumiOutput) ToStackPulumiOutput() StackPulumiOutput {
	return o
}

func (o StackPulumiOutput) ToStackPulumiOutputWithContext(ctx context.Context) StackPulumiOutput {
	return o
}

func (o StackPulumiOutput) ToStackPulumiPtrOutput() StackPulumiPtrOutput {
	return o.ToStackPulumiPtrOutputWithContext(context.Background())
}

func (o StackPulumiOutput) ToStackPulumiPtrOutputWithContext(ctx context.Context) StackPulumiPtrOutput {
	return o.ApplyT(func(v StackPulumi) *StackPulumi {
		return &v
	}).(StackPulumiPtrOutput)
}
func (o StackPulumiOutput) LoginUrl() pulumi.StringOutput {
	return o.ApplyT(func(v StackPulumi) string { return v.LoginUrl }).(pulumi.StringOutput)
}

func (o StackPulumiOutput) StackName() pulumi.StringOutput {
	return o.ApplyT(func(v StackPulumi) string { return v.StackName }).(pulumi.StringOutput)
}

type StackPulumiPtrOutput struct{ *pulumi.OutputState }

func (StackPulumiPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackPulumi)(nil)).Elem()
}

func (o StackPulumiPtrOutput) ToStackPulumiPtrOutput() StackPulumiPtrOutput {
	return o
}

func (o StackPulumiPtrOutput) ToStackPulumiPtrOutputWithContext(ctx context.Context) StackPulumiPtrOutput {
	return o
}

func (o StackPulumiPtrOutput) Elem() StackPulumiOutput {
	return o.ApplyT(func(v *StackPulumi) StackPulumi { return *v }).(StackPulumiOutput)
}

func (o StackPulumiPtrOutput) LoginUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackPulumi) *string {
		if v == nil {
			return nil
		}
		return &v.LoginUrl
	}).(pulumi.StringPtrOutput)
}

func (o StackPulumiPtrOutput) StackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StackPulumi) *string {
		if v == nil {
			return nil
		}
		return &v.StackName
	}).(pulumi.StringPtrOutput)
}

type GetModuleGitlab struct {
	Namespace string `pulumi:"namespace"`
}

// GetModuleGitlabInput is an input type that accepts GetModuleGitlabArgs and GetModuleGitlabOutput values.
// You can construct a concrete instance of `GetModuleGitlabInput` via:
//
//          GetModuleGitlabArgs{...}
type GetModuleGitlabInput interface {
	pulumi.Input

	ToGetModuleGitlabOutput() GetModuleGitlabOutput
	ToGetModuleGitlabOutputWithContext(context.Context) GetModuleGitlabOutput
}

type GetModuleGitlabArgs struct {
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (GetModuleGitlabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetModuleGitlab)(nil)).Elem()
}

func (i GetModuleGitlabArgs) ToGetModuleGitlabOutput() GetModuleGitlabOutput {
	return i.ToGetModuleGitlabOutputWithContext(context.Background())
}

func (i GetModuleGitlabArgs) ToGetModuleGitlabOutputWithContext(ctx context.Context) GetModuleGitlabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetModuleGitlabOutput)
}

// GetModuleGitlabArrayInput is an input type that accepts GetModuleGitlabArray and GetModuleGitlabArrayOutput values.
// You can construct a concrete instance of `GetModuleGitlabArrayInput` via:
//
//          GetModuleGitlabArray{ GetModuleGitlabArgs{...} }
type GetModuleGitlabArrayInput interface {
	pulumi.Input

	ToGetModuleGitlabArrayOutput() GetModuleGitlabArrayOutput
	ToGetModuleGitlabArrayOutputWithContext(context.Context) GetModuleGitlabArrayOutput
}

type GetModuleGitlabArray []GetModuleGitlabInput

func (GetModuleGitlabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetModuleGitlab)(nil)).Elem()
}

func (i GetModuleGitlabArray) ToGetModuleGitlabArrayOutput() GetModuleGitlabArrayOutput {
	return i.ToGetModuleGitlabArrayOutputWithContext(context.Background())
}

func (i GetModuleGitlabArray) ToGetModuleGitlabArrayOutputWithContext(ctx context.Context) GetModuleGitlabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetModuleGitlabArrayOutput)
}

type GetModuleGitlabOutput struct{ *pulumi.OutputState }

func (GetModuleGitlabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetModuleGitlab)(nil)).Elem()
}

func (o GetModuleGitlabOutput) ToGetModuleGitlabOutput() GetModuleGitlabOutput {
	return o
}

func (o GetModuleGitlabOutput) ToGetModuleGitlabOutputWithContext(ctx context.Context) GetModuleGitlabOutput {
	return o
}

func (o GetModuleGitlabOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetModuleGitlab) string { return v.Namespace }).(pulumi.StringOutput)
}

type GetModuleGitlabArrayOutput struct{ *pulumi.OutputState }

func (GetModuleGitlabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetModuleGitlab)(nil)).Elem()
}

func (o GetModuleGitlabArrayOutput) ToGetModuleGitlabArrayOutput() GetModuleGitlabArrayOutput {
	return o
}

func (o GetModuleGitlabArrayOutput) ToGetModuleGitlabArrayOutputWithContext(ctx context.Context) GetModuleGitlabArrayOutput {
	return o
}

func (o GetModuleGitlabArrayOutput) Index(i pulumi.IntInput) GetModuleGitlabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetModuleGitlab {
		return vs[0].([]GetModuleGitlab)[vs[1].(int)]
	}).(GetModuleGitlabOutput)
}

type GetStackCloudformation struct {
	EntryTemplateFile string `pulumi:"entryTemplateFile"`
	Region            string `pulumi:"region"`
	StackName         string `pulumi:"stackName"`
	TemplateBucket    string `pulumi:"templateBucket"`
}

// GetStackCloudformationInput is an input type that accepts GetStackCloudformationArgs and GetStackCloudformationOutput values.
// You can construct a concrete instance of `GetStackCloudformationInput` via:
//
//          GetStackCloudformationArgs{...}
type GetStackCloudformationInput interface {
	pulumi.Input

	ToGetStackCloudformationOutput() GetStackCloudformationOutput
	ToGetStackCloudformationOutputWithContext(context.Context) GetStackCloudformationOutput
}

type GetStackCloudformationArgs struct {
	EntryTemplateFile pulumi.StringInput `pulumi:"entryTemplateFile"`
	Region            pulumi.StringInput `pulumi:"region"`
	StackName         pulumi.StringInput `pulumi:"stackName"`
	TemplateBucket    pulumi.StringInput `pulumi:"templateBucket"`
}

func (GetStackCloudformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStackCloudformation)(nil)).Elem()
}

func (i GetStackCloudformationArgs) ToGetStackCloudformationOutput() GetStackCloudformationOutput {
	return i.ToGetStackCloudformationOutputWithContext(context.Background())
}

func (i GetStackCloudformationArgs) ToGetStackCloudformationOutputWithContext(ctx context.Context) GetStackCloudformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStackCloudformationOutput)
}

// GetStackCloudformationArrayInput is an input type that accepts GetStackCloudformationArray and GetStackCloudformationArrayOutput values.
// You can construct a concrete instance of `GetStackCloudformationArrayInput` via:
//
//          GetStackCloudformationArray{ GetStackCloudformationArgs{...} }
type GetStackCloudformationArrayInput interface {
	pulumi.Input

	ToGetStackCloudformationArrayOutput() GetStackCloudformationArrayOutput
	ToGetStackCloudformationArrayOutputWithContext(context.Context) GetStackCloudformationArrayOutput
}

type GetStackCloudformationArray []GetStackCloudformationInput

func (GetStackCloudformationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStackCloudformation)(nil)).Elem()
}

func (i GetStackCloudformationArray) ToGetStackCloudformationArrayOutput() GetStackCloudformationArrayOutput {
	return i.ToGetStackCloudformationArrayOutputWithContext(context.Background())
}

func (i GetStackCloudformationArray) ToGetStackCloudformationArrayOutputWithContext(ctx context.Context) GetStackCloudformationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStackCloudformationArrayOutput)
}

type GetStackCloudformationOutput struct{ *pulumi.OutputState }

func (GetStackCloudformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStackCloudformation)(nil)).Elem()
}

func (o GetStackCloudformationOutput) ToGetStackCloudformationOutput() GetStackCloudformationOutput {
	return o
}

func (o GetStackCloudformationOutput) ToGetStackCloudformationOutputWithContext(ctx context.Context) GetStackCloudformationOutput {
	return o
}

func (o GetStackCloudformationOutput) EntryTemplateFile() pulumi.StringOutput {
	return o.ApplyT(func(v GetStackCloudformation) string { return v.EntryTemplateFile }).(pulumi.StringOutput)
}

func (o GetStackCloudformationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetStackCloudformation) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetStackCloudformationOutput) StackName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStackCloudformation) string { return v.StackName }).(pulumi.StringOutput)
}

func (o GetStackCloudformationOutput) TemplateBucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetStackCloudformation) string { return v.TemplateBucket }).(pulumi.StringOutput)
}

type GetStackCloudformationArrayOutput struct{ *pulumi.OutputState }

func (GetStackCloudformationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStackCloudformation)(nil)).Elem()
}

func (o GetStackCloudformationArrayOutput) ToGetStackCloudformationArrayOutput() GetStackCloudformationArrayOutput {
	return o
}

func (o GetStackCloudformationArrayOutput) ToGetStackCloudformationArrayOutputWithContext(ctx context.Context) GetStackCloudformationArrayOutput {
	return o
}

func (o GetStackCloudformationArrayOutput) Index(i pulumi.IntInput) GetStackCloudformationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetStackCloudformation {
		return vs[0].([]GetStackCloudformation)[vs[1].(int)]
	}).(GetStackCloudformationOutput)
}

type GetStackGitlab struct {
	Namespace string `pulumi:"namespace"`
}

// GetStackGitlabInput is an input type that accepts GetStackGitlabArgs and GetStackGitlabOutput values.
// You can construct a concrete instance of `GetStackGitlabInput` via:
//
//          GetStackGitlabArgs{...}
type GetStackGitlabInput interface {
	pulumi.Input

	ToGetStackGitlabOutput() GetStackGitlabOutput
	ToGetStackGitlabOutputWithContext(context.Context) GetStackGitlabOutput
}

type GetStackGitlabArgs struct {
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (GetStackGitlabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStackGitlab)(nil)).Elem()
}

func (i GetStackGitlabArgs) ToGetStackGitlabOutput() GetStackGitlabOutput {
	return i.ToGetStackGitlabOutputWithContext(context.Background())
}

func (i GetStackGitlabArgs) ToGetStackGitlabOutputWithContext(ctx context.Context) GetStackGitlabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStackGitlabOutput)
}

// GetStackGitlabArrayInput is an input type that accepts GetStackGitlabArray and GetStackGitlabArrayOutput values.
// You can construct a concrete instance of `GetStackGitlabArrayInput` via:
//
//          GetStackGitlabArray{ GetStackGitlabArgs{...} }
type GetStackGitlabArrayInput interface {
	pulumi.Input

	ToGetStackGitlabArrayOutput() GetStackGitlabArrayOutput
	ToGetStackGitlabArrayOutputWithContext(context.Context) GetStackGitlabArrayOutput
}

type GetStackGitlabArray []GetStackGitlabInput

func (GetStackGitlabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStackGitlab)(nil)).Elem()
}

func (i GetStackGitlabArray) ToGetStackGitlabArrayOutput() GetStackGitlabArrayOutput {
	return i.ToGetStackGitlabArrayOutputWithContext(context.Background())
}

func (i GetStackGitlabArray) ToGetStackGitlabArrayOutputWithContext(ctx context.Context) GetStackGitlabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStackGitlabArrayOutput)
}

type GetStackGitlabOutput struct{ *pulumi.OutputState }

func (GetStackGitlabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStackGitlab)(nil)).Elem()
}

func (o GetStackGitlabOutput) ToGetStackGitlabOutput() GetStackGitlabOutput {
	return o
}

func (o GetStackGitlabOutput) ToGetStackGitlabOutputWithContext(ctx context.Context) GetStackGitlabOutput {
	return o
}

func (o GetStackGitlabOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetStackGitlab) string { return v.Namespace }).(pulumi.StringOutput)
}

type GetStackGitlabArrayOutput struct{ *pulumi.OutputState }

func (GetStackGitlabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStackGitlab)(nil)).Elem()
}

func (o GetStackGitlabArrayOutput) ToGetStackGitlabArrayOutput() GetStackGitlabArrayOutput {
	return o
}

func (o GetStackGitlabArrayOutput) ToGetStackGitlabArrayOutputWithContext(ctx context.Context) GetStackGitlabArrayOutput {
	return o
}

func (o GetStackGitlabArrayOutput) Index(i pulumi.IntInput) GetStackGitlabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetStackGitlab {
		return vs[0].([]GetStackGitlab)[vs[1].(int)]
	}).(GetStackGitlabOutput)
}

type GetStackPulumi struct {
	LoginUrl  string `pulumi:"loginUrl"`
	StackName string `pulumi:"stackName"`
}

// GetStackPulumiInput is an input type that accepts GetStackPulumiArgs and GetStackPulumiOutput values.
// You can construct a concrete instance of `GetStackPulumiInput` via:
//
//          GetStackPulumiArgs{...}
type GetStackPulumiInput interface {
	pulumi.Input

	ToGetStackPulumiOutput() GetStackPulumiOutput
	ToGetStackPulumiOutputWithContext(context.Context) GetStackPulumiOutput
}

type GetStackPulumiArgs struct {
	LoginUrl  pulumi.StringInput `pulumi:"loginUrl"`
	StackName pulumi.StringInput `pulumi:"stackName"`
}

func (GetStackPulumiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStackPulumi)(nil)).Elem()
}

func (i GetStackPulumiArgs) ToGetStackPulumiOutput() GetStackPulumiOutput {
	return i.ToGetStackPulumiOutputWithContext(context.Background())
}

func (i GetStackPulumiArgs) ToGetStackPulumiOutputWithContext(ctx context.Context) GetStackPulumiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStackPulumiOutput)
}

// GetStackPulumiArrayInput is an input type that accepts GetStackPulumiArray and GetStackPulumiArrayOutput values.
// You can construct a concrete instance of `GetStackPulumiArrayInput` via:
//
//          GetStackPulumiArray{ GetStackPulumiArgs{...} }
type GetStackPulumiArrayInput interface {
	pulumi.Input

	ToGetStackPulumiArrayOutput() GetStackPulumiArrayOutput
	ToGetStackPulumiArrayOutputWithContext(context.Context) GetStackPulumiArrayOutput
}

type GetStackPulumiArray []GetStackPulumiInput

func (GetStackPulumiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStackPulumi)(nil)).Elem()
}

func (i GetStackPulumiArray) ToGetStackPulumiArrayOutput() GetStackPulumiArrayOutput {
	return i.ToGetStackPulumiArrayOutputWithContext(context.Background())
}

func (i GetStackPulumiArray) ToGetStackPulumiArrayOutputWithContext(ctx context.Context) GetStackPulumiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetStackPulumiArrayOutput)
}

type GetStackPulumiOutput struct{ *pulumi.OutputState }

func (GetStackPulumiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStackPulumi)(nil)).Elem()
}

func (o GetStackPulumiOutput) ToGetStackPulumiOutput() GetStackPulumiOutput {
	return o
}

func (o GetStackPulumiOutput) ToGetStackPulumiOutputWithContext(ctx context.Context) GetStackPulumiOutput {
	return o
}

func (o GetStackPulumiOutput) LoginUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetStackPulumi) string { return v.LoginUrl }).(pulumi.StringOutput)
}

func (o GetStackPulumiOutput) StackName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStackPulumi) string { return v.StackName }).(pulumi.StringOutput)
}

type GetStackPulumiArrayOutput struct{ *pulumi.OutputState }

func (GetStackPulumiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetStackPulumi)(nil)).Elem()
}

func (o GetStackPulumiArrayOutput) ToGetStackPulumiArrayOutput() GetStackPulumiArrayOutput {
	return o
}

func (o GetStackPulumiArrayOutput) ToGetStackPulumiArrayOutputWithContext(ctx context.Context) GetStackPulumiArrayOutput {
	return o
}

func (o GetStackPulumiArrayOutput) Index(i pulumi.IntInput) GetStackPulumiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetStackPulumi {
		return vs[0].([]GetStackPulumi)[vs[1].(int)]
	}).(GetStackPulumiOutput)
}

type GetWorkerPoolsWorkerPool struct {
	Config       string `pulumi:"config"`
	Description  string `pulumi:"description"`
	Name         string `pulumi:"name"`
	WorkerPoolId string `pulumi:"workerPoolId"`
}

// GetWorkerPoolsWorkerPoolInput is an input type that accepts GetWorkerPoolsWorkerPoolArgs and GetWorkerPoolsWorkerPoolOutput values.
// You can construct a concrete instance of `GetWorkerPoolsWorkerPoolInput` via:
//
//          GetWorkerPoolsWorkerPoolArgs{...}
type GetWorkerPoolsWorkerPoolInput interface {
	pulumi.Input

	ToGetWorkerPoolsWorkerPoolOutput() GetWorkerPoolsWorkerPoolOutput
	ToGetWorkerPoolsWorkerPoolOutputWithContext(context.Context) GetWorkerPoolsWorkerPoolOutput
}

type GetWorkerPoolsWorkerPoolArgs struct {
	Config       pulumi.StringInput `pulumi:"config"`
	Description  pulumi.StringInput `pulumi:"description"`
	Name         pulumi.StringInput `pulumi:"name"`
	WorkerPoolId pulumi.StringInput `pulumi:"workerPoolId"`
}

func (GetWorkerPoolsWorkerPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkerPoolsWorkerPool)(nil)).Elem()
}

func (i GetWorkerPoolsWorkerPoolArgs) ToGetWorkerPoolsWorkerPoolOutput() GetWorkerPoolsWorkerPoolOutput {
	return i.ToGetWorkerPoolsWorkerPoolOutputWithContext(context.Background())
}

func (i GetWorkerPoolsWorkerPoolArgs) ToGetWorkerPoolsWorkerPoolOutputWithContext(ctx context.Context) GetWorkerPoolsWorkerPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkerPoolsWorkerPoolOutput)
}

// GetWorkerPoolsWorkerPoolArrayInput is an input type that accepts GetWorkerPoolsWorkerPoolArray and GetWorkerPoolsWorkerPoolArrayOutput values.
// You can construct a concrete instance of `GetWorkerPoolsWorkerPoolArrayInput` via:
//
//          GetWorkerPoolsWorkerPoolArray{ GetWorkerPoolsWorkerPoolArgs{...} }
type GetWorkerPoolsWorkerPoolArrayInput interface {
	pulumi.Input

	ToGetWorkerPoolsWorkerPoolArrayOutput() GetWorkerPoolsWorkerPoolArrayOutput
	ToGetWorkerPoolsWorkerPoolArrayOutputWithContext(context.Context) GetWorkerPoolsWorkerPoolArrayOutput
}

type GetWorkerPoolsWorkerPoolArray []GetWorkerPoolsWorkerPoolInput

func (GetWorkerPoolsWorkerPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkerPoolsWorkerPool)(nil)).Elem()
}

func (i GetWorkerPoolsWorkerPoolArray) ToGetWorkerPoolsWorkerPoolArrayOutput() GetWorkerPoolsWorkerPoolArrayOutput {
	return i.ToGetWorkerPoolsWorkerPoolArrayOutputWithContext(context.Background())
}

func (i GetWorkerPoolsWorkerPoolArray) ToGetWorkerPoolsWorkerPoolArrayOutputWithContext(ctx context.Context) GetWorkerPoolsWorkerPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkerPoolsWorkerPoolArrayOutput)
}

type GetWorkerPoolsWorkerPoolOutput struct{ *pulumi.OutputState }

func (GetWorkerPoolsWorkerPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkerPoolsWorkerPool)(nil)).Elem()
}

func (o GetWorkerPoolsWorkerPoolOutput) ToGetWorkerPoolsWorkerPoolOutput() GetWorkerPoolsWorkerPoolOutput {
	return o
}

func (o GetWorkerPoolsWorkerPoolOutput) ToGetWorkerPoolsWorkerPoolOutputWithContext(ctx context.Context) GetWorkerPoolsWorkerPoolOutput {
	return o
}

func (o GetWorkerPoolsWorkerPoolOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkerPoolsWorkerPool) string { return v.Config }).(pulumi.StringOutput)
}

func (o GetWorkerPoolsWorkerPoolOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkerPoolsWorkerPool) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetWorkerPoolsWorkerPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkerPoolsWorkerPool) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetWorkerPoolsWorkerPoolOutput) WorkerPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkerPoolsWorkerPool) string { return v.WorkerPoolId }).(pulumi.StringOutput)
}

type GetWorkerPoolsWorkerPoolArrayOutput struct{ *pulumi.OutputState }

func (GetWorkerPoolsWorkerPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkerPoolsWorkerPool)(nil)).Elem()
}

func (o GetWorkerPoolsWorkerPoolArrayOutput) ToGetWorkerPoolsWorkerPoolArrayOutput() GetWorkerPoolsWorkerPoolArrayOutput {
	return o
}

func (o GetWorkerPoolsWorkerPoolArrayOutput) ToGetWorkerPoolsWorkerPoolArrayOutputWithContext(ctx context.Context) GetWorkerPoolsWorkerPoolArrayOutput {
	return o
}

func (o GetWorkerPoolsWorkerPoolArrayOutput) Index(i pulumi.IntInput) GetWorkerPoolsWorkerPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetWorkerPoolsWorkerPool {
		return vs[0].([]GetWorkerPoolsWorkerPool)[vs[1].(int)]
	}).(GetWorkerPoolsWorkerPoolOutput)
}

func init() {
	pulumi.RegisterOutputType(ModuleGitlabOutput{})
	pulumi.RegisterOutputType(ModuleGitlabPtrOutput{})
	pulumi.RegisterOutputType(StackCloudformationOutput{})
	pulumi.RegisterOutputType(StackCloudformationPtrOutput{})
	pulumi.RegisterOutputType(StackGitlabOutput{})
	pulumi.RegisterOutputType(StackGitlabPtrOutput{})
	pulumi.RegisterOutputType(StackPulumiOutput{})
	pulumi.RegisterOutputType(StackPulumiPtrOutput{})
	pulumi.RegisterOutputType(GetModuleGitlabOutput{})
	pulumi.RegisterOutputType(GetModuleGitlabArrayOutput{})
	pulumi.RegisterOutputType(GetStackCloudformationOutput{})
	pulumi.RegisterOutputType(GetStackCloudformationArrayOutput{})
	pulumi.RegisterOutputType(GetStackGitlabOutput{})
	pulumi.RegisterOutputType(GetStackGitlabArrayOutput{})
	pulumi.RegisterOutputType(GetStackPulumiOutput{})
	pulumi.RegisterOutputType(GetStackPulumiArrayOutput{})
	pulumi.RegisterOutputType(GetWorkerPoolsWorkerPoolOutput{})
	pulumi.RegisterOutputType(GetWorkerPoolsWorkerPoolArrayOutput{})
}
