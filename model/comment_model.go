package model

type CommentRequest struct {
	PostID    string `json:"post_id"`
	ParentID  string `json:"parent_id"`
	Content   string `json:"content"`
	CommentBy uint64 `json:"comment_by"`
}

type CommentResponse struct {
	PostID    string `json:"post_id"`
	ParentID  string `json:"parent_id"`
	Content   string `json:"content"`
	CommentBy string `json:"comment_by"`
}

type ListCommentRequest struct {
	PostID string `json:"post_id"`
}

type ListCommentResponse struct {
	ID        uint                  `json:"id"`
	PostID    uint                  `json:"post_id"`
	ParentID  uint                  `json:"parent_id"`
	Content   string                `json:"content"`
	CommentBy uint                  `json:"comment_by"`
	Child     []ListCommentResponse `json:"child"`
}
