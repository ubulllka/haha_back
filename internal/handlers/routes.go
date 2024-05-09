package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *Service
}

//Role
//anon - не авторизированный
//applicant - соискатель
//employer - работодатель
//admin - админ

func NewHandler(services *Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/info", h.getAllUser)
	}

	apiAnon := r.Group("/api")
	{
		userAnon := apiAnon.Group("/user")
		{
			userAnon.GET("/:id", h.getUser)
		}
		vacAnon := apiAnon.Group("/vac")
		{
			vacAnon.GET("/", h.getAllVacancies)
			vacAnon.GET("/search", h.searchVacancies)
			vacAnon.GET("/:id", h.getVacancy)
		}
		resAnon := apiAnon.Group("/res")
		{
			resAnon.GET("/", h.getAllResumes)
			resAnon.GET("/search", h.searchResumes)
			resAnon.GET("/:id", h.getResume)
		}
	}

	api := r.Group("/api", h.userIdentity)
	{

		user := api.Group("/user")
		{
			user.GET("/", h.getInfo)
			user.PATCH("/", h.updateInfo)
		}

		appl := api.Group("/appl") //соискатель
		{
			appl.GET("/", h.getApplAllResumes)
		}

		empl := api.Group("/empl") //работодатель
		{
			empl.GET("/", h.getEmplAllVacancies)
		}

		res := api.Group("/res")
		{
			res.POST("/", h.createResume)
			res.PATCH("/:id", h.updateResume)
			res.DELETE("/:id", h.deleteResume)
			res.POST("/:id/work", h.createWork)
			work := res.Group("/work")
			{
				work.PATCH("/:id", h.updateWork)
				work.DELETE("/:id", h.deleteWork)
			}
		}

		vac := api.Group("/vac")
		{

			vac.POST("/", h.createVacancy)
			vac.PATCH("/:id", h.updateVacancy)
			vac.DELETE("/:id", h.deleteVacancy)
		}
	}

	return r
}
