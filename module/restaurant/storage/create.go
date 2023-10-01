package restaurantstorage

import (
	"context"

	restaurantmodel "fooddelivery/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
