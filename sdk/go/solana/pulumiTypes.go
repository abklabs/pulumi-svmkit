// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solana

import (
	"context"
	"reflect"

	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type Environment struct {
	RpcURL string `pulumi:"rpcURL"`
}

// EnvironmentInput is an input type that accepts EnvironmentArgs and EnvironmentOutput values.
// You can construct a concrete instance of `EnvironmentInput` via:
//
//	EnvironmentArgs{...}
type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(context.Context) EnvironmentOutput
}

type EnvironmentArgs struct {
	RpcURL pulumi.StringInput `pulumi:"rpcURL"`
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil)).Elem()
}

func (i EnvironmentArgs) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i EnvironmentArgs) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

func (i EnvironmentArgs) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return i.ToEnvironmentPtrOutputWithContext(context.Background())
}

func (i EnvironmentArgs) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput).ToEnvironmentPtrOutputWithContext(ctx)
}

// EnvironmentPtrInput is an input type that accepts EnvironmentArgs, EnvironmentPtr and EnvironmentPtrOutput values.
// You can construct a concrete instance of `EnvironmentPtrInput` via:
//
//	        EnvironmentArgs{...}
//
//	or:
//
//	        nil
type EnvironmentPtrInput interface {
	pulumi.Input

	ToEnvironmentPtrOutput() EnvironmentPtrOutput
	ToEnvironmentPtrOutputWithContext(context.Context) EnvironmentPtrOutput
}

type environmentPtrType EnvironmentArgs

func EnvironmentPtr(v *EnvironmentArgs) EnvironmentPtrInput {
	return (*environmentPtrType)(v)
}

func (*environmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *environmentPtrType) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return i.ToEnvironmentPtrOutputWithContext(context.Background())
}

func (i *environmentPtrType) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentPtrOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return o.ToEnvironmentPtrOutputWithContext(context.Background())
}

func (o EnvironmentOutput) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Environment) *Environment {
		return &v
	}).(EnvironmentPtrOutput)
}

func (o EnvironmentOutput) RpcURL() pulumi.StringOutput {
	return o.ApplyT(func(v Environment) string { return v.RpcURL }).(pulumi.StringOutput)
}

type EnvironmentPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentPtrOutput) ToEnvironmentPtrOutput() EnvironmentPtrOutput {
	return o
}

func (o EnvironmentPtrOutput) ToEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentPtrOutput {
	return o
}

func (o EnvironmentPtrOutput) Elem() EnvironmentOutput {
	return o.ApplyT(func(v *Environment) Environment {
		if v != nil {
			return *v
		}
		var ret Environment
		return ret
	}).(EnvironmentOutput)
}

func (o EnvironmentPtrOutput) RpcURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) *string {
		if v == nil {
			return nil
		}
		return &v.RpcURL
	}).(pulumi.StringPtrOutput)
}

type FaucetFlags struct {
	AllowIPs      []string `pulumi:"allowIPs"`
	PerRequestCap *int     `pulumi:"perRequestCap"`
	PerTimeCap    *int     `pulumi:"perTimeCap"`
	SliceSeconds  *int     `pulumi:"sliceSeconds"`
}

// FaucetFlagsInput is an input type that accepts FaucetFlagsArgs and FaucetFlagsOutput values.
// You can construct a concrete instance of `FaucetFlagsInput` via:
//
//	FaucetFlagsArgs{...}
type FaucetFlagsInput interface {
	pulumi.Input

	ToFaucetFlagsOutput() FaucetFlagsOutput
	ToFaucetFlagsOutputWithContext(context.Context) FaucetFlagsOutput
}

type FaucetFlagsArgs struct {
	AllowIPs      pulumi.StringArrayInput `pulumi:"allowIPs"`
	PerRequestCap pulumi.IntPtrInput      `pulumi:"perRequestCap"`
	PerTimeCap    pulumi.IntPtrInput      `pulumi:"perTimeCap"`
	SliceSeconds  pulumi.IntPtrInput      `pulumi:"sliceSeconds"`
}

func (FaucetFlagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FaucetFlags)(nil)).Elem()
}

