package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emilhein/go-aws-utility/util/services"
)

type Input struct {
	Bucket    string   `json:"bucket"`
	Filepaths []string `json:"filepaths"`
}

func GetS3Files(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var input Input
	_ = json.NewDecoder(r.Body).Decode(&input)

	fmt.Printf("Getting: %s/%s \n", input.Bucket, input.Filepaths)

	inputToFunc := services.FilesInput{Bucket: input.Bucket, FileNames: input.Filepaths}
	returnValues := services.GetS3Files(inputToFunc)
	json.NewEncoder(w).Encode(returnValues)

}
