package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

//このファイルはAIによる作成

// Date は日付のみを扱うカスタム型
type Date time.Time

const dateFormat = "2006-01-02"

// MarshalJSON は JSON エンコード時に "2025-12-21" 形式にする
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format(dateFormat) + `"`), nil
}

// UnmarshalJSON は JSON デコード時に "2025-12-21" 形式を受け入れる
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// Time は time.Time に変換する
func (d Date) Time() time.Time {
	return time.Time(d)
}

// Value は GORM で DB に保存する際の値を返す
func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// Scan は GORM で DB から読み込む際に使われる
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*d = Date(t)
		return nil
	}
	return fmt.Errorf("cannot scan %T into Date", value)
}
