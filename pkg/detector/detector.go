package detector

import (
	common "github.com/srijanone/vega/pkg/common"
)

func IsDrupal() bool {
	signature := "composer.json"
	result, _ := common.Exists(signature)
	return result
}
