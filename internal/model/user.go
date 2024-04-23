package model

type User struct {
	ID   uint32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
