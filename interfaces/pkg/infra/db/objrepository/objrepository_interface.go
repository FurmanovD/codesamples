package objrepository

import "github.com/FurmanovD/codesamples/interfaces/pkg/infra/db/model"

type Repository interface {
	GetObj(id int) (*model.Object, error)
	StoreObj(*model.Object) error
	UpdateObj(*model.Object) error
	ListObj(conditions interface{}) ([]*model.Object, error)
	DeleteObj(*model.Object) error
}
