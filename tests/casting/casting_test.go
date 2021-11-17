package test

import (
	"strconv"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestTerraformCastingExample(t *testing.T) {
	a := 1 + 2
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/casting",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.OutputMap(t, terraformOptions, "out")

	// expectedLen := 2
	expectedMap := map[string]string{
		"upper": strconv.Itoa(a),
		"lower": strconv.Itoa(a),
	}

	// require.Len(t, output, expectedLen, "Output should contain %d items", expectedLen)
	require.Equal(t, expectedMap, output, "Map %q should match %q", expectedMap, output)
}
