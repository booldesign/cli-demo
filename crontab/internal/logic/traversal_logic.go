package logic

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/25 11:23
 * @Desc:
 */

type TraversalLogic struct {
	ctx context.Context
	db  *gorm.DB
}

func NewTraversalLogic(ctx context.Context, db *gorm.DB) *TraversalLogic {
	return &TraversalLogic{
		ctx: ctx,
		db:  db,
	}
}

func (t *TraversalLogic) Traversal() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
