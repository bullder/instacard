package main

import (
	"context"
	"encoding/json"
	"github.com/Davincible/goinsta/v3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

type Profile struct {
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Images []Img `json:"images"`
}

type Img struct {
	Img  string `json:"img"`
	Caption string `json:"caption"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	name := request.QueryStringParameters["name"]
	if name == "" {
		return Response{
			StatusCode: 400,
			Body:       "Missing 'name' query parameter",
		}, nil
	}

	jsonData, err := json.Marshal(get(name))
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: 200,
		Body:       string(jsonData),
	}, nil
}

func get(name string) Profile {
	insta := goinsta.New(os.Getenv("user"), os.Getenv("password"))
	if err := insta.Login(); err != nil {
		panic(err)
	}

	profile, err := insta.VisitProfile(name)
	if err != nil {
		log.Fatal(err)
	}

	user := profile.User
	feed := profile.Feed

	images := make([]Img, len(feed.Items))
	for i, image := range feed.Items {
		images[i] = Img{
			Img: image.Images.GetBest(),
			Caption: image.Caption.Text,
		}
	}

	return Profile{
		Name:   name,
		Bio:    user.Biography,
		Images: images,
	}
}

func main() {
	lambda.Start(Handler)
}