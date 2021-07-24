package series

//
type SeriesForm struct {
	Sid   int    `binding:"-" form:"id" json:"id"`
	Name  string `binding:"required,min=2" json:"name" form:"name"`
	Slug  string `binding:"-" json:"slug" form:"slug"`
	Desc  string `binding:"-" form:"desc" json:"desc"`
	Order int    `binding:"-" form:"order,default=0" json:"order"`
}

type SeriesDelForm struct {
	Ids []int `binding:"required" form:"ids" json:"ids"`
}
