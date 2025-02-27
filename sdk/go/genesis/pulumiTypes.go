// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package genesis

import (
	"context"
	"reflect"

	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type BootstrapAccount struct {
	BalanceLamports *int   `pulumi:"balanceLamports"`
	IdentityPubkey  string `pulumi:"identityPubkey"`
	StakeLamports   *int   `pulumi:"stakeLamports"`
	StakePubkey     string `pulumi:"stakePubkey"`
	VotePubkey      string `pulumi:"votePubkey"`
}

// BootstrapAccountInput is an input type that accepts BootstrapAccountArgs and BootstrapAccountOutput values.
// You can construct a concrete instance of `BootstrapAccountInput` via:
//
//	BootstrapAccountArgs{...}
type BootstrapAccountInput interface {
	pulumi.Input

	ToBootstrapAccountOutput() BootstrapAccountOutput
	ToBootstrapAccountOutputWithContext(context.Context) BootstrapAccountOutput
}

type BootstrapAccountArgs struct {
	BalanceLamports pulumi.IntPtrInput `pulumi:"balanceLamports"`
	IdentityPubkey  pulumi.StringInput `pulumi:"identityPubkey"`
	StakeLamports   pulumi.IntPtrInput `pulumi:"stakeLamports"`
	StakePubkey     pulumi.StringInput `pulumi:"stakePubkey"`
	VotePubkey      pulumi.StringInput `pulumi:"votePubkey"`
}

func (BootstrapAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BootstrapAccount)(nil)).Elem()
}

func (i BootstrapAccountArgs) ToBootstrapAccountOutput() BootstrapAccountOutput {
	return i.ToBootstrapAccountOutputWithContext(context.Background())
}

func (i BootstrapAccountArgs) ToBootstrapAccountOutputWithContext(ctx context.Context) BootstrapAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapAccountOutput)
}

// BootstrapAccountArrayInput is an input type that accepts BootstrapAccountArray and BootstrapAccountArrayOutput values.
// You can construct a concrete instance of `BootstrapAccountArrayInput` via:
//
//	BootstrapAccountArray{ BootstrapAccountArgs{...} }
type BootstrapAccountArrayInput interface {
	pulumi.Input

	ToBootstrapAccountArrayOutput() BootstrapAccountArrayOutput
	ToBootstrapAccountArrayOutputWithContext(context.Context) BootstrapAccountArrayOutput
}

type BootstrapAccountArray []BootstrapAccountInput

func (BootstrapAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BootstrapAccount)(nil)).Elem()
}

func (i BootstrapAccountArray) ToBootstrapAccountArrayOutput() BootstrapAccountArrayOutput {
	return i.ToBootstrapAccountArrayOutputWithContext(context.Background())
}

func (i BootstrapAccountArray) ToBootstrapAccountArrayOutputWithContext(ctx context.Context) BootstrapAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapAccountArrayOutput)
}

type BootstrapAccountOutput struct{ *pulumi.OutputState }

func (BootstrapAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootstrapAccount)(nil)).Elem()
}

func (o BootstrapAccountOutput) ToBootstrapAccountOutput() BootstrapAccountOutput {
	return o
}

func (o BootstrapAccountOutput) ToBootstrapAccountOutputWithContext(ctx context.Context) BootstrapAccountOutput {
	return o
}

func (o BootstrapAccountOutput) BalanceLamports() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BootstrapAccount) *int { return v.BalanceLamports }).(pulumi.IntPtrOutput)
}

func (o BootstrapAccountOutput) IdentityPubkey() pulumi.StringOutput {
	return o.ApplyT(func(v BootstrapAccount) string { return v.IdentityPubkey }).(pulumi.StringOutput)
}

func (o BootstrapAccountOutput) StakeLamports() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BootstrapAccount) *int { return v.StakeLamports }).(pulumi.IntPtrOutput)
}

func (o BootstrapAccountOutput) StakePubkey() pulumi.StringOutput {
	return o.ApplyT(func(v BootstrapAccount) string { return v.StakePubkey }).(pulumi.StringOutput)
}

func (o BootstrapAccountOutput) VotePubkey() pulumi.StringOutput {
	return o.ApplyT(func(v BootstrapAccount) string { return v.VotePubkey }).(pulumi.StringOutput)
}

type BootstrapAccountArrayOutput struct{ *pulumi.OutputState }

func (BootstrapAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BootstrapAccount)(nil)).Elem()
}

func (o BootstrapAccountArrayOutput) ToBootstrapAccountArrayOutput() BootstrapAccountArrayOutput {
	return o
}

func (o BootstrapAccountArrayOutput) ToBootstrapAccountArrayOutputWithContext(ctx context.Context) BootstrapAccountArrayOutput {
	return o
}

