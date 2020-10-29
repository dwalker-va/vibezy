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

type SyncRequest struct {
	Settings struct {
		SyncManagerEmail string `json:"syncManagerEmail"`
		InviteNewUsers   bool   `json:"inviteNewUsers"`
	} `json:"settings"`
	Users []struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		JobTitle  string `json:"jobTitle"`
		ImageURL  string `json:"imageUrl"`
		Language  string `json:"language"`
		Gender    string `json:"Gender,omitempty"`
	} `json:"users"`
	Groups []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
	Mappings []struct {
		GroupID    string `json:"groupId"`
		UserID     string `json:"userId"`
		SubGroupID string `json:"subGroupId"`
		IsMember   bool   `json:"isMember"`
		IsManager  bool   `json:"isManager"`
	} `json:"mappings"`
}
