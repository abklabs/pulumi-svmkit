// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package agave

import (
	"context"
	"reflect"

	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type Flags struct {
	AllowPrivateAddr             *bool    `pulumi:"allowPrivateAddr"`
	BlockProductionMethod        string   `pulumi:"blockProductionMethod"`
	DynamicPortRange             string   `pulumi:"dynamicPortRange"`
	EntryPoint                   []string `pulumi:"entryPoint"`
	ExpectedGenesisHash          *string  `pulumi:"expectedGenesisHash"`
	ExtraFlags                   []string `pulumi:"extraFlags"`
	FullRpcAPI                   *bool    `pulumi:"fullRpcAPI"`
	FullSnapshotIntervalSlots    int      `pulumi:"fullSnapshotIntervalSlots"`
	GossipHost                   *string  `pulumi:"gossipHost"`
	GossipPort                   int      `pulumi:"gossipPort"`
	KnownValidator               []string `pulumi:"knownValidator"`
	LimitLedgerSize              int      `pulumi:"limitLedgerSize"`
	NoVoting                     *bool    `pulumi:"noVoting"`
	NoWaitForVoteToStartLeader   bool     `pulumi:"noWaitForVoteToStartLeader"`
	OnlyKnownRPC                 bool     `pulumi:"onlyKnownRPC"`
	PrivateRPC                   bool     `pulumi:"privateRPC"`
	RpcBindAddress               string   `pulumi:"rpcBindAddress"`
	RpcPort                      int      `pulumi:"rpcPort"`
	TvuReceiveThreads            *int     `pulumi:"tvuReceiveThreads"`
	UseSnapshotArchivesAtStartup string   `pulumi:"useSnapshotArchivesAtStartup"`
	WalRecoveryMode              string   `pulumi:"walRecoveryMode"`
}

// FlagsInput is an input type that accepts FlagsArgs and FlagsOutput values.
// You can construct a concrete instance of `FlagsInput` via:
//
//	FlagsArgs{...}
type FlagsInput interface {
	pulumi.Input

	ToFlagsOutput() FlagsOutput
	ToFlagsOutputWithContext(context.Context) FlagsOutput
}

type FlagsArgs struct {
	AllowPrivateAddr             pulumi.BoolPtrInput     `pulumi:"allowPrivateAddr"`
	BlockProductionMethod        pulumi.StringInput      `pulumi:"blockProductionMethod"`
	DynamicPortRange             pulumi.StringInput      `pulumi:"dynamicPortRange"`
	EntryPoint                   pulumi.StringArrayInput `pulumi:"entryPoint"`
	ExpectedGenesisHash          pulumi.StringPtrInput   `pulumi:"expectedGenesisHash"`
	ExtraFlags                   pulumi.StringArrayInput `pulumi:"extraFlags"`
	FullRpcAPI                   pulumi.BoolPtrInput     `pulumi:"fullRpcAPI"`
	FullSnapshotIntervalSlots    pulumi.IntInput         `pulumi:"fullSnapshotIntervalSlots"`
	GossipHost                   pulumi.StringPtrInput   `pulumi:"gossipHost"`
	GossipPort                   pulumi.IntInput         `pulumi:"gossipPort"`
	KnownValidator               pulumi.StringArrayInput `pulumi:"knownValidator"`
	LimitLedgerSize              pulumi.IntInput         `pulumi:"limitLedgerSize"`
	NoVoting                     pulumi.BoolPtrInput     `pulumi:"noVoting"`
	NoWaitForVoteToStartLeader   pulumi.BoolInput        `pulumi:"noWaitForVoteToStartLeader"`
	OnlyKnownRPC                 pulumi.BoolInput        `pulumi:"onlyKnownRPC"`
	PrivateRPC                   pulumi.BoolInput        `pulumi:"privateRPC"`
	RpcBindAddress               pulumi.StringInput      `pulumi:"rpcBindAddress"`
	RpcPort                      pulumi.IntInput         `pulumi:"rpcPort"`
	TvuReceiveThreads            pulumi.IntPtrInput      `pulumi:"tvuReceiveThreads"`
	UseSnapshotArchivesAtStartup pulumi.StringInput      `pulumi:"useSnapshotArchivesAtStartup"`
	WalRecoveryMode              pulumi.StringInput      `pulumi:"walRecoveryMode"`
}

func (FlagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Flags)(nil)).Elem()
}

