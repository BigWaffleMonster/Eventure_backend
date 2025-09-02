package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)


func M030920252053_RemoveRefreshTokenFromSession() *gormigrate.Migration {
	return &gormigrate.Migration {
		ID: "M030920252053_RemoveRefreshTokenFromSession",
		Migrate: func(tx *gorm.DB) error {
			// Проверяем существование таблицы
			var tableExists bool
			if err := tx.Raw(`
				SELECT EXISTS (
					SELECT FROM information_schema.tables 
					WHERE table_schema = 'public' 
					AND table_name = 'user_sessions'
				)
			`).Scan(&tableExists).Error; err != nil {
				return err
			}

			if !tableExists {
				// Таблицы нет, ничего делать не нужно
				return nil
			}

			// Проверяем существование столбца refresh_token
			var columnExists bool
			if err := tx.Raw(`
				SELECT EXISTS (
					SELECT FROM information_schema.columns 
					WHERE table_schema = 'public' 
					AND table_name = 'user_sessions' 
					AND column_name = 'refresh_token'
				)
			`).Scan(&columnExists).Error; err != nil {
				return err
			}

			if !columnExists {
				// Столбца уже нет, ничего делать не нужно
				return nil
			}

			// Удаляем индекс для refresh_token, если он существует
			var indexExists bool
			if err := tx.Raw(`
				SELECT EXISTS (
					SELECT FROM pg_indexes 
					WHERE schemaname = 'public' 
					AND tablename = 'user_sessions' 
					AND indexname = 'idx_user_sessions_refresh_token'
				)
			`).Scan(&indexExists).Error; err != nil {
				return err
			}

			if indexExists {
				if err := tx.Exec(`DROP INDEX idx_user_sessions_refresh_token`).Error; err != nil {
					return err
				}
			}

			// Удаляем столбец refresh_token
			if err := tx.Exec(`ALTER TABLE user_sessions DROP COLUMN refresh_token`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Проверяем существование таблицы
			var tableExists bool
			if err := tx.Raw(`
				SELECT EXISTS (
					SELECT FROM information_schema.tables 
					WHERE table_schema = 'public' 
					AND table_name = 'user_sessions'
				)
			`).Scan(&tableExists).Error; err != nil {
				return err
			}

			if !tableExists {
				return nil
			}

			// Проверяем существование столбца refresh_token
			var columnExists bool
			if err := tx.Raw(`
				SELECT EXISTS (
					SELECT FROM information_schema.columns 
					WHERE table_schema = 'public' 
					AND table_name = 'user_sessions' 
					AND column_name = 'refresh_token'
				)
			`).Scan(&columnExists).Error; err != nil {
				return err
			}

			if columnExists {
				// Столбец уже существует, ничего делать не нужно
				return nil
			}

			// Добавляем столбец refresh_token обратно
			if err := tx.Exec(`
				ALTER TABLE user_sessions 
				ADD COLUMN refresh_token VARCHAR(500) NOT NULL DEFAULT ''
			`).Error; err != nil {
				return err
			}

			// Создаем индекс для refresh_token
			if err := tx.Exec(`
				CREATE INDEX idx_user_sessions_refresh_token ON user_sessions(refresh_token)
			`).Error; err != nil {
				return err
			}

			return nil
		},
	}
}