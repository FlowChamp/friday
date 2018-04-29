package main

import (
	"github.com/emicklei/go-restful"
)

type Message struct {
	Dept   string
	Course string
}

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("/{course}").To(PrintCourse))

	return service
}

func PrintCourse(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("course")
	test := Message{"CSC", id}
	response.WriteEntity(test)
}
