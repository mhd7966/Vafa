package inputs

type GetUsersQuery struct {
	Status int `query:"status" default:"1"`
	Page   int `query:"page" default:"1"`
	Size   int `query:"size" default:"10"`
}
