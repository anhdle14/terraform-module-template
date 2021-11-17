package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/hello",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.OutputMap(t, terraformOptions, "out")

	expectedLen := 2
	expectedMap := map[string]string{
		"upper": "HELLO, WORLD!",
		"lower": "hello, world!",
	}

	require.Len(t, output, expectedLen, "Output should contain %d items", expectedLen)
	require.Equal(t, expectedMap, output, "Map %q should match %q", expectedMap, output)
}
