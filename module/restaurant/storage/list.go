package restaurantstorage

import (
	"context"

	common "fooddelivery/common"
	restaurantmodel "fooddelivery/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")

	if filter != nil {
		if filter.OwnerId > 0 {
			db = db.Where("owner_id = ?", filter.OwnerId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.Offset(offset).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
