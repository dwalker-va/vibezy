package vibezy

// HTTP JSON response bodies are stored in this file

type PingResponse struct {
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
}

type ListUsersResponse struct {
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
	Data         struct {
		UserCount int `json:"userCount"`
		Users     []struct {
			Email string `json:"email"`
			// Yes this API seems to have a bug where it will return "fistName" JSON keys, not a mistake
			FistName string `json:"fistName"`
			// Hopefully that bug is fixed and we can use this
			FirstName        string   `json:"firstName"`
			LastName         string   `json:"lastName"`
			UserName         string   `json:"userName"`
			JobTitle         string   `json:"jobTitle"`
			IsGroupManager   bool     `json:"isGroupManager"`
			IsCompanyManager bool     `json:"isCompanyManager"`
			IsAdmin          bool     `json:"isAdmin"`
			ManagedGroups    []string `json:"managedGroups"`
			MemberGroups     []string `json:"memberGroups"`
		} `json:"users"`
	} `json:"data"`
}

type GetUserResponse struct {
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
	Data         struct {
		Email            string   `json:"email"`
		FistName         string   `json:"fistName"`
		LastName         string   `json:"lastName"`
		UserName         string   `json:"userName"`
		JobTitle         string   `json:"jobTitle"`
		IsGroupManager   bool     `json:"isGroupManager"`
		IsCompanyManager bool     `json:"isCompanyManager"`
		IsAdmin          bool     `json:"isAdmin"`
		ManagedGroups    []string `json:"managedGroups"`
		MemberGroups     []string `json:"memberGroups"`
	} `json:"data"`
}

type UpdateUserResponse struct {
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
}

type DeactivateUserResponse struct {
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
}

type ListGroupsResponse struct {
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
	Data         struct {
		GroupCount int `json:"groupCount"`
		Groups     []struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			UserCount    int    `json:"userCount"`
			ManagerCount int    `json:"managerCount"`
		} `json:"groups"`
	} `json:"data"`
}

type GetGroupResponse struct {
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
	Data         struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		UserCount    int    `json:"userCount"`
		ManagerCount int    `json:"managerCount"`
		Users        []struct {
			Email            string   `json:"email"`
			FistName         string   `json:"fistName"`
			LastName         string   `json:"lastName"`
			UserName         string   `json:"userName"`
			JobTitle         string   `json:"jobTitle"`
			IsGroupManager   bool     `json:"isGroupManager"`
			IsCompanyManager bool     `json:"isCompanyManager"`
			IsAdmin          bool     `json:"isAdmin"`
			ManagedGroups    []string `json:"managedGroups"`
			MemberGroups     []string `json:"memberGroups"`
		} `json:"users"`
	} `json:"data"`
}
