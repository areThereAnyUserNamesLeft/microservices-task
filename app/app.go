package app

import (
	"github.com/gin-gonic/gin"
	//	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int64  `db:"id" json:"id"`
	UserStatus string `db:"user_status" json:"user_status"`
	UserName   string `db:"user_name" json:"user_name"`
	Date       string
}

//var dbmap = initDb()

func GetUsers(c *gin.Context) {
	var users []User
	// _, err := dbmap.Select(&users, "SELECT * FROM user")
	// if err == nil {
	// 	c.JSON(200, users)
	// } else {
	// 	c.JSON(404, gin.H{"error": "no user(s) into the table"})
	// }
	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User

	// err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
	// if err == nil {
	// 	user_id, _ := strconv.ParseInt(id, 0, 64)
	//
	// 	content := &User{
	// 		Id:         user_id,
	// 		UserStatus: user.UserStatus,
	// 		UserName:   user.UserName,
	// 	}
	//
	// 	c.JSON(200, content)
	// } else {
	// 	c.JSON(404, gin.H{"error": "user not found"})
	// }
	// curl -i http://localhost:8080/api/v1/Users/1
}

func PostUser(c *gin.Context) {
	var user User
	c.Bind(&user)

	// if user.UserStatus != "" && user.UserName != "" {
	// 	if insert, _ := dbmap.Exec(`INSERT INTO user (user_status, user_name) VALUES (?, ?)`, user.UserStatus, user.UserName); insert != nil {
	// 		user_id, err := insert.LastInsertId()
	// 		if err == nil {
	// 			content := &User{
	// 				Id:         user_id,
	// 				UserStatus: user.UserStatus,
	// 				UserName:   user.UserName,
	// 			}
	// 			c.JSON(201, content)
	// 		} else {
	// 			checkErr(err, "Insert failed")
	// 		}
	// 	}
	// } else {
	// 	c.JSON(422, gin.H{"error": "fields are empty"})
	// }
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_status\": \"83\", \"user_name\": \"100\" }" http://localhost:8080/api/v1/users
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	// err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
	//
	// if err == nil {
	// 	var json User
	// 	c.Bind(&json)
	// 	user_id, _ := strconv.ParseInt(id, 0, 64)
	// 	user := User{
	// 		Id:         user_id,
	// 		UserStatus: json.UserStatus,
	// 		UserName:   json.UserName,
	// 	}
	//
	// 	if user.UserStatus != "" && user.UserName != "" {
	// 		_, err = dbmap.Update(&user)
	//
	// 		if err == nil {
	// 			c.JSON(200, user)
	// 		} else {
	// 			checkErr(err, "Updated failed")
	// 		}
	// 	} else {
	// 		c.JSON(422, gin.H{"error": "fields are empty"})
	// 	}
	// } else {
	// 	c.JSON(404, gin.H{"error": "user not found"})
	// }
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"user_status\": \"83\", \"user_name\": \"100\" }" http://localhost:8080/api/v1/users/1
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	// err := dbmap.SelectOne(&user, "SELECT id FROM User WHERE id=?", id)
	//
	// if err == nil {
	// 	_, err = dbmap.Delete(&user)
	//
	// 	if err == nil {
	// 		c.JSON(200, gin.H{"id #" + id: " deleted"})
	// 	} else {
	// 		checkErr(err, "Delete failed")
	// 	}
	// } else {
	// 	c.JSON(404, gin.H{"error": "user not found"})
	// }
	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}
