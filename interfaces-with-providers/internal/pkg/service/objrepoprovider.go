package service

import (
	"github.com/FurmanovD/codesamples/interfaces-with-providers/pkg/infra/db/model"
	"github.com/FurmanovD/codesamples/interfaces-with-providers/pkg/infra/db/objrepository"
)

type ObjFinderProvider interface {
	GetObjFinder() ObjFinder
}

type ObjStorerProvider interface {
	GetObjStorer() ObjStorer
}

type ObjUpdaterProvider interface {
	GetObjUpdater() ObjUpdater
}

type ObjKillerProvider interface {
	GetObjKiller() ObjKiller
}

type RepoProviders struct {
	Repo objrepository.Repository

	// NoSQLRepo nosqlobjrepo.Finder

	// ?
	bothReposKiller
}

// THIS IS THE correct place to split different types of providers is it is required
func (r *RepoProviders) GetObjFinder() ObjFinder {
	return ObjFinder(r.Repo)
	// return ObjFinder(r.NoSQLRepo)

}

func (r *RepoProviders) GetObjStorer() ObjStorer {
	return ObjStorer(r.Repo)
}

func (r *RepoProviders) GetObjUpdater() ObjUpdater {
	return ObjUpdater(r.Repo)
}

func (r *RepoProviders) GetObjKiller() ObjKiller {
	return ObjKiller(r.Repo)
	// return ObjKiller(r.bothReposKiller)
}

// both repos killer
type bothReposKiller struct {
	Repo objrepository.Repository
	// NoSQLRepo nosqlobjrepo.Finder
}

func (rk *bothReposKiller) DeleteObj(*model.Object) error {
	return nil
}
