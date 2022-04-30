package inputs

type GetUsersQuery struct {
	Status   int `query:"status" default:"1"`
	Page     int `query:"page" default:"1"`
	PageSize int `query:"page_size" default:"10"`
}
