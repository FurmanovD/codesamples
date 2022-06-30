package main

import (
	"github.com/FurmanovD/codesamples/interfaces-with-providers/internal/pkg/service"
	"github.com/FurmanovD/codesamples/interfaces-with-providers/pkg/infra/db/objrepository"
)

func main() {

	repoProviders := &service.RepoProviders{
		Repo: objrepository.NewRepository(),
	}

	serviceLocator := service.NewServiceLocator()
	// V2, with Providers
	serviceLocator.ObjectFinderProvider = repoProviders
	serviceLocator.ObjectKillerProvider = repoProviders
	serviceLocator.ObjectStorerProvider = repoProviders
	serviceLocator.ObjectUpdaterProvider = repoProviders

}
