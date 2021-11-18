package firestore

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirestoreLocalClient() *firestore.Client {
	ctx := context.Background()

	_, p, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(p), "../..")
	serviceAccountPath := rootDir + "/keys/serviceAccount.json"
	if _, err := os.Stat(serviceAccountPath); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("file %s not exists.", serviceAccountPath)
	}
	serviceAccount := option.WithCredentialsFile(serviceAccountPath)

	app, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		log.Fatalf("firebase.NewApp err: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore err: %v", err)
	}

	return client
}

func NewFirestoreTestClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, "test")
	if err != nil {
		log.Fatalf("firestore.NewClient err: %v", err)
	}

	return client
}
