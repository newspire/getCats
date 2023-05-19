package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cat struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Breed  string  `json:"breed"`
	Age    int     `json:"age"`
	Weight float32 `json:"weight"`
}

var cats = []Cat{
	{ID: "1", Name: "Smokey", Breed: "Maine Coon", Age: 15, Weight: 12},
	{ID: "2", Name: "Toby", Breed: "Tux", Age: 5, Weight: 8},
	{ID: "3", Name: "Ty", Breed: "Tux", Age: 6, Weight: 7.25},
}

func main() {
	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/cats", getCats)
	router.GET("/cats/:id", getCatByID)
	router.POST("/cats/:id", updateCatByID)
	router.POST("/cats", postCats)

	router.Run("0.0.0.0:8080")
}

func getRoot(c *gin.Context) {
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("Cats!"))
}

// getCats responds with the list of all cats as JSON.
func getCats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cats)
}

func getCatByID(c *gin.Context) {
	id := c.Param("id")

	log.Printf("looking for cat %s", id)
	// Loop over the list of cats, looking for
	// a cat whose ID value matches the parameter.
	for _, a := range cats {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "cat not found"})
}

func updateCatByID(c *gin.Context) {
	id := c.Param("id")
	var upCat Cat

	// Call BindJSON to bind the received JSON to
	// newCat.
	if err := c.BindJSON(&upCat); err != nil {
		return
	}

	log.Printf("updating cat %s", id)
	// Loop over the list of cats, looking for
	// a cat whose ID value matches the parameter.
	for i := range cats {
		if cats[i].ID == id {
			cats[i].Age = upCat.Age
			cats[i].Breed = upCat.Breed
			cats[i].Name = upCat.Name
			cats[i].Weight = upCat.Weight
			c.IndentedJSON(http.StatusOK, cats[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "cat not found"})
}

func postCats(c *gin.Context) {
	var newCat Cat

	// Call BindJSON to bind the received JSON to
	// newCat.
	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	// Add the new cat to the slice.
	cats = append(cats, newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}
