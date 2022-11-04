package controllers

import (
	"creator/app/models"
	"fmt"
)

type GalleryController struct {
	Galleries []*models.Gallery
}

func (gc *GalleryController) CreateGallery(g *models.Gallery) {
	gc.Galleries = append(gc.Galleries, g)
}
func (gc *GalleryController) FindGallery(name string) *models.Gallery {
	for _, g := range gc.Galleries {
		if name == g.Name {
			return g
		}
	}
	return nil
}

func (gc *GalleryController) RegisterArtist(gallery *models.Gallery, artist *models.Artist) {
	if len(artist.Arts) > 0 {
		gc.Galleries = append(gc.Galleries, gallery)
		gallery.Artists = append(gallery.Artists, artist)
		artist.OnGallery = true
	}
	if len(artist.Arts) == 0 {
		fmt.Println("We can not register an Artist without Arts")
	}
	if artist.OnGallery {
		fmt.Println("Artist are already on gallery")
	}
}
func (gc *GalleryController) DeleteArtist(gallery *models.Gallery, artist *models.Artist) {
	for _, g := range gc.Galleries {
		if g == gallery {
			gallery.DeleteArtist(artist)
		}
	}
}
