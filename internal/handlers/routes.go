package handlers

import (
	"github.com/gin-gonic/gin"
	"haha/internal/logger"
)

type Handler struct {
	services *Service
	logg     *logger.Logger
}

func NewHandler(services *Service, logg *logger.Logger) *Handler {
	return &Handler{services: services, logg: logg}
}

//Role
//anon - не авторизированный
//applicant - соискатель
//employer - работодатель
//admin - админ

func (h *Handler) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (h *Handler) InitRouter(URL string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(h.CORSMiddleware())

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		//auth.GET("/info", h.getAllUser)
	}

	apiAnon := r.Group("/api/v1")
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

	api := r.Group("/api/v1", h.userIdentity)
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
