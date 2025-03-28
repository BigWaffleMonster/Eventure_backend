package participant

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ParticipantRepository interface{
    Create(participant *Participant)  error
	Update(participant *Participant) error
	Remove(id uuid.UUID) error
	GetByID(id uuid.UUID) (*Participant, error)
	GetCollection(eventID uuid.UUID) (*[]Participant, error)
}

type participantRepository struct {
	DB *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) ParticipantRepository {
	return &participantRepository{DB: db}
}

func (r participantRepository) Create(participant *Participant) error {
	return r.DB.Create(participant).Error
}

func (r *participantRepository) Remove(id uuid.UUID) error {
	var Participant Participant
	return r.DB.Where("id = ?", id).Delete(&Participant).Error
}

func (r *participantRepository) Update(participant *Participant) error {
	return r.DB.Save(participant).Error
}

func (r *participantRepository) GetByID(id uuid.UUID) (*Participant, error) {
	var participant Participant
	result := r.DB.Where("id = ?", id).First(&participant)
	if result.Error != nil {
		return nil, result.Error
	}
	return &participant, nil
}

func (r *participantRepository) GetCollection(eventID uuid.UUID) (*[]Participant, error){
	var participants []Participant

	result := r.DB.Where("eventID = ?", eventID).Find(&participants)

	if result.Error != nil {
		return nil, result.Error
	}
	return &participants, nil
}