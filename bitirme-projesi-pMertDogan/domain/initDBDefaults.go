package domain

import (
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/store"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/user"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/config"
)

func InitDBDefaults(cfg *config.Config) {

	user.Repo().CreateAdminIfNotExist(cfg)
	store.Repo().CreateStoreIfNotExist(cfg)
}