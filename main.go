package main

import (
	"html/template"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Name       string
	StartDate  string
	EndDate    string
	Duration   string
	Desc       string
	NodeJs     bool
	ReactJs    bool
	NextJs     bool
	TypeScript bool
}

var projectData = []Project{
	{
		Name:       "Project 1",
		StartDate:  "2020-01-15",
		EndDate:    "2020-02-15",
		Duration:   countDuration("2020-01-15", "2020-02-15"),
		Desc:       "This is the description of project 1",
		NodeJs:     true,
		ReactJs:    false,
		NextJs:     true,
		TypeScript: true,
	},
	{
		Name:       "Project 2",
		StartDate:  "2023-06-05",
		EndDate:    "2023-06-06",
		Duration:   countDuration("2023-06-05", "2023-06-06"),
		Desc:       "This is the description of project 2",
		NodeJs:     false,
		ReactJs:    false,
		NextJs:     false,
		TypeScript: false,
	},
	{
		Name:       "Project 3",
		StartDate:  "2022-06-05",
		EndDate:    "2023-06-06",
		Duration:   countDuration("2022-06-05", "2023-06-06"),
		Desc:       "This is the description of project 3",
		NodeJs:     true,
		ReactJs:    true,
		NextJs:     true,
		TypeScript: true,
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/project/:id", projectDetail)
	e.GET("/add-project", formAddProject)
	e.GET("/edit-project/:id", formEditProject)

	e.POST("/add-project", submitProject)
	e.POST("/edit-project/:id", submitEditedProject)
	e.POST("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"Projects": projectData,
	}

	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range projectData {
		if id == i {
			ProjectDetail = Project{
				Name:       data.Name,
				StartDate:  data.StartDate,
				EndDate:    data.EndDate,
				Duration:   data.Duration,
				Desc:       data.Desc,
				NodeJs:     data.NodeJs,
				ReactJs:    data.ReactJs,
				NextJs:     data.NextJs,
				TypeScript: data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"StartDate": getDateString(ProjectDetail.StartDate),
		"EndDate":   getDateString(ProjectDetail.EndDate),
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func formEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range projectData {
		if id == i {
			ProjectDetail = Project{
				Name:       data.Name,
				StartDate:  data.StartDate,
				EndDate:    data.EndDate,
				Duration:   data.Duration,
				Desc:       data.Desc,
				NodeJs:     data.NodeJs,
				ReactJs:    data.ReactJs,
				NextJs:     data.NextJs,
				TypeScript: data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
		"Id":      id,
	}

	var tmpl, err = template.ParseFiles("views/edit-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func submitProject(c echo.Context) error {
	name := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	desc := c.FormValue("description")
	nodeJs := c.FormValue("nodeJs")
	reactJs := c.FormValue("reactJs")
	nextJs := c.FormValue("nextJs")
	typescript := c.FormValue("typescript")

	var newProject = Project{
		Name:       name,
		StartDate:  startDate,
		EndDate:    endDate,
		Duration:   countDuration(startDate, endDate),
		Desc:       desc,
		NodeJs:     (nodeJs == "nodejs"),
		ReactJs:    (reactJs == "reactjs"),
		NextJs:     (nextJs == "nextjs"),
		TypeScript: (typescript == "typescript"),
	}

	projectData = append(projectData, newProject)

	id := len(projectData) - 1
	return c.Redirect(http.StatusMovedPermanently, "/project/"+strconv.Itoa(id))
}

func submitEditedProject(c echo.Context) error {
	id := getProjectIndex(c.Response(), c.Request())

	name := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	desc := c.FormValue("description")
	nodeJs := c.FormValue("nodeJs")
	reactJs := c.FormValue("reactJs")
	nextJs := c.FormValue("nextJs")
	typescript := c.FormValue("typescript")

	var editedProject = Project{
		Name:       name,
		StartDate:  startDate,
		EndDate:    endDate,
		Duration:   countDuration(startDate, endDate),
		Desc:       desc,
		NodeJs:     (nodeJs == "nodejs"),
		ReactJs:    (reactJs == "reactjs"),
		NextJs:     (nextJs == "nextjs"),
		TypeScript: (typescript == "typescript"),
	}

	i, _ := strconv.Atoi(id)
	projectData[i] = editedProject

	return c.Redirect(http.StatusMovedPermanently, "/project/"+id)
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	projectData = append(projectData[:id], projectData[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func getDateString(date string) string {
	y := date[0:4]
	m, _ := strconv.Atoi(date[5:7])
	d := date[8:10]
	if string(d[0]) == "0" {
		d = string(d[1])
	}

	mon := ""
	switch m {
	case 1:
		mon = "Jan"
	case 2:
		mon = "Feb"
	case 3:
		mon = "Mar"
	case 4:
		mon = "Apr"
	case 5:
		mon = "Mei"
	case 6:
		mon = "Jun"
	case 7:
		mon = "Jul"
	case 8:
		mon = "Agu"
	case 9:
		mon = "Sep"
	case 10:
		mon = "Okt"
	case 11:
		mon = "Nov"
	case 12:
		mon = "Des"
	}

	return d + " " + mon + " " + y
}

func countDuration(d1 string, d2 string) string {
	date1, _ := time.Parse("2006-01-02", d1)
	date2, _ := time.Parse("2006-01-02", d2)

	diff := date2.Sub(date1)
	days := int(diff.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months > 12 {
		return strconv.Itoa(months/12) + " tahun"
	}
	if months > 0 {
		return strconv.Itoa(months) + " bulan"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " minggu"
	}
	return strconv.Itoa(days) + " hari"
}

func getProjectIndex(w http.ResponseWriter, r *http.Request) string {
	// to call: getProjectIndex(c.Response(), c.Request())
	// value of url: /edit-project/0?
	// returned value: 0
	url := r.URL.String()
	lastSegment := path.Base(url)
	re := regexp.MustCompile("[0-9]+")
	return strings.Join((re.FindAllString(lastSegment, -1))[:], "")
}
