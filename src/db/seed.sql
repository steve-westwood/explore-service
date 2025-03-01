CREATE TABLE Decisions (
    DecisionId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    RecipientId UUID NOT NULL,
    ActorId UUID NOT NULL,
    Timestamp BIGINT NOT NULL,
    Liked BOOLEAN NOT NULL,
    UNIQUE (RecipientId,ActorId)
);