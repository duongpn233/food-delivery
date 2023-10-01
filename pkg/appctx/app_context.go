package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{
		db: db,
	}
}

func (ctx *appCtx) GetDBConnection() *gorm.DB {
	return ctx.db
}
