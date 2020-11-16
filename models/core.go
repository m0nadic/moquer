package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m0nadic/moquer/util"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Response struct {
	Type   string
	Value  string
	Status int
}

type Route struct {
	Method   string
	Pattern  string
	Response *Response
}

type Service struct {
	name   string
	Port   int
	Routes []*Route
}

func (s *Service) SetName(name string) {
	s.name = name
}

func (s *Service) Start() {
	log.Printf("starting %s service on %d", s.name, s.Port)

	router := gin.New()
	router.Use(gin.Recovery())

	for _, route := range s.Routes {
		router.Handle(route.Method, route.Pattern, MakeHandler(route.Response))
	}

	srv := &http.Server{
		Addr:    s.Addr(),
		Handler: router,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Println("ERROR:", err)
	}

}

func (s *Service) Addr() string {
	return fmt.Sprintf("0.0.0.0:%d", s.Port)
}


func MakeHandler(response *Response) gin.HandlerFunc {
	switch response.Type {
	case "string":
		return func(c *gin.Context) {
			_, _ = fmt.Fprintln(c.Writer, response.Value)
		}
	case "file":
		return func(c *gin.Context) {
			data, _ := ioutil.ReadFile(response.Value)
			t, _ := template.New("response").
				Funcs(util.FakerFuncs).
				Funcs(util.HelperFuncs).
				Parse(string(data))
			_ = t.Execute(c.Writer, util.Payload{c})
		}
	default:
		return func(c *gin.Context) {
			_, _ = fmt.Fprintln(c.Writer, "NOP")
		}
	}

}
