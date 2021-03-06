package main

import (
  "net/http"
  "gopkg.in/gin-gonic/gin.v1"
	_ "log"
  )

  type User struct {
    Id int64 `db:"id" json:"id"`
    Firstname string `db:"firstname" json:"firstname"`
    Lastname string `db:"lastname" json:"lastname"`
    Image string `db:"image" json:"image"`
  }

type Users struct {
	Collection []User
}
  func main() {
	 router := gin.Default()
	 router.GET("/index", func(c *gin.Context) {
		 type Users []User
	    var users = Users{
			User{Id: 1, Firstname: "Peter", Lastname: "Braun", Image: "https://fiw.fhws.de/fileadmin/_processed_/csm_Braun_DSC6980_RGB_b6d7988bbc.jpg"},
			User{Id: 2, Firstname: "Isabel", Lastname: "John", Image: "https://fiw.fhws.de/fileadmin/_processed_/csm_John_IMG4076_RGB_d3bffb5d6e.jpg"},
			User{Id: 3, Firstname: "Eberhard", Lastname: "Grötsch", Image: "https://fiw.fhws.de/fileadmin/_processed_/csm_Groetsch_IMG5513_RGB_d9bc5b45b4.jpg"},
			User{Id: 4, Firstname: "Klaus", Lastname: "Junker-Schilling", Image: "https://fiw.fhws.de/fileadmin/_processed_/csm_Junker-Schilling_IMG5093_RGB_2fc6b9361d.jpg"},
			User{Id: 5, Firstname: "Michael", Lastname: "Müßig", Image: "https://fiw.fhws.de/fileadmin/_processed_/csm_Muessig_IMG9416_RGB_59fb7e0f6e.jpg"},
			User{Id: 6, Firstname: "Markus", Lastname: "Sommer", Image: "https://www.xing.com/image/e_5_e_f9ebd2ef0_6400226_3/markus-sommer-foto.256x256.jpg"},
		}
		c.JSON(http.StatusOK, users)
	 })
	 router.Run(":8083")
  }