func (i FaucetFlagsArgs) ToFaucetFlagsOutput() FaucetFlagsOutput {
	return i.ToFaucetFlagsOutputWithContext(context.Background())
}

func (i FaucetFlagsArgs) ToFaucetFlagsOutputWithContext(ctx context.Context) FaucetFlagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FaucetFlagsOutput)
}

type FaucetFlagsOutput struct{ *pulumi.OutputState }

func (FaucetFlagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FaucetFlags)(nil)).Elem()
}

func (o FaucetFlagsOutput) ToFaucetFlagsOutput() FaucetFlagsOutput {
	return o
}

func (o FaucetFlagsOutput) ToFaucetFlagsOutputWithContext(ctx context.Context) FaucetFlagsOutput {
	return o
}

func (o FaucetFlagsOutput) AllowIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FaucetFlags) []string { return v.AllowIPs }).(pulumi.StringArrayOutput)
}

func (o FaucetFlagsOutput) PerRequestCap() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FaucetFlags) *int { return v.PerRequestCap }).(pulumi.IntPtrOutput)
}

func (o FaucetFlagsOutput) PerTimeCap() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FaucetFlags) *int { return v.PerTimeCap }).(pulumi.IntPtrOutput)
}

func (o FaucetFlagsOutput) SliceSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FaucetFlags) *int { return v.SliceSeconds }).(pulumi.IntPtrOutput)
}

type GenesisFlags struct {
	ClusterType                *string  `pulumi:"clusterType"`
	ExtraFlags                 []string `pulumi:"extraFlags"`
	FaucetLamports             *string  `pulumi:"faucetLamports"`
	FaucetPubkey               string   `pulumi:"faucetPubkey"`
	IdentityPubkey             string   `pulumi:"identityPubkey"`
	Inflation                  *string  `pulumi:"inflation"`
	LamportsPerByteYear        *string  `pulumi:"lamportsPerByteYear"`
	LedgerPath                 string   `pulumi:"ledgerPath"`
	SlotPerEpoch               *string  `pulumi:"slotPerEpoch"`
	StakePubkey                string   `pulumi:"stakePubkey"`
	TargetLamportsPerSignature *string  `pulumi:"targetLamportsPerSignature"`
	VotePubkey                 string   `pulumi:"votePubkey"`
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
	ClusterType                pulumi.StringPtrInput   `pulumi:"clusterType"`
	ExtraFlags                 pulumi.StringArrayInput `pulumi:"extraFlags"`
	FaucetLamports             pulumi.StringPtrInput   `pulumi:"faucetLamports"`
	FaucetPubkey               pulumi.StringInput      `pulumi:"faucetPubkey"`
	IdentityPubkey             pulumi.StringInput      `pulumi:"identityPubkey"`
	Inflation                  pulumi.StringPtrInput   `pulumi:"inflation"`
	LamportsPerByteYear        pulumi.StringPtrInput   `pulumi:"lamportsPerByteYear"`
	LedgerPath                 pulumi.StringInput      `pulumi:"ledgerPath"`
	SlotPerEpoch               pulumi.StringPtrInput   `pulumi:"slotPerEpoch"`
	StakePubkey                pulumi.StringInput      `pulumi:"stakePubkey"`
	TargetLamportsPerSignature pulumi.StringPtrInput   `pulumi:"targetLamportsPerSignature"`
	VotePubkey                 pulumi.StringInput      `pulumi:"votePubkey"`
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

func (o GenesisFlagsOutput) ClusterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.ClusterType }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) ExtraFlags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GenesisFlags) []string { return v.ExtraFlags }).(pulumi.StringArrayOutput)
}

func (o GenesisFlagsOutput) FaucetLamports() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.FaucetLamports }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) FaucetPubkey() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.FaucetPubkey }).(pulumi.StringOutput)
}

func (o GenesisFlagsOutput) IdentityPubkey() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.IdentityPubkey }).(pulumi.StringOutput)
}

func (o GenesisFlagsOutput) Inflation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.Inflation }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) LamportsPerByteYear() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.LamportsPerByteYear }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) LedgerPath() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.LedgerPath }).(pulumi.StringOutput)
}

