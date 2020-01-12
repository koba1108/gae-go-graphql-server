package external

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App
var firestoreClient *firestore.Client

func initFirebase() {
	var err error
	opt := option.WithCredentialsFile("secrets/development/serviceAccountKey.json")
	if app, err = firebase.NewApp(context.Background(), nil, opt); err != nil {
		panic(err)
	}
}

func initFirestore() {
	client, err := app.Firestore(context.Background())
	if err != nil {
		panic(err)
	}
	firestoreClient = client
}

func GetFirebaseApp() *firebase.App {
	return app
}

func GetFirestore() *firestore.Client {
	if firestoreClient == nil {
		initFirestore()
	}
	return firestoreClient
}
