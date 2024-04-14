package feishu

import (
	"context"
	"fmt"
	"os"

	"github.com/chyroc/lark"
)

type FC struct {
	client *lark.Lark
}

// NewFC 创建飞书客户端
func NewFC(key, secret string) *FC {
	if key == "" {
		key = os.Getenv("AppKey")
	}
	if secret == "" {
		secret = os.Getenv("AppSecret")
	}

	return &FC{client: lark.New(lark.WithAppCredential(key, secret))}
}

// GetDepts 获取部门列表
func (f *FC) GetDepts() (depts []*lark.GetDepartmentListRespItem, err error) {
	var (
		fetchChild bool  = true
		pageSize   int64 = 50
	)

	req := lark.GetDepartmentListReq{
		FetchChild:   &fetchChild,
		PageSize:     &pageSize,
		DepartmentID: "0",
	}

	for {
		res, _, err := f.client.Contact.GetDepartmentList(context.TODO(), &req)
		if err != nil {
			fmt.Printf("GetDepartmentList error: %v\n", err)
		}
		depts = append(depts, res.Items...)
		if !res.HasMore {
			break
		}
		req.PageToken = &res.PageToken
	}
	return
}

func (f *FC) GetUsers() (users []*lark.GetUserListRespItem, err error) {
	var (
		pageSize int64 = 50
	)
	depts, err := f.GetDepts()
	if err != nil {
		fmt.Printf(" get all depts failed, err:%v\n", err)
	}
	for _, dept := range depts {

		req := lark.GetUserListReq{
			PageSize:     &pageSize,
			PageToken:    new(string),
			DepartmentID: dept.OpenDepartmentID,
		}

		for {
			res, _, err := f.client.Contact.GetUserList(context.Background(), &req)
			if err != nil {
				return nil, err
			}
			users = append(users, res.Items...)
			if !res.HasMore {
				break
			}
			req.PageToken = &res.PageToken
		}
	}
	return
}

// 官方文档： https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/employee/list
// GetLeaveUserIds 获取离职人员ID列表
func (f *FC) GetLeaveUserIds() ([]string, error) {
	var ids []string
	users, _, err := f.client.EHR.GetEHREmployeeList(context.TODO(), &lark.GetEHREmployeeListReq{
		Status:     []int64{5},
		UserIDType: lark.IDTypePtr(lark.IDTypeUnionID), // 只查询unionID
	})
	if err != nil {
		return nil, err
	}
	for _, user := range users.Items {
		ids = append(ids, user.UserID)
	}
	return ids, nil
}
