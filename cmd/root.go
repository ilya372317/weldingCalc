package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "weldingCalc",
	Short: "Calculate welding expense",
	Long:  "Used for calculate expense on welding",
}
