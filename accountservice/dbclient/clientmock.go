package dbclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/yuditan/goblog/accountservice/model"
)

type MockBoltClient struct {
	mock.Mock
}


func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
	args := m.Mock.Called(accountId)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDB() {

}

func (m *MockBoltClient) Seed() {

}