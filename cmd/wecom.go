/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/eryajf/glactl/api/wecom"
	"github.com/spf13/cobra"
)

// wecomCmd represents the jenkins command
var wecomCmd = &cobra.Command{
	Use:   "wecom",
	Short: `用于测试企微的数据获取是否正常`,
}

func init() {
	rootCmd.AddCommand(wecomCmd)
	// 获取分组列表
	wecomCmd.AddCommand(wecom.GetGroupsCmd)
	wecom.GetGroupsCmd.Flags().StringP("corpid", "c", "", "指定企业ID")
	wecom.GetGroupsCmd.Flags().StringP("agentid", "a", "", "指定应用ID")
	wecom.GetGroupsCmd.Flags().StringP("secret", "s", "", "指定企业Secret")

	// 获取用户列表
	wecomCmd.AddCommand(wecom.GetUsersCmd)
	wecom.GetUsersCmd.Flags().StringP("corpid", "c", "", "指定企业ID")
	wecom.GetUsersCmd.Flags().StringP("agentid", "a", "", "指定应用ID")
	wecom.GetUsersCmd.Flags().StringP("secret", "s", "", "指定企业Secret")
}
