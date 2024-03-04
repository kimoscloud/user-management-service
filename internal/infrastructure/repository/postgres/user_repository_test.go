package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kimoscloud/user-management-service/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetById(t *testing.T) {
	sqlDB, db, mock := test.DbMock(t)
	defer sqlDB.Close()

	implObj := NewUserRepository(db)
	users := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "hash"}).
		AddRow("1", "test first name", "test last name", "testemail@testemail.com", "hash")
	mock.ExpectQuery(
		"^SELECT (.+) FROM \"Users\" WHERE id = \\$1 AND \"Users\"\\.\"deleted_at\" IS NULL ORDER BY \"Users\"\\.\"id\" LIMIT 1",
	).WillReturnRows(users)
	user, err := implObj.GetByID("1")
	assert.Nil(t, err)
	assert.Equal(t, "test first name", user.FirstName)
	assert.Nil(t, mock.ExpectationsWereMet())
}
