package main

//go:generate

import (
	"context"
	"log"
	"os"

	"github.com/yuuuutsk/gobase-backend/pkg/logger"

	"github.com/yuuuutsk/gobase-backend/app/cli_usecase"

	"github.com/yuuuutsk/gobase-backend/app/usecase"

	"github.com/yuuuutsk/gobase-backend/cmd"

	"github.com/yuuuutsk/gobase-backend/cmd/di"
	"github.com/yuuuutsk/gobase-backend/pkg"
	"github.com/yuuuutsk/gobase-backend/pkg/database"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/urfave/cli"
)

func main() {
	ctx := context.Background()

	dbName := cmd.GetEnv("DB_NAME", "gobase-backend_local")
	dbAddress := cmd.GetEnv("DB_ADDRESS", "127.0.0.1:3306")
	dbUser := cmd.GetEnv("DB_USER", "testuser")
	dbPassword := cmd.GetEnv("DB_PASSWORD", "password")
	dbNet := cmd.GetEnv("DB_NET", "tcp")
	twitterConsumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	twitterConsumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	twitterAccessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	authConfig := oauth1.NewConfig(twitterConsumerKey, twitterConsumerSecret)
	token := oauth1.NewToken(twitterAccessToken, twitterAccessSecret)
	httpClient := authConfig.Client(oauth1.NoContext, token)

	// Twitter client
	twitterClient := twitter.NewClient(httpClient)

	clients := &cmd.TwitterClients{
		Client: twitterClient,
	}

	logger := logger.NewLogger(os.Stdout, false)
	config := cmd.NewConfig()
	dbConfig := database.NewDBConfig(dbName, dbUser, dbPassword, dbAddress, dbNet)
	db, err := database.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("failed to database.NewDB: %s", err.Error())
	}
	defer db.Close()

	clock := pkg.NewClock()
	usecasses := di.InitUseCases(db, config, clients, logger, clock)
	cliUsecasses := di.InitCLIUseCases(db, config, clients, logger, clock)

	{
		app := cli.NewApp()
		app.Name = "cli"
		app.Commands = []cli.Command{
			{
				Name: "list",
				Action: func(c *cli.Context) error {
					var errs []error

					{
						input := &usecase.TodoGetInput{Text: "1"}
						_, err := usecasses.TodoUseCase.Get(ctx, input, clock)
						if err != nil {
							errs = append(errs, err)
						}
					}

					{
						input := &cli_usecase.TodoGetInput{ID: "1"}
						_, err := cliUsecasses.TodoUseCase.Get(ctx, input, clock)
						if err != nil {
							errs = append(errs, err)
						}
					}
					return err
				},
			},
		}
		app.Run(os.Args)
	}

}
