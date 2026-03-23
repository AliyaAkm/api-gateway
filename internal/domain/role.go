package domain

const (
	RoleStudent = "student"
	RoleTeacher = "teacher"
	RoleManager = "manager"
	RoleAdmin   = "admin"
)

var rolePriority = map[string]int{
	RoleStudent: 1,
	RoleTeacher: 2,
	RoleManager: 3,
	RoleAdmin:   4,
}

func IsValidRoleCode(code string) bool {
	_, ok := rolePriority[code]
	return ok
}
