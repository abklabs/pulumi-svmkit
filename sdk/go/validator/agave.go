// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package validator

import (
	"context"
	"reflect"

	"errors"
	"github.com/abklabs/pulumi-svmkit/sdk/go/agave"
	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/abklabs/pulumi-svmkit/sdk/go/ssh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Agave struct {
	pulumi.CustomResourceState

	Connection ssh.ConnectionOutput   `pulumi:"connection"`
	Flags      agave.FlagsOutput      `pulumi:"flags"`
	KeyPairs   agave.KeyPairsOutput   `pulumi:"keyPairs"`
	Metrics    agave.MetricsPtrOutput `pulumi:"metrics"`
	Variant    agave.VariantPtrOutput `pulumi:"variant"`
	Version    pulumi.StringPtrOutput `pulumi:"version"`
}

// NewAgave registers a new resource with the given unique name, arguments, and options.
func NewAgave(ctx *pulumi.Context,
	name string, args *AgaveArgs, opts ...pulumi.ResourceOption) (*Agave, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Flags == nil {
		return nil, errors.New("invalid value for required argument 'Flags'")
	}
	if args.KeyPairs == nil {
		return nil, errors.New("invalid value for required argument 'KeyPairs'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v ssh.Connection) ssh.Connection { return *v.Defaults() }).(ssh.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Agave
	err := ctx.RegisterResource("svmkit:validator:Agave", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgave gets an existing Agave resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgave(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgaveState, opts ...pulumi.ResourceOption) (*Agave, error) {
	var resource Agave
	err := ctx.ReadResource("svmkit:validator:Agave", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agave resources.
type agaveState struct {
}

type AgaveState struct {
}

func (AgaveState) ElementType() reflect.Type {
	return reflect.TypeOf((*agaveState)(nil)).Elem()
}

type agaveArgs struct {
	Connection ssh.Connection `pulumi:"connection"`
	Flags      agave.Flags    `pulumi:"flags"`
	KeyPairs   agave.KeyPairs `pulumi:"keyPairs"`
	Metrics    *agave.Metrics `pulumi:"metrics"`
	Variant    *agave.Variant `pulumi:"variant"`
	Version    *string        `pulumi:"version"`
}

// The set of arguments for constructing a Agave resource.
type AgaveArgs struct {
	Connection ssh.ConnectionInput
	Flags      agave.FlagsInput
	KeyPairs   agave.KeyPairsInput
	Metrics    agave.MetricsPtrInput
	Variant    agave.VariantPtrInput
	Version    pulumi.StringPtrInput
}

func (AgaveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agaveArgs)(nil)).Elem()
}

type AgaveInput interface {
	pulumi.Input

	ToAgaveOutput() AgaveOutput
	ToAgaveOutputWithContext(ctx context.Context) AgaveOutput
}

func (*Agave) ElementType() reflect.Type {
	return reflect.TypeOf((**Agave)(nil)).Elem()
}

func (i *Agave) ToAgaveOutput() AgaveOutput {
	return i.ToAgaveOutputWithContext(context.Background())
}

func (i *Agave) ToAgaveOutputWithContext(ctx context.Context) AgaveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgaveOutput)
}

type AgaveOutput struct{ *pulumi.OutputState }

func (AgaveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agave)(nil)).Elem()
}

func (o AgaveOutput) ToAgaveOutput() AgaveOutput {
	return o
}

func (o AgaveOutput) ToAgaveOutputWithContext(ctx context.Context) AgaveOutput {
	return o
}

func (o AgaveOutput) Connection() ssh.ConnectionOutput {
	return o.ApplyT(func(v *Agave) ssh.ConnectionOutput { return v.Connection }).(ssh.ConnectionOutput)
}

func (o AgaveOutput) Flags() agave.FlagsOutput {
	return o.ApplyT(func(v *Agave) agave.FlagsOutput { return v.Flags }).(agave.FlagsOutput)
}

func (o AgaveOutput) KeyPairs() agave.KeyPairsOutput {
	return o.ApplyT(func(v *Agave) agave.KeyPairsOutput { return v.KeyPairs }).(agave.KeyPairsOutput)
}

func (o AgaveOutput) Metrics() agave.MetricsPtrOutput {
	return o.ApplyT(func(v *Agave) agave.MetricsPtrOutput { return v.Metrics }).(agave.MetricsPtrOutput)
}

func (o AgaveOutput) Variant() agave.VariantPtrOutput {
	return o.ApplyT(func(v *Agave) agave.VariantPtrOutput { return v.Variant }).(agave.VariantPtrOutput)
}

func (o AgaveOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Agave) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgaveInput)(nil)).Elem(), &Agave{})
	pulumi.RegisterOutputType(AgaveOutput{})
}
