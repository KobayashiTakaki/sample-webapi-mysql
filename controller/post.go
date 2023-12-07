package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type PostsController struct {
	db *sql.DB
}

func NewPostsController(db *sql.DB) *PostsController {
	return &PostsController{
		db: db,
	}
}

func (c *PostsController) ListPosts(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := "" +
		"SELECT " +
		"  id, " +
		"  content, " +
		"  created_at " +
		"FROM " +
		"  posts " +
		"ORDER BY " +
		"  created_at DESC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	ret := ListPostsResponse{}
	for rows.Next() {
		var (
			id        int64
			content   string
			createdAt time.Time
		)
		err := rows.Scan(
			&id,
			&content,
			&createdAt,
		)
		if err != nil {
			return err
		}
		ret.Posts = append(ret.Posts, &Post{
			ID:        id,
			Content:   content,
			CreatedAt: createdAt,
		})
	}

	return ectx.JSON(http.StatusOK, ret)
}

func (c *PostsController) GetPost(ectx echo.Context) error {
	paramID, err := strconv.Atoi(ectx.Param("id"))
	if err != nil {
		return err
	}

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	// defer tx.Rollback()
	// うっかりCommit/Rollbackを忘れる
	query := "" +
		"SELECT " +
		"  id, " +
		"  content, " +
		"  created_at " +
		"FROM " +
		"  posts " +
		"WHERE " +
		"  id = ?"
	var (
		id        int64
		content   string
		createdAt time.Time
	)
	err = tx.QueryRow(query, paramID).Scan(
		&id,
		&content,
		&createdAt,
	)
	if err != nil {
		return err
	}

	ret := GetPostResponse{
		Post: &Post{
			ID:        id,
			Content:   content,
			CreatedAt: createdAt,
		},
	}

	return ectx.JSON(http.StatusOK, ret)
}

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type ListPostsResponse struct {
	Posts []*Post `json:"posts"`
}

type GetPostResponse struct {
	Post *Post `json:"post"`
}
