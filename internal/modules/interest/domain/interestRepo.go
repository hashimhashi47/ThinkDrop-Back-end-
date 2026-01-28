package domain

type InterestRepo interface {
	GetAll(model interface{}) error
}
