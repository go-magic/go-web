package router

import (
	"dat2/router/query"
	"github.com/gin-gonic/gin"
)

const (
	QUERY  = "/query"
	DELETE = "delete"
	UPDATE = "update"
	ADD    = "add"
)

const (
	STUDENT = "/student"
	TEACHER = "/teacher"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use()
	Query(r)
	return r
}

func Query(r *gin.Engine) {
	s := r.Group(QUERY)
	s.GET(STUDENT, query.GetStudentInfo)
	s.GET(TEACHER, query.GetTeacherInfo)
}
