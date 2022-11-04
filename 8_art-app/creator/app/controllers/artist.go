package controllers

import (
	"creator/app/models"
	"fmt"
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
			// fmt.Println("Artist Name : ", artist.Name)
			return artist
		} else {
			fmt.Println("We don't know such artist")
			break
		}
	}
	return nil
}
