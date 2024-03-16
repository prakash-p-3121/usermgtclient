package usermgtclient

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/usermodel"
)

type UserMgClient interface {
	UserCreate(req *usermodel.UserCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError)
	UserFindByID(userID string) (*usermodel.User, errorlib.AppError)
}