func (o BootstrapAccountArrayOutput) Index(i pulumi.IntInput) BootstrapAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BootstrapAccount {
		return vs[0].([]BootstrapAccount)[vs[1].(int)]
	}).(BootstrapAccountOutput)
}

type GenesisFlags struct {
	BootstrapStakeAuthorizedPubkey  *string  `pulumi:"bootstrapStakeAuthorizedPubkey"`
	BootstrapValidatorLamports      *int     `pulumi:"bootstrapValidatorLamports"`
	BootstrapValidatorStakeLamports *int     `pulumi:"bootstrapValidatorStakeLamports"`
	ClusterType                     *string  `pulumi:"clusterType"`
	CreationTime                    *string  `pulumi:"creationTime"`
	DeactivateFeatures              []string `pulumi:"deactivateFeatures"`
	EnableWarmupEpochs              *bool    `pulumi:"enableWarmupEpochs"`
	ExtraFlags                      []string `pulumi:"extraFlags"`
	FaucetLamports                  *int     `pulumi:"faucetLamports"`
	FaucetPubkey                    *string  `pulumi:"faucetPubkey"`
	FeeBurnPercentage               *int     `pulumi:"feeBurnPercentage"`
	HashesPerTick                   *string  `pulumi:"hashesPerTick"`
	IdentityPubkey                  string   `pulumi:"identityPubkey"`
	Inflation                       *string  `pulumi:"inflation"`
	LamportsPerByteYear             *int     `pulumi:"lamportsPerByteYear"`
	LedgerPath                      string   `pulumi:"ledgerPath"`
	MaxGenesisArchiveUnpackedSize   *int     `pulumi:"maxGenesisArchiveUnpackedSize"`
	RentBurnPercentage              *int     `pulumi:"rentBurnPercentage"`
	RentExemptionThreshold          *int     `pulumi:"rentExemptionThreshold"`
	SlotsPerEpoch                   *int     `pulumi:"slotsPerEpoch"`
	StakePubkey                     string   `pulumi:"stakePubkey"`
	TargetLamportsPerSignature      *int     `pulumi:"targetLamportsPerSignature"`
	TargetSignaturesPerSlot         *int     `pulumi:"targetSignaturesPerSlot"`
	TargetTickDuration              *int     `pulumi:"targetTickDuration"`
	TicksPerSlot                    *int     `pulumi:"ticksPerSlot"`
	Url                             *string  `pulumi:"url"`
	VoteCommissionPercentage        *int     `pulumi:"voteCommissionPercentage"`
	VotePubkey                      string   `pulumi:"votePubkey"`
}

// GenesisFlagsInput is an input type that accepts GenesisFlagsArgs and GenesisFlagsOutput values.
// You can construct a concrete instance of `GenesisFlagsInput` via:
//
//	GenesisFlagsArgs{...}
type GenesisFlagsInput interface {
	pulumi.Input

	ToGenesisFlagsOutput() GenesisFlagsOutput
	ToGenesisFlagsOutputWithContext(context.Context) GenesisFlagsOutput
}

type GenesisFlagsArgs struct {
	BootstrapStakeAuthorizedPubkey  pulumi.StringPtrInput   `pulumi:"bootstrapStakeAuthorizedPubkey"`
	BootstrapValidatorLamports      pulumi.IntPtrInput      `pulumi:"bootstrapValidatorLamports"`
	BootstrapValidatorStakeLamports pulumi.IntPtrInput      `pulumi:"bootstrapValidatorStakeLamports"`
	ClusterType                     pulumi.StringPtrInput   `pulumi:"clusterType"`
	CreationTime                    pulumi.StringPtrInput   `pulumi:"creationTime"`
	DeactivateFeatures              pulumi.StringArrayInput `pulumi:"deactivateFeatures"`
	EnableWarmupEpochs              pulumi.BoolPtrInput     `pulumi:"enableWarmupEpochs"`
	ExtraFlags                      pulumi.StringArrayInput `pulumi:"extraFlags"`
	FaucetLamports                  pulumi.IntPtrInput      `pulumi:"faucetLamports"`
	FaucetPubkey                    pulumi.StringPtrInput   `pulumi:"faucetPubkey"`
	FeeBurnPercentage               pulumi.IntPtrInput      `pulumi:"feeBurnPercentage"`
	HashesPerTick                   pulumi.StringPtrInput   `pulumi:"hashesPerTick"`
	IdentityPubkey                  pulumi.StringInput      `pulumi:"identityPubkey"`
	Inflation                       pulumi.StringPtrInput   `pulumi:"inflation"`
	LamportsPerByteYear             pulumi.IntPtrInput      `pulumi:"lamportsPerByteYear"`
	LedgerPath                      pulumi.StringInput      `pulumi:"ledgerPath"`
	MaxGenesisArchiveUnpackedSize   pulumi.IntPtrInput      `pulumi:"maxGenesisArchiveUnpackedSize"`
	RentBurnPercentage              pulumi.IntPtrInput      `pulumi:"rentBurnPercentage"`
	RentExemptionThreshold          pulumi.IntPtrInput      `pulumi:"rentExemptionThreshold"`
	SlotsPerEpoch                   pulumi.IntPtrInput      `pulumi:"slotsPerEpoch"`
	StakePubkey                     pulumi.StringInput      `pulumi:"stakePubkey"`
	TargetLamportsPerSignature      pulumi.IntPtrInput      `pulumi:"targetLamportsPerSignature"`
	TargetSignaturesPerSlot         pulumi.IntPtrInput      `pulumi:"targetSignaturesPerSlot"`
	TargetTickDuration              pulumi.IntPtrInput      `pulumi:"targetTickDuration"`
	TicksPerSlot                    pulumi.IntPtrInput      `pulumi:"ticksPerSlot"`
	Url                             pulumi.StringPtrInput   `pulumi:"url"`
	VoteCommissionPercentage        pulumi.IntPtrInput      `pulumi:"voteCommissionPercentage"`
	VotePubkey                      pulumi.StringInput      `pulumi:"votePubkey"`
}

