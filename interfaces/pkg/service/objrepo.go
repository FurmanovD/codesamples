package service

import "github.com/FurmanovD/codesamples/interfaces/pkg/infra/db/model"

type ObjFinder interface {
	GetObj(id int) (*model.Object, error)
	ListObj(conditions interface{}) ([]*model.Object, error)
}

type ObjStorer interface {
	StoreObj(*model.Object) error
}

type ObjUpdater interface {
	UpdateObj(*model.Object) error
}

type ObjKiller interface {
	DeleteObj(*model.Object) error
}
