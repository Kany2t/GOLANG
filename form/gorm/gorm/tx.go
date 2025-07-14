package gorm

import (
	"errors"
	"gorm.io/gorm"
)

func Transaction() {
	t := teacherTemp
	t1 := teacherTemp
	t2 := teacherTemp
	t3 := teacherTemp

	DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些数据库操作（比如创建、更新、删除等）
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		//回滚子事务，不影响大事务最终结果
		tx.Transaction(func(tx1 *gorm.DB) error {
			tx1.Create(t1)
			return errors.New("rollback t1")
		})

		tx.Transaction(func(tx2 *gorm.DB) error {
			if err := tx2.Create(&t2).Error; err != nil {
				return err
			}
			return nil
		})
		if err := tx.Create(&t3).Error; err != nil {
			return err
		}
		return nil
	})
}
