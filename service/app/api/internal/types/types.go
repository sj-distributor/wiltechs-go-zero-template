// Code generated by goctl. DO NOT EDIT.
package types

type GetUserReq struct {
	Id string `path:"id"`
}

type GetUserReply struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}