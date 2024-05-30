package database

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const connectionString = "mongodb+srv://sanketugale2003:NaVU3UJLnNlL9faP@gocluster0.lil4p1z.mongodb.net/?retryWrites=true&w=majority&appName=GoCluster0"
// const dbName = "netfix"
// const colName = "watchlist"

// // Most important
// var collection *mongo.Collection

// // connect with mongoDB
// func init() {
// 	// client options
// 	clientOptions := options.Client().ApplyURI(connectionString)

// 	// connect to mongodb
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("MongoDB connection success")

// 	collection = client.Database(dbName).Collection(colName)

// 	// collection instance
// 	fmt.Println("Collection Instance is ready")
// }
