package repository

import "github.com/ichsanputr/repository-pattern-go/model"

// define method which will you create for your table
// example i create 1 method FindById
// FindById can get all data associated with id in input
type StudentRepository interface{
	FindById(id int32) (model.Student,error)
}