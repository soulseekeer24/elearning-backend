package common

// BodyResponse represen all the data for course
type BodyResponse struct {
	Courses []CourseInfo `json:"courses"`
}