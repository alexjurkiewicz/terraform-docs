/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/print"
)

func TestFormatFactory(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
		wantErr  bool
	}{
		{
			name:     "format factory from name",
			format:   "asciidoc",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "adoc",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "asciidoc document",
			expected: "*format.AsciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "asciidoc doc",
			expected: "*format.AsciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "adoc document",
			expected: "*format.AsciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "adoc doc",
			expected: "*format.AsciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "asciidoc table",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "asciidoc tbl",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "adoc table",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "adoc tbl",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "json",
			expected: "*format.JSON",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown",
			expected: "*format.MarkdownTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "md",
			expected: "*format.MarkdownTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown document",
			expected: "*format.MarkdownDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown doc",
			expected: "*format.MarkdownDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "md document",
			expected: "*format.MarkdownDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "md doc",
			expected: "*format.MarkdownDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown table",
			expected: "*format.MarkdownTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown tbl",
			expected: "*format.MarkdownTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "md table",
			expected: "*format.MarkdownTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "md tbl",
			expected: "*format.MarkdownTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "pretty",
			expected: "*format.Pretty",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "tfvars hcl",
			expected: "*format.TfvarsHCL",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "tfvars json",
			expected: "*format.TfvarsJSON",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "toml",
			expected: "*format.TOML",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "xml",
			expected: "*format.XML",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "yaml",
			expected: "*format.YAML",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "unknown",
			expected: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := &print.Settings{}
			actual, err := Factory(tt.format, settings)
			if tt.wantErr {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, reflect.TypeOf(actual).String())
			}
		})
	}
}
