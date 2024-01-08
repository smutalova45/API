package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"main.go/models"
)

type usersRepo struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) usersRepo {
	return usersRepo{
		DB: db,
	}
}

func (u usersRepo) Insert(users models.Users) (string, error) {
	id := uuid.New()
	if _, err := u.DB.Exec(`insert into users values($1, $2, $3, $4 )`, id, users.Firstname, users.Email, users.Phone); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (u usersRepo) GetList() ([]models.Users, error) {
	rows, err := u.DB.Query(`select * from users `)
	if err != nil {
		return nil, err
	}
	u1 := []models.Users{}
	for rows.Next() {
		user1 := models.Users{}
		if err = rows.Scan(&user1.Id, &user1.Firstname, &user1.Email, &user1.Phone); err != nil {
			return nil, err
		}
		u1 = append(u1, user1)
	}
	return u1, nil

}

func (u usersRepo) Update(us models.Users) error {
	_, err := u.DB.Exec(`update users set phone = $1 where id =$2`, us.Phone, us.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u usersRepo) Delete(id uuid.UUID) error {
	if _, err := u.DB.Exec(`delete from orders where user_id =$1`, id); err != nil {
		return err
	}
	if _, err := u.DB.Query(`delete from users where id =$1`, id); err != nil {
		return err
	}
	return nil

}

func (u usersRepo) GetById(id uuid.UUID) (models.Users, error) {
	users := models.Users{}
	if err := u.DB.QueryRow(`select id, firstname,email,phone from users where id=$1`, id).Scan(
		&users.Id,
		&users.Firstname,
		&users.Email,
		&users.Phone,
	); err != nil {
		return models.Users{}, err
	}
	return users, nil
}
