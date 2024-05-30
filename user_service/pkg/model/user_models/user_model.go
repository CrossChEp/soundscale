package user_models

type UserModel struct {
	Id          string `bson:"_id"`
	Nickname    string `bson:"nickname"`
	Email       string `bson:"email"`
	PhoneNumber string `bson:"phone_number"`
	Password    string `bson:"password"`
	State       string `bson:"state"`
}