func (o GenesisFlagsOutput) SlotPerEpoch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.SlotPerEpoch }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) StakePubkey() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.StakePubkey }).(pulumi.StringOutput)
}

func (o GenesisFlagsOutput) TargetLamportsPerSignature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GenesisFlags) *string { return v.TargetLamportsPerSignature }).(pulumi.StringPtrOutput)
}

func (o GenesisFlagsOutput) VotePubkey() pulumi.StringOutput {
	return o.ApplyT(func(v GenesisFlags) string { return v.VotePubkey }).(pulumi.StringOutput)
}

type PrimorialEntry struct {
	Lamports string `pulumi:"lamports"`
	Pubkey   string `pulumi:"pubkey"`
}

// PrimorialEntryInput is an input type that accepts PrimorialEntryArgs and PrimorialEntryOutput values.
// You can construct a concrete instance of `PrimorialEntryInput` via:
//
//	PrimorialEntryArgs{...}
type PrimorialEntryInput interface {
	pulumi.Input

	ToPrimorialEntryOutput() PrimorialEntryOutput
	ToPrimorialEntryOutputWithContext(context.Context) PrimorialEntryOutput
}

type PrimorialEntryArgs struct {
	Lamports pulumi.StringInput `pulumi:"lamports"`
	Pubkey   pulumi.StringInput `pulumi:"pubkey"`
}

func (PrimorialEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimorialEntry)(nil)).Elem()
}

func (i PrimorialEntryArgs) ToPrimorialEntryOutput() PrimorialEntryOutput {
	return i.ToPrimorialEntryOutputWithContext(context.Background())
}

func (i PrimorialEntryArgs) ToPrimorialEntryOutputWithContext(ctx context.Context) PrimorialEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrimorialEntryOutput)
}

// PrimorialEntryArrayInput is an input type that accepts PrimorialEntryArray and PrimorialEntryArrayOutput values.
// You can construct a concrete instance of `PrimorialEntryArrayInput` via:
//
//	PrimorialEntryArray{ PrimorialEntryArgs{...} }
type PrimorialEntryArrayInput interface {
	pulumi.Input

	ToPrimorialEntryArrayOutput() PrimorialEntryArrayOutput
	ToPrimorialEntryArrayOutputWithContext(context.Context) PrimorialEntryArrayOutput
}

type PrimorialEntryArray []PrimorialEntryInput

func (PrimorialEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrimorialEntry)(nil)).Elem()
}

func (i PrimorialEntryArray) ToPrimorialEntryArrayOutput() PrimorialEntryArrayOutput {
	return i.ToPrimorialEntryArrayOutputWithContext(context.Background())
}

func (i PrimorialEntryArray) ToPrimorialEntryArrayOutputWithContext(ctx context.Context) PrimorialEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrimorialEntryArrayOutput)
}

type PrimorialEntryOutput struct{ *pulumi.OutputState }

func (PrimorialEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimorialEntry)(nil)).Elem()
}

func (o PrimorialEntryOutput) ToPrimorialEntryOutput() PrimorialEntryOutput {
	return o
}

func (o PrimorialEntryOutput) ToPrimorialEntryOutputWithContext(ctx context.Context) PrimorialEntryOutput {
	return o
}

func (o PrimorialEntryOutput) Lamports() pulumi.StringOutput {
	return o.ApplyT(func(v PrimorialEntry) string { return v.Lamports }).(pulumi.StringOutput)
}

func (o PrimorialEntryOutput) Pubkey() pulumi.StringOutput {
	return o.ApplyT(func(v PrimorialEntry) string { return v.Pubkey }).(pulumi.StringOutput)
}

type PrimorialEntryArrayOutput struct{ *pulumi.OutputState }

func (PrimorialEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrimorialEntry)(nil)).Elem()
}

func (o PrimorialEntryArrayOutput) ToPrimorialEntryArrayOutput() PrimorialEntryArrayOutput {
	return o
}

func (o PrimorialEntryArrayOutput) ToPrimorialEntryArrayOutputWithContext(ctx context.Context) PrimorialEntryArrayOutput {
	return o
}

