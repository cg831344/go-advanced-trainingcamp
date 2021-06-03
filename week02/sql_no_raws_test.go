package week02

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"testing"
)

var SqlNoRowError = errors.New("sql no row ")

func getRow(id int) error {
	db, mock, err := sqlmock.New()
	if err != nil {
		return err
	}
	defer db.Close()

	rs := sqlmock.NewRows([]string{"id", "title"})

	mock.ExpectQuery("SELECT (.+) FROM articles WHERE id = ?").
		WithArgs(id).
		WillReturnRows(rs)

	rows, err := db.Query("SELECT (.+) FROM articles WHERE id = ?", id)
	if err != nil {
		return err
	}

	defer func() {
		_ = rows.Close()
	}()

	if !rows.Next() {

		return SqlNoRowError
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		return err
	}
	return nil
}

func do() error {
	id := 5
	err := getRow(id)
	if err != nil {
		if errors.As(err,&SqlNoRowError) {
			return errors.Wrap(err,"数据articles不存在")
		}
		return err
	}
	return nil
}


func TestMockQuery(t *testing.T) {
	err:= do()
	if err != nil {
		fmt.Printf("%v\n",err)
		fmt.Printf("%+v\n",err)

	}

}
