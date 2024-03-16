package impl

import (
	"fmt"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/restclientlib"
	"github.com/prakash-p-3121/usermodel"
	"log"
	"net/url"
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
	baseUrl := hostPort + userFind
	log.Println("url =", baseUrl)

	params := url.Values{}
	params.Add("user-id", userID)
	encodedParams := params.Encode()
	log.Println("encodedParams=" + encodedParams)
	finalUrl := baseUrl + "?" + encodedParams
	log.Println("finalUrl=" + finalUrl)

	var resp usermodel.User
	appErr := restClient.Get(finalUrl, &resp)
	return &resp, appErr
}
