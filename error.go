package structs

// Error is the model for returning errors in the api
type Error struct {
	Message string `json:"message"`
}
