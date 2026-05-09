package domain

type EventRepository interface {
	Save(event *Event) error
	GetAll() ([]Event, error)
	GetById(id string) (*Event, error)
	Delete(id string) error
	Update(id string, data Event) error
	RegisterUserForEvent(event Event, userId int64) error
	GetRegistrationsForEvent(event Event) ([]int64, error)
	CancelUserRegistrationForEvent(event Event, userId int64) error
}
