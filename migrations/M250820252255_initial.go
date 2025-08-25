package migrations

import (
	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/participant"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain_events/domain_events_base"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func M250820252255_Initial() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "M250820252255",
		Migrate: func(tx *gorm.DB) error {

		err := tx.AutoMigrate(
			&user.User{}, 
			&event.Event{}, 
			&category.Category{}, 
			&participant.Participant{}, 
			&user.UserRefreshToken{},
			&domain_events_base.DomainEventData{})

		if err != nil {
			return err
		}

		err = tx.Exec("ALTER TABLE events " +
			"ADD CONSTRAINT fk_event_owner " +
			"FOREIGN KEY (owner_id) REFERENCES users(id) " +
			"ON DELETE NO ACTION " +
			"ON UPDATE NO ACTION;").Error

		if err != nil {
			return err
		}

		err = tx.Exec("ALTER TABLE events " +
			"ADD CONSTRAINT fk_event_category " +
			"FOREIGN KEY (category_id) REFERENCES categories(id) " +
			"ON DELETE NO ACTION " +
			"ON UPDATE NO ACTION;").Error

		if err != nil {
			return err
		}

		err = tx.Exec("ALTER TABLE participants " +
			"ADD CONSTRAINT fk_participant_event " +
			"FOREIGN KEY (event_id) REFERENCES events(id) " +
			"ON DELETE NO ACTION " +
			"ON UPDATE NO ACTION;").Error

		if err != nil {
			return err
		}

		err = tx.Exec("ALTER TABLE participants " +
			"ADD CONSTRAINT fk_participant_user " +
			"FOREIGN KEY (user_id) REFERENCES users(id) " +
			"ON DELETE NO ACTION " +
			"ON UPDATE NO ACTION;").Error

		if err != nil {
			return err
		}

		err = tx.Exec("ALTER TABLE user_refresh_tokens " +
			"ADD CONSTRAINT fk_token_user " +
			"FOREIGN KEY (user_id) REFERENCES users(id) " +
			"ON DELETE NO ACTION " +
			"ON UPDATE NO ACTION;").Error

		if err != nil {
			return err
		}

		for _, categorySeed := range categorySeeds() {
			if err := tx.Create(&categorySeed).Error; err != nil {
				return err
			}
		}

		return nil
		},

		Rollback: func(tx *gorm.DB) error {
			migrator := tx.Migrator()
			return migrator.DropTable("users", "events", "categories", "participants", "user_refresh_tokens", "domain_event_data")
		},
	}
}

func categorySeeds() []category.Category {
	return []category.Category{
		{ID: uuid.MustParse("FAC18306-89F4-42A1-A07D-478C0F5E27E0"), Title: "Свадьба"},
		{ID: uuid.MustParse("360E2CC1-4A3D-4790-96BB-C6152C298E53"), Title: "День рождения"},
		{ID: uuid.MustParse("101D7F7C-322D-41AC-BEDF-B91E8A9F2BA7"), Title: "Погулять"},
		{ID: uuid.MustParse("86942D28-ACBB-4C94-97A3-55197F3B14FD"), Title: "Настолки"},
		{ID: uuid.MustParse("B00432B6-9D8A-4B19-9D7E-C6D30F91F7F0"), Title: "В бар с Димассом"},
	}
}