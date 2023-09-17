package model

type Repository interface {
	CreateUser(*User)
}

func (db *Database) CreateUser(userDetails *User) {

}
