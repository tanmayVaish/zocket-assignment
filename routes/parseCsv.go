package routes

import (
	"encoding/csv"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Record struct {
	Name  string
	Email string
	Phone string
}

func ParseCsvRoute(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		// Render form.html from the templates directory
		c.HTML(http.StatusOK, "form.html", nil)
	})

	// Define a route to handle form submissions
	r.POST("/show", func(c *gin.Context) {
		// Read and parse the CSV file
		path := c.PostForm("path")

		// if path does not end with .csv
		if path[len(path)-4:] != ".csv" {
			c.HTML(http.StatusOK, "form.html", gin.H{
				"error": "Please enter a valid path to a csv file",
			})
			return
		}

		file, err := os.Open(path)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var data []Record
		for _, r := range records {
			// do not run first iteration
			if r[0] == "Name" {
				continue
			}
			record := Record{Name: r[0], Email: r[1], Phone: r[2]}
			data = append(data, record)
		}

		// Pass the data to the template and render it
		tmpl := template.Must(template.ParseFiles("templates/template.html"))
		c.HTML(http.StatusOK, "template.html", gin.H{
			"data": data,
			"tmpl": tmpl,
		})
	})

	// Serve static files (CSS and JavaScript)
	r.Static("/static", "./static")

}
