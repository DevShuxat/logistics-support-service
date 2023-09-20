package models
import (
	"time"
)

type RestaurantInfo struct {
		Name string `json:"name"`
}

type EaterInfo struct {
		Name string `json:"name"`
		ImageUrl string `json:"image_url"`
}
 
	type Rating struct {
		EntityID string `json:"entity_id" gorm:"index:idx_entity_rating"`
		ID string `json:"id" gorm:"primaryKey"`
		Restaurant *RestaurantInfo `json:"restaurant" gorm:"serializer:json"`
		EaterID string `json:"eater_id"`
		Eater *EaterInfo `json:"eater" gorm:"serializer:json"`
		Rating int `json:"rating"`
		Comment string `json:"comment"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
}