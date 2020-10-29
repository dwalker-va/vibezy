package vibezy

// HTTP POST JSON request bodies are stored in this file

type DeactivateUserRequest struct {
	Email string `json:"email"`
}

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type RemoveGroupRequest struct {
	Name string `json:"name"`
}

type AddUsersToGroupRequest struct {
	GroupID    string   `json:"groupId"`
	Emails     []string `json:"emails"`
	ToManagers bool     `json:"toManagers"`
	ToMembers  bool     `json:"toMembers"`
}

type RemoveUsersFromGroupRequest struct {
	GroupID      string   `json:"groupId"`
	Emails       []string `json:"emails"`
	FromMembers  bool     `json:"fromMembers"`
	FromManagers bool     `json:"fromManagers"`
}

type RemoveAllUsersFromGroupRequest struct {
	GroupID      string `json:"groupId"`
	FromMembers  bool   `json:"fromMembers"`
	FromManagers bool   `json:"fromManagers"`
}
