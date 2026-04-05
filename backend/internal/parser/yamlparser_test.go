package yamlparser_test

import (
	"testing"

	yamlparser "github.com/sapienfrom2000s/trident/backend/internal/parser"
)

func TestParsePipeline(t *testing.T) {
	flatarrayFixture := "fixtures/array.yml"
	nestedArrayFixture := "fixtures/nested-array.yml"
	dictFixture := "fixtures/dict.yml"

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Parse Flat Array",
			path:    flatarrayFixture,
			wantErr: false,
		},
		{
			name:    "Parse Nested Array",
			path:    nestedArrayFixture,
			wantErr: true,
		},
		{
			name:    "Parse Dictionary",
			path:    dictFixture,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := yamlparser.ParsePipeline(tt.path)
			gotErr := err != nil
			if gotErr != tt.wantErr {
				t.Errorf("Got: %v, Want: %v", gotErr, tt.wantErr)
			}
		})
	}
}
