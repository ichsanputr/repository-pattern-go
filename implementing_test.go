package implement_repository_pattern

import (
	"fmt"
	"testing"

	r "github.com/ichsanputr/repository-pattern-go/repository"

	_ "github.com/go-sql-driver/mysql"
)

func TestRepositoryPattern(t *testing.T) {
	db := Connection()

	repostudent := r.NewRepositoryStudentImplement(db)
	a, err := repostudent.FindById(1)	

	if err != nil {
		panic(err)
	}

	fmt.Println(a)
}