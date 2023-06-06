package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/add-project", addProject)
	e.GET("/contact", contact)
	e.GET("/project/:id", projectDetail)
	e.POST("/", submitProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
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

	data := map[string]interface{}{
		"Id":      id,
		"Title":   "Project 1",
		"Content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis ipsum suspendisse ultrices gravida dictum fusce ut. Viverra tellus in hac habitasse platea dictumst. Tellus pellentesque eu tincidunt tortor aliquam nulla facilisi cras fermentum. Est pellentesque elit ullamcorper dignissim cras tincidunt lobortis feugiat. Risus at ultrices mi tempus imperdiet nulla malesuada. Nunc lobortis mattis aliquam faucibus. Et malesuada fames ac turpis egestas integer eget aliquet nibh. Quam vulputate dignissim suspendisse in. Nunc mi ipsum faucibus vitae aliquet nec ullamcorper. Diam donec adipiscing tristique risus. Potenti nullam ac tortor vitae purus faucibus ornare suspendisse. Elit at imperdiet dui accumsan sit amet nulla facilisi. Sagittis purus sit amet volutpat consequat mauris nunc congue. Felis imperdiet proin fermentum leo vel orci porta. Eu volutpat odio facilisis mauris sit amet. Quis hendrerit dolor magna eget est lorem ipsum.",
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

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
	// nodeJs := c.FormValue("nodeJs")
	// reactJs := c.FormValue("reactJs")
	// nextJs := c.FormValue("nextJs")
	// typescript := c.FormValue("typescript")

	println("Project Name\t:", name)
	println("Start Date\t:", startDate)
	println("End Date\t:", endDate)
	println("Description\t:", desc, "\n")
	// print("Technologies\t: ")

	return c.Redirect(http.StatusMovedPermanently, "/")
}
