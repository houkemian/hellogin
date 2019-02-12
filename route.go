package main

func initializeRoutes() {


	router.GET("/", loadAllArticle)

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/:article_id", ensureLogin(), loadDetail)
		articleRoutes.POST("",ensureLogin(),  post)
		articleRoutes.GET("",ensureLogin(),  loadAllArticle)
	}

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", register)
		userRoutes.POST("/login", login)
	}

	labelScheduleRoutes := router.Group("/label")
	{
		labelScheduleRoutes.GET("/:id",loadLabelScheduleDetail)
		labelScheduleRoutes.GET("",loadAll)
		labelScheduleRoutes.POST("",postLabelSchedule)
	}
}
