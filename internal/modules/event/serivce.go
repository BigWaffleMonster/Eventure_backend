package event

import (
	"io/fs"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/internal/types"
	t "github.com/BigWaffleMonster/Eventure_backend/internal/types"
	global_utils "github.com/BigWaffleMonster/Eventure_backend/internal/utils"
	"github.com/google/uuid"
)

type EventService struct {
	repo *EventRepository
}

func NewEventService(repo *EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(req *CreateEventRequest, userDataCtx *t.UserDataCtx, coverURL *string) (*EventResponse, error) {
	if req.StartDate.After(req.EndDate) || req.StartDate.Equal(req.EndDate) {
		return nil, global_utils.NewAppError(http.StatusBadRequest, "дата начала должна быть раньше даты окончания")
	}

	if req.StartDate.Before(time.Now()) {
		return nil, global_utils.NewAppError(http.StatusBadRequest, "дата начала должна быть в будущем")
	}

	category, err := s.repo.GetCategoryForEventByID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	event, err := s.repo.CreateEvent(req, userDataCtx.UserID, category.ID, coverURL)
	if err != nil {
		return nil, err
	}

	careatedEvent := EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Capacity:    0,
		MaxCapacity: event.MaxCapacity,
		Location:    (location)(event.Location),
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,

		Cover: event.Cover,

		DateCreated: time.Now(),
		DateUpdated: time.Now(),

		Category: CategoryResponse{ID: category.ID, Title: category.Title},
		Owner:    OwnerResponse{ID: userDataCtx.UserID, Login: userDataCtx.Login, Email: userDataCtx.Email},
	}

	return &careatedEvent, nil
}

func (s *EventService) GetEvents() ([]EventResponse, error) {
	var events []EventResponse
	events_raw, err := s.repo.GetEvents()
	if err != nil {
		return nil, err
	}

	for _, event := range events_raw {
		eventDTO := EventResponse{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Capacity:    event.Capacity,
			MaxCapacity: event.MaxCapacity,
			Location:    (location)(event.Location),
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			DateCreated: event.DateCreated,
			DateUpdated: event.DateUpdated,

			Category: CategoryResponse(event.Category),
			Owner:    OwnerResponse{ID: event.Owner.ID, Login: event.Owner.Login, Email: event.Owner.Email},
		}

		events = append(events, eventDTO)
	}

	return events, nil
}

func (s *EventService) GetEventByID(eventID uuid.UUID) (*EventResponse, error) {
	event, err := s.repo.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}

	eventDTO := EventResponse{
		ID:          eventID,
		Title:       event.Title,
		Description: event.Description,
		Capacity:    event.Capacity,
		MaxCapacity: event.MaxCapacity,
		Location:    (location)(event.Location),
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		DateCreated: event.DateCreated,
		DateUpdated: event.DateUpdated,

		Category: CategoryResponse(event.Category),
		Owner:    OwnerResponse{ID: event.Owner.ID, Login: event.Owner.Login, Email: event.Owner.Email},
	}

	return &eventDTO, nil
}

func (s *EventService) GetUserCreatedEvents(userID uuid.UUID) ([]EventResponse, error) {
	var events []EventResponse
	events_raw, err := s.repo.GetUserCreatedEvents(userID)
	if err != nil {
		return nil, err
	}

	for _, event := range events_raw {
		eventDTO := EventResponse{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Capacity:    event.Capacity,
			MaxCapacity: event.MaxCapacity,
			Location:    (location)(event.Location),
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			DateCreated: event.DateCreated,
			DateUpdated: event.DateUpdated,

			Category: CategoryResponse(event.Category),
			Owner:    OwnerResponse{ID: event.Owner.ID, Login: event.Owner.Login, Email: event.Owner.Email},
		}

		events = append(events, eventDTO)
	}

	return events, nil
}

func (s *EventService) GetUserParticipantingEvents(userID uuid.UUID) ([]EventResponse, error) {
	var events []EventResponse
	events_raw, err := s.repo.GetUserParticipatingEvents(userID)
	if err != nil {
		return nil, err
	}

	for _, participant := range events_raw {
		event := participant.Event
		eventDTO := EventResponse{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Capacity:    event.Capacity,
			MaxCapacity: event.MaxCapacity,
			Location:    (location)(event.Location),
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			DateCreated: event.DateCreated,
			DateUpdated: event.DateUpdated,

			Category: CategoryResponse(event.Category),
			Owner:    OwnerResponse{ID: event.Owner.ID, Login: event.Owner.Login, Email: event.Owner.Email},
		}

		events = append(events, eventDTO)
	}

	return events, nil
}

func (s *EventService) RemoveEvent(eventID uuid.UUID) error {
	err := s.repo.RemoveEvent(eventID)
	if err != nil {
		return err
	}

	return nil
}

func (s *EventService) UpdateEvent(eventID uuid.UUID, userData *types.UserDataCtx, data *UpdateEventRequest) error {
	err := s.repo.UpdateEvent(eventID, userData, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *EventService) SaveFile(fileHeader *multipart.FileHeader, SaveUploadedFile func(file *multipart.FileHeader, dst string, perm ...fs.FileMode) error) (*string, error) {
	var coverURL string
	if fileHeader != nil {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Ошибка открытия файла", err)
		}
		defer file.Close()

		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
		if !allowedExts[ext] {
			return nil, global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Неподдерживаемый формат файла", err)
		}

		if fileHeader.Size > 5<<20 { // 5MB
			return nil, global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Файл слишком большой (макс. 5МБ)", err)
		}

		filename := uuid.New().String() + ext
		savePath := filepath.Join("uploads", "covers", filename)

		if err := SaveUploadedFile(fileHeader, savePath); err != nil {
			return nil, global_utils.NewAppErrorWithErr(http.StatusBadRequest, "Ошибка сохранения", err)
		}

		coverURL = "/uploads/covers/" + filename
	}

	return &coverURL, nil
}
