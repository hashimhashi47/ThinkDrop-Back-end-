package domain

type InterestRepo interface {
	GetAll(model interface{}) error
	FindBy(model interface{}, Query string, Any interface{}) error
	Save(model interface{}) error
	UpdateUserInterests(user interface{}, subIDs []uint) error
}
