package person

import (
	"belajar/database"
	"belajar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePerson(c *gin.Context) {
	db := database.Conn()
	defer db.Close()

	var person models.User
	id := c.Param("id")

	err := db.QueryRow("SELECT id, nama, kelas FROM user WHERE id = ?", id).Scan(&person.Id, &person.Name, &person.Kelas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	username := c.PostForm("name")
	kelas := c.PostForm("kelas")

	if username == "" || kelas == "" {
		c.JSON(404, gin.H{"error": "Nama/kelas tidak boleh kosong"})
	} else {

		_, err := db.Exec("UPDATE user SET nama = ?, kelas = ? WHERE id = ?", username, kelas, id)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data user berhasil diupdate"})
	}
}
