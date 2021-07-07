package main

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
)

// Article - Our struct for all articles
type Article struct {
    Id       string `json:"id"`
    Title    string `json:"title"`
    Intro    string `json:"intro"`
    AuthorId string `json:"authorid"`
    Content  string `json:"content"`
}

func returnAllArticles(c *gin.Context) {

    file, _ := ioutil.ReadFile("data.json")

    var Articles []Article

    _ = json.Unmarshal([]byte(file), &Articles)

    c.JSON(http.StatusOK, gin.H{"status": "success", "articles": &Articles})
    return
}

func returnSingleArticle(c *gin.Context) {
    var id = c.Param("id")

    file, _ := ioutil.ReadFile("data.json")

    var Articles []Article

    _ = json.Unmarshal([]byte(file), &Articles)

    for _, article := range Articles {
        if article.Id == id {
            c.JSON(http.StatusOK, gin.H{"status": "success", "article": &article})
            return
        }
    }

    c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid ID"})

}

func createNewArticle(c *gin.Context) {

    var article Article
    if err := c.ShouldBindJSON(&article); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid Input data"})
        return
    }

    file, _ := ioutil.ReadFile("data.json")

    var Articles []Article

    _ = json.Unmarshal([]byte(file), &Articles)

    Articles = append(Articles, article)

    data, _ := json.Marshal(Articles)

    err := ioutil.WriteFile("data.json", data, 0777)

    if err != nil {

        fmt.Println(err)
    }
    c.JSON(http.StatusOK, gin.H{"articles": &Articles})

}

func updateArticle(c *gin.Context) {

    var id = c.Param("id")

    file, _ := ioutil.ReadFile("data.json")

    var Articles []Article

    _ = json.Unmarshal([]byte(file), &Articles)

    var article Article
    if err := c.ShouldBindJSON(&article); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid Input data"})
        return
    }

    for index, article_ := range Articles {

        if article_.Id == id {
            Articles[index] = article

            data, _ := json.Marshal(Articles)
            err := ioutil.WriteFile("data.json", data, 0777)
            if err != nil {

                fmt.Println(err)
            }
            c.JSON(http.StatusOK, gin.H{"status": "success", "article": &Articles})
            return
        }
    }

    c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid ID"})

}

func main() {

    router := gin.Default()
    api := router.Group("/api/v1")
    {
        api.GET("/articles", returnAllArticles)
        api.GET("/articles/:id", returnSingleArticle)
        api.POST("/articles", createNewArticle)
        api.POST("/articles/:id/update", updateArticle)
    }
    router.NoRoute(func(c *gin.Context) {
        c.AbortWithStatus(http.StatusNotFound)
    })
    router.Run(":5000")
}
