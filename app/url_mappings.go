package app

import (
	"github.com/anuragyadav8949/bookstore_users-api/controllers/ping"
	"github.com/anuragyadav8949/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/wordCount", users.GetWordCount)
}
