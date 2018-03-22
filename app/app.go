package app

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	//	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         string `json:"id"`
	UserName   string `json:"user_name"`
	UserStatus string `json:"user_status"`
	Date       string `json:"date"`
}

//var dbmap = initDb()

func GetUsers(c *gin.Context) {
	var users []User
	csvFile, _ := os.Open("users.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			c.JSON(404, gin.H{"error": "user not found"})
		}
		users = append(users, User{
			Id:         line[0],
			UserName:   line[1],
			UserStatus: line[2],
			Date:       line[3],
		})
	}
	//usersJson, _ := json.Marshal(users)
	c.JSON(200, users)
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
	csvFile, _ := os.Open("users.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			c.JSON(404, gin.H{"error": "user not found"})
		}
		//stringId := fmt.Scanf("%s", id)
		if strings.TrimRight(id, "\n") == line[0] {
			user = User{
				Id:         line[0],
				UserName:   line[1],
				UserStatus: line[2],
				Date:       line[3],
			}
		}

	}

	c.JSON(200, user)
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
	// curl -i http://localhost:8080/api/v1/users/1
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func increment() string {
	csvFile, _ := os.Open("users.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	c := 0
	for {
		c++
		_, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
	}

	return strconv.Itoa(c)
}

func returnUsers() []User {
	var users []User
	csvFile, _ := os.Open("users.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("error: user not found")
		}
		users = append(users, User{
			Id:         line[0],
			UserName:   line[1],
			UserStatus: line[2],
			Date:       line[3],
		})
	}
	//usersJson, _ := json.Marshal(users)
	return users
	// _, err := dbmap.Select(&users, "SELECT * FROM user")
	// if err == nil {
	// 	c.JSON(200, users)
	// } else {
	// 	c.JSON(404, gin.H{"error": "no user(s) into the table"})
	// }
	// curl -i http://localhost:8080/api/v1/users
}
func PostUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	var data = []string{increment(), user.UserName, user.UserStatus, user.Date}
	var ru = returnUsers()
	var allUsers = [][]string{}
	for _, userBody := range ru {
		u := []string{userBody.Id, userBody.UserName, userBody.UserStatus, userBody.Date}
		allUsers = append(allUsers, u)
	}
	allUsers = append(allUsers, data)

	if user.UserStatus != "" && user.UserName != "" {
		file, err := os.Create("users.csv")
		checkError("Cannot create file", err)
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		for _, value := range allUsers {
			err := writer.Write(value)
			checkError("Cannot write to file", err)
		}
		c.JSON(201, user)
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}

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
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_status\": \"Chasing Windmills\", \"user_name\": \"Don Quixote\", "date": "01-07-1604" }" http://localhost:8080/api/v1/users
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	c.Bind(&user)
	var data = []string{id, user.UserName, user.UserStatus, user.Date}
	var ru = returnUsers()
	var allUsers = [][]string{}
	for _, userBody := range ru {
		if strings.TrimRight(id, "\n") == userBody.Id {
			allUsers = append(allUsers, data)
		} else {
			u := []string{userBody.Id, userBody.UserName, userBody.UserStatus, userBody.Date}
			allUsers = append(allUsers, u)
		}
	}
	if user.UserName != "" || user.UserStatus != "" || user.Date != "" {
		file, err := os.Create("users.csv")
		checkError("Cannot create file", err)
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		for _, value := range allUsers {
			err := writer.Write(value)
			checkError("Cannot write to file", err)
		}
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
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
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"user_status\": \"\", \"user_name\": \"100\" }" http://localhost:8080/api/v1/users/1

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"user_status\": \"Slaying Giants\", \"user_name\": \"Don Quixote\", \"date\": \"03-07-1604\" }" http://localhost:8080/api/v1/users/3

}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User

	var data = []string{id, user.UserName, user.UserStatus, user.Date}
	var ru = returnUsers()
	var allUsers = [][]string{}
	for _, userBody := range ru {
		if strings.TrimRight(id, "\n") == userBody.Id {
			allUsers = append(allUsers, data)
		} else {
			u := []string{userBody.Id, userBody.UserName, userBody.UserStatus, userBody.Date}
			allUsers = append(allUsers, u)
		}
	}

	file, err := os.Create("users.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range allUsers {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	c.JSON(200, user)

	//id := c.Params.ByName("id")
	//var user User
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
