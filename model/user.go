package model

type User struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"firstname,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	LoginType bool   `json:"loginType,omitempty"`
}
