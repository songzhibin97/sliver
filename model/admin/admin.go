package admin

type Admin struct {
	Id          uint64 `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	AuthorityId string `json:"authority_id"`
}

func NewAdmin() *Admin {
	return &Admin{}
}
