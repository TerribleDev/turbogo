package main

// Team is a Vercel Team object
type Team struct {
	ID        string `json:"id,omitempty"`
	Slug      string `json:"slug,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int    `json:"createdAt,omitempty"`
	Created   string `json:"created,omitempty"`
}

type Pagination struct {
	Count int `json:"count,omitempty"`
	Next  int `json:"next,omitempty"`
	Prev  int `json:"prev,omitempty"`
}

type TeamsResponse struct {
	Teams      []Team     `json:"teams,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int    `json:"createdAt,omitempty"`
}
type UserResponse struct {
	User User `json:"user,omitempty"`
}
