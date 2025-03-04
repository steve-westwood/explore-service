package persistence

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/steve-westwood/explore-service/src/app"
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

	mock.ExpectQuery("SELECT DecisionId, ActorId, Timestamp FROM Decisions").
		WithArgs("test-recipient-id", 5).
		WillReturnRows(sqlmock.NewRows([]string{"DecisionId", "ActorId", "Timestamp"}))

	recipientService := NewRecipientService(db)

	likedYouList, err := recipientService.ListLikedYou(&app.LikeYouParameters{
		RecipientId: "test-recipient-id",
	})
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

	mock.ExpectQuery("SELECT DecisionId, ActorId, Timestamp FROM Decisions").
		WithArgs("test-recipient-id", 5).
		WillReturnRows(sqlmock.NewRows([]string{"DecisionId", "ActorId", "Timestamp"}).
			AddRow("test-decision-id", "test-liker-id", "123"))

	recipientService := NewRecipientService(db)

	likedYouList, err := recipientService.ListLikedYou(&app.LikeYouParameters{
		RecipientId: "test-recipient-id",
	})
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

	if likedYouList.NextPaginationToken != "" {
		t.Errorf("Expected no pagination token , got '%s'", likedYouList.NextPaginationToken)
	}
}

func TestListLikeYou_ReturnsListOfLikersWithPagination(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT DecisionId, ActorId, Timestamp FROM Decisions").
		WithArgs("test-recipient-id", 5).
		WillReturnRows(sqlmock.NewRows([]string{"DecisionId", "ActorId", "Timestamp"}).
			AddRow("test-decision-id1", "test-liker-id1", "123").
			AddRow("test-decision-id2", "test-liker-id2", "456").
			AddRow("test-decision-id3", "test-liker-id3", "789").
			AddRow("test-decision-id4", "test-liker-id4", "101").
			AddRow("test-decision-id5", "test-liker-id5", "102"))

	recipientService := NewRecipientService(db)

	likedYouList, err := recipientService.ListLikedYou(&app.LikeYouParameters{
		RecipientId: "test-recipient-id",
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(*likedYouList.Likers) != 5 {
		t.Errorf("Expected 5 liker, got %v", len(*likedYouList.Likers))
	}

	if (*likedYouList.Likers)[4].ActorID != "test-liker-id5" {
		t.Errorf("Expected liker id to be 'test-liker-id5', got '%s'", (*likedYouList.Likers)[4].ActorID)
	}

	if (*likedYouList.Likers)[4].TimeStamp != 102 {
		t.Errorf("Expected timestamp to be '102', got '%d'", (*likedYouList.Likers)[4].TimeStamp)
	}

	if likedYouList.NextPaginationToken != "MTAyOnRlc3QtZGVjaXNpb24taWQ1" {
		t.Errorf("Expected pagination token to be 'MTAyOnRlc3QtZGVjaXNpb24taWQ1', got '%s'", likedYouList.NextPaginationToken)
	}

}

func TestListLikeYou_ReturnsListOfLikersSecondPage(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT DecisionId, ActorId, Timestamp FROM Decisions").
		WithArgs("test-recipient-id", 102, "test-decision-id5", 102, 5).
		WillReturnRows(sqlmock.NewRows([]string{"DecisionId", "ActorId", "Timestamp"}).
			AddRow("test-decision-id1", "test-liker-id6", "103"))

	recipientService := NewRecipientService(db)

	likedYouList, err := recipientService.ListLikedYou(&app.LikeYouParameters{
		RecipientId:     "test-recipient-id",
		PaginationToken: "MTAyOnRlc3QtZGVjaXNpb24taWQ1",
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(*likedYouList.Likers) != 1 {
		t.Errorf("Expected 1 liker, got %v", len(*likedYouList.Likers))
	}

	if (*likedYouList.Likers)[0].ActorID != "test-liker-id6" {
		t.Errorf("Expected liker id to be 'test-liker-id56', got '%s'", (*likedYouList.Likers)[0].ActorID)
	}

	if (*likedYouList.Likers)[0].TimeStamp != 103 {
		t.Errorf("Expected timestamp to be '103', got '%d'", (*likedYouList.Likers)[0].TimeStamp)
	}

	if likedYouList.NextPaginationToken != "" {
		t.Errorf("Expected no pagination token , got '%s'", likedYouList.NextPaginationToken)
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

func TestEncodingDecodingPaginationToken(t *testing.T) {

	decisionId := "test-id-123"
	timestamp := uint64(time.Now().Unix())
	encoded := GetPaginationToken(timestamp, decisionId)
	if encoded == "" {
		t.Error("Expected encoded token to not be empty")
	}

	decodedTimestamp, decodedDecisionId, err := GetPaginationDetailsFromToken(encoded)
	if err != nil {
		t.Errorf("Failed to decode token: %v", err)
	}

	if decodedDecisionId != decisionId {
		t.Errorf("Expected decisionId to be %s, got %s", decisionId, decodedDecisionId)
	}

	if decodedTimestamp != timestamp {
		t.Errorf("Expected timestamp to be %d, got %d", timestamp, decodedTimestamp)
	}
}
