package get_token_v2

import (
	"github.com/ondrejsika/notion-backup/cmd/root"
	"github.com/ondrejsika/notion-backup/lib/login"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "login",
	Short: "Login & get token_v2",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		login.Login()
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
