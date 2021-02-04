/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package json

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'json' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "json [PATH]",
		Short:       "Generate JSON of inputs and outputs",
		Annotations: cli.Annotations("json"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Escape, "escape", true, "escape special characters")

	// deprecation
	cmd.PersistentFlags().BoolVar(&config.Settings.Deprecated.NoEscape, "no-escape", false, "do not escape special characters")
	cmd.PersistentFlags().MarkDeprecated("no-escape", "use '--escape=false' instead") //nolint:errcheck

	return cmd
}