func (o PrimorialEntryArrayOutput) Index(i pulumi.IntInput) PrimorialEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrimorialEntry {
		return vs[0].([]PrimorialEntry)[vs[1].(int)]
	}).(PrimorialEntryOutput)
}

type StakeAccountKeyPairs struct {
	StakeAccount string `pulumi:"stakeAccount"`
	VoteAccount  string `pulumi:"voteAccount"`
}

// StakeAccountKeyPairsInput is an input type that accepts StakeAccountKeyPairsArgs and StakeAccountKeyPairsOutput values.
// You can construct a concrete instance of `StakeAccountKeyPairsInput` via:
//
//	StakeAccountKeyPairsArgs{...}
type StakeAccountKeyPairsInput interface {
	pulumi.Input

	ToStakeAccountKeyPairsOutput() StakeAccountKeyPairsOutput
	ToStakeAccountKeyPairsOutputWithContext(context.Context) StakeAccountKeyPairsOutput
}

type StakeAccountKeyPairsArgs struct {
	StakeAccount pulumi.StringInput `pulumi:"stakeAccount"`
	VoteAccount  pulumi.StringInput `pulumi:"voteAccount"`
}

func (StakeAccountKeyPairsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StakeAccountKeyPairs)(nil)).Elem()
}

func (i StakeAccountKeyPairsArgs) ToStakeAccountKeyPairsOutput() StakeAccountKeyPairsOutput {
	return i.ToStakeAccountKeyPairsOutputWithContext(context.Background())
}

func (i StakeAccountKeyPairsArgs) ToStakeAccountKeyPairsOutputWithContext(ctx context.Context) StakeAccountKeyPairsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StakeAccountKeyPairsOutput)
}

type StakeAccountKeyPairsOutput struct{ *pulumi.OutputState }

func (StakeAccountKeyPairsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StakeAccountKeyPairs)(nil)).Elem()
}

func (o StakeAccountKeyPairsOutput) ToStakeAccountKeyPairsOutput() StakeAccountKeyPairsOutput {
	return o
}

func (o StakeAccountKeyPairsOutput) ToStakeAccountKeyPairsOutputWithContext(ctx context.Context) StakeAccountKeyPairsOutput {
	return o
}

func (o StakeAccountKeyPairsOutput) StakeAccount() pulumi.StringOutput {
	return o.ApplyT(func(v StakeAccountKeyPairs) string { return v.StakeAccount }).(pulumi.StringOutput)
}

func (o StakeAccountKeyPairsOutput) VoteAccount() pulumi.StringOutput {
	return o.ApplyT(func(v StakeAccountKeyPairs) string { return v.VoteAccount }).(pulumi.StringOutput)
}

type TxnOptions struct {
	BlockHash            *string  `pulumi:"blockHash"`
	Commitment           *string  `pulumi:"commitment"`
	FeePayer             *string  `pulumi:"feePayer"`
	From                 *string  `pulumi:"from"`
	KeyPair              *string  `pulumi:"keyPair"`
	Nonce                *string  `pulumi:"nonce"`
	NonceAuthority       *string  `pulumi:"nonceAuthority"`
	Signer               []string `pulumi:"signer"`
	Url                  *string  `pulumi:"url"`
	WithComputeUnitPrice *float64 `pulumi:"withComputeUnitPrice"`
	WithMemo             *string  `pulumi:"withMemo"`
	Ws                   *string  `pulumi:"ws"`
}

// TxnOptionsInput is an input type that accepts TxnOptionsArgs and TxnOptionsOutput values.
// You can construct a concrete instance of `TxnOptionsInput` via:
//
//	TxnOptionsArgs{...}
type TxnOptionsInput interface {
	pulumi.Input

	ToTxnOptionsOutput() TxnOptionsOutput
	ToTxnOptionsOutputWithContext(context.Context) TxnOptionsOutput
}

