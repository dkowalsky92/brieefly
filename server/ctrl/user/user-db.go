package user

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/crypto"
	"github.com/dkowalsky/brieefly/ctrl/user/body"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/model"
)

// DbGet - Get user for id
func DbGet(db *db.DB, id string) (*model.User, *err.Error) {
	var user *model.User

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
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
			return db.HandleError(err)
		}

		user = &u

		return db.HandleError(err)
	})

	return user, err
}

// DbExists - Check if user exists, returns users id or nil
func DbExists(_db *db.DB, email, plainPass string) db.NullString {
	var id db.NullString
	_ = _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT u.id_user, u.password FROM User u
							WHERE u.email = ?`, email)
		var hash string

		err := row.Scan(&id, &hash)
		if err != nil {
			return nil
		}

		matches := crypto.CompareHash(plainPass, hash)

		if matches != true {
			id.String = ""
			id.Valid = false
		}

		return nil
	})

	return id
}

// DbGetAll - Get all users
func DbGetAll(db *db.DB) ([]model.User, *err.Error) {
	var users []model.User

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
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
			return db.HandleError(err)
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
				return db.HandleError(err)
			}

			users = append(users, user)
		}

		return db.HandleError(err)
	})

	return users, err
}

// DbInsert - inserts new user
func DbInsert(db *db.DB, user *model.User) (*model.User, *err.Error) {

	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
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
			return db.HandleError(err)
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

// DbChangePassword -
func DbChangePassword(db *db.DB, userid, password string) *err.Error {
	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		stmt, err := tx.Prepare("UPDATE User SET password = ? WHERE id_user = ?")
		if err != nil {
			return db.HandleError(err)
		}
		hash, hErr := crypto.Hash(password)
		if hErr != nil {
			return db.HandleError(hErr)
		}
		res, err := stmt.Exec(hash, userid)
		if err != nil {
			return db.HandleError(err)
		}

		affected, _ := res.RowsAffected()
		log.Info(fmt.Sprintf("Password updated, affected rows: %d", affected))

		return nil
	})

	return err
}

// DbUpdate - updates user's details
func DbUpdate(_db *db.DB, update body.UserUpdate, userid string) *err.Error {
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		condition := fmt.Sprintf("id_user = '%s'", userid)
		updateStmt := db.UpdateStmt(tx, update, "User", &condition)
		log.Debug(updateStmt.Stmt)
		_, sqlErr := updateStmt.Stmt.Exec(updateStmt.Args...)
		if sqlErr != nil {
			return _db.HandleError(sqlErr)
		}

		return nil
	})

	return err
}

// DbDelete - deletes user
func DbDelete(db *db.DB, id string) (bool, *err.Error) {
	err := db.WithTransaction(func(tx *sql.Tx) *err.Error {
		stmt, err := tx.Prepare("DELETE FROM user WHERE id_user = ?")

		if err != nil {
			return db.HandleError(err)
		}

		res, err := stmt.Exec(id)

		if err != nil {
			return db.HandleError(err)
		}

		affected, _ := res.RowsAffected()
		log.Info(fmt.Sprintf("User with id: %s deleted, affected rows: %d", id, affected))

		return nil
	})

	return true, err
}
