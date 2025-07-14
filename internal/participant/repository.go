package participant

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ParticipantRepository interface{
    Create(participant *Participant)  results.Result
	Update(participant *Participant) results.Result
	Delete(id uuid.UUID) results.Result
	GetByID(id uuid.UUID) (*Participant, results.Result)
	GetCollection(eventID uuid.UUID) (*[]Participant, results.Result)
	GetOwnedCollection(currentUserID uuid.UUID) (*[]Participant, results.Result)
}

type participantRepository struct {
	DB *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) ParticipantRepository {
	return &participantRepository{DB: db}
}

func (r participantRepository) Create(participant *Participant) results.Result {
	err := r.DB.Create(participant).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *participantRepository) Delete(id uuid.UUID) results.Result {
	var Participant Participant
	err := r.DB.Where("id = ?", id).Delete(&Participant).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *participantRepository) Update(participant *Participant) results.Result {
	err := r.DB.Save(participant).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *participantRepository) GetByID(id uuid.UUID) (*Participant, results.Result) {
	var participant Participant
	result := r.DB.Where("id = ?", id).First(&participant)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &participant, results.NewResultOk()
}

func (r *participantRepository) GetCollection(eventID uuid.UUID) (*[]Participant, results.Result){
	var participants []Participant

	result := r.DB.Find(&participants, "event_id = ?", eventID)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &participants, results.NewResultOk()
}

func (r *participantRepository) GetOwnedCollection(currentUserID uuid.UUID) (*[]Participant, results.Result){
	var participants []Participant

	result := r.DB.Find(&participants, "user_id=?", currentUserID)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &participants, results.NewResultOk()
}