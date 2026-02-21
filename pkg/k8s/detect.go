package k8s

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

type Resource struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

func IsK8sResource(data []byte) bool {
	documents := splitDocuments(data)
	if len(documents) == 0 {
		return false
	}

	for _, doc := range documents {
		doc = bytes.TrimSpace(doc)
		if len(doc) == 0 {
			continue
		}

		var resource Resource
		if err := yaml.Unmarshal(doc, &resource); err != nil {
			return false
		}
		if resource.ApiVersion == "" || resource.Kind == "" {
			return false
		}
	}
	return true
}

func splitDocuments(input []byte) [][]byte {
	var documents [][]byte
	lines := bytes.Split(input, []byte("\n"))
	var currentDoc [][]byte

	for _, line := range lines {
		trimmed := bytes.TrimSpace(line)
		if bytes.Equal(trimmed, []byte("---")) {
			if len(currentDoc) > 0 {
				documents = append(documents, bytes.Join(currentDoc, []byte("\n")))
				currentDoc = nil
			}
			continue
		}
		currentDoc = append(currentDoc, line)
	}

	if len(currentDoc) > 0 {
		documents = append(documents, bytes.Join(currentDoc, []byte("\n")))
	}

	return documents
}

func GetResourceInfo(data []byte) (apiVersion, kind string, ok bool) {
	var resource Resource
	if err := yaml.Unmarshal(data, &resource); err != nil {
		return "", "", false
	}
	if resource.ApiVersion == "" || resource.Kind == "" {
		return "", "", false
	}
	return resource.ApiVersion, resource.Kind, true
}
