package participant

import (
	// "errors"
	// "net/http"
	// "time"

	"errors"
	"net/http"

	schema "github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
	"github.com/BigWaffleMonster/Eventure_backend/internal/utils"

	// "github.com/BigWaffleMonster/Eventure_backend/internal/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) *ParticipantRepository {
	return &ParticipantRepository{db: db}
}

func (r *ParticipantRepository) GetParticipantsFromEvent(eventID uuid.UUID) ([]schema.Participant, error) {
	var participants []schema.Participant

	result := r.db.
		Preload("User").
		Where("event_id = ?", eventID).
		Find(&participants)

	if result.Error != nil {
		return nil, utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to fetch participants",
			result.Error,
		)
	}

	return participants, nil
}

func (r *ParticipantRepository) AddParticipantToEvent(userID, eventID uuid.UUID) error {
	participant := schema.Participant{
		ID:      uuid.New(),
		UserID:  userID,
		EventID: eventID,
	}

	if err := r.db.Create(&participant).Error; err != nil {
		return utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to register participant",
			err,
		)
	}

	return nil
}

func (r *ParticipantRepository) RemoveParticipantFromEvent(userID, eventID uuid.UUID) error {
	result := r.db.
		Where("user_id = ? AND event_id = ?", userID, eventID).
		Delete(&schema.Participant{})

	if result.Error != nil {
		return utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to remove participant",
			result.Error,
		)
	}

	if result.RowsAffected == 0 {
		return utils.ErrNotFound
	}

	return nil
}

func (r *ParticipantRepository) RemoveAllParticipantsFromEvent(eventID uuid.UUID) error {
	result := r.db.
		Where("event_id = ?", eventID).
		Delete(&schema.Participant{})

	if result.Error != nil {
		return utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to remove all participants",
			result.Error,
		)
	}

	// Note: RowsAffected == 0 is not an error here (event might have had no participants)
	return nil
}

func (r *ParticipantRepository) GetEventCapacity(eventID uuid.UUID) (capacity, max_capacity *int, err error) {
	var event schema.Event
	if err := r.db.First(&event, "id = ?", eventID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, utils.ErrNotFound
		}
		return nil, nil, utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to fetch event",
			err,
		)
	}

	return event.Capacity, event.MaxCapacity, nil
}

func (r *ParticipantRepository) CheckUserExistence(userID uuid.UUID) error {
	var user schema.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewAppErrorWithErr(
				http.StatusBadRequest,
				"User not found",
				err,
			)
		}
		return utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to fetch user",
			err,
		)
	}
	return nil
}

func (r *ParticipantRepository) CheckIfUserParticipant(userID, eventID uuid.UUID) error {
	var existing schema.Participant
	if err := r.db.Where("user_id = ? AND event_id = ?", userID, eventID).
		First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to check if user is participant",
			err,
		)
	}

	return utils.ErrConflict
}
