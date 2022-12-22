package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yuuuutsk/gobase-backend/pkg/logger"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rs/cors"
	"github.com/yuuuutsk/gobase-backend/cmd"
	"github.com/yuuuutsk/gobase-backend/cmd/di"
	"github.com/yuuuutsk/gobase-backend/pkg"
	"github.com/yuuuutsk/gobase-backend/pkg/database"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yuuuutsk/gobase-backend/graph"
	"github.com/yuuuutsk/gobase-backend/graph/generated"
)

const defaultPort = "8082"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

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

	config := cmd.NewConfig()
	dbConfig := database.NewDBConfig(dbName, dbUser, dbPassword, dbAddress, dbNet)
	db, err := database.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("failed to database.NewDB: %s", err.Error())
	}
	defer db.Close()

	logger := logger.NewLogger(os.Stdout, false)
	clock := pkg.NewClock()
	usecasses := di.InitUseCases(db, config, clients, logger, clock)
	graph.SetUsecases(usecasses)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	handler := cors.Default().Handler(srv) // ★CORS レスポンス対応

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
