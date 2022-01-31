
package root

import (
	"github.com/ondrejsika/notion-backup/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "notion-backup",
	Short: "notion-backup, " + version.Version,
}
