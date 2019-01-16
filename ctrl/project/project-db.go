package project

import (
	"database/sql"
	"fmt"

	"github.com/brieefly/ctrl/project/body"
	"github.com/brieefly/db"
	"github.com/brieefly/err"
	"github.com/brieefly/model"
)

// DbGetProjectForID -
func DbGetProjectForID(_db *db.DB, id string) (*model.Project, *err.Error) {
	var project *model.Project
	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT u.id_project,
									u.name,
									u.type,
									u.description,
									u.url_name,
									u.overall_progress,
									u.language,
									u.budget_min,
									u.budget_max,
									u.subpage_count,
									u.image_url,
									u.date_deadline,
									u.date_created,
									u.date_last_modified 
									FROM Project u 
									WHERE u.id_project = ?;`, id)
		var p model.Project
		err := row.Scan(&p.ID,
			&p.Name,
			&p.Type,
			&p.Description,
			&p.NameURL,
			&p.OverallProgress,
			&p.Language,
			&p.BudgetMin,
			&p.BudgetMax,
			&p.SubpageCount,
			&p.ImageURL,
			&p.DateDeadline,
			&p.DateCreated,
			&p.DateLastModified)
		fmt.Printf("printing err - %+v", err)
		if err != nil {
			return _db.HandleError(err)
		}
		fmt.Printf("printing p - %+v", p)
		project = &p

		return nil
	})

	return project, err
}

// DbGetNameForID - gets the project's name for project id
func DbGetNameForID(_db *db.DB, id string) (*db.NullString, *err.Error) {
	var name db.NullString

	err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT name FROM Project WHERE id_project = ?;`, id)
		err := row.Scan(&name)

		if err != nil {
			return _db.HandleError(err)
		}

		return nil
	})

	return &name, err
}

// DbExists - Check if user exists, returns users id or nil
func DbExists(_db *db.DB, name, userid string) db.NullString {
	var id db.NullString
	_ = _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		row := tx.QueryRow(`SELECT u.id_project FROM Project p
							INNER JOIN Client_project cp ON cp.id_project = p.id_project
							WHERE p.name = ? AND cp.id_user = ?;`, name, userid)
		_ = row.Scan(&id)

		return nil
	})

	return id
}

// DbInsert -
func DbInsert(_db *db.DB, userid string, b *body.Body) (*model.Project, *err.Error) {
	var projectID string
	var finalProject *model.Project

	_err := _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		id := DbExists(_db, b.Project.Name, userid)
		if id.Valid {
			return _db.HandleTypedError(nil, db.ErrAlreadyExists)
		}

		status, sErr := DbGetStatusForName(_db, "Pending")
		if sErr != nil {
			return sErr
		}

		project := body.NewProject(b.Project, b.Project.CmsID.String, status.ID)
		projectStmt := db.InsertStmt(tx, project, "project")
		_, err := projectStmt.Stmt.Exec(projectStmt.Args...)
		if err != nil {
			return _db.HandleError(err)
		}

		clientProject := body.NewClientProject(body.ClientProjectBody{UserID: userid, ProjectID: project.ID})
		clientProjectStmt := db.InsertStmt(tx, clientProject, "client_project")
		_, err = clientProjectStmt.Stmt.Exec(clientProjectStmt.Args...)
		if err != nil {
			return _db.HandleError(err)
		}

		visualIdentity := body.NewVisualIdentity(b.VisualIdentity, project.ID)
		visualIdentityStmt := db.InsertStmt(tx, visualIdentity, "visual_identity")
		_, err = visualIdentityStmt.Stmt.Exec(visualIdentityStmt.Args...)
		if err != nil {
			return _db.HandleError(err)
		}

		for _, v := range b.Colors {
			color := body.NewColor(v, project.ID)
			colorStmt := db.InsertStmt(tx, color, "color")
			_, err = colorStmt.Stmt.Exec(colorStmt.Args...)
			if err != nil {
				return _db.HandleError(err)
			}
		}

		for _, v := range b.TargetGroups {
			targetGroup := body.NewTargetGroup(v, project.ID)
			targetGroupStmt := db.InsertStmt(tx, targetGroup, "target_group")
			_, err = targetGroupStmt.Stmt.Exec(targetGroupStmt.Args...)
			if err != nil {
				return _db.HandleError(err)
			}
		}

		for _, v := range b.CustomFeatures {
			customFeature := body.NewCustomFeature(v, project.ID)
			customFeatureStmt := db.InsertStmt(tx, customFeature, "custom_feature")
			_, err = customFeatureStmt.Stmt.Exec(customFeatureStmt.Args...)
			if err != nil {
				return _db.HandleError(err)
			}
		}

		for _, v := range b.SimilarProjects {
			similarProject := body.NewSimilarProject(v, project.ID)
			similarProjectStmt := db.InsertStmt(tx, similarProject, "similar_project")
			_, err = similarProjectStmt.Stmt.Exec(similarProjectStmt.Args...)
			if err != nil {
				return _db.HandleError(err)
			}
		}

		for _, v := range b.Features {
			feature := body.NewFeature(v, project.ID)
			featureStmt := db.InsertStmt(tx, feature, "project_feature")
			_, err = featureStmt.Stmt.Exec(featureStmt.Args...)
			if err != nil {
				return _db.HandleError(err)
			}
		}

		projectID = project.ID

		return nil
	})

	if _err != nil {
		return nil, _err
	}

	_err = _db.WithTransaction(func(tx *sql.Tx) *err.Error {
		fp, dErr := DbGetProjectForID(_db, projectID)
		if dErr != nil {
			return dErr
		}
		finalProject = fp

		return nil
	})

	return finalProject, _err
}
