/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/internal/testutil"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("markdown", "table")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		ShowRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-WithRequired")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-SortByName")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-SortByRequired")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Name:     true,
			Required: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByType(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByType: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-SortByType")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Type: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableEscapeCharacters(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		EscapeCharacters: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-EscapeCharacters")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationBelowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 0,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-IndentationBelowAllowed")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 10,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-IndentationAboveAllowed")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 4,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-IndentationOfFour")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OutputValues")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableHeaderFromFile(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		file   string
	}{
		{
			name:   "load module header from .adoc",
			golden: "table-HeaderFromADOCFile",
			file:   "doc.adoc",
		},
		{
			name:   "load module header from .md",
			golden: "table-HeaderFromMDFile",
			file:   "doc.md",
		},
		{
			name:   "load module header from .tf",
			golden: "table-HeaderFromTFFile",
			file:   "doc.tf",
		},
		{
			name:   "load module header from .txt",
			golden: "table-HeaderFromTXTFile",
			file:   "doc.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := testutil.Settings().WithSections().Build()

			expected, err := testutil.GetExpected("markdown", tt.golden)
			assert.Nil(err)

			options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
				HeaderFromFile: tt.file,
			})
			assert.Nil(err)

			module, err := testutil.GetModule(options)
			assert.Nil(err)

			printer := NewMarkdownTable(settings)
			actual, err := printer.Print(module, settings)

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}

func TestTableOutputValuesNoSensitivity(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OutputValuesNoSensitivity")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableEmpty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
		HeaderFromFile: "bad.tf",
	})
	options.ShowHeader = false // Since we don't show the header, the file won't be loaded at all
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewMarkdownTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal("", actual)
}
