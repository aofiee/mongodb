package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aofiee/mongodb/graph/generated"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func GraphqlHandler(c *fiber.Ctx) error {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{}))
	fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.ServeHTTP(writer, request)
	})(c.Context())
	return nil
}

func PlaygroundHandler(c *fiber.Ctx) error {
	h := playground.Handler("GraphQL", "/query")
	fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.ServeHTTP(writer, request)
	})(c.Context())
	return nil
}
