package main

import (
	"github.com/FurmanovD/codesamples/interfaces/pkg/infra/db/objrepository"
	"github.com/FurmanovD/codesamples/interfaces/pkg/service"
)

func main() {

	repo := objrepository.NewRepository()

	serviceLocator := service.NewServiceLocator()
	// specificKiller :=
	// V1, with no Providers
	serviceLocator.ObjectFinder = repo // or set here another finder, e.g. in a NoSQL DB
	serviceLocator.ObjectKiller = repo // or set here another object to kill in different places at once.
	serviceLocator.ObjectStorer = repo
	serviceLocator.ObjectUpdater = repo

}
