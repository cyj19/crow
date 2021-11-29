/**
 * @Author: cyj19
 * @Date: 2021/11/23 10:27
 */

package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

// Execute 创建自定义命令
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
}
