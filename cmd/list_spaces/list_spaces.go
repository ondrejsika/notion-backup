package list_spaces

import (
	"fmt"

	"github.com/ondrejsika/notion-backup/cmd/root"
	"github.com/ondrejsika/notion-backup/lib/client"
	"github.com/spf13/cobra"
)

var FlagToken string

var Cmd = &cobra.Command{
	Use:     "list-spaces",
	Short:   "List Spaces",
	Args:    cobra.NoArgs,
	Aliases: []string{"list", "ls"},
	Run: func(c *cobra.Command, args []string) {
		api := client.New(FlagToken)
		spaces, _ := api.GetSpaces()
		for spaceID, spaceName := range spaces {
			fmt.Printf("%s: %s\n", spaceID, spaceName)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&FlagToken,
		"token",
		"t",
		"",
		"Notion's token_v2",
	)
	Cmd.MarkPersistentFlagRequired("token")
}
