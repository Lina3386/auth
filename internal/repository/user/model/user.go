package model

import (
	"database/sql"
	"github.com/Lina3386/auth/pkg/user"
	"time"
)

type User struct {
	Id        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `ab:"email"`
	Role      user.Role    `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
