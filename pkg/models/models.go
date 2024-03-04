package models

type GrpcAddUser struct {
	Person string
}

type GrpcGetUser struct {
	UserId int64
}

type DataUser struct {
	UserID int64  `db:"user_id"`
	User   string `db:"user_name"`
}
