package model

type CategoryResponse struct {
	ID       uint               `json:"id"`
	ParentID uint               `json:"parent_id"`
	Title    string             `json:"title"`
	Slug     string             `json:"slug"`
	Child    []CategoryResponse `json:"child"`
}
