package persistence

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func SetupMockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}

func TestListLikeYou_ReturnsEmptyListWhenNoDecisions(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ActorId, Timestamp FROM Decisions").
		WithArgs("test-recipient-id").
		WillReturnRows(sqlmock.NewRows([]string{"ActorId", "Timestamp"}))

	recipientService := NewRecipientService(db)

	likedYouList, err := recipientService.ListLikedYou("test-recipient-id")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(*likedYouList.Likers) != 0 {
		t.Errorf("Expected empty list, got %v likers", len(*likedYouList.Likers))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestListLikeYou_ReturnsListOfLikers(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ActorId, Timestamp FROM Decisions").
		WithArgs("test-recipient-id").
		WillReturnRows(sqlmock.NewRows([]string{"ActorId", "Timestamp"}).
			AddRow("test-liker-id", "123"))

	recipientService := NewRecipientService(db)

	likedYouList, err := recipientService.ListLikedYou("test-recipient-id")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(*likedYouList.Likers) != 1 {
		t.Errorf("Expected 1 liker, got %v", len(*likedYouList.Likers))
	}

	if (*likedYouList.Likers)[0].ActorID != "test-liker-id" {
		t.Errorf("Expected liker id to be 'test-liker-id', got '%s'", (*likedYouList.Likers)[0].ActorID)
	}

	if (*likedYouList.Likers)[0].TimeStamp != 123 {
		t.Errorf("Expected timestamp to be '123', got '%d'", (*likedYouList.Likers)[0].TimeStamp)
	}
}

func TestPutDecision_ActorlikeRecipientButNotReciprocated(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT Liked FROM Decisions").
		WithArgs("test-recipient-id", "test-actor-id").
		WillReturnRows(sqlmock.NewRows([]string{"Liked"}).AddRow(false))

	mock.ExpectExec("INSERT INTO Decisions").
		WithArgs("test-recipient-id", "test-actor-id", sqlmock.AnyArg(), true).
		WillReturnResult(sqlmock.NewResult(1, 1))

	recipientService := NewRecipientService(db)

	reciprocated, err := recipientService.PutDecision("test-actor-id", "test-recipient-id", true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if reciprocated {
		t.Errorf("Expected reciprocated to be false, got true")
	}
}

func TestPutDecision_ActorLikeRecipientAndReciprocated(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT Liked FROM Decisions").
		WithArgs("test-recipient-id", "test-actor-id").
		WillReturnRows(sqlmock.NewRows([]string{"Liked"}).AddRow(true))

	mock.ExpectExec("INSERT INTO Decisions").
		WithArgs("test-recipient-id", "test-actor-id", sqlmock.AnyArg(), true).
		WillReturnResult(sqlmock.NewResult(1, 1))

	recipientService := NewRecipientService(db)

	reciprocated, err := recipientService.PutDecision("test-actor-id", "test-recipient-id", true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !reciprocated {
		t.Errorf("Expected reciprocated to be true, got false")
	}
}
