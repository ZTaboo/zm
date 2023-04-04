package cmd

import (
	"ZM/router"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行ZM",
	Long:  `运行ZM.`,
	Run:   router.Run,
}
var runBackCmd = &cobra.Command{
	Use:   "backend",
	Short: "后台运行ZM",
	Long:  `后台运行ZM.`,
	Run:   router.BackRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.AddCommand(runBackCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
