package posts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"query-api/pkg/common/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func UpdatePostRegisters(db *mongo.Database) {
	reqUrl := fmt.Sprintf("%s/posts", os.Getenv("POSTS_ENDPOINT"))
	res, err := http.Get(reqUrl)

	if err != nil {
		fmt.Println("Error getting posts from posts service", err.Error())
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body) // response body is []byte
	if err != nil {
		fmt.Println("Error reading body", err.Error())
		return
	}

	var data []models.Post
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling posts from posts service", err.Error())
		return
	}

	fmt.Println(data)
}
