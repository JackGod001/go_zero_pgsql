// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1

package types

type GetUserRequest struct {
	Id int64 `form:"id"`
}

type GetUserResponse struct {
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}