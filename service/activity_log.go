package service

import (
	"co-studio-e-commerce/repo"
)

type activityLog struct {
	repo repo.IRepo
}

type IActivityLog interface {
}

func NewActivityLog(repo repo.IRepo) *activityLog {
	return &activityLog{repo: repo}
}
