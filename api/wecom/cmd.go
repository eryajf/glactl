package wecom

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
	Short: "获取企微部门分组列表",
	Run: func(cmd *cobra.Command, args []string) {
		corpid, _ := cmd.Flags().GetString("corpid")
		agentid, _ := cmd.Flags().GetString("agentid")
		secret, _ := cmd.Flags().GetString("secret")
		wc := NewWC(corpid, agentid, secret)
		depts, err := wc.GetDepts()
		if err != nil {
			logger.Error("获取部门列表失败", err)
		}

		table, err := gotable.Create("ID", "Name", "ParentId")
		if err != nil {
			log.Fatal("创建表格失败: ", err)
		}

		for _, dept := range depts {
			_ = table.AddRow([]string{
				strconv.Itoa(dept.ID),
				dept.Name,
				strconv.Itoa(dept.ParentID),
			})
		}
		fmt.Println(table)
	},
}

var GetUsersCmd = &cobra.Command{
	Use:   "getusers",
	Short: "获取企微用户列表",
	Run: func(cmd *cobra.Command, args []string) {
		corpid, _ := cmd.Flags().GetString("corpid")
		agentid, _ := cmd.Flags().GetString("agentid")
		secret, _ := cmd.Flags().GetString("secret")
		wc := NewWC(corpid, agentid, secret)

		users, err := wc.GetUsers()
		if err != nil {
			logger.Error("获取用户列表失败", err)
		}

		table, err := gotable.Create("UserID", "Name", "Mobile", "Department", "Order", "Position", "Gender", "Email", "BizMail", "IsLeaderInDept", "Avatar", "ThumbAvatar", "Telephone", "Alias", "ExtAttr", "Status", "QrCode", "ExternalProfile", "ExternalPosition", "Address", "OpenUserID", "MainDepartment", "EnglishName", "HideMobile")
		if err != nil {
			log.Fatal("创建表格失败: ", err)
		}

		for _, user := range users {
			_ = table.AddRow([]string{
				user.UserID,
				user.Name,
				user.Mobile,
				fmt.Sprintf("%v", user.Department),
				fmt.Sprintf("%v", user.Order),
				user.Position,
				user.Gender,
				user.Email,
				user.BizMail,
				fmt.Sprintf("%v", user.IsLeaderInDept),
				user.Avatar,
				user.ThumbAvatar,
				user.Telephone,
				user.Alias,
				fmt.Sprintf("%v", user.ExtAttr),
				strconv.Itoa(user.Status),
				user.QrCode,
				fmt.Sprintf("%v", user.ExternalProfile),
				user.ExternalPosition,
				user.Address,
				user.OpenUserID,
				strconv.Itoa(user.MainDepartment),
				user.EnglishName,
				strconv.Itoa(user.HideMobile),
			})
		}
		fmt.Println(table)
	},
}
