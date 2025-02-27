package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Project struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

var projects = []Project{
	{Id: "1", Name: "Project 1", Desc: "Description 1", Completed: false, Created: "2022-01-01", Updated: "2022-01-01"},
	{Id: "2", Name: "Project 2", Desc: "Description 2", Completed: true, Created: "2022-02-01", Updated: "2022-02-01"},
	{Id: "3", Name: "Project 3", Desc: "Description 3", Completed: false, Created: "2022-03-01", Updated: "2022-03-01"},
	{Id: "4", Name: "Project 4", Desc: "Description 4", Completed: true, Created: "2022-04-01", Updated: "2022-04-01"},
	{Id: "5", Name: "Project 5", Desc: "Description 5", Completed: false, Created: "2022-05-01", Updated: "2022-05-01"},
}

func main() {
	r := gin.Default()

	r.GET("/projects", getAllPrj)

	r.GET("/projects/:id", getPrjById)

	r.GET("/projects/name/:name", getPrjByName)

	r.POST("/projects/new", createPrj)

	r.PUT("/projects/:id", updateProjectById)

	r.DELETE("/projects/:id", deleteProjectById)

	r.Run(":8080")
}

func getAllPrj(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit := 10
	startIndex := (pageInt - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= len(projects) {
		c.JSON(http.StatusOK, []Project{})
		return
	}

	if endIndex > len(projects) {
		endIndex = len(projects)
	}

	c.JSON(http.StatusOK, projects[startIndex:endIndex])

}

func getPrjById(c *gin.Context) {
	id := c.Param("id")
	for _, prj := range projects {
		if prj.Id == id {
			c.JSON(http.StatusOK, prj)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func getPrjByName(c *gin.Context) {
	name := c.Param("name")
	for _, prj := range projects {
		if prj.Name == name {
			c.JSON(http.StatusOK, prj)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func createPrj(c *gin.Context) {
	var prj Project
	if err := c.ShouldBindJSON(&prj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	projects = append(projects, prj)
	c.JSON(http.StatusCreated, prj)
}

func updateProjectById(c *gin.Context) {
	id := c.Param("id")
	for i, prj := range projects {
		if prj.Id == id {
			var updatedPrj Project
			if err := c.ShouldBindJSON(&updatedPrj); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			projects[i] = updatedPrj
			c.JSON(http.StatusOK, updatedPrj)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func deleteProjectById(c *gin.Context) {
	id := c.Param("id")
	for i, prj := range projects {
		if prj.Id == id {
			projects = append(projects[:i], projects[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}
