// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apt

import (
	"context"
	"reflect"

	"github.com/abklabs/pulumi-svmkit/sdk/go/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type Config struct {
	ExcludeDefaultSources *bool    `pulumi:"excludeDefaultSources"`
	Sources               []Source `pulumi:"sources"`
}

// ConfigInput is an input type that accepts ConfigArgs and ConfigOutput values.
// You can construct a concrete instance of `ConfigInput` via:
//
//	ConfigArgs{...}
type ConfigInput interface {
	pulumi.Input

	ToConfigOutput() ConfigOutput
	ToConfigOutputWithContext(context.Context) ConfigOutput
}

type ConfigArgs struct {
	ExcludeDefaultSources pulumi.BoolPtrInput `pulumi:"excludeDefaultSources"`
	Sources               SourceArrayInput    `pulumi:"sources"`
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil)).Elem()
}

func (i ConfigArgs) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i ConfigArgs) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

func (i ConfigArgs) ToConfigPtrOutput() ConfigPtrOutput {
	return i.ToConfigPtrOutputWithContext(context.Background())
}

func (i ConfigArgs) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput).ToConfigPtrOutputWithContext(ctx)
}

// ConfigPtrInput is an input type that accepts ConfigArgs, ConfigPtr and ConfigPtrOutput values.
// You can construct a concrete instance of `ConfigPtrInput` via:
//
//	        ConfigArgs{...}
//
//	or:
//
//	        nil
type ConfigPtrInput interface {
	pulumi.Input

	ToConfigPtrOutput() ConfigPtrOutput
	ToConfigPtrOutputWithContext(context.Context) ConfigPtrOutput
}

type configPtrType ConfigArgs

func ConfigPtr(v *ConfigArgs) ConfigPtrInput {
	return (*configPtrType)(v)
}

func (*configPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (i *configPtrType) ToConfigPtrOutput() ConfigPtrOutput {
	return i.ToConfigPtrOutputWithContext(context.Background())
}

func (i *configPtrType) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigPtrOutput)
}

type ConfigOutput struct{ *pulumi.OutputState }

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil)).Elem()
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigPtrOutput() ConfigPtrOutput {
	return o.ToConfigPtrOutputWithContext(context.Background())
}

func (o ConfigOutput) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Config) *Config {
		return &v
	}).(ConfigPtrOutput)
}

func (o ConfigOutput) ExcludeDefaultSources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Config) *bool { return v.ExcludeDefaultSources }).(pulumi.BoolPtrOutput)
}

func (o ConfigOutput) Sources() SourceArrayOutput {
	return o.ApplyT(func(v Config) []Source { return v.Sources }).(SourceArrayOutput)
}

type ConfigPtrOutput struct{ *pulumi.OutputState }

func (ConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (o ConfigPtrOutput) ToConfigPtrOutput() ConfigPtrOutput {
	return o
}

func (o ConfigPtrOutput) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return o
}

func (o ConfigPtrOutput) Elem() ConfigOutput {
	return o.ApplyT(func(v *Config) Config {
		if v != nil {
			return *v
		}
		var ret Config
		return ret
	}).(ConfigOutput)
}

func (o ConfigPtrOutput) ExcludeDefaultSources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Config) *bool {
		if v == nil {
			return nil
		}
		return v.ExcludeDefaultSources
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigPtrOutput) Sources() SourceArrayOutput {
	return o.ApplyT(func(v *Config) []Source {
		if v == nil {
			return nil
		}
		return v.Sources
	}).(SourceArrayOutput)
}

type SignedBy struct {
	Paths     []string `pulumi:"paths"`
	PublicKey *string  `pulumi:"publicKey"`
}

// SignedByInput is an input type that accepts SignedByArgs and SignedByOutput values.
// You can construct a concrete instance of `SignedByInput` via:
//
//	SignedByArgs{...}
type SignedByInput interface {
	pulumi.Input

	ToSignedByOutput() SignedByOutput
	ToSignedByOutputWithContext(context.Context) SignedByOutput
}

type SignedByArgs struct {
	Paths     pulumi.StringArrayInput `pulumi:"paths"`
	PublicKey pulumi.StringPtrInput   `pulumi:"publicKey"`
}

func (SignedByArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignedBy)(nil)).Elem()
}

func (i SignedByArgs) ToSignedByOutput() SignedByOutput {
	return i.ToSignedByOutputWithContext(context.Background())
}

func (i SignedByArgs) ToSignedByOutputWithContext(ctx context.Context) SignedByOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignedByOutput)
}

func (i SignedByArgs) ToSignedByPtrOutput() SignedByPtrOutput {
	return i.ToSignedByPtrOutputWithContext(context.Background())
}

func (i SignedByArgs) ToSignedByPtrOutputWithContext(ctx context.Context) SignedByPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignedByOutput).ToSignedByPtrOutputWithContext(ctx)
}

