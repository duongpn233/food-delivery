package ginrestaurant

import (
	"fooddelivery/common"
	restaurantbiz "fooddelivery/module/restaurant/biz"
	restaurantmodel "fooddelivery/module/restaurant/model"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	appctx "fooddelivery/pkg/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetDBConnection()

		var pagingData common.Paging
		if err := ctx.ShouldBind(&pagingData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		pagingData.Fulfill()

		var filter restaurantmodel.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &pagingData)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRes(result, pagingData, filter))
	}
}
