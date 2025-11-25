package model

import (
	"database/sql"
	"time"

	"github.com/Lina3386/auth/pkg/user"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Role      user.Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserToCreate struct {
	Name     string
	Email    string
	Password string
	Role     user.Role
}

type UserToUpdate struct {
	Id    int64
	Name  *wrapperspb.StringValue
	Email *wrapperspb.StringValue
	Role  user.Role
}
