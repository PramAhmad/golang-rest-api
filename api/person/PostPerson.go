package person

import (
	"belajar/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostPerson(c *gin.Context) {
	db := database.Conn()
	defer db.Close()

	username := c.PostForm("name")
	kelas := c.PostForm("kelas")

	if username == "" || kelas == "" {
		c.JSON(404, gin.H{"error": "Nama/kelas tidak boleh kosong"})
	} else {

		result, err := db.Exec("INSERT INTO user(nama, kelas) VALUES(?, ?)", username, kelas)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userID, err := result.LastInsertId()
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data user berhasil disimpan", "user_id": userID})
	}
}
