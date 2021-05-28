package service

import (
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/dao"
)

type Service struct {
	ctx context.Context // connect to validate
	dao *dao.Dao        // connect to dao
}

func New(ctx context.Context) *Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine)) // add Jaeger trace.
	// svc.dao = dao.New(global.DBEngine)
	return &svc
}
