package users

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/olegvn88/gostudy/app/main/app/tests/mock/mocks"
	"github.com/olegvn88/gostudy/app/main/app/tests/mock/users"
	"testing"
)

func doSomeWork(u users.UserInterface) {
	u.SetName("Oleg Nesterov")
	name := u.GetName()
	fmt.Println(name)
}

func TestDoSomethingWithUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := mocks.NewMockUserInterface(ctrl)

	testUser.EXPECT().SetName("Oleg Nesterov")
	testUser.EXPECT().GetName().Return("Oleg Nesterov")

	doSomeWork(testUser)
}
