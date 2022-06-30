package service

// BAD ! should be an interface as well!
type ServiceLocator struct {
	ObjectFinderProvider  ObjFinderProvider
	ObjectStorerProvider  ObjStorerProvider
	ObjectUpdaterProvider ObjUpdaterProvider
	ObjectKillerProvider  ObjKillerProvider
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{}
}
