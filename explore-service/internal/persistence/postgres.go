package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/steve-westwood/explore-service/explore-service/internal/app"
)

type RecipientService struct {
	db *sql.DB
}

func NewDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	return db
}

func NewRecipientService(db *sql.DB) *RecipientService {
	return &RecipientService{db: db}
}

func (s *RecipientService) ListLikedYou(recipientId string) (*app.LikedYouList, error) {
	var likers []app.Liker
	rows, err := s.db.Query(`SELECT ActorId, Timestamp FROM Decisions
						WHERE RecipientId = $1 AND Liked = TRUE
						ORDER BY Timestamp DESC;`, recipientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var liker app.Liker
		if err := rows.Scan(&liker.ActorID, &liker.TimeStamp); err != nil {
			return nil, err
		}
		likers = append(likers, liker)
	}

	return &app.LikedYouList{
		Likers: &likers,
	}, nil
}

func (s *RecipientService) ListNewLikedYou(recipientId string) (*app.LikedYouList, error) {
	var likers []app.Liker
	rows, err := s.db.Query(`SELECT d.ActorId, d.Timestamp FROM Decisions d
						LEFT JOIN Decisions r ON d.RecipientId = r.ActorId AND d.ActorId = r.RecipientId
						WHERE d.RecipientID=$1 AND d.Liked = TRUE AND r.DecisionId IS NULL
						ORDER BY d.Timestamp DESC;`, recipientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var liker app.Liker
		if err := rows.Scan(&liker.ActorID, &liker.TimeStamp); err != nil {
			return nil, err
		}
		likers = append(likers, liker)
	}

	return &app.LikedYouList{
		Likers: &likers,
	}, nil
}

func (s *RecipientService) CountLikedYou(recipientId string) (int64, error) {
	row := s.db.QueryRow(`SELECT COUNT(*) FROM Decisions
						WHERE RecipientId = $1 AND Liked = TRUE;`, recipientId)
	var count int64
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (s *RecipientService) PutDecision(actorId string, recipientId string, liked bool) (bool, error) {
	reciprocated := false
	if liked {
		likedRow := s.db.QueryRow(`SELECT Liked FROM Decisions WHERE ActorId = $1 AND RecipientId = $2;`, recipientId, actorId)
		if err := likedRow.Scan(&reciprocated); err != nil && err != sql.ErrNoRows {
			log.Fatalf("select statement error: %s", err)
		}
	}
	_, err := s.db.Exec(`INSERT INTO Decisions (RecipientId, ActorId, Timestamp, Liked)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (RecipientId, ActorId) 
			DO UPDATE SET Timestamp = EXCLUDED.Timestamp, Liked = EXCLUDED.Liked;`, recipientId, actorId, time.Now().Unix(), liked)
	if err != nil {
		log.Fatalf("insert error: %s", err)
	}
	return reciprocated, nil
}