func (GenesisFlagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GenesisFlags)(nil)).Elem()
}

func (i GenesisFlagsArgs) ToGenesisFlagsOutput() GenesisFlagsOutput {
	return i.ToGenesisFlagsOutputWithContext(context.Background())
}

func (i GenesisFlagsArgs) ToGenesisFlagsOutputWithContext(ctx context.Context) GenesisFlagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenesisFlagsOutput)
}

type GenesisFlagsOutput struct{ *pulumi.OutputState }

func (GenesisFlagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GenesisFlags)(nil)).Elem()
}

func (o GenesisFlagsOutput) ToGenesisFlagsOutput() GenesisFlagsOutput {
	return o
}

func (o GenesisFlagsOutput) ToGenesisFlagsOutputWithContext(ctx context.Context) GenesisFlagsOutput {
	return o
}

func (o GenesisFlagsOutput) BootstrapStakeAuthorizedPubkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.BootstrapStakeAuthorizedPubkey }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) BootstrapValidatorLamports() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.BootstrapValidatorLamports }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) BootstrapValidatorStakeLamports() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.BootstrapValidatorStakeLamports }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) ClusterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.ClusterType }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) DeactivateFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GenesisFlags) []string { return v.DeactivateFeatures }).(pulumi.StringArrayOutput)
}

func (o GenesisFlagsOutput) EnableWarmupEpochs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *bool { return v.EnableWarmupEpochs }).(pulumi.BoolPtrOutput)
}

func (o GenesisFlagsOutput) ExtraFlags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GenesisFlags) []string { return v.ExtraFlags }).(pulumi.StringArrayOutput)
}

func (o GenesisFlagsOutput) FaucetLamports() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.FaucetLamports }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) FaucetPubkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.FaucetPubkey }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) FeeBurnPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.FeeBurnPercentage }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) HashesPerTick() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.HashesPerTick }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) IdentityPubkey() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.IdentityPubkey }).(pulumi.StringOutput)
}

func (o GenesisFlagsOutput) Inflation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.Inflation }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) LamportsPerByteYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.LamportsPerByteYear }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) LedgerPath() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.LedgerPath }).(pulumi.StringOutput)
}

func (o GenesisFlagsOutput) MaxGenesisArchiveUnpackedSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.MaxGenesisArchiveUnpackedSize }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) RentBurnPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.RentBurnPercentage }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) RentExemptionThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.RentExemptionThreshold }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) SlotsPerEpoch() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.SlotsPerEpoch }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) StakePubkey() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.StakePubkey }).(pulumi.StringOutput)
}

func (o GenesisFlagsOutput) TargetLamportsPerSignature() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.TargetLamportsPerSignature }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) TargetSignaturesPerSlot() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.TargetSignaturesPerSlot }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) TargetTickDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.TargetTickDuration }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) TicksPerSlot() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.TicksPerSlot }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) VoteCommissionPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *int { return v.VoteCommissionPercentage }).(pulumi.IntPtrOutput)
}

func (o GenesisFlagsOutput) VotePubkey() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.VotePubkey }).(pulumi.StringOutput)
}

type PrimordialAccount struct {
	Data       *string `pulumi:"data"`
	Executable *bool   `pulumi:"executable"`
	Lamports   int     `pulumi:"lamports"`
	Owner      *string `pulumi:"owner"`
	Pubkey     string  `pulumi:"pubkey"`
}

