package cmd

import (
	"fmt"
	"os"

	"github.com/hive-io/hive-go-client/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var taskWaitCmd = &cobra.Command{
	Use:   "wait",
	Short: "wait for a task to complete",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("id", cmd.Flags().Lookup("id"))
		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		viper.BindPFlag("progress", cmd.Flags().Lookup("progress"))
		viper.BindPFlag("progress-bar", cmd.Flags().Lookup("progress-bar"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		var task *rest.Task
		var err error
		switch {
		case cmd.Flags().Changed("id"):
			task, err = restClient.GetTask(viper.GetString("id"))
		case cmd.Flags().Changed("name"):
			task, err = restClient.GetTaskByName(viper.GetString("name"))
		default:
			cmd.Usage()
			os.Exit(1)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newTask := task.WaitForTask(restClient, viper.GetBool("progress"))
		if newTask.State == "completed" {
			fmt.Println(formatString("Task Complete"))
		}
		if newTask.State == "failed" {
			fmt.Println(formatString("Task Failed: " + task.Message))
			os.Exit(1)
		}
	},
}

func init() {
	taskCmd.AddCommand(taskWaitCmd)
	taskWaitCmd.Flags().StringP("id", "i", "", "task id")
	taskWaitCmd.Flags().StringP("name", "n", "", "task name")
	taskWaitCmd.Flags().Bool("progress", false, "print progress")
	taskWaitCmd.Flags().Bool("progress-bar", false, "print progress-bar")
}