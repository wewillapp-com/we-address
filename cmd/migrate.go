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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wewillapp-com/we-address/internal/database"
	"github.com/wewillapp-com/we-address/pkg/address"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectDatabase()
		confirmDatabase()
		defer database.Close()
		if answer.Confirmed {
			database.DB.AutoMigrate(&address.ProvinceModel{}, &address.AmphurModel{}, &address.DistrictModel{})
			if cmd.Flag("seed").Value.String() == "true" {
				seedCmd.Run(cmd, args)
			}
			fmt.Println("✅ migrate finished")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Bool("seed", false, "seed the database")
	h := address.GetConfig("DB_HOST")
	migrateCmd.Flags().String("host", h, "database host")
}
