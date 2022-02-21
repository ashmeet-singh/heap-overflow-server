package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	// "io"
	"net/http"
	// "net/url"
	"context"
	"os"
	// "path"
	// "strconv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://localhost:27017"

type JSON_T1 struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	http.Handle("/", http.FileServer(http.Dir(os.Getenv("HO_DIR_WEBSITE"))))

	http.HandleFunc("/api/signup", func(w http.ResponseWriter, r *http.Request) {
		b := make([]byte, 1000)
		l, _ := r.Body.Read(b)
		m := JSON_T1{}
		json.Unmarshal(b[:l], &m)
		fmt.Println(string(b))
		fmt.Println(m.Name)
		fmt.Println(m.Password)

		salt := make([]byte, 32)
		rand.Read(salt)
		hashed_key := pbkdf2.Key([]byte(m.Password), salt, 100000, 32, sha256.New)
		fmt.Println(string(hashed_key))
	})

	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		fmt.Println("Unable to listen and serve")
	}
}
