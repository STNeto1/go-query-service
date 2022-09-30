package posts

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"query-api/pkg/common/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdatePostRegisters(db *mongo.Database) {
	fmt.Println("==== Updating posts ====")
	col := db.Collection("posts")

	postsReqUrl := fmt.Sprintf("%s/posts", os.Getenv("POSTS_ENDPOINT"))
	postRes, err := http.Get(postsReqUrl)

	if err != nil {
		fmt.Println("Error getting posts from posts service", err.Error())
		return
	}

	defer postRes.Body.Close()

	postsBody, err := io.ReadAll(postRes.Body) // response body is []byte
	if err != nil {
		fmt.Println("Error reading body", err.Error())
		return
	}

	var posts []models.Post
	if err := json.Unmarshal(postsBody, &posts); err != nil {
		fmt.Println("Error unmarshalling posts from posts service", err.Error())
		return
	}

	commentsReqUrl := fmt.Sprintf("%s/comments", os.Getenv("COMMENTS_ENDPOINT"))
	commentsRes, err := http.Get(commentsReqUrl)
	if err != nil {
		fmt.Println("Error getting comments from comments service", err.Error())
		return
	}

	defer commentsRes.Body.Close()
	commentsBody, err := io.ReadAll(commentsRes.Body) // response body is []byte
	if err != nil {
		fmt.Println("Error reading body", err.Error())
		return
	}

	var comments []models.Comment
	if err := json.Unmarshal(commentsBody, &comments); err != nil {
		fmt.Println("Error unmarshalling comments from comments service", err.Error())
		return
	}

	for _, post := range posts {
		dbPost := models.DBPost{
			Ref:       post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Comments:  []models.Comment{},
		}

		for _, comment := range comments {
			if comment.PostID == post.ID {
				dbPost.Comments = append(dbPost.Comments, models.Comment{
					ID:        comment.ID,
					Comment:   comment.Comment,
					PostID:    comment.PostID,
					CreatedAt: comment.CreatedAt,
					UpdatedAt: comment.UpdatedAt,
				})
			}
		}

		var result models.DBPost
		err := col.FindOne(context.Background(), bson.D{{Key: "ref", Value: post.ID}}).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				dbPost.ID = primitive.NewObjectID()
				_, err := col.InsertOne(context.Background(), dbPost)
				if err != nil {
					fmt.Println("Error inserting post", err.Error())
				}
			}
		}

		if result.ID != primitive.NilObjectID {
			dbPost.ID = result.ID
			_, err := col.ReplaceOne(context.Background(), bson.D{{Key: "ref", Value: post.ID}}, dbPost)
			if err != nil {
				fmt.Println("Error updating post", err.Error())
			}
		}
	}

	fmt.Println("==== Done updating posts ====")
}
