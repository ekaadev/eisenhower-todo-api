package web

type TodoPatchRequest struct {
	Id          int    `validate:"required, int" json:"id"`
	Title       string `validate:"required, min=1, max=200" json:"title"`
	Description string `validate:"required, min=1" json:"description"`
	Type        string `validate:"required, min=1, contains=urgent_important|not_urgent_important|urgent_not_important|not_urgent_not_important" json:"type"`
	IsDone      bool   `validate:"required, bool" json:"is_done"`
}
