package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M250820252258_AddDomainEvents() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "M250820252258_AddDomainEvents",
		Migrate: func(tx *gorm.DB) error {
			// Проверяем существование таблицы
			var tableExists bool
			if err := tx.Raw(`
				SELECT EXISTS (
					SELECT FROM information_schema.tables 
					WHERE table_schema = 'public' 
					AND table_name = 'domain_event_data'
				)
			`).Scan(&tableExists).Error; err != nil {
				return err
			}

			// Если таблица уже существует, ничего не делаем
			if tableExists {
				return nil
			}

			// Создаем таблицу domain_event_data
			if err := tx.Exec(`
				CREATE TABLE domain_event_data (
					id UUID PRIMARY KEY,
					type VARCHAR(50) NOT NULL,
					content TEXT NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				)
			`).Error; err != nil {
				return err
			}

			// Создаем индекс для поля type (если планируются частые поиски по типу)
			if err := tx.Exec(`
				CREATE INDEX idx_domain_event_data_type ON domain_event_data(type)
			`).Error; err != nil {
				return err
			}

			// Создаем индекс для поля created_at (для сортировки по дате)
			if err := tx.Exec(`
				CREATE INDEX idx_domain_event_data_created_at ON domain_event_data(created_at)
			`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Удаляем таблицу domain_event_data
			if err := tx.Exec("DROP TABLE IF EXISTS domain_event_data").Error; err != nil {
				return err
			}
			return nil
		},
	}
}