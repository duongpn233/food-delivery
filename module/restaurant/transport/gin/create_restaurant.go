package ginrestaurant

import (
	common "fooddelivery/common"
	restaurantbiz "fooddelivery/module/restaurant/biz"
	restaurantmodel "fooddelivery/module/restaurant/model"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	appctx "fooddelivery/pkg/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetDBConnection()

		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(data.Id))
	}
}
