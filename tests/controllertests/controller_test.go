package controllertests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/wewebplus/wewebapi/api/controllers"
	"github.com/wewebplus/wewebapi/api/models"
)

var server = controllers.Server{}
var userInstance = models.SysStf{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
	if TestDbDriver == "sqlite3" {
		//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		testDbName := os.Getenv("TestDbName")
		server.DB, err = gorm.Open(TestDbDriver, testDbName)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
		server.DB.Exec("PRAGMA foreign_keys = ON")
	}
}

func refreshUserTable() error {
	/*
		err := server.DB.DropTableIfExists(&models.User{}).Error
		if err != nil {
			return err
		}
		err = server.DB.AutoMigrate(&models.User{}).Error
		if err != nil {
			return err
		}
	*/
	err := server.DB.DropTableIfExists(&models.Post{}, &models.SysStf{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.SysStf{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed table(s)")
	return nil
}

func seedOneUser() (models.SysStf, error) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.SysStf{
		SyStfFnameeng: "Pet",
		SyStfEmail:    "pet@gmail.com",
		SyStfPassword: "password",
	}

	err = server.DB.Model(&models.SysStf{}).Create(&user).Error
	if err != nil {
		return models.SysStf{}, err
	}
	return user, nil
}

func seedUsers() ([]models.SysStf, error) {

	var err error
	if err != nil {
		return nil, err
	}
	users := []models.SysStf{
		models.SysStf{
			SyStfFnameeng: "Steven victor",
			SyStfEmail:    "steven@gmail.com",
			SyStfPassword: "password",
		},
		models.SysStf{
			SyStfFnameeng: "Kenny Morris",
			SyStfEmail:    "kenny@gmail.com",
			SyStfPassword: "password",
		},
	}
	for i, _ := range users {
		err := server.DB.Model(&models.SysStf{}).Create(&users[i]).Error
		if err != nil {
			return []models.SysStf{}, err
		}
	}
	return users, nil
}

func refreshUserAndPostTable() error {

	err := server.DB.DropTableIfExists(&models.Post{}, &models.SysStf{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.SysStf{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}
	user := models.SysStf{
		SyStfFnameeng: "Sam Phil",
		SyStfEmail:    "sam@gmail.com",
		SyStfPassword: "password",
	}
	err = server.DB.Model(&models.SysStf{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.SyStfId,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedUsersAndPosts() ([]models.SysStf, []models.Post, error) {

	var err error

	if err != nil {
		return []models.SysStf{}, []models.Post{}, err
	}
	var users = []models.SysStf{
		models.SysStf{
			SyStfFnameeng: "Steven victor",
			SyStfEmail:    "steven@gmail.com",
			SyStfPassword: "password",
		},
		models.SysStf{
			SyStfFnameeng: "Magu Frank",
			SyStfEmail:    "magu@gmail.com",
			SyStfPassword: "password",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title:   "Title 1",
			Content: "Hello world 1",
		},
		models.Post{
			Title:   "Title 2",
			Content: "Hello world 2",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.SysStf{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].SyStfId

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, posts, nil
}
