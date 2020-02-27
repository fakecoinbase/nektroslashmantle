package db

import (
	"database/sql"
	"strconv"

	"github.com/nektro/go-util/alias"
	"github.com/nektro/go-util/util"
	dbstorage "github.com/nektro/go.dbstorage"
)

type User struct {
	ID         int64  `json:"id"`
	Provider   string `json:"provider" sqlite:"text"`
	Snowflake  string `json:"snowflake" sqlite:"text"`
	UUID       string `json:"uuid" sqlite:"text"`
	IsMember   bool   `json:"is_member" sqlite:"tinyint(1)"`
	IsBanned   bool   `json:"is_banned" sqlite:"tinyint(1)"`
	Name       string `json:"name" sqlite:"text"`
	Nickname   string `json:"nickname" sqlite:"text"`
	JoindedOn  string `json:"joined_on" sqlite:"text"`
	LastActive string `json:"last_active" sqlite:"text"`
	Roles      string `json:"roles" sqlite:"text"`
}

//
//

func QueryUserByUUID(uid string) (*User, bool) {
	ch, ok := dbstorage.ScanFirst(db.Build().Se("*").Fr(cTableUsers).Wh("uuid", uid), User{}).(*User)
	return ch, ok
}

func QueryUserBySnowflake(provider string, flake string, name string) *User {
	us := dbstorage.ScanFirst(db.Build().Se("*").Fr(cTableUsers).Wh("provider", provider).Wh("snowflake", flake), User{})
	if us != nil {
		return us.(*User)
	}
	// else
	id := db.QueryNextID(cTableUsers)
	uid := newUUID()
	now := alias.T()
	roles := ""
	if id == 1 {
		roles += "o"
		Props.Set("owner", uid)
	}
	db.Build().Ins(cTableUsers, id, provider, flake, uid, 0, 0, name, "", now, now, roles).Exe()
	return QueryUserBySnowflake(provider, flake, name)
}

//
//

func (v User) Scan(rows *sql.Rows) dbstorage.Scannable {
	rows.Scan(&v.ID, &v.Provider, &v.Snowflake, &v.UUID, &v.IsMember, &v.IsBanned, &v.Name, &v.Nickname, &v.JoindedOn, &v.LastActive, &v.Roles)
	return &v
}

func (u *User) SetAsMember(b bool) {
	db.Build().Up(cTableUsers, "is_member", strconv.Itoa(util.Btoi(b))).Wh("uuid", u.UUID).Exe()
	u.IsMember = b
}

func (u *User) SetName(s string) {
	db.Build().Up(cTableUsers, "name", s).Wh("uuid", u.UUID).Exe()
	u.Name = s
}
