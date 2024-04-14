package wecom

import (
	"os"
	"strconv"

	"github.com/wenerme/go-wecom/wecom"
)

type WC struct {
	client *wecom.Client
}

// NewWC 创建企业微信客户端
func NewWC(CorpID, AgentID, CorpSecret string) *WC {
	if CorpID == "" {
		CorpID = os.Getenv("CORP_ID")
	}
	if AgentID == "" {
		AgentID = os.Getenv("AGENT_ID")
	}
	if CorpSecret == "" {
		CorpSecret = os.Getenv("CORP_SECRET")
	}

	AgentIDI, _ := strconv.Atoi(AgentID)
	// token store - 默认内存 Map - 可以使用数据库实现
	store := &wecom.SyncMapStore{}
	// 加载缓存 - 复用之前的 Token
	if bytes, err := os.ReadFile("wecom-cache.json"); err == nil {
		_ = store.Restore(bytes)
	}
	// 当 Token 变化时生成缓存文件
	store.OnChange = func(s *wecom.SyncMapStore) {
		_ = os.WriteFile("wecom-cache.json", s.Dump(), 0o600)
	}

	client := wecom.NewClient(wecom.Conf{
		CorpID:     CorpID,
		AgentID:    AgentIDI,
		CorpSecret: CorpSecret,
		// 不配置默认使用 内存缓存
		TokenProvider: &wecom.TokenCache{
			Store: store,
		},
	})

	return &WC{client: client}
}

// GetDepts 获取部门列表
func (w *WC) GetDepts() ([]wecom.ListDepartmentResponseItem, error) {
	depts, err := w.client.ListDepartment(
		&wecom.ListDepartmentRequest{},
	)
	if err != nil {
		return nil, err
	}
	return depts.Department, nil
}

// GetUsers 获取用户列表
func (w *WC) GetUsers() ([]wecom.ListUserResponseItem, error) {
	depts, err := w.GetDepts()
	if err != nil {
		return nil, err
	}
	var us []wecom.ListUserResponseItem
	for _, dept := range depts {
		users, err := w.client.ListUser(
			&wecom.ListUserRequest{
				DepartmentID: strconv.Itoa(dept.ID),
				FetchChild:   "1",
			},
		)
		if err != nil {
			return nil, err
		}
		us = append(us, users.UserList...)
	}
	return us, nil
}
