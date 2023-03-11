package domain

type Users struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name"`
}

type UserResponse struct {
	ID       *string `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string  `bson:"email" json:"email"`
	Password string  `bson:"password" json:"password"`
	Name     string  `bson:"name" json:"name"`
}