func (i FlagsArgs) ToFlagsOutput() FlagsOutput {
	return i.ToFlagsOutputWithContext(context.Background())
}

func (i FlagsArgs) ToFlagsOutputWithContext(ctx context.Context) FlagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlagsOutput)
}

type FlagsOutput struct{ *pulumi.OutputState }

func (FlagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Flags)(nil)).Elem()
}

func (o FlagsOutput) ToFlagsOutput() FlagsOutput {
	return o
}

func (o FlagsOutput) ToFlagsOutputWithContext(ctx context.Context) FlagsOutput {
	return o
}

func (o FlagsOutput) AllowPrivateAddr() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Flags) *bool { return v.AllowPrivateAddr }).(pulumi.BoolPtrOutput)
}

func (o FlagsOutput) BlockProductionMethod() pulumi.StringOutput {
	return o.ApplyT(func(v Flags) string { return v.BlockProductionMethod }).(pulumi.StringOutput)
}

func (o FlagsOutput) DynamicPortRange() pulumi.StringOutput {
	return o.ApplyT(func(v Flags) string { return v.DynamicPortRange }).(pulumi.StringOutput)
}

func (o FlagsOutput) EntryPoint() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Flags) []string { return v.EntryPoint }).(pulumi.StringArrayOutput)
}

func (o FlagsOutput) ExpectedGenesisHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Flags) *string { return v.ExpectedGenesisHash }).(pulumi.StringPtrOutput)
}

func (o FlagsOutput) ExtraFlags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Flags) []string { return v.ExtraFlags }).(pulumi.StringArrayOutput)
}

func (o FlagsOutput) FullRpcAPI() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Flags) *bool { return v.FullRpcAPI }).(pulumi.BoolPtrOutput)
}

func (o FlagsOutput) FullSnapshotIntervalSlots() pulumi.IntOutput {
	return o.ApplyT(func(v Flags) int { return v.FullSnapshotIntervalSlots }).(pulumi.IntOutput)
}

func (o FlagsOutput) GossipHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Flags) *string { return v.GossipHost }).(pulumi.StringPtrOutput)
}

func (o FlagsOutput) GossipPort() pulumi.IntOutput {
	return o.ApplyT(func(v Flags) int { return v.GossipPort }).(pulumi.IntOutput)
}

func (o FlagsOutput) KnownValidator() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Flags) []string { return v.KnownValidator }).(pulumi.StringArrayOutput)
}

func (o FlagsOutput) LimitLedgerSize() pulumi.IntOutput {
	return o.ApplyT(func(v Flags) int { return v.LimitLedgerSize }).(pulumi.IntOutput)
}

func (o FlagsOutput) NoVoting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Flags) *bool { return v.NoVoting }).(pulumi.BoolPtrOutput)
}

func (o FlagsOutput) NoWaitForVoteToStartLeader() pulumi.BoolOutput {
	return o.ApplyT(func(v Flags) bool { return v.NoWaitForVoteToStartLeader }).(pulumi.BoolOutput)
}

func (o FlagsOutput) OnlyKnownRPC() pulumi.BoolOutput {
	return o.ApplyT(func(v Flags) bool { return v.OnlyKnownRPC }).(pulumi.BoolOutput)
}

func (o FlagsOutput) PrivateRPC() pulumi.BoolOutput {
	return o.ApplyT(func(v Flags) bool { return v.PrivateRPC }).(pulumi.BoolOutput)
}

func (o FlagsOutput) RpcBindAddress() pulumi.StringOutput {
	return o.ApplyT(func(v Flags) string { return v.RpcBindAddress }).(pulumi.StringOutput)
}

func (o FlagsOutput) RpcPort() pulumi.IntOutput {
	return o.ApplyT(func(v Flags) int { return v.RpcPort }).(pulumi.IntOutput)
}

func (o FlagsOutput) TvuReceiveThreads() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Flags) *int { return v.TvuReceiveThreads }).(pulumi.IntPtrOutput)
}

