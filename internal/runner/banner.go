// Package runner executes the enumeration process.
package runner

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/nuclei/v3/internal/pdcp"
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/config"
	updateutils "github.com/projectdiscovery/utils/update"
)

var banner = fmt.Sprintf(`
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/   %s
`, config.Version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tprojectdiscovery.io\n\n")
}

// NucleiToolUpdateCallback updates nuclei binary/tool to latest version 更新nuclei
func NucleiToolUpdateCallback() {
	showBanner()
	updateutils.GetUpdateToolCallback(config.BinaryName, config.Version)()
}

// AuthWithPDCP is used to authenticate with PDCP AuthWithPDCP用于使用PDCP进行身份验证
func AuthWithPDCP() {
	showBanner()
	pdcp.CheckNValidateCredentials(config.BinaryName)
}
