package vega

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"

	updater "github.com/srijanone/vega/pkg/updater"
	version "github.com/srijanone/vega/pkg/version"
)

const (
	vegaRepo = "srijanone/vega"
)

func CheckForLatestVersion() {
	currentVersionStr := version.New().FormatVersion(true)
	fmt.Println("currentVersionStr --", currentVersionStr)

	u := updater.NewUpdater(vegaRepo, currentVersionStr)

	latestVersion, available, err := u.IsLatestAvailable()
	if err != nil {
		fmt.Printf("Error in checking latest version: %v", err)
	}
	fmt.Println("latestVersion --", latestVersion)
	fmt.Println("available --", available)


	if !available {
		fmt.Println("Already latest version")
		return
	}

	update := false
	prompt := &survey.Confirm{
		Message: "Do you want to update Vega to " + latestVersion + "?",
	}
	survey.AskOne(prompt, &update, nil)

	if !update {
		return
	}

	err = u.SelfUpdate()
	if err != nil {
		fmt.Println("Error in updating Vega")
	}
}
