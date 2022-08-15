package repository

import (
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) Create(email string, password string, name string) (*User, error) {

	// query := "insert into user_table (user_email,user_password,user_name) values (?,?,?)"
	// result, err := r.db.Exec(
	// 	query,
	// 	email,
	// 	password,
	// 	name,
	// )
	user := &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	id := user.ID

	user, err := r.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// user := User{
	// 	Id:       int(id),
	// 	Email:    email,
	// 	Password: password,
	// 	Name:     name,
	// }

	return user, nil
}

func (r userRepositoryDB) GetAll() ([]User, error) {

	users := []User{}
	tx := r.db.Order("ID").Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r userRepositoryDB) Login() (*User, error) {

	return nil, nil
}

func (r userRepositoryDB) GetUserByID(userID uint) (*User, error) {

	user := &User{}
	tx := r.db.First(user, userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (r userRepositoryDB) GetUserByEmail(email string) (*User, error) {

	user := &User{}
	tx := r.db.First(user, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

/*

Exam gorm command
	- tx := db.Model(&Gender{}).Where("id = @myid", sql.Named("myid", id)).Updates(gender)
	- tx := db.Order("id").Find(&genders)
	- tx := db.First(&gender, id)
	- tx := db.First(&gender, "name = ?", name)
	- tx := db.Create(&gender)
	- tx := db.Order("id").Find(&genders)

*/
