package repository_test

import (
	"azl-vote-api/internal/domain/entities"
	"azl-vote-api/internal/infrastructure/repository"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewRepositoryMock() (*sql.DB, sqlmock.Sqlmock) {
	dbConnection, mock, error := sqlmock.New()
	if error != nil {
		log.Fatalf("Ocorreu uma falha %s ao estabelece a conexação com o mock do banco de dados!", error)
	}

	return dbConnection, mock
}

func TestShouldBeCreateContest(test *testing.T) {
	db, mock := NewRepositoryMock()
	contestRepository := repository.NewContestRepository(db)

	input := entities.CreateContestEntity(time.Now(), time.Now().Add(time.Hour*24))

	dateLayout := "2006-01-02 15:04:05 -07:00"
	parsedInitialDate := input.InitialDate.Format(dateLayout)
	parsedFinalDate := input.FinalDate.Format(dateLayout)

	mock.ExpectPrepare("INSERT INTO Contests").ExpectExec().WithArgs(input.Id, parsedInitialDate, parsedFinalDate, input.Active).WillReturnResult(sqlmock.NewResult(1, 1))

	error := contestRepository.Create(input)
	if error != nil {
		test.Errorf("Ocorreu um erro inesperado ao inserir o contest: %s", error)
	}

	error = mock.ExpectationsWereMet()
	if error != nil {
		test.Errorf("Asserts falharam: %s", error)
	}
}

func TestShouldBeFindAllContests(test *testing.T) {
	db, mock := NewRepositoryMock()
	contestRepository := repository.NewContestRepository(db)

	input := entities.CreateContestEntity(time.Now(), time.Now())
	//dateLayout := "2006-01-02 15:04:05 -07:00"

	mockedRows := mock.NewRows([]string{"id", "initial_date", "final_date", "active"}).
		AddRow(input.Id, input.InitialDate, input.FinalDate, input.Active)
		//AddRow(uuid.Generate().String(), "2023-01-02 15:04:05.0000000 -08:00", "2023-01-02 15:04:05.0000000 -08:00", false)
	mock.ExpectQuery("SELECT id, initial_date, final_date, active FROM Contests").WillReturnRows(mockedRows)

	contests, error := contestRepository.FindAll()
	if error != nil {
		test.Errorf("Ocorreu um erro inesperado ao buscar os constests: %s", error)
	}

	error = mock.ExpectationsWereMet()
	if error != nil {
		test.Errorf("Asserts falharam: %s", error)
	}

	assert.Equal(test, input, contests[0])
}

func TestShouldBeDeleteContestById(test *testing.T) {
	db, mock := NewRepositoryMock()
	contestRepository := repository.NewContestRepository(db)

	input := entities.CreateContestEntity(time.Now(), time.Now())

	mock.NewRows([]string{"id", "initial_date", "final_date", "active"}).
		AddRow(input.Id, input.InitialDate, input.FinalDate, input.Active)
	mock.ExpectPrepare("DELETE FROM Contests WHERE id=@p1").
		ExpectExec().
		WithArgs(input.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	error := contestRepository.DeleteById(input)
	if error != nil {
		test.Errorf("Ocorreu um erro inesperado ao deletar o constest: %s", error)
	}

	error = mock.ExpectationsWereMet()
	if error != nil {
		test.Errorf("Asserts falharam: %s", error)
	}
}

func TestUnexistsIdErrorDeleteContestById(test *testing.T) {
	//pending
}
