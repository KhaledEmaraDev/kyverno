package patch

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
	"gomodules.xyz/jsonpatch/v2"
	"sigs.k8s.io/kustomize/api/filters/patchstrategicmerge"
	filtersutil "sigs.k8s.io/kustomize/kyaml/filtersutil"
	yaml "sigs.k8s.io/kustomize/kyaml/yaml"
)

// ProcessStrategicMergePatch ...
func ProcessStrategicMergePatch(logger logr.Logger, overlay interface{}, resource resource) (patches, error) {
	overlayBytes, err := json.Marshal(overlay)
	if err != nil {
		logger.Error(err, "failed to marshal overlay")
		return [][]byte{}, err
	}
	resourceBytes, err := resource.MarshalJSON()
	if err != nil {
		logger.Error(err, "failed to marshal resource")
		return [][]byte{}, err
	}
	patchedBytes, err := strategicMergePatch(logger, string(resourceBytes), string(overlayBytes))
	if err != nil {
		logger.Error(err, "failed to apply patchStrategicMerge")
		return [][]byte{}, err
	}
	patches, err := jsonpatch.CreatePatch(resourceBytes, patchedBytes)
	if err != nil {
		return [][]byte{}, err
	}
	return ConvertPatches(patches...), nil
}

func strategicMergePatch(logger logr.Logger, base, overlay string) ([]byte, error) {
	preprocessedYaml, err := preProcessStrategicMergePatch(logger, overlay, base)
	if err != nil {
		_, isConditionError := err.(ConditionError)
		_, isGlobalConditionError := err.(GlobalConditionError)

		if isConditionError || isGlobalConditionError {
			if err = preprocessedYaml.UnmarshalJSON([]byte(`{}`)); err != nil {
				return []byte{}, err
			}
		} else {
			return []byte{}, fmt.Errorf("failed to preProcess rule: %+v", err)
		}
	}

	patchStr, _ := preprocessedYaml.String()
	logger.V(3).Info("applying strategic merge patch", "patch", patchStr)
	f := patchstrategicmerge.Filter{
		Patch: preprocessedYaml,
	}

	baseObj := buffer{Buffer: bytes.NewBufferString(base)}
	err = filtersutil.ApplyToJSON(f, baseObj)

	return baseObj.Bytes(), err
}

func preProcessStrategicMergePatch(logger logr.Logger, pattern, resource string) (*yaml.RNode, error) {
	patternNode := yaml.MustParse(pattern)
	resourceNode := yaml.MustParse(resource)

	err := PreProcessPattern(logger, patternNode, resourceNode)

	return patternNode, err
}
