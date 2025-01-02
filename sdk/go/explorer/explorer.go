// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package explorer

import (
	"context"
	"reflect"

	"errors"
	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/abklabs/pulumi-svmkit/sdk/go/solana"
	"github.com/abklabs/pulumi-svmkit/sdk/go/ssh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Explorer struct {
	pulumi.CustomResourceState

	Connection  ssh.ConnectionOutput       `pulumi:"connection"`
	Environment solana.EnvironmentOutput   `pulumi:"environment"`
	Flags       solana.ExplorerFlagsOutput `pulumi:"flags"`
	Version     pulumi.StringPtrOutput     `pulumi:"version"`
}

// NewExplorer registers a new resource with the given unique name, arguments, and options.
func NewExplorer(ctx *pulumi.Context,
	name string, args *ExplorerArgs, opts ...pulumi.ResourceOption) (*Explorer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Flags == nil {
		return nil, errors.New("invalid value for required argument 'Flags'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v ssh.Connection) ssh.Connection { return *v.Defaults() }).(ssh.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Explorer
	err := ctx.RegisterResource("svmkit:explorer:Explorer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExplorer gets an existing Explorer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExplorer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExplorerState, opts ...pulumi.ResourceOption) (*Explorer, error) {
	var resource Explorer
	err := ctx.ReadResource("svmkit:explorer:Explorer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Explorer resources.
type explorerState struct {
}

type ExplorerState struct {
}

func (ExplorerState) ElementType() reflect.Type {
	return reflect.TypeOf((*explorerState)(nil)).Elem()
}

type explorerArgs struct {
	Connection  ssh.Connection       `pulumi:"connection"`
	Environment solana.Environment   `pulumi:"environment"`
	Flags       solana.ExplorerFlags `pulumi:"flags"`
	Version     *string              `pulumi:"version"`
}

// The set of arguments for constructing a Explorer resource.
type ExplorerArgs struct {
	Connection  ssh.ConnectionInput
	Environment solana.EnvironmentInput
	Flags       solana.ExplorerFlagsInput
	Version     pulumi.StringPtrInput
}

func (ExplorerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*explorerArgs)(nil)).Elem()
}

type ExplorerInput interface {
	pulumi.Input

	ToExplorerOutput() ExplorerOutput
	ToExplorerOutputWithContext(ctx context.Context) ExplorerOutput
}

func (*Explorer) ElementType() reflect.Type {
	return reflect.TypeOf((**Explorer)(nil)).Elem()
}

func (i *Explorer) ToExplorerOutput() ExplorerOutput {
	return i.ToExplorerOutputWithContext(context.Background())
}

func (i *Explorer) ToExplorerOutputWithContext(ctx context.Context) ExplorerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExplorerOutput)
}

type ExplorerOutput struct{ *pulumi.OutputState }

func (ExplorerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Explorer)(nil)).Elem()
}

func (o ExplorerOutput) ToExplorerOutput() ExplorerOutput {
	return o
}

func (o ExplorerOutput) ToExplorerOutputWithContext(ctx context.Context) ExplorerOutput {
	return o
}

func (o ExplorerOutput) Connection() ssh.ConnectionOutput {
	return o.ApplyT(func(v *Explorer) ssh.ConnectionOutput { return v.Connection }).(ssh.ConnectionOutput)
}

func (o ExplorerOutput) Environment() solana.EnvironmentOutput {
	return o.ApplyT(func(v *Explorer) solana.EnvironmentOutput { return v.Environment }).(solana.EnvironmentOutput)
}

func (o ExplorerOutput) Flags() solana.ExplorerFlagsOutput {
	return o.ApplyT(func(v *Explorer) solana.ExplorerFlagsOutput { return v.Flags }).(solana.ExplorerFlagsOutput)
}

func (o ExplorerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Explorer) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExplorerInput)(nil)).Elem(), &Explorer{})
	pulumi.RegisterOutputType(ExplorerOutput{})
}
