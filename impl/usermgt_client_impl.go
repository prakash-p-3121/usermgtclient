package impl

import (
	"fmt"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/restclientlib"
	"github.com/prakash-p-3121/usermodel"
	"log"
)

type UserMgtClientImpl struct {
	Host string
	Port uint
}

func (client *UserMgtClientImpl) HostPort() string {
	return fmt.Sprintf("%s:%d", client.Host, client.Port)
}

func (client *UserMgtClientImpl) UserCreate(req *usermodel.UserCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	restClient := restclientlib.NewRestClient()
	hostPort := client.HostPort()
	url := hostPort + userCreate
	log.Println("url =", url)
	var resp idgenmodel.IDGenResp
	appErr := restClient.Post(url, req, &resp)
	return &resp, appErr
}

func (client *UserMgtClientImpl) UserFindByID(userID string) (*usermodel.User, errorlib.AppError) {
	restClient := restclientlib.NewRestClient()
	hostPort := client.HostPort()
	url := hostPort + userFind
	log.Println("url =", url)
	var resp usermodel.User
	appErr := restClient.Get(url, &resp)
	return &resp, appErr
}
