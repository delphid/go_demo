package common

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type JSONMap map[string]string

func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	ba, err := m.MarshalJSON()
	return string(ba), err
}

func (m *JSONMap) Scan(val interface{}) error {
	if val == nil {
		*m = make(JSONMap)
		return nil
	}
	var bytes []byte
	switch v := val.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to unmarshal JSONMap value: %s", val)
	}
	buf := map[string]string{}
	err := json.Unmarshal(bytes, &buf)
	*m = buf
	return err
}

func (m JSONMap) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	buf := (map[string]string)(m)
	return json.Marshal(buf)
}

func (m *JSONMap) UnmarshalJSON(b []byte) error {
	buf := map[string]string{}
	err := json.Unmarshal(b, &buf)
	*m = buf
	return err
}

func (m JSONMap) GormDataType() string {
	return "jsonmap"
}

func (JSONMap) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "VARCHAR(2000)"
}

func (m JSONMap) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	data, _ := m.MarshalJSON()
	return gorm.Expr("?", string(data))
}