// SignedByPtrInput is an input type that accepts SignedByArgs, SignedByPtr and SignedByPtrOutput values.
// You can construct a concrete instance of `SignedByPtrInput` via:
//
//	        SignedByArgs{...}
//
//	or:
//
//	        nil
type SignedByPtrInput interface {
	pulumi.Input

	ToSignedByPtrOutput() SignedByPtrOutput
	ToSignedByPtrOutputWithContext(context.Context) SignedByPtrOutput
}

type signedByPtrType SignedByArgs

func SignedByPtr(v *SignedByArgs) SignedByPtrInput {
	return (*signedByPtrType)(v)
}

func (*signedByPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SignedBy)(nil)).Elem()
}

func (i *signedByPtrType) ToSignedByPtrOutput() SignedByPtrOutput {
	return i.ToSignedByPtrOutputWithContext(context.Background())
}

func (i *signedByPtrType) ToSignedByPtrOutputWithContext(ctx context.Context) SignedByPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignedByPtrOutput)
}

type SignedByOutput struct{ *pulumi.OutputState }

func (SignedByOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignedBy)(nil)).Elem()
}

func (o SignedByOutput) ToSignedByOutput() SignedByOutput {
	return o
}

func (o SignedByOutput) ToSignedByOutputWithContext(ctx context.Context) SignedByOutput {
	return o
}

func (o SignedByOutput) ToSignedByPtrOutput() SignedByPtrOutput {
	return o.ToSignedByPtrOutputWithContext(context.Background())
}

func (o SignedByOutput) ToSignedByPtrOutputWithContext(ctx context.Context) SignedByPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignedBy) *SignedBy {
		return &v
	}).(SignedByPtrOutput)
}

func (o SignedByOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SignedBy) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o SignedByOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SignedBy) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

type SignedByPtrOutput struct{ *pulumi.OutputState }

func (SignedByPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignedBy)(nil)).Elem()
}

func (o SignedByPtrOutput) ToSignedByPtrOutput() SignedByPtrOutput {
	return o
}

func (o SignedByPtrOutput) ToSignedByPtrOutputWithContext(ctx context.Context) SignedByPtrOutput {
	return o
}

func (o SignedByPtrOutput) Elem() SignedByOutput {
	return o.ApplyT(func(v *SignedBy) SignedBy {
		if v != nil {
			return *v
		}
		var ret SignedBy
		return ret
	}).(SignedByOutput)
}

func (o SignedByPtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SignedBy) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

func (o SignedByPtrOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignedBy) *string {
		if v == nil {
			return nil
		}
		return v.PublicKey
	}).(pulumi.StringPtrOutput)
}

type Source struct {
	URIs                     []string  `pulumi:"URIs"`
	AllowDowngradeToInsecure *bool     `pulumi:"allowDowngradeToInsecure"`
	AllowInsecure            *bool     `pulumi:"allowInsecure"`
	AllowWeak                *bool     `pulumi:"allowWeak"`
	Architectures            []string  `pulumi:"architectures"`
	CheckDate                *bool     `pulumi:"checkDate"`
	CheckValidUntil          *bool     `pulumi:"checkValidUntil"`
	Components               []string  `pulumi:"components"`
	DateMaxFuture            *int      `pulumi:"dateMaxFuture"`
	ExtraLines               []string  `pulumi:"extraLines"`
	InReleasePath            *string   `pulumi:"inReleasePath"`
	SignedBy                 *SignedBy `pulumi:"signedBy"`
	Snapshot                 *string   `pulumi:"snapshot"`
	Suites                   []string  `pulumi:"suites"`
	Trusted                  *bool     `pulumi:"trusted"`
	Types                    []string  `pulumi:"types"`
	ValidUntilMax            *int      `pulumi:"validUntilMax"`
	ValidUntilMin            *int      `pulumi:"validUntilMin"`
}

// SourceInput is an input type that accepts SourceArgs and SourceOutput values.
// You can construct a concrete instance of `SourceInput` via:
//
//	SourceArgs{...}
type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(context.Context) SourceOutput
}

