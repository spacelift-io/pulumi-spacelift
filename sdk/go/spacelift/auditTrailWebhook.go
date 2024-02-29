// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spacelift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/spacelift-io/pulumi-spacelift/sdk/v2/go/spacelift/internal"
)

// `AuditTrailWebhook` represents a webhook endpoint to which Spacelift sends POST requests about audit events.
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
//			_, err := spacelift.NewAuditTrailWebhook(ctx, "example", &spacelift.AuditTrailWebhookArgs{
//				Enabled:  pulumi.Bool(true),
//				Endpoint: pulumi.String("https://example.com"),
//				Secret:   pulumi.String("mysecretkey"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AuditTrailWebhook struct {
	pulumi.CustomResourceState

	// `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// `endpoint` is the URL to which Spacelift will send POST requests about audit events.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// `includeRuns` determines whether the webhook should include information about the run that triggered the event.
	IncludeRuns pulumi.BoolPtrOutput `pulumi:"includeRuns"`
	// `secret` is a secret that Spacelift will send with the request
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewAuditTrailWebhook registers a new resource with the given unique name, arguments, and options.
func NewAuditTrailWebhook(ctx *pulumi.Context,
	name string, args *AuditTrailWebhookArgs, opts ...pulumi.ResourceOption) (*AuditTrailWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.Secret == nil {
		return nil, errors.New("invalid value for required argument 'Secret'")
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuditTrailWebhook
	err := ctx.RegisterResource("spacelift:index/auditTrailWebhook:AuditTrailWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuditTrailWebhook gets an existing AuditTrailWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuditTrailWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditTrailWebhookState, opts ...pulumi.ResourceOption) (*AuditTrailWebhook, error) {
	var resource AuditTrailWebhook
	err := ctx.ReadResource("spacelift:index/auditTrailWebhook:AuditTrailWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuditTrailWebhook resources.
type auditTrailWebhookState struct {
	// `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
	Enabled *bool `pulumi:"enabled"`
	// `endpoint` is the URL to which Spacelift will send POST requests about audit events.
	Endpoint *string `pulumi:"endpoint"`
	// `includeRuns` determines whether the webhook should include information about the run that triggered the event.
	IncludeRuns *bool `pulumi:"includeRuns"`
	// `secret` is a secret that Spacelift will send with the request
	Secret *string `pulumi:"secret"`
}

type AuditTrailWebhookState struct {
	// `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
	Enabled pulumi.BoolPtrInput
	// `endpoint` is the URL to which Spacelift will send POST requests about audit events.
	Endpoint pulumi.StringPtrInput
	// `includeRuns` determines whether the webhook should include information about the run that triggered the event.
	IncludeRuns pulumi.BoolPtrInput
	// `secret` is a secret that Spacelift will send with the request
	Secret pulumi.StringPtrInput
}

func (AuditTrailWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditTrailWebhookState)(nil)).Elem()
}

type auditTrailWebhookArgs struct {
	// `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
	Enabled bool `pulumi:"enabled"`
	// `endpoint` is the URL to which Spacelift will send POST requests about audit events.
	Endpoint string `pulumi:"endpoint"`
	// `includeRuns` determines whether the webhook should include information about the run that triggered the event.
	IncludeRuns *bool `pulumi:"includeRuns"`
	// `secret` is a secret that Spacelift will send with the request
	Secret string `pulumi:"secret"`
}

// The set of arguments for constructing a AuditTrailWebhook resource.
type AuditTrailWebhookArgs struct {
	// `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
	Enabled pulumi.BoolInput
	// `endpoint` is the URL to which Spacelift will send POST requests about audit events.
	Endpoint pulumi.StringInput
	// `includeRuns` determines whether the webhook should include information about the run that triggered the event.
	IncludeRuns pulumi.BoolPtrInput
	// `secret` is a secret that Spacelift will send with the request
	Secret pulumi.StringInput
}

func (AuditTrailWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditTrailWebhookArgs)(nil)).Elem()
}

type AuditTrailWebhookInput interface {
	pulumi.Input

	ToAuditTrailWebhookOutput() AuditTrailWebhookOutput
	ToAuditTrailWebhookOutputWithContext(ctx context.Context) AuditTrailWebhookOutput
}

func (*AuditTrailWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditTrailWebhook)(nil)).Elem()
}

func (i *AuditTrailWebhook) ToAuditTrailWebhookOutput() AuditTrailWebhookOutput {
	return i.ToAuditTrailWebhookOutputWithContext(context.Background())
}

func (i *AuditTrailWebhook) ToAuditTrailWebhookOutputWithContext(ctx context.Context) AuditTrailWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditTrailWebhookOutput)
}

// AuditTrailWebhookArrayInput is an input type that accepts AuditTrailWebhookArray and AuditTrailWebhookArrayOutput values.
// You can construct a concrete instance of `AuditTrailWebhookArrayInput` via:
//
//	AuditTrailWebhookArray{ AuditTrailWebhookArgs{...} }
type AuditTrailWebhookArrayInput interface {
	pulumi.Input

	ToAuditTrailWebhookArrayOutput() AuditTrailWebhookArrayOutput
	ToAuditTrailWebhookArrayOutputWithContext(context.Context) AuditTrailWebhookArrayOutput
}

