package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// createUserHandler - expects JSON body
func createUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Users
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := insertUserData(db, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func updateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}

		var payload map[string]interface{}
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		allowed := map[string]bool{
			"name": true, "lastname": true, "age": true,
			"height": true, "weight": true, "gender": true,
			"skilllevel": true, "plan": true, "email": true,
		}
		setParts := []string{}
		args := []interface{}{}
		i := 1
		for k, v := range payload {
			if allowed[k] {
				setParts = append(setParts, fmt.Sprintf("%s = $%d", k, i))
				args = append(args, v)
				i++
			}
		}
		if len(setParts) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no valid fields"})
			return
		}
		args = append(args, id)
		query := fmt.Sprintf("UPDATE user_account SET %s WHERE id = $%d",
			strings.Join(setParts, ", "), i)
		if _, err := db.Exec(query, args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "user updated"})
	}
}

func deleteUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}
		if _, err := db.Exec("DELETE FROM user_account WHERE id=$1", id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "user deleted"})
	}
}

// createCoachHandler - accepts multipart/form-data (fields + optional file "profile_pic")
func createCoachHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse form (max 10MB file)
		if err := c.Request.ParseMultipartForm(10 << 20); err != nil && err != http.ErrNotMultipart {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		name := c.Request.FormValue("name")
		lastname := c.Request.FormValue("lastname")
		age, _ := strconv.Atoi(c.Request.FormValue("age"))
		gender := c.Request.FormValue("gender")
		email := c.Request.FormValue("email")
		cost, _ := strconv.ParseFloat(c.Request.FormValue("cost"), 64)
		bio := c.Request.FormValue("bio")

		coach := Coach{
			Name:     name,
			LastName: lastname,
			Age:      age,
			Gender:   gender,
			Email:    email,
			Cost:     cost,
			Bio:      bio,
		}

		// handle optional file upload
		fileHeader, err := c.FormFile("profile_pic")
		if err == nil && fileHeader != nil {
			ensureUploadsDir("uploads/coaches")
			dst := filepath.Join("uploads/coaches", fileHeader.Filename)
			if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			coach.ProfilePic = dst
		}

		id, err := insertCoachData(db, coach)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func updateCoachHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}
		var payload map[string]interface{}
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		allowed := map[string]bool{
			"name": true, "lastname": true, "age": true,
			"gender": true, "email": true, "cost": true, "verified": true,
			"profile_pic": true, "bio": true,
		}
		setParts := []string{}
		args := []interface{}{}
		i := 1
		for k, v := range payload {
			if allowed[k] {
				setParts = append(setParts, fmt.Sprintf("%s = $%d", k, i))
				args = append(args, v)
				i++
			}
		}
		if len(setParts) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no valid fields"})
			return
		}
		args = append(args, id)
		query := fmt.Sprintf("UPDATE coach_account SET %s WHERE id = $%d", strings.Join(setParts, ", "), i)
		if _, err := db.Exec(query, args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "coach updated"})
	}
}

func deleteCoachHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}
		if _, err := db.Exec("DELETE FROM coach_account WHERE id=$1", id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "coach deleted"})
	}
}

// uploadCoachProfilePic - separate upload endpoint; saves file and updates DB
func uploadCoachProfilePic(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		coachIDStr := c.Param("id")
		coachID, err := uuid.Parse(coachIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}

		fileHeader, err := c.FormFile("profilePic")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "profile picture is required"})
			return
		}

		// simple validation
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "only .jpg .jpeg .png allowed"})
			return
		}

		ensureUploadsDir("uploads/coaches")
		path := fmt.Sprintf("uploads/coaches/%s%s", coachID.String(), ext)

		if err := c.SaveUploadedFile(fileHeader, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rows, err := updateCoachProfilePic(db, coachID, path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":        "profile picture uploaded",
			"path":          path,
			"rows_affected": rows,
		})
	}
}
