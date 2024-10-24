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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrimorialEntryInput)(nil)).Elem(), PrimorialEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrimorialEntryArrayInput)(nil)).Elem(), PrimorialEntryArray{})
	pulumi.RegisterOutputType(PrimorialEntryOutput{})
	pulumi.RegisterOutputType(PrimorialEntryArrayOutput{})
}
