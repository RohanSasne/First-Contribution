package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){

	for event:= range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5103732507206-5107423134357-PVhtmroPCLbP0FQFYXBEkF31") //for demonstation i have used this in the code, always declare all api keys in .env file
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A053Y0WTFK2-5110218401155-41634dc2e9690a67c38ba43f45ec506fe3de89ec36a620106a66855a52ebe3d3")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},

	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil{
		log.Fatal(err)
	}


}