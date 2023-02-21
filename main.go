package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
	"zocket-assignment/models"
	"zocket-assignment/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Record struct {
	Name  string
	Email string
	Phone string
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the CSV data into a slice of Record structs
	var data []Record
	for _, r := range records {
		record := Record{Name: r[0], Email: r[1], Phone: r[2]}
		data = append(data, record)
	}

	// Define a template to render the data as a table
	tmpl := template.Must(template.New("table").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>CSV Data</title>
			</head>
			<body>
				<table>
					<thead>
						<tr>
							<th>Name</th>
							<th>Email</th>
							<th>Phone</th>
						</tr>
					</thead>
					<tbody>
						{{range .}}
							<tr>
								<td>{{.Name}}</td>
								<td>{{.Email}}</td>
								<td>{{.Phone}}</td>
							</tr>
						{{end}}
					</tbody>
				</table>
			</body>
		</html>
	`))

	// Create a new gin instance
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		tmpl.Execute(c.Writer, data)
	})

	// Load .env file and Create a new connection to the database
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Load the routes
	routes.CrudRoute(r)

	// Run the server
	r.Run(":8080")
}