type TxnOptionsArgs struct {
	BlockHash            pulumi.StringPtrInput   `pulumi:"blockHash"`
	Commitment           pulumi.StringPtrInput   `pulumi:"commitment"`
	FeePayer             pulumi.StringPtrInput   `pulumi:"feePayer"`
	From                 pulumi.StringPtrInput   `pulumi:"from"`
	KeyPair              pulumi.StringPtrInput   `pulumi:"keyPair"`
	Nonce                pulumi.StringPtrInput   `pulumi:"nonce"`
	NonceAuthority       pulumi.StringPtrInput   `pulumi:"nonceAuthority"`
	Signer               pulumi.StringArrayInput `pulumi:"signer"`
	Url                  pulumi.StringPtrInput   `pulumi:"url"`
	WithComputeUnitPrice pulumi.Float64PtrInput  `pulumi:"withComputeUnitPrice"`
	WithMemo             pulumi.StringPtrInput   `pulumi:"withMemo"`
	Ws                   pulumi.StringPtrInput   `pulumi:"ws"`
}

func (TxnOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxnOptions)(nil)).Elem()
}

func (i TxnOptionsArgs) ToTxnOptionsOutput() TxnOptionsOutput {
	return i.ToTxnOptionsOutputWithContext(context.Background())
}

func (i TxnOptionsArgs) ToTxnOptionsOutputWithContext(ctx context.Context) TxnOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxnOptionsOutput)
}

type TxnOptionsOutput struct{ *pulumi.OutputState }

func (TxnOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxnOptions)(nil)).Elem()
}

func (o TxnOptionsOutput) ToTxnOptionsOutput() TxnOptionsOutput {
	return o
}

func (o TxnOptionsOutput) ToTxnOptionsOutputWithContext(ctx context.Context) TxnOptionsOutput {
	return o
}

func (o TxnOptionsOutput) BlockHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.BlockHash }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) Commitment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.Commitment }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) FeePayer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.FeePayer }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.From }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.KeyPair }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) Nonce() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.Nonce }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) NonceAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.NonceAuthority }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) Signer() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxnOptions) []string { return v.Signer }).(pulumi.StringArrayOutput)
}

func (o TxnOptionsOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) WithComputeUnitPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v TxnOptions) *float64 { return v.WithComputeUnitPrice }).(pulumi.Float64PtrOutput)
}

func (o TxnOptionsOutput) WithMemo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.WithMemo }).(pulumi.StringPtrOutput)
}

func (o TxnOptionsOutput) Ws() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TxnOptions) *string { return v.Ws }).(pulumi.StringPtrOutput)
}

type ValidatorInfo struct {
	Details *string `pulumi:"details"`
	IconURL *string `pulumi:"iconURL"`
	Name    string  `pulumi:"name"`
	Website *string `pulumi:"website"`
}

// ValidatorInfoInput is an input type that accepts ValidatorInfoArgs and ValidatorInfoOutput values.
// You can construct a concrete instance of `ValidatorInfoInput` via:
//
//	ValidatorInfoArgs{...}
type ValidatorInfoInput interface {
	pulumi.Input

	ToValidatorInfoOutput() ValidatorInfoOutput
	ToValidatorInfoOutputWithContext(context.Context) ValidatorInfoOutput
}

