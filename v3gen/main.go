package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
)

const (
	v2Fame      = "api/docs/docs.swagger.json"
	v3FnameJSON = "api/docs/v3/openapi3.json"
	v3FnameYAML = "api/docs/v3/openapi3.yaml"
)

func main() {
	docV2, err := loadV2()
	if err != nil {
		log.Fatalf("unable to load openapi v2: %v", err)
	}

	docV3, err := convertToV3(docV2)
	if err != nil {
		log.Fatalf("unable to convert to openapi v3 standards: %v", err)
	}

	err = writeOutDocV3(docV3)
	if err != nil {
		log.Fatalf("unable to write openapi v3 files: %+v", err)
	}
}

func loadV2() (*openapi2.T, error) {
	f, err := os.Open(v2Fame)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", v2Fame, err)
	}
	defer f.Close()

	docV2 := &openapi2.T{}
	err = json.NewDecoder(f).Decode(docV2)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", v2Fame, err)
	}

	return docV2, nil
}

func convertToV3(docV2 *openapi2.T) (*openapi3.T, error) {
	docV3, err := openapi2conv.ToV3(docV2)
	if err != nil {
		return nil, fmt.Errorf("failed to generate docV3: %w", err)
	}

	return docV3, nil
}

func writeOutDocV3(docV3 *openapi3.T) error {
	b, err := json.MarshalIndent(docV3, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to indent docV3 json: %w", err)
	}

	if err := os.WriteFile(v3FnameJSON, b, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", v3FnameJSON, err)
	}

	fmt.Println("created", v3FnameJSON)

	yamlBuf, err := yaml.JSONToYAML(b)
	if err != nil {
		return fmt.Errorf("failed to convert json docV3 to yaml: %w", err)
	}

	err = os.WriteFile(v3FnameYAML, yamlBuf, 0644)
	if err != nil {
		return fmt.Errorf("failed to write %s: %w ", v3FnameYAML, err)
	}

	fmt.Println("created", v3FnameYAML)

	return nil

}
