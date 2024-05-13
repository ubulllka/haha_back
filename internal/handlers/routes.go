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
			userAnon.GET("/:id/listpag", h.getListPag)
		}
		vacAnon := apiAnon.Group("/vaca")
		{
			vacAnon.GET("/", h.getAllVacancies)
			vacAnon.GET("/search", h.searchVacanciesAnon)
			vacAnon.GET("/:id", h.getVacancyAnon)
		}
		resAnon := apiAnon.Group("/resa")
		{
			resAnon.GET("/", h.getAllResumes)
			resAnon.GET("/search", h.searchResumesAnon)
			resAnon.GET("/:id", h.getResumeAnon)
			resAnon.GET("/:id/work", h.getListWork)
		}
	}

	api := r.Group("/api", h.userIdentity)
	{

		user := api.Group("/user")
		{
			user.GET("/", h.getInfo)
			user.GET("/:id/is", h.isUser)
			user.PATCH("/", h.updateInfo)
			user.GET("/listmypag", h.getMyListPag)
			user.GET("/list", h.getList)
		}

		respond := api.Group("/respond")
		{
			respond.POST("/", h.createRespond)
			respond.PATCH("/:id", h.updateRespond)
			respond.GET("/my", h.getMyAllResponds)
			respond.GET("/other", h.getOtherAllResponds)
			respond.GET("/:id/my", h.getMyRespond)
			respond.GET("/:id/other", h.getOtherRespond)
			respond.DELETE("/:id/my", h.deleteMyRespond)
			respond.DELETE("/:id/other", h.deleteOtherRespond)

		}

		res := api.Group("/res")
		{
			res.GET("/search", h.searchResumes)
			res.GET("/:id", h.getResume)
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
			vac.GET("/search", h.searchVacancies)
			vac.GET("/:id", h.getVacancy)
			vac.POST("/", h.createVacancy)
			vac.PATCH("/:id", h.updateVacancy)
			vac.DELETE("/:id", h.deleteVacancy)
		}
	}

	return r
}
