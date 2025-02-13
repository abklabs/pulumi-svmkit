// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tuner

import (
	"context"
	"reflect"

	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDefaultTunerParams(ctx *pulumi.Context, args *GetDefaultTunerParamsArgs, opts ...pulumi.InvokeOption) (*GetDefaultTunerParamsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDefaultTunerParamsResult
	err := ctx.Invoke("svmkit:tuner:getDefaultTunerParams", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDefaultTunerParamsArgs struct {
	Variant TunerVariant `pulumi:"variant"`
}

type GetDefaultTunerParamsResult struct {
	CpuGovernor *CpuGovernor       `pulumi:"cpuGovernor"`
	Kernel      *TunerKernelParams `pulumi:"kernel"`
	Net         *TunerNetParams    `pulumi:"net"`
	Variant     TunerVariant       `pulumi:"variant"`
	Vm          *TunerVmParams     `pulumi:"vm"`
}

func GetDefaultTunerParamsOutput(ctx *pulumi.Context, args GetDefaultTunerParamsOutputArgs, opts ...pulumi.InvokeOption) GetDefaultTunerParamsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDefaultTunerParamsResultOutput, error) {
			args := v.(GetDefaultTunerParamsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("svmkit:tuner:getDefaultTunerParams", args, GetDefaultTunerParamsResultOutput{}, options).(GetDefaultTunerParamsResultOutput), nil
		}).(GetDefaultTunerParamsResultOutput)
}

type GetDefaultTunerParamsOutputArgs struct {
	Variant TunerVariantInput `pulumi:"variant"`
}

func (GetDefaultTunerParamsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefaultTunerParamsArgs)(nil)).Elem()
}

type GetDefaultTunerParamsResultOutput struct{ *pulumi.OutputState }

func (GetDefaultTunerParamsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefaultTunerParamsResult)(nil)).Elem()
}

func (o GetDefaultTunerParamsResultOutput) ToGetDefaultTunerParamsResultOutput() GetDefaultTunerParamsResultOutput {
	return o
}

func (o GetDefaultTunerParamsResultOutput) ToGetDefaultTunerParamsResultOutputWithContext(ctx context.Context) GetDefaultTunerParamsResultOutput {
	return o
}

func (o GetDefaultTunerParamsResultOutput) CpuGovernor() CpuGovernorPtrOutput {
	return o.ApplyT(func(v GetDefaultTunerParamsResult) *CpuGovernor { return v.CpuGovernor }).(CpuGovernorPtrOutput)
}

func (o GetDefaultTunerParamsResultOutput) Kernel() TunerKernelParamsPtrOutput {
	return o.ApplyT(func(v GetDefaultTunerParamsResult) *TunerKernelParams { return v.Kernel }).(TunerKernelParamsPtrOutput)
}

func (o GetDefaultTunerParamsResultOutput) Net() TunerNetParamsPtrOutput {
	return o.ApplyT(func(v GetDefaultTunerParamsResult) *TunerNetParams { return v.Net }).(TunerNetParamsPtrOutput)
}

func (o GetDefaultTunerParamsResultOutput) Variant() TunerVariantOutput {
	return o.ApplyT(func(v GetDefaultTunerParamsResult) TunerVariant { return v.Variant }).(TunerVariantOutput)
}

func (o GetDefaultTunerParamsResultOutput) Vm() TunerVmParamsPtrOutput {
	return o.ApplyT(func(v GetDefaultTunerParamsResult) *TunerVmParams { return v.Vm }).(TunerVmParamsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDefaultTunerParamsResultOutput{})
}
