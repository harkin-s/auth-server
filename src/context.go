package main

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

var datastoreClient *datastore.Client

func getDbConnection() *datastore.Client {

	if datastoreClient == nil {
		ctx := context.Background()

		// Set your Google Cloud Platform project ID.
		projectID := "fisean-223219"

		// Creates a client.
		client, err := datastore.NewClient(ctx, projectID)

		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		datastoreClient = client

	}

	return datastoreClient
}
