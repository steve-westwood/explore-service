package app

type Liker struct {
	ActorID   string
	TimeStamp uint64
}

type LikedYouList struct {
	Likers              *[]Liker
	NextPaginationToken string
}

type RecipientService interface {
	ListLikedYou(recipientId string) (*LikedYouList, error)
	ListNewLikedYou(recipientId string) (*LikedYouList, error)
	CountLikedYou(recipientId string) (int64, error)
	PutDecision(actorId string, recipientId string, liked bool) (bool, error)
}
