package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func FirestoreInit() (*firestore.Client, error) {
	ctx := context.Background()
	serviceAccount := option.WithCredentialsFile("../../keys/serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return client, nil
}
