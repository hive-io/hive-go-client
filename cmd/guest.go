package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-cmp/cmp"
	"github.com/hive-io/hive-go-client/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var guestCmd = &cobra.Command{
	Use:   "guest",
	Short: "guest operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
		os.Exit(0)
	},
}

var guestAssignCmd = &cobra.Command{
	Use:   "assign [GuestName]",
	Short: "assign guest to a user",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("guest-user")
		cmd.MarkFlagRequired("guest-realm")
	},
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res, err := restClient.AssignGuest(guest.PoolID, viper.GetString("guest-user"), viper.GetString("guest-realm"), guest.Name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(formatString(res))
	},
}

var guestDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete guest pool",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = guest.Delete(restClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestDiffCmd = &cobra.Command{
	Use:   "diff [guest1] [guest2]",
	Short: "compare 2 guests",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		guest1, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		guest2, err := restClient.GetGuest(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(cmp.Diff(guest1, guest2))
	},
}

var guestGetCmd = &cobra.Command{
	Use:   "get [Name]",
	Short: "get guest details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(formatString(guest))
	},
}

var guestListCmd = &cobra.Command{
	Use:   "list",
	Short: "list guests",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("filter", cmd.Flags().Lookup("filter"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		guests, err := restClient.ListGuests(viper.GetString("filter"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if cmd.Flags().Changed("details") {
			fmt.Println(formatString(guests))
		} else {
			guestList := []string{}
			for _, guest := range guests {
				guestList = append(guestList, guest.Name)
			}
			fmt.Println(formatString(guestList))
		}
	},
}

var guestUpdateCmd = &cobra.Command{
	Use:   "update [file]",
	Short: "update a guest",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var file *os.File
		var err error
		if args[0] == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		defer file.Close()
		data, _ := ioutil.ReadAll(file)
		var guest rest.Guest
		err = unmarshal(data, &guest)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		msg, err := guest.Update(restClient)
		fmt.Println(msg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestPoweroffCmd = &cobra.Command{
	Use:   "poweroff [Name]",
	Short: "force power off guest",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		guest.Poweroff(restClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestRebootCmd = &cobra.Command{
	Use:   "reboot [Name]",
	Short: "reboot guest",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		guest.Reboot(restClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestReleaseCmd = &cobra.Command{
	Use:   "release [Name]",
	Short: "release guest assignment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = restClient.ReleaseGuest(guest.PoolID, guest.Username, guest.Name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestResetCmd = &cobra.Command{
	Use:   "reset [Name]",
	Short: "force reset guest",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		guest.Reset(restClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestShutdownCmd = &cobra.Command{
	Use:   "shutdown [Name]",
	Short: "shutdown guest",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		guest.Shutdown(restClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestBackupCmd = &cobra.Command{
	Use:   "backup [Name]",
	Short: "start guest backup",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		bindTaskFlags(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		guest.StartBackup(restClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var guestRestoreCmd = &cobra.Command{
	Use:   "backup [Name]",
	Short: "start guest backup",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		guest, err := restClient.GetGuest(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		handleTask(guest.Restore(restClient))

	},
}

func init() {
	RootCmd.AddCommand(guestCmd)

	guestCmd.AddCommand(guestAssignCmd)
	guestAssignCmd.Flags().String("guest-user", "", "user to assign to this guest")
	guestAssignCmd.Flags().String("guest-realm", "", "user's realm")

	guestCmd.AddCommand(guestDeleteCmd)
	guestCmd.AddCommand(guestDiffCmd)
	guestCmd.AddCommand(guestGetCmd)

	//list
	guestCmd.AddCommand(guestListCmd)
	guestListCmd.Flags().Bool("details", false, "show details")
	guestListCmd.Flags().String("filter", "", "filter query string")

	guestCmd.AddCommand(guestPoweroffCmd)
	guestCmd.AddCommand(guestRebootCmd)
	guestCmd.AddCommand(guestReleaseCmd)
	guestCmd.AddCommand(guestResetCmd)
	guestCmd.AddCommand(guestShutdownCmd)
	guestCmd.AddCommand(guestUpdateCmd)

	guestCmd.AddCommand(guestBackupCmd)
	guestCmd.AddCommand(guestRestoreCmd)
	addTaskFlags(guestRestoreCmd)
}
