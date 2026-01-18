package routes

import (
	"golang_mvc_starterpack/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouteApi(h *controllers.PeopleHandler) *gin.Engine {
	r := gin.Default()
	// Routing people api
	{
		people := r.Group("/api/people")
		people.GET("", h.IndexPeople)        // all data
		people.POST("", h.StorePeople)       // store data
		people.GET("/:id", h.ShowPeople)     // update data
		people.PATCH("/:id", h.UpdatePeople) // update data partial
		people.DELETE("/:id", nil)           // deleted data
	}
	return r
}
