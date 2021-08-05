package model

type Supplier struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Coordinates string `json:"location"`
	Address     string `json:"address"`
	Type        string `json:"type"`
	IconURL     string `json:"iconUrl"`
	Description string `json:"description"`
}
