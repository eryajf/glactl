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
	"github.com/eryajf/glactl/api/feishu"
	"github.com/spf13/cobra"
)

// feishuCmd represents the jenkins command
var feishuCmd = &cobra.Command{
	Use:   "feishu",
	Short: `用于测试钉钉的数据获取是否正常`,
}

func init() {
	rootCmd.AddCommand(feishuCmd)
	// 获取分组列表
	feishuCmd.AddCommand(feishu.GetGroupsCmd)
	feishu.GetGroupsCmd.Flags().StringP("key", "k", "", "指定AppKey")
	feishu.GetGroupsCmd.Flags().StringP("secret", "s", "", "指定AppSecret")

	// 获取用户列表
	feishuCmd.AddCommand(feishu.GetUsersCmd)
	feishu.GetUsersCmd.Flags().StringP("key", "k", "", "指定AppKey")
	feishu.GetUsersCmd.Flags().StringP("secret", "s", "", "指定AppSecret")
}
