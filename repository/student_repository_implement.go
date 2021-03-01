package repository

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/ichsanputr/repository-pattern-go/model"
)

// implement interface
type RepositoryStudentImplement struct{
	DB *sql.DB
	Err error
}

// create relation
func NewRepositoryStudentImplement(db *sql.DB) StudentRepository {
	return &RepositoryStudentImplement{DB:db}
}

func (r *RepositoryStudentImplement) FindById(id int32) (model.Student,error){
	rows, err := r.DB.Query("SELECT id,nama,school FROM student WHERE id = ?",id)
	
	if err != nil {
		r.Err = err
	}
	
	var data = model.Student{}

	if rows.Next(){
		rows.Scan(&data.Id,&data.Name,&data.School)
		r.Err = nil
	}else{
		r.Err = errors.New("Data for id "+strconv.Itoa(int(id))+" not found..")
	}

	return data,r.Err
}