package yamlparser

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func ParsePipeline(p string) ([]string, error) {
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var result []string
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, errors.New("expected a flat string array")
	}
	return result, nil
}
