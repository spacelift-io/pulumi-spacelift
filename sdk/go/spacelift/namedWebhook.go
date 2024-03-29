// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift/internal"
)

// `NamedWebhook` represents a named webhook endpoint used for creating webhookswhich are referred to in Notification policies to route messages.
type NamedWebhook struct {
	pulumi.CustomResourceState

	// enables or disables sending webhooks.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// endpoint to send the requests to
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// labels for the webhook to use when referring in policies or filtering them
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// the name for the webhook which will also be used to generate the id
	Name pulumi.StringOutput `pulumi:"name"`
	// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// ID of the space the webhook is in
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
}

// NewNamedWebhook registers a new resource with the given unique name, arguments, and options.
func NewNamedWebhook(ctx *pulumi.Context,
	name string, args *NamedWebhookArgs, opts ...pulumi.ResourceOption) (*NamedWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.SpaceId == nil {
		return nil, errors.New("invalid value for required argument 'SpaceId'")
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NamedWebhook
	err := ctx.RegisterResource("spacelift:index/namedWebhook:NamedWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamedWebhook gets an existing NamedWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamedWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamedWebhookState, opts ...pulumi.ResourceOption) (*NamedWebhook, error) {
	var resource NamedWebhook
	err := ctx.ReadResource("spacelift:index/namedWebhook:NamedWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamedWebhook resources.
type namedWebhookState struct {
	// enables or disables sending webhooks.
	Enabled *bool `pulumi:"enabled"`
	// endpoint to send the requests to
	Endpoint *string `pulumi:"endpoint"`
	// labels for the webhook to use when referring in policies or filtering them
	Labels []string `pulumi:"labels"`
	// the name for the webhook which will also be used to generate the id
	Name *string `pulumi:"name"`
	// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret *string `pulumi:"secret"`
	// ID of the space the webhook is in
	SpaceId *string `pulumi:"spaceId"`
}

type NamedWebhookState struct {
	// enables or disables sending webhooks.
	Enabled pulumi.BoolPtrInput
	// endpoint to send the requests to
	Endpoint pulumi.StringPtrInput
	// labels for the webhook to use when referring in policies or filtering them
	Labels pulumi.StringArrayInput
	// the name for the webhook which will also be used to generate the id
	Name pulumi.StringPtrInput
	// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret pulumi.StringPtrInput
	// ID of the space the webhook is in
	SpaceId pulumi.StringPtrInput
}

func (NamedWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*namedWebhookState)(nil)).Elem()
}

type namedWebhookArgs struct {
	// enables or disables sending webhooks.
	Enabled bool `pulumi:"enabled"`
	// endpoint to send the requests to
	Endpoint string `pulumi:"endpoint"`
	// labels for the webhook to use when referring in policies or filtering them
	Labels []string `pulumi:"labels"`
	// the name for the webhook which will also be used to generate the id
	Name *string `pulumi:"name"`
	// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret *string `pulumi:"secret"`
	// ID of the space the webhook is in
	SpaceId string `pulumi:"spaceId"`
}

// The set of arguments for constructing a NamedWebhook resource.
type NamedWebhookArgs struct {
	// enables or disables sending webhooks.
	Enabled pulumi.BoolInput
	// endpoint to send the requests to
	Endpoint pulumi.StringInput
	// labels for the webhook to use when referring in policies or filtering them
	Labels pulumi.StringArrayInput
	// the name for the webhook which will also be used to generate the id
	Name pulumi.StringPtrInput
	// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret pulumi.StringPtrInput
	// ID of the space the webhook is in
	SpaceId pulumi.StringInput
}

func (NamedWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namedWebhookArgs)(nil)).Elem()
}

type NamedWebhookInput interface {
	pulumi.Input

	ToNamedWebhookOutput() NamedWebhookOutput
	ToNamedWebhookOutputWithContext(ctx context.Context) NamedWebhookOutput
}

func (*NamedWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedWebhook)(nil)).Elem()
}

func (i *NamedWebhook) ToNamedWebhookOutput() NamedWebhookOutput {
	return i.ToNamedWebhookOutputWithContext(context.Background())
}

func (i *NamedWebhook) ToNamedWebhookOutputWithContext(ctx context.Context) NamedWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedWebhookOutput)
}

func (i *NamedWebhook) ToOutput(ctx context.Context) pulumix.Output[*NamedWebhook] {
	return pulumix.Output[*NamedWebhook]{
		OutputState: i.ToNamedWebhookOutputWithContext(ctx).OutputState,
	}
}

// NamedWebhookArrayInput is an input type that accepts NamedWebhookArray and NamedWebhookArrayOutput values.
// You can construct a concrete instance of `NamedWebhookArrayInput` via:
//
//	NamedWebhookArray{ NamedWebhookArgs{...} }
type NamedWebhookArrayInput interface {
	pulumi.Input

	ToNamedWebhookArrayOutput() NamedWebhookArrayOutput
	ToNamedWebhookArrayOutputWithContext(context.Context) NamedWebhookArrayOutput
}

type NamedWebhookArray []NamedWebhookInput

