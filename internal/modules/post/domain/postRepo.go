package domain

type PostRepo interface {
	FindAnything(model interface{}, Query, Any interface{}) error
	Insert(model interface{}) error
	AllowPost(userID uint) (bool, error)
	FindAnyWithpreload(model interface{}, Query, AnyData interface{}, Preload string) error
	FindByUser(model interface{}, Query string, Any interface{}) error
	FindAll(model interface{}) error
}
