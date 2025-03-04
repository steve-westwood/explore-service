package persistence

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/steve-westwood/explore-service/src/app"
)

const (
	limit = 5
)

type RecipientService struct {
	db *sql.DB
}

func NewDB(datasource string) *sql.DB {
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewRecipientService(db *sql.DB) RecipientService {
	return RecipientService{db: db}
}

func (s *RecipientService) ListLikedYou(likeYouParams *app.LikeYouParameters) (*app.LikedYouList, error) {
	var likers []app.Liker
	var rows *sql.Rows
	var err2 error
	if likeYouParams.PaginationToken != "" {
		timestamp, decisionId, err := GetPaginationDetailsFromToken(likeYouParams.PaginationToken)
		if err != nil {
			return nil, err
		}

		rows, err2 = s.db.Query(`SELECT DecisionId, ActorId, Timestamp FROM Decisions
						WHERE RecipientId = $1 AND Liked = TRUE
						AND ((Timestamp = $2 AND DecisionId < $3) OR Timestamp < $4)
						ORDER BY Timestamp, DecisionId DESC
						LIMIT $5;`, likeYouParams.RecipientId, timestamp, decisionId, timestamp, limit)
	} else {
		rows, err2 = s.db.Query(`SELECT DecisionId, ActorId, Timestamp FROM Decisions
						WHERE RecipientId = $1 AND Liked = TRUE
						ORDER BY Timestamp, DecisionId DESC
						LIMIT $2;`, likeYouParams.RecipientId, limit)
	}

	if err2 != nil {
		return nil, err2
	}

	defer rows.Close()

	decisionId := ""
	lastTimestamp := uint64(0)

	for rows.Next() {
		var liker app.Liker
		if err := rows.Scan(&decisionId, &liker.ActorID, &liker.TimeStamp); err != nil {
			return nil, err
		}
		lastTimestamp = liker.TimeStamp
		likers = append(likers, liker)
	}
	likedYouList := &app.LikedYouList{
		Likers: &likers,
	}
	if len(likers) == limit {
		likedYouList.NextPaginationToken = GetPaginationToken(lastTimestamp, decisionId)
	}
	return likedYouList, nil
}

func (s *RecipientService) ListNewLikedYou(likeYouParams *app.LikeYouParameters) (*app.LikedYouList, error) {
	var likers []app.Liker
	var rows *sql.Rows
	var err2 error
	if likeYouParams.PaginationToken != "" {
		timestamp, decisionId, err := GetPaginationDetailsFromToken(likeYouParams.PaginationToken)
		if err != nil {
			return nil, err
		}

		rows, err2 = s.db.Query(`SELECT d.DecisionId, d.ActorId, d.Timestamp FROM Decisions d
						LEFT JOIN Decisions r ON d.RecipientId = r.ActorId AND d.ActorId = r.RecipientId
						WHERE d.RecipientID=$1 AND d.Liked = TRUE AND r.DecisionId IS NULL
						AND ((d.Timestamp = $2 AND d.DecisionId < $3) OR d.Timestamp < $4)
						ORDER BY d.Timestamp, d.DecisionId DESC
						LIMIT $5;`, likeYouParams.RecipientId, timestamp, decisionId, timestamp, limit)
	} else {
		rows, err2 = s.db.Query(`SELECT d.DecisionId, d.ActorId, d.Timestamp FROM Decisions d
						LEFT JOIN Decisions r ON d.RecipientId = r.ActorId AND d.ActorId = r.RecipientId
						WHERE d.RecipientID=$1 AND d.Liked = TRUE AND r.DecisionId IS NULL 
						ORDER BY d.Timestamp, d.DecisionId DESC
						LIMIT $2;`, likeYouParams.RecipientId, limit)
	}

	if err2 != nil {
		return nil, err2
	}

	defer rows.Close()

	decisionId := ""
	lastTimestamp := uint64(0)

	for rows.Next() {
		var liker app.Liker
		if err := rows.Scan(&decisionId, &liker.ActorID, &liker.TimeStamp); err != nil {
			return nil, err
		}
		lastTimestamp = liker.TimeStamp
		likers = append(likers, liker)
	}

	likedYouList := &app.LikedYouList{
		Likers: &likers,
	}
	if len(likers) == limit {
		likedYouList.NextPaginationToken = GetPaginationToken(lastTimestamp, decisionId)
	}
	return likedYouList, nil
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

func GetPaginationToken(timestamp uint64, recipientId string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d:%s", timestamp, recipientId)))
}

func GetPaginationDetailsFromToken(paginationToken string) (uint64, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(paginationToken)
	if err != nil {
		return 0, "", err
	}
	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return 0, "", errors.New("invalid pagination token")
	}
	timestamp, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, "", err
	}
	return uint64(timestamp), parts[1], nil
}
