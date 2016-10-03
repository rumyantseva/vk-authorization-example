package models

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type authTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *authTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("auth").
func (v *authTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *authTableType) Columns() []string {
	return []string{"id", "email", "password", "vk_user_id", "vk_token", "vk_token_lifetime", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *authTableType) NewStruct() reform.Struct {
	return new(Auth)
}

// NewRecord makes a new record for that table.
func (v *authTableType) NewRecord() reform.Record {
	return new(Auth)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *authTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// AuthTable represents auth view or table in SQL database.
var AuthTable = &authTableType{
	s: parse.StructInfo{Type: "Auth", SQLSchema: "", SQLName: "auth", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int", Column: "id"}, {Name: "Email", PKType: "", Column: "email"}, {Name: "Password", PKType: "", Column: "password"}, {Name: "VKUserID", PKType: "", Column: "vk_user_id"}, {Name: "VKToken", PKType: "", Column: "vk_token"}, {Name: "VKTokenLifetime", PKType: "", Column: "vk_token_lifetime"}, {Name: "CreatedAt", PKType: "", Column: "created_at"}, {Name: "UpdatedAt", PKType: "", Column: "updated_at"}}, PKFieldIndex: 0},
	z: new(Auth).Values(),
}

// String returns a string representation of this struct or record.
func (s Auth) String() string {
	res := make([]string, 8)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Email: " + reform.Inspect(s.Email, true)
	res[2] = "Password: " + reform.Inspect(s.Password, true)
	res[3] = "VKUserID: " + reform.Inspect(s.VKUserID, true)
	res[4] = "VKToken: " + reform.Inspect(s.VKToken, true)
	res[5] = "VKTokenLifetime: " + reform.Inspect(s.VKTokenLifetime, true)
	res[6] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[7] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Auth) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Email,
		s.Password,
		s.VKUserID,
		s.VKToken,
		s.VKTokenLifetime,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Auth) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Email,
		&s.Password,
		&s.VKUserID,
		&s.VKToken,
		&s.VKTokenLifetime,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Auth) View() reform.View {
	return AuthTable
}

// Table returns Table object for that record.
func (s *Auth) Table() reform.Table {
	return AuthTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Auth) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Auth) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Auth) HasPK() bool {
	return s.ID != AuthTable.z[AuthTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Auth) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int(i64)
	} else {
		s.ID = pk.(int)
	}
}

// check interfaces
var (
	_ reform.View   = AuthTable
	_ reform.Struct = new(Auth)
	_ reform.Table  = AuthTable
	_ reform.Record = new(Auth)
	_ fmt.Stringer  = new(Auth)
)

func init() {
	parse.AssertUpToDate(&AuthTable.s, new(Auth))
}
