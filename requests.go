package vibezy

// HTTP POST JSON request bodies are stored in this file

type DeactivateUserRequest struct {
	Email string `json:"email"`
}