type AuditTrailWebhookArray []AuditTrailWebhookInput

func (AuditTrailWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditTrailWebhook)(nil)).Elem()
}

func (i AuditTrailWebhookArray) ToAuditTrailWebhookArrayOutput() AuditTrailWebhookArrayOutput {
	return i.ToAuditTrailWebhookArrayOutputWithContext(context.Background())
}

func (i AuditTrailWebhookArray) ToAuditTrailWebhookArrayOutputWithContext(ctx context.Context) AuditTrailWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditTrailWebhookArrayOutput)
}

// AuditTrailWebhookMapInput is an input type that accepts AuditTrailWebhookMap and AuditTrailWebhookMapOutput values.
// You can construct a concrete instance of `AuditTrailWebhookMapInput` via:
//
//	AuditTrailWebhookMap{ "key": AuditTrailWebhookArgs{...} }
type AuditTrailWebhookMapInput interface {
	pulumi.Input

	ToAuditTrailWebhookMapOutput() AuditTrailWebhookMapOutput
	ToAuditTrailWebhookMapOutputWithContext(context.Context) AuditTrailWebhookMapOutput
}

type AuditTrailWebhookMap map[string]AuditTrailWebhookInput

func (AuditTrailWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditTrailWebhook)(nil)).Elem()
}

func (i AuditTrailWebhookMap) ToAuditTrailWebhookMapOutput() AuditTrailWebhookMapOutput {
	return i.ToAuditTrailWebhookMapOutputWithContext(context.Background())
}

func (i AuditTrailWebhookMap) ToAuditTrailWebhookMapOutputWithContext(ctx context.Context) AuditTrailWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditTrailWebhookMapOutput)
}

type AuditTrailWebhookOutput struct{ *pulumi.OutputState }

func (AuditTrailWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditTrailWebhook)(nil)).Elem()
}

func (o AuditTrailWebhookOutput) ToAuditTrailWebhookOutput() AuditTrailWebhookOutput {
	return o
}

func (o AuditTrailWebhookOutput) ToAuditTrailWebhookOutputWithContext(ctx context.Context) AuditTrailWebhookOutput {
	return o
}

// `enabled` determines whether the webhook is enabled. If it is not, Spacelift will not send any requests to the endpoint.
func (o AuditTrailWebhookOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AuditTrailWebhook) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// `endpoint` is the URL to which Spacelift will send POST requests about audit events.
func (o AuditTrailWebhookOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditTrailWebhook) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// `includeRuns` determines whether the webhook should include information about the run that triggered the event.
func (o AuditTrailWebhookOutput) IncludeRuns() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuditTrailWebhook) pulumi.BoolPtrOutput { return v.IncludeRuns }).(pulumi.BoolPtrOutput)
}

// `secret` is a secret that Spacelift will send with the request
func (o AuditTrailWebhookOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditTrailWebhook) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

type AuditTrailWebhookArrayOutput struct{ *pulumi.OutputState }

func (AuditTrailWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditTrailWebhook)(nil)).Elem()
}

func (o AuditTrailWebhookArrayOutput) ToAuditTrailWebhookArrayOutput() AuditTrailWebhookArrayOutput {
	return o
}

func (o AuditTrailWebhookArrayOutput) ToAuditTrailWebhookArrayOutputWithContext(ctx context.Context) AuditTrailWebhookArrayOutput {
	return o
}

func (o AuditTrailWebhookArrayOutput) Index(i pulumi.IntInput) AuditTrailWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuditTrailWebhook {
		return vs[0].([]*AuditTrailWebhook)[vs[1].(int)]
	}).(AuditTrailWebhookOutput)
}

type AuditTrailWebhookMapOutput struct{ *pulumi.OutputState }

func (AuditTrailWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditTrailWebhook)(nil)).Elem()
}

func (o AuditTrailWebhookMapOutput) ToAuditTrailWebhookMapOutput() AuditTrailWebhookMapOutput {
	return o
}

func (o AuditTrailWebhookMapOutput) ToAuditTrailWebhookMapOutputWithContext(ctx context.Context) AuditTrailWebhookMapOutput {
	return o
}

func (o AuditTrailWebhookMapOutput) MapIndex(k pulumi.StringInput) AuditTrailWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuditTrailWebhook {
		return vs[0].(map[string]*AuditTrailWebhook)[vs[1].(string)]
	}).(AuditTrailWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditTrailWebhookInput)(nil)).Elem(), &AuditTrailWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditTrailWebhookArrayInput)(nil)).Elem(), AuditTrailWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditTrailWebhookMapInput)(nil)).Elem(), AuditTrailWebhookMap{})
	pulumi.RegisterOutputType(AuditTrailWebhookOutput{})
	pulumi.RegisterOutputType(AuditTrailWebhookArrayOutput{})
	pulumi.RegisterOutputType(AuditTrailWebhookMapOutput{})
}
