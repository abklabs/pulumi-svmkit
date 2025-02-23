// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package genesis

import (
	"context"
	"reflect"

	"errors"
	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/abklabs/pulumi-svmkit/sdk/go/runner"
	"github.com/abklabs/pulumi-svmkit/sdk/go/ssh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Solana struct {
	pulumi.CustomResourceState

	Accounts     BootstrapAccountArrayOutput  `pulumi:"accounts"`
	Connection   ssh.ConnectionOutput         `pulumi:"connection"`
	Flags        GenesisFlagsOutput           `pulumi:"flags"`
	GenesisHash  pulumi.StringOutput          `pulumi:"genesisHash"`
	Primordial   PrimordialAccountArrayOutput `pulumi:"primordial"`
	RunnerConfig runner.ConfigPtrOutput       `pulumi:"runnerConfig"`
	Version      pulumi.StringPtrOutput       `pulumi:"version"`
}

// NewSolana registers a new resource with the given unique name, arguments, and options.
func NewSolana(ctx *pulumi.Context,
	name string, args *SolanaArgs, opts ...pulumi.ResourceOption) (*Solana, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Flags == nil {
		return nil, errors.New("invalid value for required argument 'Flags'")
	}
	if args.Primordial == nil {
		return nil, errors.New("invalid value for required argument 'Primordial'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v ssh.Connection) ssh.Connection { return *v.Defaults() }).(ssh.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Solana
	err := ctx.RegisterResource("svmkit:genesis:Solana", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSolana gets an existing Solana resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSolana(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SolanaState, opts ...pulumi.ResourceOption) (*Solana, error) {
	var resource Solana
	err := ctx.ReadResource("svmkit:genesis:Solana", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Solana resources.
type solanaState struct {
}

type SolanaState struct {
}

func (SolanaState) ElementType() reflect.Type {
	return reflect.TypeOf((*solanaState)(nil)).Elem()
}

type solanaArgs struct {
	Accounts     []BootstrapAccount  `pulumi:"accounts"`
	Connection   ssh.Connection      `pulumi:"connection"`
	Flags        GenesisFlags        `pulumi:"flags"`
	Primordial   []PrimordialAccount `pulumi:"primordial"`
	RunnerConfig *runner.Config      `pulumi:"runnerConfig"`
	Version      *string             `pulumi:"version"`
}

// The set of arguments for constructing a Solana resource.
type SolanaArgs struct {
	Accounts     BootstrapAccountArrayInput
	Connection   ssh.ConnectionInput
	Flags        GenesisFlagsInput
	Primordial   PrimordialAccountArrayInput
	RunnerConfig runner.ConfigPtrInput
	Version      pulumi.StringPtrInput
}

func (SolanaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*solanaArgs)(nil)).Elem()
}

type SolanaInput interface {
	pulumi.Input

	ToSolanaOutput() SolanaOutput
	ToSolanaOutputWithContext(ctx context.Context) SolanaOutput
}

func (*Solana) ElementType() reflect.Type {
	return reflect.TypeOf((**Solana)(nil)).Elem()
}

func (i *Solana) ToSolanaOutput() SolanaOutput {
	return i.ToSolanaOutputWithContext(context.Background())
}

func (i *Solana) ToSolanaOutputWithContext(ctx context.Context) SolanaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolanaOutput)
}

type SolanaOutput struct{ *pulumi.OutputState }

func (SolanaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Solana)(nil)).Elem()
}

func (o SolanaOutput) ToSolanaOutput() SolanaOutput {
	return o
}

func (o SolanaOutput) ToSolanaOutputWithContext(ctx context.Context) SolanaOutput {
	return o
}

func (o SolanaOutput) Accounts() BootstrapAccountArrayOutput {
	return o.ApplyT(func(v *Solana) BootstrapAccountArrayOutput { return v.Accounts }).(BootstrapAccountArrayOutput)
}

func (o SolanaOutput) Connection() ssh.ConnectionOutput {
	return o.ApplyT(func(v *Solana) ssh.ConnectionOutput { return v.Connection }).(ssh.ConnectionOutput)
}

func (o SolanaOutput) Flags() GenesisFlagsOutput {
	return o.ApplyT(func(v *Solana) GenesisFlagsOutput { return v.Flags }).(GenesisFlagsOutput)
}

func (o SolanaOutput) GenesisHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Solana) pulumi.StringOutput { return v.GenesisHash }).(pulumi.StringOutput)
}

func (o SolanaOutput) Primordial() PrimordialAccountArrayOutput {
	return o.ApplyT(func(v *Solana) PrimordialAccountArrayOutput { return v.Primordial }).(PrimordialAccountArrayOutput)
}

func (o SolanaOutput) RunnerConfig() runner.ConfigPtrOutput {
	return o.ApplyT(func(v *Solana) runner.ConfigPtrOutput { return v.RunnerConfig }).(runner.ConfigPtrOutput)
}

func (o SolanaOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Solana) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SolanaInput)(nil)).Elem(), &Solana{})
	pulumi.RegisterOutputType(SolanaOutput{})
}
