package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}
func main() {

	c := http.Client{Timeout: 10 * time.Second}
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	//AUTHORIZATION
	req, err := http.NewRequest(`GET`, `https://api.github.com/user`, nil)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	req.Header.Add(`Accept`, `application/vnd.github+json`)
	req.Header.Add(`Authorization`, `Bearer `+token)

	//GET ALL PUBLIC REPO
	//req, err := http.NewRequest(`GET`, `https://api.github.com/users/Alena557miles/repos`, nil)
	//if err != nil {
	//	fmt.Printf("Error: %s\\n", err)
	//	return
	//}

	//  GET PRIVATE REPO
	//req, err := http.NewRequest(`GET`, `https://api.github.com/repos/Alena557miles/GO_simple_app`, nil)
	//if err != nil {
	//	fmt.Printf("Error: %s\\n", err)
	//	return
	//}
	//req.Header.Add(`Accept`, `application/vnd.github+json`)
	//req.Header.Add(`Authorization`, `Bearer `+token)

	//CREATE A REPO FOR THE AUTH USER
	//newRepo := bytes.NewBuffer([]byte(`{
	//	"name":"Hello-World",
	//	"description":"This is your first repo!",
	//	"homepage":"https://github.com",
	//	"private":false,
	//	"is_template":true
	//}`))
	//
	//req, err := http.NewRequest(`POST`, `https://api.github.com/user/repos`, newRepo)
	//if err != nil {
	//	fmt.Printf("Error: %s\\n", err)
	//	return
	//}
	//req.Header.Add(`Accept`, `application/vnd.github+json`)
	//req.Header.Add(`Authorization`, `Bearer `+token)

	// DELETE REPO
	//req, err := http.NewRequest(`DELETE`, `https://api.github.com/repos/Alena557miles/Hello-World`, nil)
	//if err != nil {
	//	fmt.Printf("Error: %s\\n", err)
	//	return
	//}
	//req.Header.Add(`Accept`, `application/vnd.github+json`)
	//req.Header.Add(`Authorization`, `Bearer `+token)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body: %s\n", body)

}
