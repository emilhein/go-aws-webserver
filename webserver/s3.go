package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"net/http"
)

type Input struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

const S3_REGION = "eu-west-1"

type S3Handler struct {
	Session *session.Session
	Bucket  string
}

func (h S3Handler) ReadFile(key string) ([]byte, error) {
	results, err := s3.New(h.Session).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func extractMovieData(byteData []byte) ([]*Movie, error) {
	var moviesData []*Movie
	err := json.Unmarshal(byteData, &moviesData)
	if err != nil {
		return nil, err
	}
	return moviesData, nil
}

func GetS3File(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		// Handle error
	}
	w.Header().Set("Content-Type", "application/json")
	var input Input
	_ = json.NewDecoder(r.Body).Decode(&input)

	fmt.Printf("Getting: %s/%s \n", input.Bucket, input.Key)

	handler := S3Handler{
		Session: sess,
		Bucket:  input.Bucket,
	}
	fileString := input.Key // "movies.json"
	contents, err := handler.ReadFile(fileString)
	if err != nil {
		fmt.Println(err)
	}

	// toJson, error := extractMovieData(contents)
	// if err != nil {
	// 	fmt.Println(error)
	// }
	// for _, p := range toJson {
	// 	log.Printf("Name: %s , adsense: %s \n", p.Title, p.Rating)
	// }

	// fmt.Fprintf(w, contents)
	// var openReplacement interface{}
	// err := json.Marshal(contents, &openReplacement)
	var jsons interface{}
	error := json.Unmarshal(contents, &jsons)
	if error != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(jsons)
	// return contents

}
