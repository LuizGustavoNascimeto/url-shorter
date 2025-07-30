package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	// environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	MONGO_URL := os.Getenv("MONGO_URL")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URL).
		SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// ping pong
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Define a struct for the url shortening
	type urlSchema struct {
		OriginalUrl string `bson:"original_url"`
		ShortUrl    string `bson:"short_url"`
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Define o banco de dados e a coleção (MONGO DB THINGS)
	urlsColl := client.Database("shortener").Collection("urls")

	// Endpoint para inserir
	router.POST("/api/shorten", func(c *gin.Context) {

		//Isso aqui é parte da mágica do gin
		//Ele faz o bind do JSON recebido no corpo da requisição para a struct definida
		var reqBody map[string]interface{}
		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}

		shortId, err := shortid.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar ID curto"})
			return
		}

		newUrl := urlSchema{
			OriginalUrl: reqBody["original_url"].(string),
			ShortUrl:    shortId,
		}

		_, err = urlsColl.InsertOne(context.TODO(), newUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao inserir no banco"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("URL: %s encurtada com sucesso para %s", newUrl.OriginalUrl, newUrl.ShortUrl)})
	})

	// Endpoint para buscar url encurtada
	router.GET("/:shortUrl", func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")
		var url urlSchema
		err := urlsColl.FindOne(context.TODO(), gin.H{"short_url": shortUrl}).Decode(&url)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL não encontrada"})
			return
		}
		// Redireciona para a URL original
		if url.OriginalUrl != "" {
			c.Redirect(http.StatusFound, url.OriginalUrl)
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL original não encontrada"})
			return
		}

	})

	router.Run(":" + os.Getenv("PORT"))
}