// PrimordialAccountInput is an input type that accepts PrimordialAccountArgs and PrimordialAccountOutput values.
// You can construct a concrete instance of `PrimordialAccountInput` via:
//
//	PrimordialAccountArgs{...}
type PrimordialAccountInput interface {
	pulumi.Input

	ToPrimordialAccountOutput() PrimordialAccountOutput
	ToPrimordialAccountOutputWithContext(context.Context) PrimordialAccountOutput
}

type PrimordialAccountArgs struct {
	Data       pulumi.StringPtrInput `pulumi:"data"`
	Executable pulumi.BoolPtrInput   `pulumi:"executable"`
	Lamports   pulumi.IntInput       `pulumi:"lamports"`
	Owner      pulumi.StringPtrInput `pulumi:"owner"`
	Pubkey     pulumi.StringInput    `pulumi:"pubkey"`
}

func (PrimordialAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimordialAccount)(nil)).Elem()
}

func (i PrimordialAccountArgs) ToPrimordialAccountOutput() PrimordialAccountOutput {
	return i.ToPrimordialAccountOutputWithContext(context.Background())
}

func (i PrimordialAccountArgs) ToPrimordialAccountOutputWithContext(ctx context.Context) PrimordialAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrimordialAccountOutput)
}

// PrimordialAccountArrayInput is an input type that accepts PrimordialAccountArray and PrimordialAccountArrayOutput values.
// You can construct a concrete instance of `PrimordialAccountArrayInput` via:
//
//	PrimordialAccountArray{ PrimordialAccountArgs{...} }
type PrimordialAccountArrayInput interface {
	pulumi.Input

	ToPrimordialAccountArrayOutput() PrimordialAccountArrayOutput
	ToPrimordialAccountArrayOutputWithContext(context.Context) PrimordialAccountArrayOutput
}

type PrimordialAccountArray []PrimordialAccountInput

func (PrimordialAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrimordialAccount)(nil)).Elem()
}

func (i PrimordialAccountArray) ToPrimordialAccountArrayOutput() PrimordialAccountArrayOutput {
	return i.ToPrimordialAccountArrayOutputWithContext(context.Background())
}

func (i PrimordialAccountArray) ToPrimordialAccountArrayOutputWithContext(ctx context.Context) PrimordialAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrimordialAccountArrayOutput)
}

type PrimordialAccountOutput struct{ *pulumi.OutputState }

func (PrimordialAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimordialAccount)(nil)).Elem()
}

func (o PrimordialAccountOutput) ToPrimordialAccountOutput() PrimordialAccountOutput {
	return o
}

func (o PrimordialAccountOutput) ToPrimordialAccountOutputWithContext(ctx context.Context) PrimordialAccountOutput {
	return o
}

func (o PrimordialAccountOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrimordialAccount) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o PrimordialAccountOutput) Executable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PrimordialAccount) *bool { return v.Executable }).(pulumi.BoolPtrOutput)
}

func (o PrimordialAccountOutput) Lamports() pulumi.IntOutput {
	return o.ApplyT(func(v PrimordialAccount) int { return v.Lamports }).(pulumi.IntOutput)
}

func (o PrimordialAccountOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrimordialAccount) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o PrimordialAccountOutput) Pubkey() pulumi.StringOutput {
	return o.ApplyT(func(v PrimordialAccount) string { return v.Pubkey }).(pulumi.StringOutput)
}

type PrimordialAccountArrayOutput struct{ *pulumi.OutputState }

func (PrimordialAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrimordialAccount)(nil)).Elem()
}

func (o PrimordialAccountArrayOutput) ToPrimordialAccountArrayOutput() PrimordialAccountArrayOutput {
	return o
}

func (o PrimordialAccountArrayOutput) ToPrimordialAccountArrayOutputWithContext(ctx context.Context) PrimordialAccountArrayOutput {
	return o
}

func (o PrimordialAccountArrayOutput) Index(i pulumi.IntInput) PrimordialAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrimordialAccount {
		return vs[0].([]PrimordialAccount)[vs[1].(int)]
	}).(PrimordialAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapAccountInput)(nil)).Elem(), BootstrapAccountArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapAccountArrayInput)(nil)).Elem(), BootstrapAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenesisFlagsInput)(nil)).Elem(), GenesisFlagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrimordialAccountInput)(nil)).Elem(), PrimordialAccountArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrimordialAccountArrayInput)(nil)).Elem(), PrimordialAccountArray{})
	pulumi.RegisterOutputType(BootstrapAccountOutput{})
	pulumi.RegisterOutputType(BootstrapAccountArrayOutput{})
	pulumi.RegisterOutputType(GenesisFlagsOutput{})
	pulumi.RegisterOutputType(PrimordialAccountOutput{})
	pulumi.RegisterOutputType(PrimordialAccountArrayOutput{})
}
