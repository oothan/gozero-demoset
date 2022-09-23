// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Stock  uint64 `json:"stock"`
	Amount uint64 `json:"amount"`
	Status uint64 `json:"status"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
}

type UpdateRequest struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name,optional"`
	Desc   string `json:"desc,optional"`
	Stock  uint64 `json:"stock"`
	Amount uint64 `json:"amount,optional"`
	Status uint   `json:"status,optional"`
}

type UpdateResponse struct {
}

type RemoveRequest struct {
	Id uint64 `json:"id"`
}

type RemoveResponse struct {
}

type DetailRequest struct {
	Id uint64 `json:"id"`
}

type DetailResponse struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Stock  uint64 `json:"stock"`
	Amount uint64 `json:"amount"`
	Status uint64 `json:"status"`
}

type ListRequest struct {
	Page int64  `json:"page"`
	Size int64  `json:"size"`
	Name string `json:"name"`
}

type ListResponse struct {
	List []DetailResponse `json:"list"`
}
