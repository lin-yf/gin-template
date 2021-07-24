package post

import "mime/multipart"

type UploadFiles struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// MarkdownPost 上传的格式
type MarkdownPost struct {
	Author       string   `binding:"required" json:"author" form:"author"`
	PlainContent string   `binding:"required"`
	Content      string   `binding:"required"`
	Toc          string   `binding:"required"`
	Markdown     string   `binding:"required"`
	Title        string   `binding:"required,max=72"`
	Slug         string   `binding:"required,max=72"`
	Desc         string   `binding:"required"`
	Excerpt      string   `binding:"required"`
	IsNote       bool     `form:"is_note,default=false"`
	IsDraft      bool     `form:"is_draft,default=false"`
	SeriesName   string   `form:"series_name"`
	Categories   []string `binding:"required,omitempty"`
	Tags         []string `binding:"required,omitempty"`
	Date         string   `binding:"omitempty"`
}

// SearchParam 搜索参数
type SearchParam struct {
	Query string `binding:"-" json:"query" form:"q"`
	Page  string `binding:"-" json:"page" form:"p"`
}

type PostForm struct {
	ID       int64    `binding:"omitempty" json:"pid"`
	Do       string   `binding:"required" json:"do"`
	Title    string   `binding:"required,max=72" json:"title"`
	Slug     string   `binding:"required,max=72" json:"slug"`
	SeriesID string   `binding:"required" json:"sid"`
	Cid      []string `binding:"required,omitempty" json:"cidList"`
	IsNote   bool     `form:"is_note,default=false"`
	Tags     []string `binding:"required,omitempty" json:"tags"`
	IsTop    string   `binding:"required" json:"isTop"`
	Content  string   `binding:"required" json:"content"`
	Toc      string   `binding:"required" json:"toc"`
	Markdown string   `binding:"required" json:"markdown"`
	Desc     string   `binding:"required" json:"desc"`
	Excerpt  string   `binding:"required" json:"excerpt"`
	IsDraft  bool     `form:"is_draft,default=false" json:"isDraft"`
	Date     string   `binding:"omitempty" json:"date"`
}
