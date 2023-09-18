package person

import (
	"belajar/database"
	"belajar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getperson(c *gin.Context) {
	// enableCors(&w)
	acces := true

	db := database.Conn()
	if !acces {
		c.JSON(500, gin.H{"error": "kamu ditak memiliki akses"})
	} else {

		rows, err := db.Query("SELECT id, nama, kelas FROM user ")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.Id, &user.Name, &user.Kelas); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			users = append(users, user)

		}

		if users == nil {
			c.JSON(405, gin.H{"error": "data tidak ada"})
		} else {

			c.JSON(http.StatusOK, users)
		}

	}

}
