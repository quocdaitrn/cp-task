package entity

// Filter condition to filter tasks.
type Filter struct {
	UserID *uint   `json:"user_id,omitempty"`
	Status *string `json:"status,omitempty"`
}
