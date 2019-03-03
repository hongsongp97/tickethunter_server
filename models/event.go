package models

type Event struct {
	Id           uint   `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Des          string `json:"des,omitempty"`
	AvatarUrl    string `json:"avatarUrl,omitempty"`
	Location     string `json:"loginType,omitempty"`
	Category     string `json:"category,omitempty"`
	ListImageUrl string `json:"listImageUrl,omitempty"`
	IDUserowner  uint   `json:"idUserowner,omitempty"`
	EmailOwner   string `json:"emailOwner,omitempty"`
	PhoneOwner   string `json:"phoneowner,omitempty"`
	Time         string `json:"time,omitempty"`
	AmountSold   string `json:"amountSold,omitempty"`
	Price        string `json:"price,omitempty"`
}

type EventsByUserID struct {
	Owner     []Event `json:"owner,omitempty"`
	Following []Event `json:"following,omitempty"`
	Joinning  []Event `json:"joinning,omitempty"`
}
