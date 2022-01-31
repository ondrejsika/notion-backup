package cmd

import (
	"github.com/ondrejsika/notion-backup/cmd/root"
	_ "github.com/ondrejsika/notion-backup/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
