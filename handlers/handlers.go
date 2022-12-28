package handlers

import (
	"github.com/MatheusLasserre/go-fiber-docker-sqloback/database"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.SendString("Here's the Index!")
}

func GetCourses(c *fiber.Ctx) error {

	type queryRowStruct struct {
		Id               int    `json:"id" db:"id"`
		Name             string `json:"name" db:"name"`
		Description      string `json:"description" db:"description"`
		OrganizationId   int    `json:"organizationId" db:"organizationId"`
		OrganizationName string `json:"organizationName" db:"organizationName"`
		ClassesId        int    `json:"classesId" db:"classesId"`
		ClassesName      string `json:"classesName" db:"classesName"`
		ClassesLink      string `json:"classesLink" db:"classesLink"`
		ClassesPosition  int    `json:"classesPosition" db:"classesPosition"`
	}

	type Organization struct {
		Id   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}

	type Class struct {
		Id       int    `json:"id" db:"id"`
		Name     string `json:"name" db:"name"`
		Link     string `json:"link" db:"link"`
		Position int    `json:"position" db:"position"`
	}

	type queryDataStruct struct {
		Id           int          `json:"id" db:"id"`
		Name         string       `json:"name" db:"name"`
		Description  string       `json:"description" db:"description"`
		Organization Organization `json:"organization" db:"organization"`
		Classes      []Class      `json:"classes" db:"classes"`
	}

	rows, err := database.Db.Queryx(`
		SELECT
		courses.id "id",
		courses.name "name",
		courses.description "description",
		organizations.name "organizationName",
		organizations.id "organizationId",
		classes.id "classesId",
		classes.name "classesName",
		classes.link "classesLink",
		classes.position "classesPosition"


		FROM courses
		JOIN organizations ON courses.organizationid = organizations.id
		JOIN classes ON courses.id = classes.courseid
	

		/* application='GO-FIBER' */

	`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var queryData []queryDataStruct

	isCourseInSlice := func(courseId int, slice []queryDataStruct) int {
		for i, course := range slice {
			if course.Id == courseId {
				return i
			}
		}
		return -1
	}

	for rows.Next() {
		var queryRow queryRowStruct
		err = rows.StructScan(&queryRow)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var class Class
		class.Id = queryRow.ClassesId
		class.Name = queryRow.ClassesName
		class.Link = queryRow.ClassesLink
		class.Position = queryRow.ClassesPosition
		if i := isCourseInSlice(queryRow.Id, queryData); i != -1 {
			queryData[i].Classes = append(queryData[i].Classes, class)
			continue
		}
		var organization Organization
		organization.Id = queryRow.OrganizationId
		organization.Name = queryRow.OrganizationName

		var course queryDataStruct
		course.Id = queryRow.Id
		course.Name = queryRow.Name
		course.Description = queryRow.Description
		course.Organization = organization
		course.Classes = append(course.Classes, class)

		queryData = append(queryData, course)
	}

	return c.Status(fiber.StatusOK).JSON(queryData)

}

type GetCourse struct {
	Id             int     `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	Description    string  `json:"description" db:"description"`
	OrganizationId int     `json:"organizationId" db:"organizationId"`
	Price          float64 `json:"price" db:"price"`
}

func PostCourses(c *fiber.Ctx) error {
	return c.SendString("Here's the POST Courses!")
}

// Database Models

type User struct {
	ID            string
	Name          string
	Email         string
	EmailVerified string
	Image         string
	Courses       []CourseAccesses
	Organizations []OrganizationsMembers
	Password      string
}

type Organizations struct {
	ID      int                      `json:"id" db:"id"`
	Name    string                   `json:"name" db:"name"`
	Members []OrganizationsMembers   `json:"members" db:"members"`
	Courses []Courses                `json:"courses" db:"courses"`
	Library []OrganizationsLibraries `json:"library" db:"library"`
}

type OrganizationsMembers struct {
	ID             int
	OrganizationId int
	Organization   Organizations
	UserId         string
	User           User
}

type OrganizationsLibraries struct {
	ID             int           `json:"id" db:"id"`
	OrganizationId int           `json:"organizationId" db:"organizationId"`
	Organization   Organizations `json:"organization" db:"organization"`
	LibraryId      int           `json:"libraryId" db:"libraryId"`
}

type CourseAccesses struct {
	ID       int
	UserID   string
	CourseID int
	User     User
	Course   Courses
}

type Courses struct {
	Id             int              `json:"id" db:"id"`
	Name           string           `json:"name" db:"name"`
	Description    string           `json:"description" db:"description"`
	OrganizationId int              `json:"organizationId" db:"organizationId"`
	Organization   Organizations    `json:"organization" db:"organization"`
	Classes        []Classes        `json:"classes" db:"classes"`
	Users          []CourseAccesses `json:"users" db:"users"`
	Enabled        bool             `json:"enabled" db:"enabled"`
	Price          float64          `json:"price" db:"price"`
}

type Classes struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Link       string  `json:"link" db:"link"`
	Position   int     `json:"position" db:"position"`
	UploadedAt string  `json:"uploadedAt" db:"uploadedAt"`
	Duration   string  `json:"duration" db:"duration"`
	CourseID   int     `json:"courseId" db:"courseId"`
	Course     Courses `json:"course" db:"course"`
}
