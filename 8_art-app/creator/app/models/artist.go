package models

type Artist struct {
	Id        int
	Name      string `json:"name"`
	Arts      []*Art `json:"art"`
	OnGallery bool   `json:"ongallery"`
}
