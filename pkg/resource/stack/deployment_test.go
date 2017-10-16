// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
)

// TestDeploymentSerialization creates a basic
func TestDeploymentSerialization(t *testing.T) {
	res := resource.NewState(
		tokens.Type("Test"),
		resource.NewURN(
			tokens.QName("test"),
			tokens.PackageName("resource/test"),
			tokens.Type("Test"),
			tokens.QName("resource-x"),
		),
		true,
		resource.ID("test-resource-x"),
		resource.NewPropertyMapFromMap(map[string]interface{}{
			"in-nil":         nil,
			"in-bool":        true,
			"in-float64":     float64(1.5),
			"in-string":      "lumilumilo",
			"in-array":       []interface{}{"a", true, float64(32)},
			"in-empty-array": []interface{}{},
			"in-map": map[string]interface{}{
				"a": true,
				"b": float64(88),
				"c": "c-see-saw",
				"d": "d-dee-daw",
			},
			"in-empty-map": map[string]interface{}{},
		}),
		make(resource.PropertyMap),
		resource.NewPropertyMapFromMap(map[string]interface{}{
			"out-nil":         nil,
			"out-bool":        false,
			"out-float64":     float64(76),
			"out-string":      "loyolumiloom",
			"out-array":       []interface{}{false, "zzxx"},
			"out-empty-array": []interface{}{},
			"out-map": map[string]interface{}{
				"x": false,
				"y": "z-zee-zaw",
				"z": float64(999.9),
			},
			"out-empty-map": map[string]interface{}{},
		}),
		nil,
	)

	dep := SerializeResource(res)

	// assert some things about the deployment record:
	assert.NotNil(t, dep)
	assert.NotNil(t, dep.ID)
	assert.Equal(t, resource.ID("test-resource-x"), dep.ID)
	assert.Equal(t, tokens.Type("Test"), dep.Type)

	// assert some things about the inputs:
	assert.NotNil(t, dep.Inputs)
	assert.Nil(t, dep.Inputs["in-nil"])
	assert.NotNil(t, dep.Inputs["in-bool"])
	assert.True(t, dep.Inputs["in-bool"].(bool))
	assert.NotNil(t, dep.Inputs["in-float64"])
	assert.Equal(t, float64(1.5), dep.Inputs["in-float64"].(float64))
	assert.NotNil(t, dep.Inputs["in-string"])
	assert.Equal(t, "lumilumilo", dep.Inputs["in-string"].(string))
	assert.NotNil(t, dep.Inputs["in-array"])
	assert.Equal(t, 3, len(dep.Inputs["in-array"].([]interface{})))
	assert.Equal(t, "a", dep.Inputs["in-array"].([]interface{})[0])
	assert.Equal(t, true, dep.Inputs["in-array"].([]interface{})[1])
	assert.Equal(t, float64(32), dep.Inputs["in-array"].([]interface{})[2])
	assert.NotNil(t, dep.Inputs["in-empty-array"])
	assert.Equal(t, 0, len(dep.Inputs["in-empty-array"].([]interface{})))
	assert.NotNil(t, dep.Inputs["in-map"])
	inmap := dep.Inputs["in-map"].(map[string]interface{})
	assert.Equal(t, 4, len(inmap))
	assert.NotNil(t, inmap["a"])
	assert.Equal(t, true, inmap["a"].(bool))
	assert.NotNil(t, inmap["b"])
	assert.Equal(t, float64(88), inmap["b"].(float64))
	assert.NotNil(t, inmap["c"])
	assert.Equal(t, "c-see-saw", inmap["c"].(string))
	assert.NotNil(t, inmap["d"])
	assert.Equal(t, "d-dee-daw", inmap["d"].(string))
	assert.NotNil(t, dep.Inputs["in-empty-map"])
	assert.Equal(t, 0, len(dep.Inputs["in-empty-map"].(map[string]interface{})))

	// assert some things about the outputs:
	assert.NotNil(t, dep.Outputs)
	assert.Nil(t, dep.Outputs["out-nil"])
	assert.NotNil(t, dep.Outputs["out-bool"])
	assert.False(t, dep.Outputs["out-bool"].(bool))
	assert.NotNil(t, dep.Outputs["out-float64"])
	assert.Equal(t, float64(76), dep.Outputs["out-float64"].(float64))
	assert.NotNil(t, dep.Outputs["out-string"])
	assert.Equal(t, "loyolumiloom", dep.Outputs["out-string"].(string))
	assert.NotNil(t, dep.Outputs["out-array"])
	assert.Equal(t, 2, len(dep.Outputs["out-array"].([]interface{})))
	assert.Equal(t, false, dep.Outputs["out-array"].([]interface{})[0])
	assert.Equal(t, "zzxx", dep.Outputs["out-array"].([]interface{})[1])
	assert.NotNil(t, dep.Outputs["out-empty-array"])
	assert.Equal(t, 0, len(dep.Outputs["out-empty-array"].([]interface{})))
	assert.NotNil(t, dep.Outputs["out-map"])
	outmap := dep.Outputs["out-map"].(map[string]interface{})
	assert.Equal(t, 3, len(outmap))
	assert.NotNil(t, outmap["x"])
	assert.Equal(t, false, outmap["x"].(bool))
	assert.NotNil(t, outmap["y"])
	assert.Equal(t, "z-zee-zaw", outmap["y"].(string))
	assert.NotNil(t, outmap["z"])
	assert.Equal(t, float64(999.9), outmap["z"].(float64))
	assert.NotNil(t, dep.Outputs["out-empty-map"])
	assert.Equal(t, 0, len(dep.Outputs["out-empty-map"].(map[string]interface{})))
}