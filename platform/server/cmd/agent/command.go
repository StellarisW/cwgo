/*
 *
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package agent

import (
	"github.com/cloudwego/cwgo/platform/server/shared/args"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	aArgs := args.NewAgentArgs()
	cmd := &cobra.Command{
		Use:   "agent",
		Short: "agent service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := run(aArgs); err != nil {
				return err
			}
			return nil
		},
	}
	aArgs.AddFlags(cmd)
	return cmd
}
