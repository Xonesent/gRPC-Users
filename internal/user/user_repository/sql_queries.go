package user_repository

const (
	UserTableName = "user_schema.users"
	UserId        = "user_id"
	UserName      = "user_name"
)

var (
	InsertUserColumns = []string{
		UserName,
	}
	GetUserColumns = []string{
		UserId,
		UserName,
	}
)
