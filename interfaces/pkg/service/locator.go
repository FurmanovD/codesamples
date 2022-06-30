package service

// BAD ! should be an interface as well!
type ServiceLocator struct {
	ObjectFinder  ObjFinder
	ObjectStorer  ObjStorer
	ObjectUpdater ObjUpdater
	ObjectKiller  ObjKiller
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{}
}
