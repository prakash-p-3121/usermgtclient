package usermgtclient

import "github.com/prakash-p-3121/usermgtclient/impl"

func NewUserMgtClient(host string, port uint) UserMgClient {
	return &impl.UserMgtClientImpl{Host: host, Port: port}
}
