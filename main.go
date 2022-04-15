/*
Copyright © 2022 natakorn

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"embed"
	"os"

	"github.com/wewillapp-com/we-address/cmd"
	"github.com/wewillapp-com/we-address/pkg/address"
)

//go:embed static/raw_data/*.csv
var rawData embed.FS

func main() {
	address.InitConfig(&address.Config{})
	//check if user try to run command
	if len(os.Args) > 1 {
		cmd.RawData = rawData
		cmd.Execute()
	}
}