func (o FlagsOutput) UseSnapshotArchivesAtStartup() pulumi.StringOutput {
	return o.ApplyT(func(v Flags) string { return v.UseSnapshotArchivesAtStartup }).(pulumi.StringOutput)
}

func (o FlagsOutput) WalRecoveryMode() pulumi.StringOutput {
	return o.ApplyT(func(v Flags) string { return v.WalRecoveryMode }).(pulumi.StringOutput)
}

type KeyPairs struct {
	Identity    string `pulumi:"identity"`
	VoteAccount string `pulumi:"voteAccount"`
}

// KeyPairsInput is an input type that accepts KeyPairsArgs and KeyPairsOutput values.
// You can construct a concrete instance of `KeyPairsInput` via:
//
//	KeyPairsArgs{...}
type KeyPairsInput interface {
	pulumi.Input

	ToKeyPairsOutput() KeyPairsOutput
	ToKeyPairsOutputWithContext(context.Context) KeyPairsOutput
}

type KeyPairsArgs struct {
	Identity    pulumi.StringInput `pulumi:"identity"`
	VoteAccount pulumi.StringInput `pulumi:"voteAccount"`
}

func (KeyPairsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPairs)(nil)).Elem()
}

func (i KeyPairsArgs) ToKeyPairsOutput() KeyPairsOutput {
	return i.ToKeyPairsOutputWithContext(context.Background())
}

func (i KeyPairsArgs) ToKeyPairsOutputWithContext(ctx context.Context) KeyPairsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairsOutput)
}

type KeyPairsOutput struct{ *pulumi.OutputState }

func (KeyPairsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPairs)(nil)).Elem()
}

func (o KeyPairsOutput) ToKeyPairsOutput() KeyPairsOutput {
	return o
}

func (o KeyPairsOutput) ToKeyPairsOutputWithContext(ctx context.Context) KeyPairsOutput {
	return o
}

func (o KeyPairsOutput) Identity() pulumi.StringOutput {
	return o.ApplyT(func(v KeyPairs) string { return v.Identity }).(pulumi.StringOutput)
}

func (o KeyPairsOutput) VoteAccount() pulumi.StringOutput {
	return o.ApplyT(func(v KeyPairs) string { return v.VoteAccount }).(pulumi.StringOutput)
}

type Metrics struct {
	Database string `pulumi:"database"`
	Password string `pulumi:"password"`
	Url      string `pulumi:"url"`
	User     string `pulumi:"user"`
}

// MetricsInput is an input type that accepts MetricsArgs and MetricsOutput values.
// You can construct a concrete instance of `MetricsInput` via:
//
//	MetricsArgs{...}
type MetricsInput interface {
	pulumi.Input

	ToMetricsOutput() MetricsOutput
	ToMetricsOutputWithContext(context.Context) MetricsOutput
}

type MetricsArgs struct {
	Database pulumi.StringInput `pulumi:"database"`
	Password pulumi.StringInput `pulumi:"password"`
	Url      pulumi.StringInput `pulumi:"url"`
	User     pulumi.StringInput `pulumi:"user"`
}

func (MetricsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Metrics)(nil)).Elem()
}

func (i MetricsArgs) ToMetricsOutput() MetricsOutput {
	return i.ToMetricsOutputWithContext(context.Background())
}

func (i MetricsArgs) ToMetricsOutputWithContext(ctx context.Context) MetricsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsOutput)
}

func (i MetricsArgs) ToMetricsPtrOutput() MetricsPtrOutput {
	return i.ToMetricsPtrOutputWithContext(context.Background())
}

func (i MetricsArgs) ToMetricsPtrOutputWithContext(ctx context.Context) MetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsOutput).ToMetricsPtrOutputWithContext(ctx)
}

// MetricsPtrInput is an input type that accepts MetricsArgs, MetricsPtr and MetricsPtrOutput values.
// You can construct a concrete instance of `MetricsPtrInput` via:
//
//	        MetricsArgs{...}
//
//	or:
//
//	        nil
type MetricsPtrInput interface {
	pulumi.Input

	ToMetricsPtrOutput() MetricsPtrOutput
	ToMetricsPtrOutputWithContext(context.Context) MetricsPtrOutput
}

