package backup

import (
	"github.com/ondrejsika/notion-backup/cmd/root"
	"github.com/ondrejsika/notion-backup/lib/backup"
	"github.com/spf13/cobra"
)

var FlagToken string
var FlagSpaceID string
var FlagName string
var FlagBackupDir string

var Cmd = &cobra.Command{
	Use:     "backup",
	Short:   "Backup space in HTML and MD + CSV",
	Aliases: []string{"b"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		backup.BackupAll(FlagToken, FlagSpaceID, FlagName, FlagBackupDir)
	},
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
	Cmd.PersistentFlags().StringVarP(
		&FlagSpaceID,
		"space-id",
		"s",
		"",
		"Space ID",
	)
	Cmd.MarkPersistentFlagRequired("space-id")
	Cmd.PersistentFlags().StringVarP(
		&FlagName,
		"name",
		"n",
		"",
		"backup name",
	)
	Cmd.MarkPersistentFlagRequired("name")
	Cmd.PersistentFlags().StringVarP(
		&FlagBackupDir,
		"backup-dir",
		"d",
		".",
		"backup dir (default is current dir)",
	)
}
