package user

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/model"
)

// Get - Get user for id
func Get(db *db.DB, id string) (*model.User, error) {
	var user *model.User

	err := db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT u.id_user,
							u.login,
							u.email, 
							u.name, 
							u.surname, 
							u.phone, 
							u.website_url, 
							u.image_url, 
							u.description, 
							u.date_of_birth, 
							u.date_last_logged, 
							u.date_created, 
							u.date_last_modified FROM User u
							WHERE u.id_user = ?;`, id)

		var u model.User

		err := row.Scan(&u.ID,
			&u.Login,
			&u.Email,
			&u.Name,
			&u.Surname,
			&u.Phone,
			&u.WebsiteURL,
			&u.ImageURL,
			&u.Description,
			&u.DateOfBirth,
			&u.DateLastLogged,
			&u.DateCreated,
			&u.DateLastModified)

		if err != nil {
			switch err {
			case sql.ErrNoRows:
				log.Error(fmt.Sprintf("No rows found for id: %s", id))
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		user = &u

		return err
	})

	return user, err
}

// GetAll - Get all users
func GetAll(db *db.DB) ([]model.User, error) {
	var users []model.User

	err := db.WithTransaction(func(tx *sql.Tx) error {
		rows, err := tx.Query(`SELECT u.id_user,
								u.login,
								u.email, 
								u.name, 
								u.surname, 
								u.phone, 
								u.website_url, 
								u.image_url, 
								u.description, 
								u.date_of_birth, 
								u.date_last_logged, 
								u.date_created, 
								u.date_last_modified FROM User u;`)
		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %v", err))
			}
			return err
		}

		for rows.Next() {
			var user model.User

			err := rows.Scan(&user.ID,
				&user.Login,
				&user.Email,
				&user.Name,
				&user.Surname,
				&user.Phone,
				&user.WebsiteURL,
				&user.ImageURL,
				&user.Description,
				&user.DateOfBirth,
				&user.DateLastLogged,
				&user.DateCreated,
				&user.DateLastModified)

			if err != nil {
				switch err {
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			users = append(users, user)
		}

		return nil
	})

	return users, err
}

// Insert - inserts new user
func Insert(db *db.DB, user *model.User) (*model.User, error) {
	err := db.WithTransaction(func(tx *sql.Tx) error {
		stmt, err := tx.Prepare(`INSERT INTO user (id_user,
								  email, 
								  password_fail_attempts, 
								  login, 
								  name, 
								  surname, 
								  phone, 
								  website_url, 
								  image_url, 
								  description, 
								  date_of_birth, 
								  date_last_logged, 
								  date_last_modified) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {

		}

		_, err = stmt.Exec(user.ID,
			user.Email,
			user.Login,
			user.Name,
			user.Surname,
			user.Phone,
			user.WebsiteURL,
			user.ImageURL,
			user.Description,
			user.DateOfBirth,
			user.DateLastLogged,
			user.DateCreated)

		// lastInserted := res

		return nil
	})

	return user, err
}

// Update - updates user's details
func Update(db *db.DB, update *model.User) (*model.User, error) {
	// TODO: implement
	return nil, nil
}

// Delete - deletes user
func Delete(db *db.DB, id string) (bool, error) {
	err := db.WithTransaction(func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("DELETE FROM user WHERE id_user = ?")

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		res, err := stmt.Exec(id)

		if err != nil {
			switch err {
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		affected, _ := res.RowsAffected()
		log.Info(fmt.Sprintf("User with id: %s deleted, affected rows: %d", id, affected))

		return nil
	})

	return true, err
}
