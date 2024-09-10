package provider

import (
	"encoding/json"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KeyPairArgs struct{}

type KeyPair struct {
	pulumi.ResourceState

	PublicKey  pulumi.StringOutput   `pulumi:"publicKey"`
	PrivateKey pulumi.IntArrayOutput `pulumi:"privateKey"`
	Json       pulumi.StringOutput   `pulumi:"json"`
}

type KeyPairInput struct {
	pulumi.ResourceInput

	PublicKey  string `pulumi:"publicKey"`
	PrivateKey string `pulumi:"privateKey" provider:"secret"`
}

func NewKeyPair(ctx *pulumi.Context, name string, args *KeyPairArgs, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	keyPair := &KeyPair{}
	err := ctx.RegisterComponentResource("svmkit:index:KeyPair", name, keyPair, opts...)
	if err != nil {
		return nil, err
	}
	account := solana.NewWallet()
	publicKey := account.PublicKey().String()

	privateKey := make([]int, len(account.PrivateKey))

	for i, b := range account.PrivateKey {
		privateKey[i] = int(b)
	}

	json, err := json.Marshal(privateKey)

	if err != nil {
		return nil, err
	}

	fmt.Printf("publicKey: %s\n", publicKey)
	fmt.Printf("privateKey: %v\n", privateKey)
	fmt.Printf("json: %s\n", string(json))

	keyPair.PublicKey = pulumi.ToOutput(publicKey).(pulumi.StringOutput)
	keyPair.PrivateKey = pulumi.ToSecret(pulumi.Any(privateKey).AsIntArrayOutput()).(pulumi.IntArrayOutput)
	keyPair.Json = pulumi.ToSecret(pulumi.String(string(json))).(pulumi.StringOutput)

	if err := ctx.RegisterResourceOutputs(keyPair, pulumi.Map{
		"publicKey":  keyPair.PublicKey,
		"privateKey": keyPair.PrivateKey,
		"json":       keyPair.Json,
	}); err != nil {
		return nil, err
	}

	return keyPair, nil
}
