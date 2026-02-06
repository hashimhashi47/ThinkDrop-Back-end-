package domain

type RewardRepo interface {
	Find(model interface{}, query string, args interface{}, preloads ...string) error
	Create(model interface{}) error
	FindAll(model interface{}) error
	Save(model interface{}) error
	Update(model interface{}, query string, arg interface{}, updates map[string]interface{}) error
}
