package models

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	//ID       primitive.ObjectID `json:"id, omitempty" bson:"_id, omitempty"`
	// Username string `json:"username,omitempty" bson:"username,omitempty"`
	// Password string `json:"password,omitempty" bson:"password,omitempty"`
	// FullName string `json:"fullname,omitempty" bson:"fullname,omitempty"`
	//Email    string             `json:"size:50:not null;unique" bson:email"`
	//Password string             `json:size:60;not null" bson:"password`
	//CreatedAt time.Time          `json:"default":current_timestamp()" bson:"created_at"`
	//UpdatedAt time.Time          `json:"default":current_timestamp()" bson:"updated_at"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

// func (u *User) BeforeSave() error {
// 	hashedPassword, err := security.Hash(u.Password)
// 	if err != nil {
// 		return err
// 	}

// 	u.Password = string(hashedPassword)
// 	return nil
// }
