package user_repository

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	"users/pkg/models"
)

type UserRepo struct {
	db       *sqlx.DB
	txGetter *trmsqlx.CtxGetter
}

func NewClientPGRepository(db *sqlx.DB, txGetter *trmsqlx.CtxGetter) *UserRepo {
	return &UserRepo{db: db, txGetter: txGetter}
}

func (r *UserRepo) CreateUser(ctx context.Context, UserParams models.GrpcAddUser) (int64, error) {
	query, args, err := sq.Insert(UserTableName).
		Columns(InsertUserColumns...).
		Values(
			UserParams.Person,
		).
		Suffix("RETURNING " + UserId).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	var userID int64
	tr := r.txGetter.DefaultTrOrDB(ctx, r.db)
	err = tr.QueryRowContext(ctx, query, args...).Scan(&userID)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (r *UserRepo) GetUser(ctx context.Context, UserParams models.GrpcGetUser) (string, error) {
	query, args, err := sq.Select(GetUserColumns...).
		From(UserTableName).
		Where(sq.Eq{UserId: UserParams.UserId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return "", err
	}

	var result string

	tr := r.txGetter.DefaultTrOrDB(ctx, r.db)
	err = tr.GetContext(
		ctx,
		&result,
		query,
		args...,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", sql.ErrNoRows
		}

		return "", err
	}
	return result, nil
}
