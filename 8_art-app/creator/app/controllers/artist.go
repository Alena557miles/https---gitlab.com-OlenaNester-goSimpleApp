package controllers

import (
	"creator/app/models"
)

type ArtistController struct {
	artists []*models.Artist
}

func (ac *ArtistController) CreateArtist(a *models.Artist) {
	ac.artists = append(ac.artists, a)
}
func (ac *ArtistController) FindArtist(name string) *models.Artist {
	for _, artist := range ac.artists {
		if artist.Name == name {
			return artist
		}
	}
	return nil
}
