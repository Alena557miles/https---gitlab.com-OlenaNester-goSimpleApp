package main

import (
	"creator/app/controllers"
	"creator/app/models"
	"creator/app/views"
	"fmt"
)

func main() {
	var id, idArt int
	id = 0

	artistC := &controllers.ArtistController{}
	artC := &controllers.ArtController{}
	galleryC := &controllers.GalleryController{}

START:
	fmt.Printf("APP 'CREATOR' CHOOSE WHAT TO DO:" +
		"\nRegister an Artist: ra" +
		"\nCreate an art: ca" +
		"\nAdd art to the Artist: aa" +
		"\nCreate a Gallery: cg" +
		"\nRegister an Artist in Gallery: rg" +
		"\nDelete an Artist from the Gallery: da" +
		"\nYou want to exit: e\n")

	var command string
	_, _ = fmt.Scan(&command)

	switch command {
	case `ra`:
		// ARTIST REGISTRATION
		fmt.Printf("Artist name:\n")
		var name string
		_, _ = fmt.Scan(&name)
		id++
		artist := &models.Artist{Id: id, Name: name, OnGallery: false}
		// create an Artist and print information about user
		artistC.CreateArtist(artist)
		if err := views.PrintArtist(artist); err != nil {
			fmt.Println(err)
		}
		fmt.Println("---Artist registrated successfuly---")
		fmt.Println("\n\n\n\n\n\n")
		goto START

	case `ca`:
		// CREATE AN ART
		fmt.Printf("Art name:\n")
		var name string
		_, _ = fmt.Scan(&name)
		fmt.Printf("Width, cm (only numbers): \n")
		var w int
		_, _ = fmt.Scan(&w)
		fmt.Printf("Height, cm (only numbers): \n")
		var h int
		_, _ = fmt.Scan(&h)
		idArt++
		art := &models.Art{Id: idArt, Name: name, Width: w, Height: h}
		// create an art and print information about art
		artC.CreateArt(art)
		if err := views.PrintArt(art); err != nil {
			fmt.Println(err)
		}
		fmt.Println("---________Art created successfuly________---")
		fmt.Println("\n\n\n\n\n\n")
		goto START

	case `cg`:
		// CREATE A GALLERY
		fmt.Printf("Gallery name:\n")
		var name string
		_, _ = fmt.Scan(&name)
		fmt.Printf("Gallery location (city): \n")
		var l string
		_, _ = fmt.Scan(&l)
		gallery := &models.Gallery{Name: name, Location: l}
		// create an gallery and print information about it
		galleryC.CreateGallery(gallery)
		if err := views.PrintGallery(gallery); err != nil {
			fmt.Println(err)
		}
		fmt.Println("---Gallery created successfuly---")
		fmt.Println("\n\n\n\n\n\n")
		goto START

	case `aa`:
		//ASSIGN AN ART TO THE ARTIST (BY NAME)
		//asking and searching an art by name
		fmt.Printf("Which art do you want to assign? (Print name of the art):\n")
		var name string
		_, _ = fmt.Scan(&name)
		art := artC.FindArt(name)
		if err := artC.FindArt(name); err != nil {
			//asking and searching an artist by name
			fmt.Printf("To which artist do you want to assign this art? (Print name of the artist):\n")
			var artistName string
			_, _ = fmt.Scan(&artistName)
			artist := artistC.FindArtist(artistName)
			if err := artistC.FindArtist(artistName); err != nil {
				//assign an art to the Artist and
				artC.AssignedArtToArtist(art, artist)
			}
		}

		// print information about art
		if err := views.PrintArt(art); err != nil {
			fmt.Println(err)
		}
		fmt.Println("---Art was assigned to Artist successfuly---")
		fmt.Println("\n\n\n\n\n\n")
		goto START

	case `rg`:
		//REGISTRATION AN ARTIST ON THE GALLERY
		//asking and searching an artist by name
		fmt.Printf("Which artist do you want to registrate on the gallery? (Print name of the artist):\n")
		var artistName string
		_, _ = fmt.Scan(&artistName)
		artist := artistC.FindArtist(artistName)
		if err := artistC.FindArtist(artistName); err != nil {
			fmt.Printf("To which Gallery ? (Print name of the Gallery):\n")
			var galleryName string
			_, _ = fmt.Scan(&galleryName)
			//searching gallery
			gallery := galleryC.FindGallery(galleryName)
			galleryC.RegisterArtist(gallery, artist)
			// Artist registration on the Gallery and print information about artist
			if err := views.PrintArtist(artist); err != nil {
				fmt.Println(err)
			}
			fmt.Println("---Artist was registrated on the Gallery successfuly---")
		}
		fmt.Println("\n\n\n\n\n\n")
		goto START

	case `da`:
		// DELETE AN ARTIST FROM GALLERY
		//finding the gallery
		fmt.Printf("From which Gallery ? (Print name of the Gallery):\n")
		var galleryName string
		_, _ = fmt.Scan(&galleryName)
		gallery := galleryC.FindGallery(galleryName)
		// finding the artist
		fmt.Printf("Which Artist do you want to delete from the Gallery? (Print name of the artist):\n")
		var artistName string
		_, _ = fmt.Scan(&artistName)
		artist := artistC.FindArtist(artistName)
		//delete
		galleryC.DeleteArtist(gallery, artist)
		fmt.Println("---Artist was removed from the Gallery successfuly---")
		fmt.Println("\n\n\n\n\n\n")
		goto START

	case `e`:
		break
	}
	fmt.Printf("as you wish...bye...")
}
