package main

import (
	"log"
	"os"

	middleware "fooddelivery/middleware"
	restaurantgin "fooddelivery/module/restaurant/transport/gin"
	appctx "fooddelivery/pkg/appctx"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	appCtx := appctx.NewAppContext(db)

	router := gin.Default()
	router.Use(middleware.Recover(appCtx))

	v1 := router.Group("/v1")

	restaurant := v1.Group("/restaurants")
	restaurant.POST("/create", restaurantgin.CreateRestaurant(appCtx))
	restaurant.DELETE("/delete/:id", restaurantgin.DeleteRestaurant(appCtx))
	restaurant.GET("", restaurantgin.ListRestaurant(appCtx))

	router.Run()
}
