package routes

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Record struct {
	Name  string
	Email string
	Phone string
}

func ParseCsvRoute(r *gin.Engine) {
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
				<input type="file"/>
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

	r.GET("/", func(c *gin.Context) {
		tmpl.Execute(c.Writer, data)
	})

}
