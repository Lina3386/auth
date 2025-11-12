package user

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Lina3386/auth/internal/model"
	modelRepo "github.com/Lina3386/auth/internal/repository/user/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableName = "user"

	idColumn         = "id"
	nameColumn       = "name"
	emailColumn      = "email"
	passwordCollumun = "password"
	roleColumn       = "role"
	createdAtColumn  = "created_at"
	updatedAtColumn  = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func (r *repo) Create(ctx context.Context, req *model.UserToCreate) (int64, error) {
	builder := sq.Insert(tableName).PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, createdAtColumn, updatedAtColumn).
		Values(req.Name, req.Email, genPassHash(req.Password), req.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) Get(ctx context.Context, req *model.UserToUpdate) (int64, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var user modelRepo.User
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, req *model.UserToUpdate) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: req.Id})
	if req.Name != nil {
		builder = builder.Set(nameColumn, req.Name.Value)
	}

	if req.Email != nil {
		builder = builder.Set(emailColumn, req.Email.Value)
	}

	if req.Role != nil {
		builder = builder.Set(roleColumn, req.Role.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": id})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func genPassHash(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return fmt.Sprintf("%x", h.Sum(nil))
}
