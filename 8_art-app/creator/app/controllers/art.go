package controllers

import (
	"creator/app/models"
	"fmt"
)

type ArtController struct {
	arts []*models.Art
}

func (ac *ArtController) CreateArt(a *models.Art) {
	ac.arts = append(ac.arts, a)
}

func (ac *ArtController) FindArt(name string) *models.Art {
	for _, art := range ac.arts {
		if art.Name == name {
			// fmt.Println("Art Name : ", art.Name)
			return art
		}
	}
	return nil
}

func (ac *ArtController) AssignedArtToArtist(art *models.Art, artist *models.Artist) *models.Art {
	if art.IsntAssigned() {
		art.Owner = artist
		artist.Arts = append(artist.Arts, art)
		return art
	} else {
		fmt.Println("This art already has an owner! You'd better make your own art")
	}
	return nil
}
