// Copyright © 2022 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PluginInfo collects information about plugins
type PluginInfo struct {
	Name        string
	Status      string
	Description string
	URL         string
	Release     PluginRelease
}

// PluginRelease collects info about a release for a plugin
type PluginRelease struct {
	Download string
	Shas     PluginShas
}

type PluginShas struct {
	DawrwinAmd64 string `yaml:"darwin-amd64"`
	LinuxAmd64   string `yaml:"linux-amd64"`
}

// NewSearchCommand implements 'kn plugin search' command
func NewSearchCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Prints the plugin version",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "Search results here\n")
			return nil
		},
	}
}
