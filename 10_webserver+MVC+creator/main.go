package main

import (
	"creator/app/controllers"
	"creator/app/sl"
	"fmt"

	"github.com/Alena557miles/gowebserver"
)

func main() {
	artistC := &controllers.ArtistController{}
	artC := &controllers.ArtController{}
	galleryC := &controllers.GalleryController{}
	sl := &sl.ServiceLocator{}
	sl.Register(artC)
	sl.Register(artistC)
	sl.Register(galleryC)

	gowebserver.StartServer(artistC)

	// switch command {
	// case `ra`:
	// 	// ARTIST REGISTRATION
	// 	artistC.Registration()
	// 	goto START

	// case `ca`:
	// 	// CREATE AN ART
	// 	artC.ArtCreation()
	// 	goto START

	// case `cg`:
	// 	// CREATE A GALLERY
	// 	galleryC.GalleryCreation()
	// 	goto START

	// case `aa`:
	// 	//ASSIGN AN ART TO THE ARTIST (BY NAME)
	// 	sl.Register(artistC)
	// 	artC.AssignArt(sl)
	// 	goto START

	// case `rg`:
	// 	//REGISTRATION AN ARTIST ON THE GALLERY
	// 	sl.Register(galleryC)
	// 	artistC.ArtistRegistration(sl)
	// 	goto START

	// case `da`:
	// 	// DELETE AN ARTIST FROM GALLERY
	// 	sl.Register(artistC)
	// 	galleryC.RemoveArtistFromGal(sl)
	// 	goto START
	// case `e`:
	// 	break
	// }
	fmt.Printf("as you wish...bye...")
}
