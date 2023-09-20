package Repositories

import (
	"context"
	
)

type RatingRepository interface {
SaveRating(ctx context.Context, rating *models.Rating) error
UpdateRating(ctx context.Context, rating *models.Rating) error
ListRatingsByEntity(ctx context.Context, entityID, sort string, page, pageSize int) ([]*models.Rating, error)
}