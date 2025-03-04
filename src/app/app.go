package app

type LikeYouParameters struct {
	RecipientId     string
	PaginationToken string
}

type Liker struct {
	ActorID   string
	TimeStamp uint64
}

type LikedYouList struct {
	Likers              *[]Liker
	NextPaginationToken string
}

type RecipientService interface {
	ListLikedYou(likeYouParams *LikeYouParameters) (*LikedYouList, error)
	ListNewLikedYou(likeYouParams *LikeYouParameters) (*LikedYouList, error)
	CountLikedYou(recipientId string) (int64, error)
	PutDecision(actorId string, recipientId string, liked bool) (bool, error)
}
