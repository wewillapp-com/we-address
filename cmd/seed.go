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
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/wewillapp-com/we-address/internal/database"
	"github.com/wewillapp-com/we-address/pkg/address"
)

type Answer struct {
	Selected  []string
	Confirmed bool
}

var answer Answer

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Add address data to database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectDatabase()
		defer database.Close()
		p := &survey.MultiSelect{
			Message: "What do you want to seed?",
			Options: []string{"province", "amphur", "district"},
			Default: []string{"province", "amphur", "district"},
		}
		survey.AskOne(p, &answer.Selected)
		if !answer.Confirmed {
			confirmDatabase()
		}
		if answer.Confirmed {
			for _, v := range answer.Selected {
				switch v {
				case "province":
					if err := seedProvince(); err != nil {
						log.Fatal(err)
					}
				case "amphur":
					if err := seedAmphur(); err != nil {
						log.Fatal(err)
					}
				case "district":
					if err := seedDistrict(); err != nil {
						log.Fatal(err)
					}
				}
			}
			fmt.Println("✅ database seed completed")
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func seedProvince() error {
	csvFile, err := RawData.Open("static/raw_data/provinces.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return fmt.Errorf("read csv file error >> %v", err)
	}
	var provinces []address.ProvinceModel
	for _, line := range csvLines[1:] {
		_id, _ := strconv.ParseUint(line[0], 10, 64)
		id := uint(_id)
		mo := address.ProvinceModel{
			ID:     id,
			NameTh: line[1],
			NameEn: line[2],
		}
		provinces = append(provinces, mo)
		// if err := database.DB.Save(&mo).Error; err != nil {
		// 	return fmt.Errorf("error, save data into database: %v", err.Error())
		// }
	}
	if err := database.DB.Save(&provinces).Error; err != nil {
		return fmt.Errorf("error, save data into database: %v", err.Error())
	}
	return nil
}

func seedAmphur() error {
	csvFile, err := RawData.Open("static/raw_data/amphurs.csv")
	if err != nil {
		return fmt.Errorf("open file error >> %v", err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return fmt.Errorf("read csv file error >> %v", err)
	}
	var amphurs []address.AmphurModel
	for _, line := range csvLines[1:] {
		id := strToUInt(line[0])
		pId := strToUInt(line[3])
		mo := address.AmphurModel{
			ID:         id,
			NameTh:     line[1],
			NameEn:     line[2],
			ProvinceID: pId,
		}
		amphurs = append(amphurs, mo)
	}
	if err := database.DB.Save(&amphurs).Error; err != nil {
		return fmt.Errorf("error, save data into database: %v", err.Error())
	}
	return nil
}

func seedDistrict() error {
	csvFile, err := RawData.Open("static/raw_data/districts.csv")
	if err != nil {
		return fmt.Errorf("open file error >> %v", err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return fmt.Errorf("read csv file error >> %v", err)
	}
	var districts []address.DistrictModel
	for _, line := range csvLines[1:] {
		id := strToUInt(line[0])
		id2 := strToUInt(line[4])
		mo := address.DistrictModel{
			ID:       id,
			ZipCode:  line[1],
			NameTh:   line[2],
			NameEn:   line[3],
			AmphurID: id2,
		}
		districts = append(districts, mo)

	}
	if err := database.DB.Save(&districts).Error; err != nil {
		return fmt.Errorf("error, save data into database: %v", err.Error())
	}
	return nil
}

func strToUInt(str string) uint {
	i, _ := strconv.ParseUint(str, 10, 64)
	return uint(i)
}
