package person

import (
	"belajar/database"
	"belajar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SinglePerson(c *gin.Context) {
	db := database.Conn()
	defer db.Close()

	var person models.User
	id := c.Param("id")

	err := db.QueryRow("SELECT id, nama, kelas FROM user WHERE id = ?", id).Scan(&person.Id, &person.Name, &person.Kelas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"person": person})
}
