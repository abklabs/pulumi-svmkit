package tests

import (
	"testing"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	svmkit "github.com/abklabs/pulumi-svmkit"
)

func TestKeyPairCreate(t *testing.T) {
	prov := provider()

	response, err := prov.Create(p.CreateRequest{
		Urn:        urn("KeyPair"),
		Properties: resource.PropertyMap{},
		Preview:    false,
	})

	require.NoError(t, err)

	publicKey := response.Properties["publicKey"].StringValue()
	privateKey := response.Properties["privateKey"].ArrayValue()
	jsonKey := response.Properties["json"].StringValue()

	assert.IsType(t, "", publicKey)
	assert.IsType(t, "", jsonKey)
	assert.IsType(t, []resource.PropertyValue{}, privateKey)
}

// urn is a helper function to build an urn for running integration tests.
func urn(typ string) resource.URN {
	return resource.NewURN("stack", "proj", "",
		tokens.Type("svm:svm:"+typ), "name")
}

// Create a test server.
func provider() integration.Server {
	return integration.NewServer(svmkit.Name, semver.MustParse("0.0.1"), svmkit.Provider())

}
