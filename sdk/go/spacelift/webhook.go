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

// `Webhook` represents a webhook endpoint to which Spacelift sends the POST request about run state changes.
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
//			_, err := spacelift.NewWebhook(ctx, "webhook", &spacelift.WebhookArgs{
//				Endpoint: pulumi.String("https://example.com/webhooks"),
//				StackId:  pulumi.String("k8s-core"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import spacelift:index/webhook:Webhook webhook stack/$STACK_ID/$WEBHOOK_ID
//
// ```
type Webhook struct {
	pulumi.CustomResourceState

	// enables or disables sending webhooks. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// endpoint to send the POST request to
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// ID of the module which triggers the webhooks
	ModuleId pulumi.StringPtrOutput `pulumi:"moduleId"`
	// secret used to sign each POST request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// ID of the stack which triggers the webhooks
	StackId pulumi.StringPtrOutput `pulumi:"stackId"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Webhook
	err := ctx.RegisterResource("spacelift:index/webhook:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("spacelift:index/webhook:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	// enables or disables sending webhooks. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// endpoint to send the POST request to
	Endpoint *string `pulumi:"endpoint"`
	// ID of the module which triggers the webhooks
	ModuleId *string `pulumi:"moduleId"`
	// secret used to sign each POST request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret *string `pulumi:"secret"`
	// ID of the stack which triggers the webhooks
	StackId *string `pulumi:"stackId"`
}

type WebhookState struct {
	// enables or disables sending webhooks. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// endpoint to send the POST request to
	Endpoint pulumi.StringPtrInput
	// ID of the module which triggers the webhooks
	ModuleId pulumi.StringPtrInput
	// secret used to sign each POST request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret pulumi.StringPtrInput
	// ID of the stack which triggers the webhooks
	StackId pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// enables or disables sending webhooks. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// endpoint to send the POST request to
	Endpoint string `pulumi:"endpoint"`
	// ID of the module which triggers the webhooks
	ModuleId *string `pulumi:"moduleId"`
	// secret used to sign each POST request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret *string `pulumi:"secret"`
	// ID of the stack which triggers the webhooks
	StackId *string `pulumi:"stackId"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// enables or disables sending webhooks. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// endpoint to send the POST request to
	Endpoint pulumi.StringInput
	// ID of the module which triggers the webhooks
	ModuleId pulumi.StringPtrInput
	// secret used to sign each POST request so you're able to verify that the request comes from us. Defaults to an empty value.
	Secret pulumi.StringPtrInput
	// ID of the stack which triggers the webhooks
	StackId pulumi.StringPtrInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

func (i *Webhook) ToOutput(ctx context.Context) pulumix.Output[*Webhook] {
	return pulumix.Output[*Webhook]{
		OutputState: i.ToWebhookOutputWithContext(ctx).OutputState,
	}
}

// WebhookArrayInput is an input type that accepts WebhookArray and WebhookArrayOutput values.
// You can construct a concrete instance of `WebhookArrayInput` via:
//
//	WebhookArray{ WebhookArgs{...} }
type WebhookArrayInput interface {
	pulumi.Input

	ToWebhookArrayOutput() WebhookArrayOutput
	ToWebhookArrayOutputWithContext(context.Context) WebhookArrayOutput
}

type WebhookArray []WebhookInput

func (WebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Webhook)(nil)).Elem()
}

func (i WebhookArray) ToWebhookArrayOutput() WebhookArrayOutput {
	return i.ToWebhookArrayOutputWithContext(context.Background())
}

func (i WebhookArray) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookArrayOutput)
}

func (i WebhookArray) ToOutput(ctx context.Context) pulumix.Output[[]*Webhook] {
	return pulumix.Output[[]*Webhook]{
		OutputState: i.ToWebhookArrayOutputWithContext(ctx).OutputState,
	}
}

// WebhookMapInput is an input type that accepts WebhookMap and WebhookMapOutput values.
// You can construct a concrete instance of `WebhookMapInput` via:
//
//	WebhookMap{ "key": WebhookArgs{...} }
type WebhookMapInput interface {
	pulumi.Input

	ToWebhookMapOutput() WebhookMapOutput
	ToWebhookMapOutputWithContext(context.Context) WebhookMapOutput
}

type WebhookMap map[string]WebhookInput

func (WebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Webhook)(nil)).Elem()
}

func (i WebhookMap) ToWebhookMapOutput() WebhookMapOutput {
	return i.ToWebhookMapOutputWithContext(context.Background())
}

func (i WebhookMap) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookMapOutput)
}

func (i WebhookMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Webhook] {
	return pulumix.Output[map[string]*Webhook]{
		OutputState: i.ToWebhookMapOutputWithContext(ctx).OutputState,
	}
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func (o WebhookOutput) ToOutput(ctx context.Context) pulumix.Output[*Webhook] {
	return pulumix.Output[*Webhook]{
		OutputState: o.OutputState,
	}
}

// enables or disables sending webhooks. Defaults to `true`.
func (o WebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// endpoint to send the POST request to
func (o WebhookOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// ID of the module which triggers the webhooks
func (o WebhookOutput) ModuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.ModuleId }).(pulumi.StringPtrOutput)
}

// secret used to sign each POST request so you're able to verify that the request comes from us. Defaults to an empty value.
func (o WebhookOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

// ID of the stack which triggers the webhooks
func (o WebhookOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.StackId }).(pulumi.StringPtrOutput)
}

type WebhookArrayOutput struct{ *pulumi.OutputState }

func (WebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Webhook)(nil)).Elem()
}

func (o WebhookArrayOutput) ToWebhookArrayOutput() WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Webhook] {
	return pulumix.Output[[]*Webhook]{
		OutputState: o.OutputState,
	}
}

func (o WebhookArrayOutput) Index(i pulumi.IntInput) WebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Webhook {
		return vs[0].([]*Webhook)[vs[1].(int)]
	}).(WebhookOutput)
}

type WebhookMapOutput struct{ *pulumi.OutputState }

func (WebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Webhook)(nil)).Elem()
}

func (o WebhookMapOutput) ToWebhookMapOutput() WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Webhook] {
	return pulumix.Output[map[string]*Webhook]{
		OutputState: o.OutputState,
	}
}

func (o WebhookMapOutput) MapIndex(k pulumi.StringInput) WebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Webhook {
		return vs[0].(map[string]*Webhook)[vs[1].(string)]
	}).(WebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookInput)(nil)).Elem(), &Webhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookArrayInput)(nil)).Elem(), WebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookMapInput)(nil)).Elem(), WebhookMap{})
	pulumi.RegisterOutputType(WebhookOutput{})
	pulumi.RegisterOutputType(WebhookArrayOutput{})
	pulumi.RegisterOutputType(WebhookMapOutput{})
}