type ValidatorInfoArgs struct {
	Details pulumi.StringPtrInput `pulumi:"details"`
	IconURL pulumi.StringPtrInput `pulumi:"iconURL"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Website pulumi.StringPtrInput `pulumi:"website"`
}

func (ValidatorInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatorInfo)(nil)).Elem()
}

func (i ValidatorInfoArgs) ToValidatorInfoOutput() ValidatorInfoOutput {
	return i.ToValidatorInfoOutputWithContext(context.Background())
}

func (i ValidatorInfoArgs) ToValidatorInfoOutputWithContext(ctx context.Context) ValidatorInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatorInfoOutput)
}

func (i ValidatorInfoArgs) ToValidatorInfoPtrOutput() ValidatorInfoPtrOutput {
	return i.ToValidatorInfoPtrOutputWithContext(context.Background())
}

func (i ValidatorInfoArgs) ToValidatorInfoPtrOutputWithContext(ctx context.Context) ValidatorInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatorInfoOutput).ToValidatorInfoPtrOutputWithContext(ctx)
}

// ValidatorInfoPtrInput is an input type that accepts ValidatorInfoArgs, ValidatorInfoPtr and ValidatorInfoPtrOutput values.
// You can construct a concrete instance of `ValidatorInfoPtrInput` via:
//
//	        ValidatorInfoArgs{...}
//
//	or:
//
//	        nil
type ValidatorInfoPtrInput interface {
	pulumi.Input

	ToValidatorInfoPtrOutput() ValidatorInfoPtrOutput
	ToValidatorInfoPtrOutputWithContext(context.Context) ValidatorInfoPtrOutput
}

type validatorInfoPtrType ValidatorInfoArgs

func ValidatorInfoPtr(v *ValidatorInfoArgs) ValidatorInfoPtrInput {
	return (*validatorInfoPtrType)(v)
}

func (*validatorInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatorInfo)(nil)).Elem()
}

func (i *validatorInfoPtrType) ToValidatorInfoPtrOutput() ValidatorInfoPtrOutput {
	return i.ToValidatorInfoPtrOutputWithContext(context.Background())
}

func (i *validatorInfoPtrType) ToValidatorInfoPtrOutputWithContext(ctx context.Context) ValidatorInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatorInfoPtrOutput)
}

type ValidatorInfoOutput struct{ *pulumi.OutputState }

func (ValidatorInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatorInfo)(nil)).Elem()
}

func (o ValidatorInfoOutput) ToValidatorInfoOutput() ValidatorInfoOutput {
	return o
}

func (o ValidatorInfoOutput) ToValidatorInfoOutputWithContext(ctx context.Context) ValidatorInfoOutput {
	return o
}

func (o ValidatorInfoOutput) ToValidatorInfoPtrOutput() ValidatorInfoPtrOutput {
	return o.ToValidatorInfoPtrOutputWithContext(context.Background())
}

func (o ValidatorInfoOutput) ToValidatorInfoPtrOutputWithContext(ctx context.Context) ValidatorInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ValidatorInfo) *ValidatorInfo {
		return &v
	}).(ValidatorInfoPtrOutput)
}

func (o ValidatorInfoOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ValidatorInfo) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o ValidatorInfoOutput) IconURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ValidatorInfo) *string { return v.IconURL }).(pulumi.StringPtrOutput)
}

func (o ValidatorInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ValidatorInfo) string { return v.Name }).(pulumi.StringOutput)
}

func (o ValidatorInfoOutput) Website() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ValidatorInfo) *string { return v.Website }).(pulumi.StringPtrOutput)
}

type ValidatorInfoPtrOutput struct{ *pulumi.OutputState }

func (ValidatorInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatorInfo)(nil)).Elem()
}

func (o ValidatorInfoPtrOutput) ToValidatorInfoPtrOutput() ValidatorInfoPtrOutput {
	return o
}

func (o ValidatorInfoPtrOutput) ToValidatorInfoPtrOutputWithContext(ctx context.Context) ValidatorInfoPtrOutput {
	return o
}

func (o ValidatorInfoPtrOutput) Elem() ValidatorInfoOutput {
	return o.ApplyT(func(v *ValidatorInfo) ValidatorInfo {
		if v != nil {
			return *v
		}
		var ret ValidatorInfo
		return ret
	}).(ValidatorInfoOutput)
}

func (o ValidatorInfoPtrOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatorInfo) *string {
		if v == nil {
			return nil
		}
		return v.Details
	}).(pulumi.StringPtrOutput)
}

func (o ValidatorInfoPtrOutput) IconURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatorInfo) *string {
		if v == nil {
			return nil
		}
		return v.IconURL
	}).(pulumi.StringPtrOutput)
}

func (o ValidatorInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatorInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ValidatorInfoPtrOutput) Website() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatorInfo) *string {
		if v == nil {
			return nil
		}
		return v.Website
	}).(pulumi.StringPtrOutput)
}

type VoteAccountKeyPairs struct {
	AuthWithdrawer string `pulumi:"authWithdrawer"`
	Identity       string `pulumi:"identity"`
	VoteAccount    string `pulumi:"voteAccount"`
}

// VoteAccountKeyPairsInput is an input type that accepts VoteAccountKeyPairsArgs and VoteAccountKeyPairsOutput values.
// You can construct a concrete instance of `VoteAccountKeyPairsInput` via:
//
//	VoteAccountKeyPairsArgs{...}
type VoteAccountKeyPairsInput interface {
	pulumi.Input

	ToVoteAccountKeyPairsOutput() VoteAccountKeyPairsOutput
	ToVoteAccountKeyPairsOutputWithContext(context.Context) VoteAccountKeyPairsOutput
}

type VoteAccountKeyPairsArgs struct {
	AuthWithdrawer pulumi.StringInput `pulumi:"authWithdrawer"`
	Identity       pulumi.StringInput `pulumi:"identity"`
	VoteAccount    pulumi.StringInput `pulumi:"voteAccount"`
}

func (VoteAccountKeyPairsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VoteAccountKeyPairs)(nil)).Elem()
}

func (i VoteAccountKeyPairsArgs) ToVoteAccountKeyPairsOutput() VoteAccountKeyPairsOutput {
	return i.ToVoteAccountKeyPairsOutputWithContext(context.Background())
}

func (i VoteAccountKeyPairsArgs) ToVoteAccountKeyPairsOutputWithContext(ctx context.Context) VoteAccountKeyPairsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoteAccountKeyPairsOutput)
}

type VoteAccountKeyPairsOutput struct{ *pulumi.OutputState }

func (VoteAccountKeyPairsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VoteAccountKeyPairs)(nil)).Elem()
}

func (o VoteAccountKeyPairsOutput) ToVoteAccountKeyPairsOutput() VoteAccountKeyPairsOutput {
	return o
}

func (o VoteAccountKeyPairsOutput) ToVoteAccountKeyPairsOutputWithContext(ctx context.Context) VoteAccountKeyPairsOutput {
	return o
}

func (o VoteAccountKeyPairsOutput) AuthWithdrawer() pulumi.StringOutput {
	return o.ApplyT(func(v VoteAccountKeyPairs) string { return v.AuthWithdrawer }).(pulumi.StringOutput)
}

func (o VoteAccountKeyPairsOutput) Identity() pulumi.StringOutput {
	return o.ApplyT(func(v VoteAccountKeyPairs) string { return v.Identity }).(pulumi.StringOutput)
}

func (o VoteAccountKeyPairsOutput) VoteAccount() pulumi.StringOutput {
	return o.ApplyT(func(v VoteAccountKeyPairs) string { return v.VoteAccount }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), EnvironmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentPtrInput)(nil)).Elem(), EnvironmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FaucetFlagsInput)(nil)).Elem(), FaucetFlagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenesisFlagsInput)(nil)).Elem(), GenesisFlagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrimorialEntryInput)(nil)).Elem(), PrimorialEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrimorialEntryArrayInput)(nil)).Elem(), PrimorialEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StakeAccountKeyPairsInput)(nil)).Elem(), StakeAccountKeyPairsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxnOptionsInput)(nil)).Elem(), TxnOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatorInfoInput)(nil)).Elem(), ValidatorInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatorInfoPtrInput)(nil)).Elem(), ValidatorInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoteAccountKeyPairsInput)(nil)).Elem(), VoteAccountKeyPairsArgs{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentPtrOutput{})
	pulumi.RegisterOutputType(FaucetFlagsOutput{})
	pulumi.RegisterOutputType(GenesisFlagsOutput{})
	pulumi.RegisterOutputType(PrimorialEntryOutput{})
	pulumi.RegisterOutputType(PrimorialEntryArrayOutput{})
	pulumi.RegisterOutputType(StakeAccountKeyPairsOutput{})
	pulumi.RegisterOutputType(TxnOptionsOutput{})
	pulumi.RegisterOutputType(ValidatorInfoOutput{})
	pulumi.RegisterOutputType(ValidatorInfoPtrOutput{})
	pulumi.RegisterOutputType(VoteAccountKeyPairsOutput{})
}