type metricsPtrType MetricsArgs

func MetricsPtr(v *MetricsArgs) MetricsPtrInput {
	return (*metricsPtrType)(v)
}

func (*metricsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Metrics)(nil)).Elem()
}

func (i *metricsPtrType) ToMetricsPtrOutput() MetricsPtrOutput {
	return i.ToMetricsPtrOutputWithContext(context.Background())
}

func (i *metricsPtrType) ToMetricsPtrOutputWithContext(ctx context.Context) MetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsPtrOutput)
}

type MetricsOutput struct{ *pulumi.OutputState }

func (MetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Metrics)(nil)).Elem()
}

func (o MetricsOutput) ToMetricsOutput() MetricsOutput {
	return o
}

func (o MetricsOutput) ToMetricsOutputWithContext(ctx context.Context) MetricsOutput {
	return o
}

func (o MetricsOutput) ToMetricsPtrOutput() MetricsPtrOutput {
	return o.ToMetricsPtrOutputWithContext(context.Background())
}

func (o MetricsOutput) ToMetricsPtrOutputWithContext(ctx context.Context) MetricsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Metrics) *Metrics {
		return &v
	}).(MetricsPtrOutput)
}

func (o MetricsOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v Metrics) string { return v.Database }).(pulumi.StringOutput)
}

func (o MetricsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v Metrics) string { return v.Password }).(pulumi.StringOutput)
}

func (o MetricsOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v Metrics) string { return v.Url }).(pulumi.StringOutput)
}

func (o MetricsOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v Metrics) string { return v.User }).(pulumi.StringOutput)
}

type MetricsPtrOutput struct{ *pulumi.OutputState }

func (MetricsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Metrics)(nil)).Elem()
}

func (o MetricsPtrOutput) ToMetricsPtrOutput() MetricsPtrOutput {
	return o
}

func (o MetricsPtrOutput) ToMetricsPtrOutputWithContext(ctx context.Context) MetricsPtrOutput {
	return o
}

func (o MetricsPtrOutput) Elem() MetricsOutput {
	return o.ApplyT(func(v *Metrics) Metrics {
		if v != nil {
			return *v
		}
		var ret Metrics
		return ret
	}).(MetricsOutput)
}

func (o MetricsPtrOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metrics) *string {
		if v == nil {
			return nil
		}
		return &v.Database
	}).(pulumi.StringPtrOutput)
}

func (o MetricsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metrics) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o MetricsPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metrics) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

func (o MetricsPtrOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metrics) *string {
		if v == nil {
			return nil
		}
		return &v.User
	}).(pulumi.StringPtrOutput)
}

type TimeoutConfig struct {
	RpcServiceTimeout *int `pulumi:"rpcServiceTimeout"`
}

// TimeoutConfigInput is an input type that accepts TimeoutConfigArgs and TimeoutConfigOutput values.
// You can construct a concrete instance of `TimeoutConfigInput` via:
//
//	TimeoutConfigArgs{...}
type TimeoutConfigInput interface {
	pulumi.Input

	ToTimeoutConfigOutput() TimeoutConfigOutput
	ToTimeoutConfigOutputWithContext(context.Context) TimeoutConfigOutput
}

type TimeoutConfigArgs struct {
	RpcServiceTimeout pulumi.IntPtrInput `pulumi:"rpcServiceTimeout"`
}

func (TimeoutConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeoutConfig)(nil)).Elem()
}

func (i TimeoutConfigArgs) ToTimeoutConfigOutput() TimeoutConfigOutput {
	return i.ToTimeoutConfigOutputWithContext(context.Background())
}

func (i TimeoutConfigArgs) ToTimeoutConfigOutputWithContext(ctx context.Context) TimeoutConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeoutConfigOutput)
}

func (i TimeoutConfigArgs) ToTimeoutConfigPtrOutput() TimeoutConfigPtrOutput {
	return i.ToTimeoutConfigPtrOutputWithContext(context.Background())
}

func (i TimeoutConfigArgs) ToTimeoutConfigPtrOutputWithContext(ctx context.Context) TimeoutConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeoutConfigOutput).ToTimeoutConfigPtrOutputWithContext(ctx)
}

