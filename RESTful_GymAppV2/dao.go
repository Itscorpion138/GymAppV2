package main

import (
	"database/sql"

	"github.com/google/uuid"
)

// insertUserData inserts a user and returns generated UUID.
func insertUserData(db *sql.DB, user Users) (uuid.UUID, error) {
	var pk uuid.UUID
	query := `
    INSERT INTO user_account (name, lastname, age, height, weight, gender, skilllevel, plan, email)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id`
	err := db.QueryRow(query,
		user.Name, user.LastName, user.Age, user.Height, user.Weight,
		user.Gender, user.SkillLevel, user.Plan, user.Email,
	).Scan(&pk)
	return pk, err
}

// insertCoachData inserts a coach and returns generated UUID.
func insertCoachData(db *sql.DB, coach Coach) (uuid.UUID, error) {
	var pk uuid.UUID
	query := `
    INSERT INTO coach_account (name, lastname, gender, age, email, cost, profile_pic, bio )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`
	err := db.QueryRow(query,
		coach.Name, coach.LastName, coach.Gender, coach.Age, coach.Email, coach.Cost, coach.ProfilePic,
	coach.Bio).Scan(&pk)
	return pk, err
}

func verifyCoach(db *sql.DB, coachID uuid.UUID) error {
	_, err := db.Exec("UPDATE coach_account SET verified = true WHERE id = $1", coachID)
	return err
}

func updateCoachProfilePic(db *sql.DB, coachID uuid.UUID, path string) (int64, error) {
	res, err := db.Exec("UPDATE coach_account SET profile_pic=$1 WHERE id=$2", path, coachID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
