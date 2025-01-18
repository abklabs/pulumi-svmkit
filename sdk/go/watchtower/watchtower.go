// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package watchtower

import (
	"context"
	"reflect"

	"errors"
	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/abklabs/pulumi-svmkit/sdk/go/solana"
	"github.com/abklabs/pulumi-svmkit/sdk/go/ssh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Watchtower struct {
	pulumi.CustomResourceState

	Connection    ssh.ConnectionOutput     `pulumi:"connection"`
	Environment   solana.EnvironmentOutput `pulumi:"environment"`
	Flags         WatchtowerFlagsOutput    `pulumi:"flags"`
	Notifications NotificationConfigOutput `pulumi:"notifications"`
}

// NewWatchtower registers a new resource with the given unique name, arguments, and options.
func NewWatchtower(ctx *pulumi.Context,
	name string, args *WatchtowerArgs, opts ...pulumi.ResourceOption) (*Watchtower, error) {
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
	if args.Notifications == nil {
		return nil, errors.New("invalid value for required argument 'Notifications'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v ssh.Connection) ssh.Connection { return *v.Defaults() }).(ssh.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Watchtower
	err := ctx.RegisterResource("svmkit:watchtower:Watchtower", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatchtower gets an existing Watchtower resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatchtower(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatchtowerState, opts ...pulumi.ResourceOption) (*Watchtower, error) {
	var resource Watchtower
	err := ctx.ReadResource("svmkit:watchtower:Watchtower", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Watchtower resources.
type watchtowerState struct {
}

type WatchtowerState struct {
}

func (WatchtowerState) ElementType() reflect.Type {
	return reflect.TypeOf((*watchtowerState)(nil)).Elem()
}

type watchtowerArgs struct {
	Connection    ssh.Connection     `pulumi:"connection"`
	Environment   solana.Environment `pulumi:"environment"`
	Flags         WatchtowerFlags    `pulumi:"flags"`
	Notifications NotificationConfig `pulumi:"notifications"`
}

// The set of arguments for constructing a Watchtower resource.
type WatchtowerArgs struct {
	Connection    ssh.ConnectionInput
	Environment   solana.EnvironmentInput
	Flags         WatchtowerFlagsInput
	Notifications NotificationConfigInput
}

func (WatchtowerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watchtowerArgs)(nil)).Elem()
}

type WatchtowerInput interface {
	pulumi.Input

	ToWatchtowerOutput() WatchtowerOutput
	ToWatchtowerOutputWithContext(ctx context.Context) WatchtowerOutput
}

func (*Watchtower) ElementType() reflect.Type {
	return reflect.TypeOf((**Watchtower)(nil)).Elem()
}

func (i *Watchtower) ToWatchtowerOutput() WatchtowerOutput {
	return i.ToWatchtowerOutputWithContext(context.Background())
}

func (i *Watchtower) ToWatchtowerOutputWithContext(ctx context.Context) WatchtowerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchtowerOutput)
}

type WatchtowerOutput struct{ *pulumi.OutputState }

func (WatchtowerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Watchtower)(nil)).Elem()
}

func (o WatchtowerOutput) ToWatchtowerOutput() WatchtowerOutput {
	return o
}

func (o WatchtowerOutput) ToWatchtowerOutputWithContext(ctx context.Context) WatchtowerOutput {
	return o
}

func (o WatchtowerOutput) Connection() ssh.ConnectionOutput {
	return o.ApplyT(func(v *Watchtower) ssh.ConnectionOutput { return v.Connection }).(ssh.ConnectionOutput)
}

func (o WatchtowerOutput) Environment() solana.EnvironmentOutput {
	return o.ApplyT(func(v *Watchtower) solana.EnvironmentOutput { return v.Environment }).(solana.EnvironmentOutput)
}

func (o WatchtowerOutput) Flags() WatchtowerFlagsOutput {
	return o.ApplyT(func(v *Watchtower) WatchtowerFlagsOutput { return v.Flags }).(WatchtowerFlagsOutput)
}

func (o WatchtowerOutput) Notifications() NotificationConfigOutput {
	return o.ApplyT(func(v *Watchtower) NotificationConfigOutput { return v.Notifications }).(NotificationConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WatchtowerInput)(nil)).Elem(), &Watchtower{})
	pulumi.RegisterOutputType(WatchtowerOutput{})
}