// TimeoutConfigPtrInput is an input type that accepts TimeoutConfigArgs, TimeoutConfigPtr and TimeoutConfigPtrOutput values.
// You can construct a concrete instance of `TimeoutConfigPtrInput` via:
//
//	        TimeoutConfigArgs{...}
//
//	or:
//
//	        nil
type TimeoutConfigPtrInput interface {
	pulumi.Input

	ToTimeoutConfigPtrOutput() TimeoutConfigPtrOutput
	ToTimeoutConfigPtrOutputWithContext(context.Context) TimeoutConfigPtrOutput
}

type timeoutConfigPtrType TimeoutConfigArgs

func TimeoutConfigPtr(v *TimeoutConfigArgs) TimeoutConfigPtrInput {
	return (*timeoutConfigPtrType)(v)
}

func (*timeoutConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeoutConfig)(nil)).Elem()
}

func (i *timeoutConfigPtrType) ToTimeoutConfigPtrOutput() TimeoutConfigPtrOutput {
	return i.ToTimeoutConfigPtrOutputWithContext(context.Background())
}

func (i *timeoutConfigPtrType) ToTimeoutConfigPtrOutputWithContext(ctx context.Context) TimeoutConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeoutConfigPtrOutput)
}

type TimeoutConfigOutput struct{ *pulumi.OutputState }

func (TimeoutConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeoutConfig)(nil)).Elem()
}

func (o TimeoutConfigOutput) ToTimeoutConfigOutput() TimeoutConfigOutput {
	return o
}

func (o TimeoutConfigOutput) ToTimeoutConfigOutputWithContext(ctx context.Context) TimeoutConfigOutput {
	return o
}

func (o TimeoutConfigOutput) ToTimeoutConfigPtrOutput() TimeoutConfigPtrOutput {
	return o.ToTimeoutConfigPtrOutputWithContext(context.Background())
}

func (o TimeoutConfigOutput) ToTimeoutConfigPtrOutputWithContext(ctx context.Context) TimeoutConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeoutConfig) *TimeoutConfig {
		return &v
	}).(TimeoutConfigPtrOutput)
}

func (o TimeoutConfigOutput) RpcServiceTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TimeoutConfig) *int { return v.RpcServiceTimeout }).(pulumi.IntPtrOutput)
}

type TimeoutConfigPtrOutput struct{ *pulumi.OutputState }

func (TimeoutConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeoutConfig)(nil)).Elem()
}

func (o TimeoutConfigPtrOutput) ToTimeoutConfigPtrOutput() TimeoutConfigPtrOutput {
	return o
}

func (o TimeoutConfigPtrOutput) ToTimeoutConfigPtrOutputWithContext(ctx context.Context) TimeoutConfigPtrOutput {
	return o
}

func (o TimeoutConfigPtrOutput) Elem() TimeoutConfigOutput {
	return o.ApplyT(func(v *TimeoutConfig) TimeoutConfig {
		if v != nil {
			return *v
		}
		var ret TimeoutConfig
		return ret
	}).(TimeoutConfigOutput)
}

func (o TimeoutConfigPtrOutput) RpcServiceTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeoutConfig) *int {
		if v == nil {
			return nil
		}
		return v.RpcServiceTimeout
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlagsInput)(nil)).Elem(), FlagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairsInput)(nil)).Elem(), KeyPairsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricsInput)(nil)).Elem(), MetricsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricsPtrInput)(nil)).Elem(), MetricsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimeoutConfigInput)(nil)).Elem(), TimeoutConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimeoutConfigPtrInput)(nil)).Elem(), TimeoutConfigArgs{})
	pulumi.RegisterOutputType(FlagsOutput{})
	pulumi.RegisterOutputType(KeyPairsOutput{})
	pulumi.RegisterOutputType(MetricsOutput{})
	pulumi.RegisterOutputType(MetricsPtrOutput{})
	pulumi.RegisterOutputType(TimeoutConfigOutput{})
	pulumi.RegisterOutputType(TimeoutConfigPtrOutput{})
}
