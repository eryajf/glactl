package dingding

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eryajf/glactl/public/logger"
	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
)

var GetGroupsCmd = &cobra.Command{
	Use:   "getgroups",
	Short: "获取钉钉部门分组列表",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		secret, _ := cmd.Flags().GetString("secret")
		dc := NewDC(key, secret)
		depts, err := dc.GetDepts()
		if err != nil {
			logger.Error("获取部门列表失败", err)
		}

		table, err := gotable.Create("ID", "Name", "ParentId", "Remark")
		if err != nil {
			log.Fatal("创建表格失败: ", err)
		}

		for _, dept := range depts {
			_ = table.AddRow([]string{
				strconv.Itoa(dept.Id),
				dept.Name,
				strconv.Itoa(dept.ParentId),
				dept.Remark,
			})
		}
		fmt.Println(table)
	},
}

var GetUsersCmd = &cobra.Command{
	Use:   "getusers",
	Short: "获取钉钉用户列表",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		secret, _ := cmd.Flags().GetString("secret")
		dc := NewDC(key, secret)
		users, err := dc.GetUsers()
		if err != nil {
			logger.Error("获取用户列表失败", err)
		}

		table, err := gotable.Create("UserId", "UnionId", "Name", "Avatar", "StateCode", "ManagerUserId", "Mobile", "HideMobile", "Telephone", "JobNumber", "Title", "WorkPlace", "Remark", "LoginId", "DeptIds", "DeptOrder", "Extension", "HiredDate", "Active", "Admin", "Boss", "ExclusiveAccount", "Leader", "ExclusiveAccountType", "OrgEmail", "Email")
		if err != nil {
			log.Fatal("创建表格失败: ", err)
		}

		for _, user := range users {
			_ = table.AddRow([]string{
				user.UserId,
				user.UnionId,
				user.Name,
				user.Avatar,
				user.StateCode,
				user.ManagerUserId,
				user.Mobile,
				fmt.Sprintf("%t", user.HideMobile),
				user.Telephone,
				user.JobNumber,
				user.Title,
				user.WorkPlace,
				user.Remark,
				user.LoginId,
				fmt.Sprint(user.DeptIds),
				fmt.Sprint(user.DeptOrder),
				user.Extension,
				fmt.Sprint(user.HiredDate),
				strconv.FormatBool(user.Active),
				strconv.FormatBool(user.Admin),
				strconv.FormatBool(user.Boss),
				strconv.FormatBool(user.ExclusiveAccount),
				strconv.FormatBool(user.Leader),
				user.ExclusiveAccountType,
				user.OrgEmail,
				user.Email,
			})
		}
		fmt.Println(table)
	},
}
