package user

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/model"
)

// Get - Get user for id
func Get(db *db.DB, id string) (*model.User, error) {
	var user model.User
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow("SELECT * FROM user WHERE id_user = ?", id)
		err = row.Scan(&user.ID,
			&user.Login,
			&user.Password,
			&user.PasswordFailAttempts,
			&user.Email,
			&user.Name,
			&user.Surname,
			&user.Phone,
			&user.WebsiteURL,
			&user.ImageURL,
			&user.Description,
			&user.DateOfBirth,
			&user.DateLastLogged,
			&user.DateCreated)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				log.Error(fmt.Sprintf("No rows found for id: %s", id))
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		return nil
	})

	return &user, err
}

// GetAll - Get all users
func GetAll(db *db.DB) ([]model.User, error) {
	var users []model.User
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		rows, qerr := tx.Query("SELECT * FROM user")
		err = qerr
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
				&user.Password,
				&user.PasswordFailAttempts,
				&user.Email,
				&user.Name,
				&user.Surname,
				&user.Phone,
				&user.Phone,
				&user.WebsiteURL,
				&user.ImageURL,
				&user.Description,
				&user.DateOfBirth,
				&user.DateLastLogged,
				&user.DateCreated)
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
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		stmt, ierr := tx.Prepare("INSERT INTO user (id_user, email, password_fail_attempts, login, name, surname, phone, website_url, image_url, description, date_of_birth, date_last_logged, date_last_modified) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		err = ierr
		if err != nil {

		}

		_, err = stmt.Exec(user.ID,
			user.Email,
			user.PasswordFailAttempts,
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
	var err error
	db.WithTransaction(func(tx *sql.Tx) error {
		stmt, perr := tx.Prepare("DELETE FROM user WHERE id_user = ?")
		err = perr
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

// CREATE TABLE User (
//     id_user int NOT NULL AUTO_INCREMENT,
//     login varchar(20) NULL,
//     password varchar(20) NOT NULL,
//     password_fail_attempts int NOT NULL DEFAULT 0,
//     email varchar(75) NOT NULL,
//     name varchar(20) NULL,
//     surname varchar(20) NULL,
//     phone varchar(14) NULL,
//     website_url varchar(300) NULL,
//     image_url varchar(200) NULL,
//     description varchar(300) NULL,
//     date_of_birth date NULL,
//     date_last_logged date NULL,
//     date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
//     CONSTRAINT User_pk PRIMARY KEY (id_user)
// );
