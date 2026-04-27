package routes

import (
	"fmt"
	"net/http"

	"bangkok-brand/app/modules"

	"github.com/gin-gonic/gin"
)

func WarpH(router *gin.RouterGroup, prefix string, handler http.Handler) {
	router.Any(fmt.Sprintf("%s/*w", prefix), gin.WrapH(http.StripPrefix(fmt.Sprintf("%s%s", router.BasePath(), prefix), handler)))
}

func api(r *gin.RouterGroup, mod *modules.Modules) {
	r.GET("/example/:id", mod.Example.Ctl.Get)
	r.GET("/example-http", mod.Example.Ctl.GetHttpReq)
	r.POST("/example", mod.Example.Ctl.Create)
}

func apiSystem(r *gin.RouterGroup, mod *modules.Modules) {
	systems := r.Group("/systems")
	{
		genders := systems.Group("/genders")
		{
			genders.GET("", mod.Gender.Ctl.List)
			genders.GET("/:id", mod.Gender.Ctl.Info)
			genders.PATCH("/:id", mod.Gender.Ctl.Update)
			genders.DELETE("/:id", mod.Gender.Ctl.Delete)
		}
		prefixes := systems.Group("/prefixes")
		{
			prefixes.GET("", mod.Prefix.Ctl.List)
			prefixes.GET("/:id", mod.Prefix.Ctl.Info)
			prefixes.PATCH("/:id", mod.Prefix.Ctl.Update)
			prefixes.DELETE("/:id", mod.Prefix.Ctl.Delete)
		}
		provinces := systems.Group("/provinces")
		{
			provinces.GET("", mod.Province.Ctl.List)
			provinces.GET("/:id", mod.Province.Ctl.Info)
			provinces.PATCH("/:id", mod.Province.Ctl.Update)
			provinces.DELETE("/:id", mod.Province.Ctl.Delete)
		}
		districts := systems.Group("/districts")
		{
			districts.GET("", mod.District.Ctl.List)
			districts.GET("/:id", mod.District.Ctl.Info)
			districts.PATCH("/:id", mod.District.Ctl.Update)
			districts.DELETE("/:id", mod.District.Ctl.Delete)
		}
		subdistricts := systems.Group("/subdistricts")
		{
			subdistricts.GET("", mod.Subdistrict.Ctl.List)
			subdistricts.GET("/:id", mod.Subdistrict.Ctl.Info)
			subdistricts.PATCH("/:id", mod.Subdistrict.Ctl.Update)
			subdistricts.DELETE("/:id", mod.Subdistrict.Ctl.Delete)
		}
		zipcodes := systems.Group("/zipcodes")
		{
			zipcodes.GET("", mod.Zipcode.Ctl.List)
			zipcodes.GET("/:id", mod.Zipcode.Ctl.Info)
			zipcodes.PATCH("/:id", mod.Zipcode.Ctl.Update)
			zipcodes.DELETE("/:id", mod.Zipcode.Ctl.Delete)
		}
	}
}