type SourceArgs struct {
	URIs                     pulumi.StringArrayInput `pulumi:"URIs"`
	AllowDowngradeToInsecure pulumi.BoolPtrInput     `pulumi:"allowDowngradeToInsecure"`
	AllowInsecure            pulumi.BoolPtrInput     `pulumi:"allowInsecure"`
	AllowWeak                pulumi.BoolPtrInput     `pulumi:"allowWeak"`
	Architectures            pulumi.StringArrayInput `pulumi:"architectures"`
	CheckDate                pulumi.BoolPtrInput     `pulumi:"checkDate"`
	CheckValidUntil          pulumi.BoolPtrInput     `pulumi:"checkValidUntil"`
	Components               pulumi.StringArrayInput `pulumi:"components"`
	DateMaxFuture            pulumi.IntPtrInput      `pulumi:"dateMaxFuture"`
	ExtraLines               pulumi.StringArrayInput `pulumi:"extraLines"`
	InReleasePath            pulumi.StringPtrInput   `pulumi:"inReleasePath"`
	SignedBy                 SignedByPtrInput        `pulumi:"signedBy"`
	Snapshot                 pulumi.StringPtrInput   `pulumi:"snapshot"`
	Suites                   pulumi.StringArrayInput `pulumi:"suites"`
	Trusted                  pulumi.BoolPtrInput     `pulumi:"trusted"`
	Types                    pulumi.StringArrayInput `pulumi:"types"`
	ValidUntilMax            pulumi.IntPtrInput      `pulumi:"validUntilMax"`
	ValidUntilMin            pulumi.IntPtrInput      `pulumi:"validUntilMin"`
}

func (SourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (i SourceArgs) ToSourceOutput() SourceOutput {
	return i.ToSourceOutputWithContext(context.Background())
}

func (i SourceArgs) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutput)
}

// SourceArrayInput is an input type that accepts SourceArray and SourceArrayOutput values.
// You can construct a concrete instance of `SourceArrayInput` via:
//
//	SourceArray{ SourceArgs{...} }
type SourceArrayInput interface {
	pulumi.Input

	ToSourceArrayOutput() SourceArrayOutput
	ToSourceArrayOutputWithContext(context.Context) SourceArrayOutput
}

type SourceArray []SourceInput

func (SourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Source)(nil)).Elem()
}

func (i SourceArray) ToSourceArrayOutput() SourceArrayOutput {
	return i.ToSourceArrayOutputWithContext(context.Background())
}

func (i SourceArray) ToSourceArrayOutputWithContext(ctx context.Context) SourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceArrayOutput)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

func (o SourceOutput) URIs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.URIs }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) AllowDowngradeToInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Source) *bool { return v.AllowDowngradeToInsecure }).(pulumi.BoolPtrOutput)
}

func (o SourceOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Source) *bool { return v.AllowInsecure }).(pulumi.BoolPtrOutput)
}

func (o SourceOutput) AllowWeak() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Source) *bool { return v.AllowWeak }).(pulumi.BoolPtrOutput)
}

func (o SourceOutput) Architectures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.Architectures }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) CheckDate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Source) *bool { return v.CheckDate }).(pulumi.BoolPtrOutput)
}

func (o SourceOutput) CheckValidUntil() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Source) *bool { return v.CheckValidUntil }).(pulumi.BoolPtrOutput)
}

func (o SourceOutput) Components() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.Components }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) DateMaxFuture() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Source) *int { return v.DateMaxFuture }).(pulumi.IntPtrOutput)
}

func (o SourceOutput) ExtraLines() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.ExtraLines }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) InReleasePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.InReleasePath }).(pulumi.StringPtrOutput)
}

func (o SourceOutput) SignedBy() SignedByPtrOutput {
	return o.ApplyT(func(v Source) *SignedBy { return v.SignedBy }).(SignedByPtrOutput)
}

func (o SourceOutput) Snapshot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.Snapshot }).(pulumi.StringPtrOutput)
}

func (o SourceOutput) Suites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.Suites }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) Trusted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Source) *bool { return v.Trusted }).(pulumi.BoolPtrOutput)
}

func (o SourceOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.Types }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) ValidUntilMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Source) *int { return v.ValidUntilMax }).(pulumi.IntPtrOutput)
}

func (o SourceOutput) ValidUntilMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Source) *int { return v.ValidUntilMin }).(pulumi.IntPtrOutput)
}

type SourceArrayOutput struct{ *pulumi.OutputState }

func (SourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Source)(nil)).Elem()
}

func (o SourceArrayOutput) ToSourceArrayOutput() SourceArrayOutput {
	return o
}

func (o SourceArrayOutput) ToSourceArrayOutputWithContext(ctx context.Context) SourceArrayOutput {
	return o
}

func (o SourceArrayOutput) Index(i pulumi.IntInput) SourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Source {
		return vs[0].([]Source)[vs[1].(int)]
	}).(SourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInput)(nil)).Elem(), ConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigPtrInput)(nil)).Elem(), ConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SignedByInput)(nil)).Elem(), SignedByArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SignedByPtrInput)(nil)).Elem(), SignedByArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceInput)(nil)).Elem(), SourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceArrayInput)(nil)).Elem(), SourceArray{})
	pulumi.RegisterOutputType(ConfigOutput{})
	pulumi.RegisterOutputType(ConfigPtrOutput{})
	pulumi.RegisterOutputType(SignedByOutput{})
	pulumi.RegisterOutputType(SignedByPtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourceArrayOutput{})
}
