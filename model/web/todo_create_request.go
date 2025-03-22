package web

// TodoCreateRequest struct
// Use for request body when create new todo
type TodoCreateRequest struct {
	Title       string `validate:"required,min=1,max=200" json:"title"`
	Description string `json:"description"`
	Type        string `validate:"required,min=1,oneof=urgent_important not_urgent_important urgent_not_important not_urgent_not_important" json:"type"`
}
