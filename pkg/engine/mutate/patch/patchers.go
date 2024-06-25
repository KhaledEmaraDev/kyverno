package patch

import (
	"github.com/go-logr/logr"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type (
	resource     = unstructured.Unstructured
	resource_raw = []byte
	patches      = [][]byte
)

// Patcher patches the resource
type Patcher interface {
	Patch(logr.Logger, resource) (patches, error)
}

// patchStrategicMergeHandler
type patchStrategicMergeHandler struct {
	patch apiextensions.JSON
}

func NewPatchStrategicMerge(patch apiextensions.JSON) Patcher {
	return patchStrategicMergeHandler{
		patch: patch,
	}
}

func (h patchStrategicMergeHandler) Patch(logger logr.Logger, resource resource) (patches, error) {
	return ProcessStrategicMergePatch(logger, h.patch, resource)
}

// patchesJSON6902Handler
type patchesJSON6902Handler struct {
	patches string
}

func NewPatchesJSON6902(patches string) Patcher {
	return patchesJSON6902Handler{
		patches: patches,
	}
}

func (h patchesJSON6902Handler) Patch(logger logr.Logger, _ resource) (patches, error) {
	patchJSON6902, err := convertPatchesToJSON(h.patches)
	if err != nil {
		logger.Error(err, "error in type conversion")
		return nil, err
	}

	return [][]byte{patchJSON6902}, nil
}
