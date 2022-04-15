package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/wewillapp-com/we-address/internal/database"
)

func confirmDatabase() {
	dbName := database.DB.Migrator().CurrentDatabase()
	p := &survey.Confirm{
		Message: fmt.Sprintf("migrate to database \u001b[32m%s\u001b[0m", dbName),
		Default: true,
	}
	survey.AskOne(p, &answer.Confirmed)
}
