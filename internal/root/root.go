// Copyright © 2021 The Knative Authors
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

package root

import (
	"github.com/spf13/cobra"
	"knative.dev/kn-plugin-sample/internal/command"
)

// NewSourceKafkaCommand represents the plugin's entrypoint
func NewRootCommand() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "kn-sample",
		Short: "Sample kn plugin printing out a nice message",
		Long:  `Longer description of this fantastic plugin that can go over several lines.`,
	}

	rootCmd.AddCommand(command.NewPrintCommand())
	rootCmd.AddCommand(command.NewVersionCommand())

	return rootCmd
}
