package repositories

import (
	"LEARN_API_SPEC_SWAGGER/api/models"
	"context"
	"database/sql"
)

type Repositories struct {
	db *sql.DB
}

func (r Repositories) GetAllBlog(ctx context.Context) []models.Blogs {

}
