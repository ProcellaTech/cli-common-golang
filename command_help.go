/*
 * Copyright 2018. Akamai Technologies, Inc
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
 */

package akamai_cli

import (
	"os"

	"github.com/urfave/cli"
)

func CmdHelp(c *cli.Context) error {
	if c.Args().Present() {
		if sub_cmd := c.Args().Get(1); sub_cmd != "" {
			args := []string{os.Args[0]}
			args = append(args, os.Args[2:]...)
			args = append(args, "--help")
			os.Args = args
			App.Run(os.Args)
			return nil
		}

		cmd := c.Args().First()
		return cli.ShowCommandHelp(c, cmd)
	}

	return cli.ShowAppHelp(c)
}
