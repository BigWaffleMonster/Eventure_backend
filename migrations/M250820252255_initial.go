package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M250820252255_Initial() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "M250820252255_Initial",
		Migrate: func(tx *gorm.DB) error {
			// Создание таблицы categories
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS categories (
					id UUID PRIMARY KEY,
					title VARCHAR(255)
				)
			`).Error; err != nil {
				return err
			}

			// Создание таблицы users
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS users (
					id UUID PRIMARY KEY,
					user_name VARCHAR(255),
					email VARCHAR(255) NOT NULL,
					password VARCHAR(255),
					date_created TIMESTAMP,
					date_updated TIMESTAMP,
					is_email_confirmed BOOLEAN DEFAULT false
				)
			`).Error; err != nil {
				return err
			}

			// Создание уникальных индексов для users
			if err := tx.Exec(`
				CREATE UNIQUE INDEX IF NOT EXISTS idx_users_user_name_unique ON users(user_name)
			`).Error; err != nil {
				return err
			}

			if err := tx.Exec(`
				CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email_unique ON users(email)
			`).Error; err != nil {
				return err
			}

			// Создание таблицы events
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS events (
					id UUID PRIMARY KEY,
					owner_id UUID,
					title VARCHAR(255) NOT NULL,
					description TEXT NOT NULL,
					max_qty_participants INTEGER,
					location VARCHAR(255),
					private BOOLEAN DEFAULT false,
					start_date TIMESTAMP,
					end_date TIMESTAMP,
					date_created TIMESTAMP,
					date_updated TIMESTAMP,
					category_id UUID,
					CONSTRAINT fk_events_owner FOREIGN KEY (owner_id) REFERENCES users(id),
					CONSTRAINT fk_events_category FOREIGN KEY (category_id) REFERENCES categories(id)
				)
			`).Error; err != nil {
				return err
			}

			// Для events
			if err := tx.Exec(`
				CREATE INDEX IF NOT EXISTS idx_events_owner_id ON events(owner_id)
			`).Error; err != nil {
				return err
			}

			if err := tx.Exec(`
				CREATE INDEX IF NOT EXISTS idx_events_category_id ON events(category_id)
			`).Error; err != nil {
				return err
			}

			// Создание таблицы participants
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS participants (
					id UUID PRIMARY KEY,
					user_id UUID,
					event_id UUID,
					status VARCHAR(10) CHECK (status IN ('Yes', 'No', 'Maybe')),
					ticket VARCHAR(255),
					date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					CONSTRAINT fk_participants_user FOREIGN KEY (user_id) REFERENCES users(id),
					CONSTRAINT fk_participants_event FOREIGN KEY (event_id) REFERENCES events(id)
				)
			`).Error; err != nil {
				return err
			}

			// Создание индексов для participants
			if err := tx.Exec(`
				CREATE INDEX IF NOT EXISTS idx_participants_user_id ON participants(user_id)
			`).Error; err != nil {
				return err
			}

			if err := tx.Exec(`
				CREATE INDEX IF NOT EXISTS idx_participants_event_id ON participants(event_id)
			`).Error; err != nil {
				return err
			}

			// Создание таблицы user_refresh_tokens (старое название)
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS user_refresh_tokens (
					user_id UUID,
					refresh_token VARCHAR(500) NOT NULL,
					CONSTRAINT fk_user_refresh_tokens_user FOREIGN KEY (user_id) REFERENCES users(id)
				)
			`).Error; err != nil {
				return err
			}

			// Создание индекса для refresh_token в старой таблице
			if err := tx.Exec(`
				CREATE INDEX IF NOT EXISTS idx_user_refresh_tokens_token ON user_refresh_tokens(refresh_token)
			`).Error; err != nil {
				return err
			}

			if err := tx.Exec(`
				CREATE INDEX IF NOT EXISTS idx_user_refresh_tokens_user_id ON user_refresh_tokens(user_id)
			`).Error; err != nil {
				return err
			}

			// Сидинг категорий (добавьте в самый конец)
			if err := tx.Exec(`
				INSERT INTO categories (id, title)
				VALUES 
					('FAC18306-89F4-42A1-A07D-478C0F5E27E0', 'Свадьба'),
					('360E2CC1-4A3D-4790-96BB-C6152C298E53', 'День рождения'),
					('101D7F7C-322D-41AC-BEDF-B91E8A9F2BA7', 'Погулять'),
					('86942D28-ACBB-4C94-97A3-55197F3B14FD', 'Настолки'),
					('B00432B6-9D8A-4B19-9D7E-C6D30F91F7F0', 'В бар с Димассом')
				ON CONFLICT (id) DO NOTHING
			`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Удаление таблиц в обратном порядке (сначала дочерние, потом родительские)
			tables := []string{
				"user_refresh_tokens",
				"participants", 
				"events",
				"users",
				"categories",
			}

			for _, table := range tables {
				if err := tx.Exec("DROP TABLE IF EXISTS " + table + " CASCADE").Error; err != nil {
					return err
				}
			}
			return nil
		},
	}
}