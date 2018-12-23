package project

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/model"
)

// GetAllForUserID - Get all projects for user id
func GetAllForUserID(db *db.DB, id string) ([]model.Project, error) {
	projects := []model.Project{}
	var err error

	db.WithTransaction(func(tx *sql.Tx) error {
		rows, qerr := tx.Query(`SELECT p.id_project,
									   p.name,
									   p.type,
									   o.salary,
									   p.date_created,
									   o.date_deadline, 
									   p.subpage_count, 
									   p.overall_progress, 
									   p.image_url, 
									   c.name, 
									   c.image_url, 
									   s.name FROM Project p
									   INNER JOIN Client_project up ON p.id_project = up.id_project
									   INNER JOIN Offer o ON p.id_project = o.id_project 
									   INNER JOIN Agency a ON o.id_company = a.id_company 
									   INNER JOIN Status s ON p.id_status = s.id_status 
									   INNER JOIN Company c ON c.id_company = a.id_company
									   WHERE up.id_user = ?`, id)
		err = qerr
		if err != nil {
			log.Error(fmt.Sprintf("Error occurred: %+v", err))
			return err
		}
		for rows.Next() {
			var p model.Project
			var o model.Offer
			var c model.Company
			var s model.ProjectStatus
			err = rows.Scan(&p.ID,
				&p.Name,
				&p.Type,
				&o.Salary,
				&p.DateCreated,
				&o.DateDeadline,
				&p.SubpageCount,
				&p.OverallProgress,
				&p.ImageURL,
				&c.Name,
				&c.ImageURL,
				&s.Name)

			if err != nil {
				switch err {
				case sql.ErrNoRows:
					log.Error(fmt.Sprintf("No rows found"))
				default:
					log.Error(fmt.Sprintf("Error occurred: %+v", err))
				}
				return err
			}

			projects = append(projects, p)
		}

		return nil
	})

	return projects, err
}

// GetNameForID - gets the project's name for project id
func GetNameForID(_db *db.DB, id string) (*db.NullString, error) {
	var name db.NullString
	var err error

	_db.WithTransaction(func(tx *sql.Tx) error {
		row := tx.QueryRow(`SELECT name FROM Project WHERE id_project = ?;`, id)
		err = row.Scan(&name)

		if err != nil {
			switch err {
			case sql.ErrNoRows:
				log.Error(fmt.Sprintf("No rows found"))
			default:
				log.Error(fmt.Sprintf("Error occurred: %+v", err))
			}
			return err
		}

		return nil
	})

	return &name, err
}
