package formatter

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Options struct {
	Indent int
}

type Formatter struct {
	options Options
}

func New(opts Options) *Formatter {
	if opts.Indent <= 0 {
		opts.Indent = 2
	}
	return &Formatter{options: opts}
}

func (f *Formatter) Format(input []byte) ([]byte, error) {
	documents := f.splitDocuments(input)
	if len(documents) == 0 {
		return nil, fmt.Errorf("no YAML documents found")
	}

	var formattedDocs []string
	for i, doc := range documents {
		doc = bytes.TrimSpace(doc)
		if len(doc) == 0 {
			continue
		}

		formatted, err := f.formatSingleDocument(doc)
		if err != nil {
			return nil, fmt.Errorf("failed to format document %d: %w", i+1, err)
		}
		formattedDocs = append(formattedDocs, strings.TrimSpace(formatted))
	}

	if len(formattedDocs) == 0 {
		return nil, fmt.Errorf("no valid documents to format")
	}

	return []byte(strings.Join(formattedDocs, "\n---\n") + "\n"), nil
}

func (f *Formatter) splitDocuments(input []byte) [][]byte {
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

func (f *Formatter) formatSingleDocument(input []byte) (string, error) {
	var node yaml.Node
	if err := yaml.Unmarshal(input, &node); err != nil {
		return "", fmt.Errorf("failed to parse YAML: %w", err)
	}

	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		f.processNode(node.Content[0])
	}

	var buf strings.Builder
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(f.options.Indent)

	if err := encoder.Encode(&node); err != nil {
		return "", fmt.Errorf("failed to encode YAML: %w", err)
	}
	encoder.Close()

	return buf.String(), nil
}

func (f *Formatter) processNode(node *yaml.Node) {
	f.processNodeWithContext(node, "")
}

func (f *Formatter) processNodeWithContext(node *yaml.Node, parentKey string) {
	if node == nil {
		return
	}

	switch node.Kind {
	case yaml.MappingNode:
		f.formatMappingWithContext(node, parentKey)
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 < len(node.Content) {
				key := node.Content[i].Value
				f.processNodeWithContext(node.Content[i+1], key)
			}
		}
	case yaml.SequenceNode:
		for _, child := range node.Content {
			f.processNodeWithContext(child, parentKey)
		}
	}
}

func (f *Formatter) formatMappingWithContext(node *yaml.Node, parentKey string) {
	if len(node.Content) == 0 {
		return
	}

	type keyValue struct {
		key   *yaml.Node
		value *yaml.Node
		order int
	}

	var pairs []keyValue
	for i := 0; i < len(node.Content); i += 2 {
		if i+1 < len(node.Content) {
			key := node.Content[i].Value
			pairs = append(pairs, keyValue{
				key:   node.Content[i],
				value: node.Content[i+1],
				order: f.getFieldOrder(key, parentKey),
			})
		}
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].order != pairs[j].order {
			return pairs[i].order < pairs[j].order
		}
		return pairs[i].key.Value < pairs[j].key.Value
	})

	node.Content = make([]*yaml.Node, 0, len(pairs)*2)
	for _, pair := range pairs {
		node.Content = append(node.Content, pair.key, pair.value)
	}
}

func (f *Formatter) formatMapping(node *yaml.Node) {
	f.formatMappingWithContext(node, "")
}

func (f *Formatter) getFieldOrder(key, parentKey string) int {
	topLevelOrder := map[string]int{
		"apiVersion": 1,
		"kind":       2,
		"metadata":   3,
		"spec":       4,
		"data":       5,
		"status":     6,
	}

	metadataOrder := map[string]int{
		"name":        1,
		"namespace":   2,
		"labels":      3,
		"annotations": 4,
	}

	containerOrder := map[string]int{
		"name":            1,
		"image":           2,
		"imagePullPolicy": 3,
		"command":         4,
		"args":            5,
		"workingDir":      6,
		"ports":           7,
		"env":             8,
		"resources":       9,
		"volumeMounts":    10,
		"livenessProbe":   11,
		"readinessProbe":  12,
		"startupProbe":    13,
		"lifecycle":       14,
		"securityContext": 15,
	}

	podSpecOrder := map[string]int{
		"replicas":                      1,
		"selector":                      2,
		"template":                      3,
		"serviceName":                   4,
		"serviceAccountName":            5,
		"serviceAccount":                6,
		"automountServiceAccountToken":  7,
		"nodeSelector":                  8,
		"affinity":                      9,
		"tolerations":                   10,
		"initContainers":                11,
		"containers":                    12,
		"volumes":                       13,
		"restartPolicy":                 14,
		"terminationGracePeriodSeconds": 15,
	}

	switch parentKey {
	case "metadata":
		if order, ok := metadataOrder[key]; ok {
			return order
		}
	case "containers", "initContainers":
		if order, ok := containerOrder[key]; ok {
			return order
		}
	case "spec":
		if order, ok := podSpecOrder[key]; ok {
			return order
		}
	}

	if order, ok := topLevelOrder[key]; ok {
		return order
	}

	return 100
}
