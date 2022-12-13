package service

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// UserServiceInterface 用户业务接口
type UserServiceInterface interface {
	GetUserSessionContext(userId string) string
	SetUserSessionContext(userId string, question, reply string)
	ClearUserSessionContext(userId string)
}

var _ UserServiceInterface = (*UserService)(nil)

// UserService 用戶业务
type UserService struct {
	// 缓存
	cache *cache.Cache
}

// NewUserService 创建新的业务层
func NewUserService() UserServiceInterface {
	return &UserService{cache: cache.New(time.Minute*5, time.Minute*10)}
}

// GetUserSessionContext 获取用户会话上下文文本
func (s *UserService) GetUserSessionContext(userId string) string {
	sessionContext, ok := s.cache.Get(userId)
	if !ok {
		return ""
	}
	return sessionContext.(string)
}

// SetUserSessionContext 设置用户会话上下文文本，question用户提问内容，GTP回复内容
func (s *UserService) SetUserSessionContext(userId string, question, reply string) {
	value := question + "\n" + reply
	s.cache.Set(userId, value, time.Minute*5)
}

// ClearUserSessionContext 清理用户会话上下文文本，用于超时时清理
func (s *UserService) ClearUserSessionContext(userId string) {
	s.cache.Delete(userId)
}
