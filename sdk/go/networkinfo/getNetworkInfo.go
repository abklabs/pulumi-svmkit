// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkinfo

import (
	"context"
	"reflect"

	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/abklabs/pulumi-svmkit/sdk/go/solana"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNetworkInfo(ctx *pulumi.Context, args *GetNetworkInfoArgs, opts ...pulumi.InvokeOption) (*GetNetworkInfoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNetworkInfoResult
	err := ctx.Invoke("svmkit:networkinfo:getNetworkInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetNetworkInfoArgs struct {
	NetworkName solana.NetworkName `pulumi:"networkName"`
}

type GetNetworkInfoResult struct {
	EntryPoint     []string           `pulumi:"entryPoint"`
	GenesisHash    string             `pulumi:"genesisHash"`
	KnownValidator []string           `pulumi:"knownValidator"`
	NetworkName    solana.NetworkName `pulumi:"networkName"`
	RpcURL         []string           `pulumi:"rpcURL"`
}

func GetNetworkInfoOutput(ctx *pulumi.Context, args GetNetworkInfoOutputArgs, opts ...pulumi.InvokeOption) GetNetworkInfoResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNetworkInfoResultOutput, error) {
			args := v.(GetNetworkInfoArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("svmkit:networkinfo:getNetworkInfo", args, GetNetworkInfoResultOutput{}, options).(GetNetworkInfoResultOutput), nil
		}).(GetNetworkInfoResultOutput)
}

type GetNetworkInfoOutputArgs struct {
	NetworkName solana.NetworkNameInput `pulumi:"networkName"`
}

func (GetNetworkInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkInfoArgs)(nil)).Elem()
}

type GetNetworkInfoResultOutput struct{ *pulumi.OutputState }

func (GetNetworkInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkInfoResult)(nil)).Elem()
}

func (o GetNetworkInfoResultOutput) ToGetNetworkInfoResultOutput() GetNetworkInfoResultOutput {
	return o
}

func (o GetNetworkInfoResultOutput) ToGetNetworkInfoResultOutputWithContext(ctx context.Context) GetNetworkInfoResultOutput {
	return o
}

func (o GetNetworkInfoResultOutput) EntryPoint() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetworkInfoResult) []string { return v.EntryPoint }).(pulumi.StringArrayOutput)
}

func (o GetNetworkInfoResultOutput) GenesisHash() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkInfoResult) string { return v.GenesisHash }).(pulumi.StringOutput)
}

func (o GetNetworkInfoResultOutput) KnownValidator() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetworkInfoResult) []string { return v.KnownValidator }).(pulumi.StringArrayOutput)
}

func (o GetNetworkInfoResultOutput) NetworkName() solana.NetworkNameOutput {
	return o.ApplyT(func(v GetNetworkInfoResult) solana.NetworkName { return v.NetworkName }).(solana.NetworkNameOutput)
}

func (o GetNetworkInfoResultOutput) RpcURL() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetworkInfoResult) []string { return v.RpcURL }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworkInfoResultOutput{})
}