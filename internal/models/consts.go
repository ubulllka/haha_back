package models

const (
	AuthorizationHeader = "Authorization"
	UserIdCtx           = "userId"
	UserRoleCtx         = "userRole"

	ANON      = "ANON"
	APPLICANT = "APPLICANT"
	EMPLOYER  = "EMPLOYER"
	ADMIN     = "ADMIN"

	ACTIVE  = "ACTIVE"
	PASSIVE = "PASSIVE"
	NO      = "NO"

	ACCEPT  = "ACCEPT"
	DECLINE = "DECLINE"
	WAIT    = "WAIT"

	PARSEDATE = "02-01-2006"
)

//Role
//anon - не авторизированный
//applicant - соискатель
//employer - работодатель
//admin - админ

//Status find work
//active - активно ищу работу
//passive - рассматриваю предложения
//no - не ищу работу

//Status answer
//accept - принято
//decline - отклонено
//wait - ожидание
