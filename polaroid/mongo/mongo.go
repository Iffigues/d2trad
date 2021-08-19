package mongo

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "time"
    "context"
)


func Open()(a context.Context, b context.CancelFunc)  {
	return context.WithTimeout(context.Background(), 10*time.Second)
}


func Client(ctx context.Context, b context.CancelFunc)(a *mongo.Client, err error) {
	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:8081"))
}

func deco(client *mongo.Client, ctx context.Context) {
	client.Disconnect(ctx)
}

func Ping(client *mongo.Client, ctx context.Context) {
	client.Ping(ctx, readpref.Primary())
}

func Collection(client *mongo.Client, c string) (e *mongo.Collection) {
	return  client.Database("stuff").Collection(c)
}
