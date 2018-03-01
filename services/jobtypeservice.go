package services

import (
	"github.com/ducngocvnk57/circleci-demo/app"
	"github.com/ducngocvnk57/circleci-demo/models"
)

var db = app.Db()

func GetAllJobType() []models.JobType {
	var jobtypes []models.JobType
	db.Find(&jobtypes)
	return jobtypes
}
