package web

// TodoPatchRequest struct
// Use for request body when update some todo
type TodoPatchRequest struct {
	Id          int    `validate:"required" json:"id"`
	Title       string `validate:"max=200" json:"title"`
	Description string `json:"description"`
	Type        string `validate:"omitempty,oneof=urgent_important not_urgent_important urgent_not_important not_urgent_not_important" json:"type"`
	IsDone      bool   `validate:"boolean" json:"is_done"`
}
