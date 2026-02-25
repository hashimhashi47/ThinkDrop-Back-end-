package domain

type AdminRepo interface {
	GetCounts(Data string) (int64, error)
	FindTopUser(data interface{}) error
	FindbyArgs(data interface{}, query string, args interface{}) error
	FindAll(model interface{}, limit, offset int, preload1, preload2 string) error
	Find(model interface{}) error
	FindWithoutPreload(model interface{}, limit, offset int) error
	Count(model interface{}) (int64, error)
	UpdateColumn(model interface{}, query string, id interface{}, column string, value interface{}) error
	FindReportedPosts(model interface{}, limit, offset int, preloads ...string) error
	FindWithPreload(model interface{}, limit, offset int, preload string) error
	Delete(model interface{}, query string, arg interface{}) error
	Insert(model interface{}) error
	DeletePostWithRelations(postID uint) error
}
