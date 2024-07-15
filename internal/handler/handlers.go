package handler

type Handlers struct {
	HealthHandler HealthHandler
	UserHandler   UserHandler
	GradeHandler  GradeHandler
}
