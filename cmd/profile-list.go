package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "list profiles",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("filter", cmd.Flags().Lookup("filter"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		profiles, err := restClient.ListProfiles(viper.GetString("filter"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if cmd.Flags().Changed("details") {
			fmt.Println(formatString(profiles))
		} else {
			list := []map[string]string{}
			for _, profile := range profiles {
				var info = map[string]string{"id": profile.ID, "name": profile.Name}
				list = append(list, info)
			}
			fmt.Println(formatString(list))
		}
	},
}

func init() {
	profileCmd.AddCommand(profileListCmd)
	profileListCmd.Flags().Bool("details", false, "show details")
	profileListCmd.Flags().String("filter", "", "filter query string")
}
