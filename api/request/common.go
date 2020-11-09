package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id uint64 `json:"id,string"`
}

type IdsReq struct {
	Ids []uint64 `json:"ids,string"`
}