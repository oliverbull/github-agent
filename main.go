package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v70/github"
	"github.com/joho/godotenv"
)

func main() {

	// pull in the env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file:", err)
	}

	// pull the access token
	githubToken, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		log.Fatalln("no github token")
	}

	// setup the github client
	client := github.NewClient(nil).WithAuthToken(githubToken)
	if client == nil {
		log.Fatalln("error creating github client with token")
	}

	// list my repos
	optsUser := &github.RepositoryListByUserOptions{Type: "public"}
	repos, resp, err := client.Repositories.ListByUser(context.Background(), "oliverbull", optsUser)
	if err != nil {
		log.Fatalln("error Repositories.ListByUser():", err)
	}
	log.Println("gh resp: " + resp.Status)

	log.Println("public repos:")
	for _, repo := range repos {
		log.Println(*repo.Name)
	}

	log.Println("first repo contents:")
	//log.Println(*repos[0])

	// list all repos
	optsList := &github.RepositoryListAllOptions{Since: 905176388}
	repos, resp, err = client.Repositories.ListAll(context.Background(), optsList)
	if err != nil {
		log.Fatalln("error Repositories.ListAll():", err)
	}
	log.Println("gh resp: " + resp.Status)

	log.Println("num repos: ", len(repos))
	for idx, repo := range repos {
		log.Println(*repo.Name)
		log.Println(*repo.HTMLURL)
		log.Println(*repo.ID)
		log.Println(*repo.Private)
		//log.Println(repo.UpdatedAt.String())
		if idx > 5 {
			break
		}
	}
	//log.Println(*repos[2])

	// get the full details
	repo, _, err := client.Repositories.GetByID(context.Background(), *repos[2].ID)
	if err != nil {
		log.Fatalln("error Repositories.GetByID():", err)
	}
	log.Println(*repo.PushedAt)

	// list latest repos by creation
	// client.Repositories.Tr

}
