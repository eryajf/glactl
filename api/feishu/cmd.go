package feishu

import (
	"fmt"
	"log"

	"github.com/eryajf/glactl/public/logger"
	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
)

var GetGroupsCmd = &cobra.Command{
	Use:   "getgroups",
	Short: "获取飞书部门分组列表",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		secret, _ := cmd.Flags().GetString("secret")
		fc := NewFC(key, secret)
		depts, err := fc.GetDepts()
		if err != nil {
			logger.Error("获取部门列表失败", err)
		}

		table, err := gotable.Create("Name", "I18nName", "ParentDepartmentID", "DepartmentID", "OpenDepartmentID", "LeaderUserID", "ChatID", "Order", "UnitIDs", "MemberCount", "Status", "CreateGroupChat", "Leaders", "GroupChatEmployeeTypes", "DepartmentHrbps", "PrimaryMemberCount")
		if err != nil {
			log.Fatal("创建表格失败: ", err)
		}

		for _, dept := range depts {
			_ = table.AddRow([]string{
				dept.Name,
				dept.I18nName.ZhCn,
				dept.ParentDepartmentID,
				dept.DepartmentID,
				dept.OpenDepartmentID,
				dept.LeaderUserID,
				dept.ChatID,
				dept.Order,
				fmt.Sprintf("%v", dept.UnitIDs),
				fmt.Sprint(dept.MemberCount),
				fmt.Sprint(dept.Status),
				fmt.Sprintf("%v", dept.CreateGroupChat),
				fmt.Sprintf("%v", dept.Leaders),
				fmt.Sprintf("%v", dept.GroupChatEmployeeTypes),
				fmt.Sprintf("%v", dept.DepartmentHrbps),
				fmt.Sprintf("%v", dept.PrimaryMemberCount),
			})
		}
		fmt.Println(table)
	},
}

var GetUsersCmd = &cobra.Command{
	Use:   "getusers",
	Short: "获取飞书用户列表",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		secret, _ := cmd.Flags().GetString("secret")
		fc := NewFC(key, secret)

		users, err := fc.GetUsers()
		if err != nil {
			logger.Error("获取用户列表失败", err)
		}

		table, err := gotable.Create("UnionID", "UserID", "OpenID", "Name", "EnName", "Nickname", "Email", "Mobile", "MobileVisible", "Gender", "AvatarKey", "Avatar", "Status", "DepartmentIDs", "LeaderUserID", "City", "Country", "WorkStation", "JoinTime", "IsTenantManager", "EmployeeNo", "EmployeeType", "Orders", "CustomAttrs", "EnterpriseEmail", "JobTitle", "IsFrozen", "Geo", "JobLevelID", "JobFamilyID", "DepartmentPath", "DottedLineLeaderUserIDs")
		if err != nil {
			log.Fatal("创建表格失败: ", err)
		}

		for _, user := range users {
			_ = table.AddRow([]string{
				user.UnionID,
				user.UserID,
				user.OpenID,
				user.Name,
				user.EnName,
				user.Nickname,
				user.Email,
				user.Mobile,
				fmt.Sprintf("%v", user.MobileVisible),
				fmt.Sprint(user.Gender),
				user.AvatarKey,
				fmt.Sprintf("%v", user.Avatar),
				fmt.Sprint(user.Status),
				fmt.Sprint(user.DepartmentIDs),
				user.LeaderUserID,
				user.City,
				user.Country,
				user.WorkStation,
				fmt.Sprintf("%v", user.JoinTime),
				fmt.Sprintf("%v", user.IsTenantManager),
				user.EmployeeNo,
				fmt.Sprintf("%v", user.EmployeeType),
				fmt.Sprint(user.Orders),
				fmt.Sprintf("%v", user.CustomAttrs),
				user.EnterpriseEmail,
				user.JobTitle,
				fmt.Sprintf("%v", user.IsFrozen),
				user.Geo,
				user.JobLevelID,
				user.JobFamilyID,
				fmt.Sprintf("%v", user.DepartmentPath),
				fmt.Sprintf("%v", user.DottedLineLeaderUserIDs),
			})
		}
		fmt.Println(table)
	},
}
