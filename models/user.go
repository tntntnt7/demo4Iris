package models

type User struct {
	Id 			string	`bson:"_id"`
	Name		string	`bson:"Name"`
	Password	string	`bson:"Password"`
	Age 		int		`bson:"Age"`
	Gender		string	`bson:"Gender"`
}
