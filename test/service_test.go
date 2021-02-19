package test

import (
	"strings"
	"testing"

	repositorymemory "bands-api/repository/memory"
	"bands-api/user"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load("../.env.test")
var repo, err = repositorymemory.NewMemoryRepository()
var service = user.NewUserService(repo)

func Test_ShouldGenerateUserID(t *testing.T) {
	
	user := user.User{}
	user.Name = "Paolo"
	user.Password = "123456"
	user.Email = "paolo@paolo.com"

	service.Register(&user)
	if user.ID == "" {
		t.Fail()
	}
}

func Test_ShouldFailUserValidation(t *testing.T) {
	user := user.User{}
	user.Name = "Paolo"
	user.Password = "123456"
	user.Email = "paolo@paolo.com"

	err = service.Register(&user)

	if err == nil {
		t.Fail()
	}
}

var loginToken string = ""
func Test_ShouldPerformLogin(t *testing.T) {
	token, err := service.Login("paolo@paolo.com", "123456")
	loginToken = token
	if token == "" || err != nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldFailLogin(t *testing.T) {
	token, err := service.Login("paolo@paolo.com", "12345asd")
	if token != "" || err == nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldReceiveTokenError(t *testing.T) {
	_, err = service.CheckLoginWithToken(strings.Replace(loginToken, "a", "b", 1))
	if err == nil {
		t.Error("should have expired token")
	}
}

func Test_ShouldValidateTokenOk(t *testing.T) {
	token, err := service.CheckLoginWithToken(loginToken)
	if err != nil {
		t.Error(err)
	}
	if token == nil {
		t.Error("Should have returned valid token")
	}
}