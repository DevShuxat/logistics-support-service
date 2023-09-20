package rating

import "gorm.io/gorm"

const (
	tableRating = "support.ratings"
)

type ratingRepoImpl struct {
	db *gorm.DB
} 

func NewRatingRepository(db *gorm.DB) *ratingRepoImpl {
	return &ratingRepoImpl{
		db: db,
	}
}

func (r *ratingRepoImpl) SaveRating(ctx context.Context, rating *models.Rating) error {
	result := r.db.WithContext(ctx).Table(tableRating).Create(rating)
	if result.Error != nil{
		return result.Error 
	return nil
	}
}