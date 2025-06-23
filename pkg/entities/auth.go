package entities

type Auth struct {
	TokenType string `bson:"token_type"`
	Token     string `bson:"token"`
	UserUuid  string `bson:"user_uuid"`
	CreatedAt string `bson:"created_at"`
}