func (NamedWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamedWebhook)(nil)).Elem()
}

func (i NamedWebhookArray) ToNamedWebhookArrayOutput() NamedWebhookArrayOutput {
	return i.ToNamedWebhookArrayOutputWithContext(context.Background())
}

func (i NamedWebhookArray) ToNamedWebhookArrayOutputWithContext(ctx context.Context) NamedWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedWebhookArrayOutput)
}

func (i NamedWebhookArray) ToOutput(ctx context.Context) pulumix.Output[[]*NamedWebhook] {
	return pulumix.Output[[]*NamedWebhook]{
		OutputState: i.ToNamedWebhookArrayOutputWithContext(ctx).OutputState,
	}
}

// NamedWebhookMapInput is an input type that accepts NamedWebhookMap and NamedWebhookMapOutput values.
// You can construct a concrete instance of `NamedWebhookMapInput` via:
//
//	NamedWebhookMap{ "key": NamedWebhookArgs{...} }
type NamedWebhookMapInput interface {
	pulumi.Input

	ToNamedWebhookMapOutput() NamedWebhookMapOutput
	ToNamedWebhookMapOutputWithContext(context.Context) NamedWebhookMapOutput
}

type NamedWebhookMap map[string]NamedWebhookInput

func (NamedWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamedWebhook)(nil)).Elem()
}

func (i NamedWebhookMap) ToNamedWebhookMapOutput() NamedWebhookMapOutput {
	return i.ToNamedWebhookMapOutputWithContext(context.Background())
}

func (i NamedWebhookMap) ToNamedWebhookMapOutputWithContext(ctx context.Context) NamedWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedWebhookMapOutput)
}

func (i NamedWebhookMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NamedWebhook] {
	return pulumix.Output[map[string]*NamedWebhook]{
		OutputState: i.ToNamedWebhookMapOutputWithContext(ctx).OutputState,
	}
}

type NamedWebhookOutput struct{ *pulumi.OutputState }

func (NamedWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedWebhook)(nil)).Elem()
}

func (o NamedWebhookOutput) ToNamedWebhookOutput() NamedWebhookOutput {
	return o
}

func (o NamedWebhookOutput) ToNamedWebhookOutputWithContext(ctx context.Context) NamedWebhookOutput {
	return o
}

func (o NamedWebhookOutput) ToOutput(ctx context.Context) pulumix.Output[*NamedWebhook] {
	return pulumix.Output[*NamedWebhook]{
		OutputState: o.OutputState,
	}
}

// enables or disables sending webhooks.
func (o NamedWebhookOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *NamedWebhook) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// endpoint to send the requests to
func (o NamedWebhookOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *NamedWebhook) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// labels for the webhook to use when referring in policies or filtering them
func (o NamedWebhookOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamedWebhook) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// the name for the webhook which will also be used to generate the id
func (o NamedWebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamedWebhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value.
func (o NamedWebhookOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamedWebhook) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

// ID of the space the webhook is in
func (o NamedWebhookOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NamedWebhook) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

type NamedWebhookArrayOutput struct{ *pulumi.OutputState }

func (NamedWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamedWebhook)(nil)).Elem()
}

func (o NamedWebhookArrayOutput) ToNamedWebhookArrayOutput() NamedWebhookArrayOutput {
	return o
}

func (o NamedWebhookArrayOutput) ToNamedWebhookArrayOutputWithContext(ctx context.Context) NamedWebhookArrayOutput {
	return o
}

func (o NamedWebhookArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NamedWebhook] {
	return pulumix.Output[[]*NamedWebhook]{
		OutputState: o.OutputState,
	}
}

func (o NamedWebhookArrayOutput) Index(i pulumi.IntInput) NamedWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NamedWebhook {
		return vs[0].([]*NamedWebhook)[vs[1].(int)]
	}).(NamedWebhookOutput)
}

type NamedWebhookMapOutput struct{ *pulumi.OutputState }

func (NamedWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamedWebhook)(nil)).Elem()
}

func (o NamedWebhookMapOutput) ToNamedWebhookMapOutput() NamedWebhookMapOutput {
	return o
}

func (o NamedWebhookMapOutput) ToNamedWebhookMapOutputWithContext(ctx context.Context) NamedWebhookMapOutput {
	return o
}

func (o NamedWebhookMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NamedWebhook] {
	return pulumix.Output[map[string]*NamedWebhook]{
		OutputState: o.OutputState,
	}
}

func (o NamedWebhookMapOutput) MapIndex(k pulumi.StringInput) NamedWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NamedWebhook {
		return vs[0].(map[string]*NamedWebhook)[vs[1].(string)]
	}).(NamedWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamedWebhookInput)(nil)).Elem(), &NamedWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamedWebhookArrayInput)(nil)).Elem(), NamedWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamedWebhookMapInput)(nil)).Elem(), NamedWebhookMap{})
	pulumi.RegisterOutputType(NamedWebhookOutput{})
	pulumi.RegisterOutputType(NamedWebhookArrayOutput{})
	pulumi.RegisterOutputType(NamedWebhookMapOutput{})
}
