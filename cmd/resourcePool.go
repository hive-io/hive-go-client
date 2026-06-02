package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/hive-io/hive-go-client/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var resourcePoolCmd = &cobra.Command{
	Use:   "resource-pool",
	Short: "resource pool operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
		os.Exit(0)
	},
}

var resourcePoolListCmd = &cobra.Command{
	Use:   "list",
	Short: "list resource pools",
	PreRun: func(cmd *cobra.Command, args []string) {
		bindListFlags(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		pools, err := restClient.ListResourcePools(listFlagsToQuery())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if cmd.Flags().Changed("details") {
			fmt.Println(formatString(pools))
		} else {
			list := []map[string]string{}
			for _, pool := range pools {
				info := map[string]string{"id": pool.ID, "name": pool.Name, "type": pool.Type}
				list = append(list, info)
			}
			fmt.Println(formatString(list))
		}
	},
}

var resourcePoolGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get resource pool details",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("id", cmd.Flags().Lookup("id"))
		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		var pool *rest.ResourcePool
		var err error
		switch {
		case cmd.Flags().Changed("id"):
			pool, err = restClient.GetResourcePool(viper.GetString("id"))
		case cmd.Flags().Changed("name"):
			pool, err = restClient.GetResourcePoolByName(viper.GetString("name"))
		default:
			cmd.Usage()
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(formatString(pool))
	},
}

var resourcePoolCreateCmd = &cobra.Command{
	Use:   "create [file]",
	Short: "create a new resource pool",
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
		data, _ := io.ReadAll(file)
		var pool rest.ResourcePool
		err = unmarshal(data, &pool)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		msg, err := pool.Create(restClient)
		fmt.Println(msg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var resourcePoolDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a resource pool",
	Run: func(cmd *cobra.Command, args []string) {
		var pool *rest.ResourcePool
		var err error
		switch {
		case cmd.Flags().Changed("id"):
			id, _ := cmd.Flags().GetString("id")
			pool, err = restClient.GetResourcePool(id)
		case cmd.Flags().Changed("name"):
			name, _ := cmd.Flags().GetString("name")
			pool, err = restClient.GetResourcePoolByName(name)
		default:
			cmd.Usage()
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = pool.Delete(restClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var resourcePoolUpdateConfigCmd = &cobra.Command{
	Use:   "update-config",
	Short: "update resource pool name, description, or tags",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("id", cmd.Flags().Lookup("id"))
		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		var pool *rest.ResourcePool
		var err error
		switch {
		case cmd.Flags().Changed("id"):
			pool, err = restClient.GetResourcePool(viper.GetString("id"))
		case cmd.Flags().Changed("name"):
			pool, err = restClient.GetResourcePoolByName(viper.GetString("name"))
		default:
			cmd.Usage()
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		config := rest.ResourcePoolConfig{}
		if cmd.Flags().Changed("new-name") {
			config.Name, _ = cmd.Flags().GetString("new-name")
		}
		if cmd.Flags().Changed("description") {
			config.Description, _ = cmd.Flags().GetString("description")
		}
		if cmd.Flags().Changed("tags") {
			config.Tags, _ = cmd.Flags().GetStringArray("tags")
		}
		msg, err := pool.UpdateConfiguration(restClient, config)
		fmt.Println(msg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var resourcePoolAddMemberCmd = &cobra.Command{
	Use:   "add-member",
	Short: "add a member to a resource pool",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("id", cmd.Flags().Lookup("id"))
		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		var pool *rest.ResourcePool
		var err error
		switch {
		case cmd.Flags().Changed("id"):
			pool, err = restClient.GetResourcePool(viper.GetString("id"))
		case cmd.Flags().Changed("name"):
			pool, err = restClient.GetResourcePoolByName(viper.GetString("name"))
		default:
			cmd.Usage()
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		memberID, _ := cmd.Flags().GetString("member-id")
		source, _ := cmd.Flags().GetString("source")
		memberName, _ := cmd.Flags().GetString("member-name")
		msg, err := pool.AddMember(restClient, memberID, source, memberName)
		fmt.Println(msg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var resourcePoolRemoveMemberCmd = &cobra.Command{
	Use:   "remove-member",
	Short: "remove a member from a resource pool",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("id", cmd.Flags().Lookup("id"))
		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		var pool *rest.ResourcePool
		var err error
		switch {
		case cmd.Flags().Changed("id"):
			pool, err = restClient.GetResourcePool(viper.GetString("id"))
		case cmd.Flags().Changed("name"):
			pool, err = restClient.GetResourcePoolByName(viper.GetString("name"))
		default:
			cmd.Usage()
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		memberID, _ := cmd.Flags().GetString("member-id")
		msg, err := pool.RemoveMember(restClient, memberID)
		fmt.Println(msg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(resourcePoolCmd)

	resourcePoolCmd.AddCommand(resourcePoolListCmd)
	addListFlags(resourcePoolListCmd)

	resourcePoolCmd.AddCommand(resourcePoolGetCmd)
	resourcePoolGetCmd.Flags().StringP("id", "i", "", "resource pool id")
	resourcePoolGetCmd.Flags().StringP("name", "n", "", "resource pool name")

	resourcePoolCmd.AddCommand(resourcePoolCreateCmd)

	resourcePoolCmd.AddCommand(resourcePoolDeleteCmd)
	resourcePoolDeleteCmd.Flags().StringP("id", "i", "", "resource pool id")
	resourcePoolDeleteCmd.Flags().StringP("name", "n", "", "resource pool name")

	resourcePoolCmd.AddCommand(resourcePoolUpdateConfigCmd)
	resourcePoolUpdateConfigCmd.Flags().StringP("id", "i", "", "resource pool id")
	resourcePoolUpdateConfigCmd.Flags().StringP("name", "n", "", "resource pool name")
	resourcePoolUpdateConfigCmd.Flags().String("new-name", "", "new name for the resource pool")
	resourcePoolUpdateConfigCmd.Flags().String("description", "", "description")
	resourcePoolUpdateConfigCmd.Flags().StringArray("tags", nil, "tags (repeatable)")

	resourcePoolCmd.AddCommand(resourcePoolAddMemberCmd)
	resourcePoolAddMemberCmd.Flags().StringP("id", "i", "", "resource pool id")
	resourcePoolAddMemberCmd.Flags().StringP("name", "n", "", "resource pool name")
	resourcePoolAddMemberCmd.Flags().String("member-id", "", "id of the member to add")
	resourcePoolAddMemberCmd.Flags().String("source", "", "source cluster for the member")
	resourcePoolAddMemberCmd.Flags().String("member-name", "", "display name for the member")
	resourcePoolAddMemberCmd.MarkFlagRequired("member-id")

	resourcePoolCmd.AddCommand(resourcePoolRemoveMemberCmd)
	resourcePoolRemoveMemberCmd.Flags().StringP("id", "i", "", "resource pool id")
	resourcePoolRemoveMemberCmd.Flags().StringP("name", "n", "", "resource pool name")
	resourcePoolRemoveMemberCmd.Flags().String("member-id", "", "id of the member to remove")
	resourcePoolRemoveMemberCmd.MarkFlagRequired("member-id")
}
