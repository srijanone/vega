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

// CheckForLatestVersion checks for latest releases
func CheckForLatestVersion() {
	currentVersionStr := version.New().FormatVersion(true)

	u := updater.NewUpdater(vegaRepo, currentVersionStr)

	latestVersion, available, err := u.IsLatestAvailable()
	if err != nil {
		fmt.Printf("Error in checking latest version: %v\n", err)
	}

	if !available {
		return
	}

	update := false
	prompt := &survey.Confirm{
		Message: "New update available, Do you want to update Vega to " + latestVersion + " version?",
	}
	survey.AskOne(prompt, &update, nil)

	if !update {
		return
	}

	err = u.SelfUpdate()
	if err != nil {
		fmt.Printf("Error in updating Vega: %v\n", err)
	}
}
