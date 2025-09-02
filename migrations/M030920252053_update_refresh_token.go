package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func M030920252053_UpdateRefreshToken() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "M030920252053_UpdateRefreshToken",
        Migrate: func(tx *gorm.DB) error {
            // Сначала проверяем существование старой таблицы
            var oldTableExists bool
            if err := tx.Raw(`
                SELECT EXISTS (
                    SELECT FROM information_schema.tables 
                    WHERE table_schema = 'public' 
                    AND table_name = 'user_refresh_tokens'
                )
            `).Scan(&oldTableExists).Error; err != nil {
                return err
            }

            // Проверяем существование новой таблицы
            var newTableExists bool
            if err := tx.Raw(`
                SELECT EXISTS (
                    SELECT FROM information_schema.tables 
                    WHERE table_schema = 'public' 
                    AND table_name = 'user_sessions'
                )
            `).Scan(&newTableExists).Error; err != nil {
                return err
            }

            // Если новая таблица уже существует, ничего не делаем
            if newTableExists {
                return nil
            }

            // Если старая таблица существует, переносим данные
            if oldTableExists {
                // Создаем новую таблицу
                if err := tx.Exec(`
                    CREATE TABLE user_sessions (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        user_id UUID NOT NULL,
                        refresh_token VARCHAR(500) NOT NULL,
                        user_agent VARCHAR(255) NOT NULL,
                        ip_address VARCHAR(45) NOT NULL,
                        fingerprint VARCHAR(255) NOT NULL,
                        expires_at TIMESTAMP NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        CONSTRAINT fk_user_sessions_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
                    )
                `).Error; err != nil {
                    return err
                }

                // Создаем индексы для новой таблицы
                indexes := []string{
                    "CREATE INDEX idx_user_sessions_user_id ON user_sessions(user_id)",
                    "CREATE INDEX idx_user_sessions_refresh_token ON user_sessions(refresh_token)",
                    "CREATE INDEX idx_user_sessions_expires_at ON user_sessions(expires_at)",
                }

                for _, indexSQL := range indexes {
                    if err := tx.Exec(indexSQL).Error; err != nil {
                        return err
                    }
                }

                // Переносим данные из старой таблицы (если есть)
                if err := tx.Exec(`
                    INSERT INTO user_sessions (user_id, refresh_token, user_agent, ip_address, fingerprint, expires_at, created_at)
                    SELECT 
                        user_id, 
                        refresh_token,
                        '' as user_agent,
                        '' as ip_address, 
                        '' as fingerprint,
                        NOW() + INTERVAL '30 days' as expires_at,
                        NOW() as created_at
                    FROM user_refresh_tokens
                `).Error; err != nil {
                    return err
                }

                // Удаляем старую таблицу
                if err := tx.Exec("DROP TABLE user_refresh_tokens CASCADE").Error; err != nil {
                    return err
                }

            } else {
                // Если старой таблицы нет, просто создаем новую
                if err := tx.Exec(`
                    CREATE TABLE IF NOT EXISTS user_sessions (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        user_id UUID NOT NULL,
                        refresh_token VARCHAR(500) NOT NULL,
                        user_agent VARCHAR(255) NOT NULL,
                        ip_address VARCHAR(45) NOT NULL,
                        fingerprint VARCHAR(255) NOT NULL,
                        expires_at TIMESTAMP NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        CONSTRAINT fk_user_sessions_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
                    )
                `).Error; err != nil {
                    return err
                }

                // Создаем индексы для новой таблицы
                indexes := []string{
                    "CREATE INDEX IF NOT EXISTS idx_user_sessions_user_id ON user_sessions(user_id)",
                    "CREATE INDEX IF NOT EXISTS idx_user_sessions_refresh_token ON user_sessions(refresh_token)",
                    "CREATE INDEX IF NOT EXISTS idx_user_sessions_expires_at ON user_sessions(expires_at)",
                }

                for _, indexSQL := range indexes {
                    if err := tx.Exec(indexSQL).Error; err != nil {
                        return err
                    }
                }
            }

            return nil
        },
        Rollback: func(tx *gorm.DB) error {
            // Проверяем существование старой таблицы
            var oldTableExists bool
            if err := tx.Raw(`
                SELECT EXISTS (
                    SELECT FROM information_schema.tables 
                    WHERE table_schema = 'public' 
                    AND table_name = 'user_refresh_tokens'
                )
            `).Scan(&oldTableExists).Error; err != nil {
                return err
            }

            // Если старая таблица не существует, создаем ее
            if !oldTableExists {
                if err := tx.Exec(`
                    CREATE TABLE user_refresh_tokens (
                        user_id UUID,
                        refresh_token VARCHAR(500) NOT NULL,
                        CONSTRAINT fk_user_refresh_tokens_user FOREIGN KEY (user_id) REFERENCES users(id)
                    )
                `).Error; err != nil {
                    return err
                }

                // Создаем индексы для старой таблицы
                if err := tx.Exec(`
                    CREATE INDEX idx_user_refresh_tokens_user_id ON user_refresh_tokens(user_id)
                `).Error; err != nil {
                    return err
                }

                if err := tx.Exec(`
                    CREATE INDEX idx_user_refresh_tokens_refresh_token ON user_refresh_tokens(refresh_token)
                `).Error; err != nil {
                    return err
                }
            }

            // Удаляем новую таблицу
            if err := tx.Exec("DROP TABLE IF EXISTS user_sessions CASCADE").Error; err != nil {
                return err
            }

            return nil
        },
    }
}