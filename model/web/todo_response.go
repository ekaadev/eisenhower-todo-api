package web

type TodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	IsDone      bool   `json:"is_done"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"update_at"`
}
