package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Insua/dm"
)

type JSON json.RawMessage

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	bytes, err := json.RawMessage(j).MarshalJSON()
	return string(bytes), err
}

func (j *JSON) Scan(value interface{}) error {
	clob, ok := value.(*dm.DmClob)
	if !ok {
		return errors.New(fmt.Sprint("Failed to parse dm clob value:", value))
	}
	length, err := clob.GetLength()
	if err != nil {
		return errors.New(fmt.Sprint("Failed to get dm clob length:", value))
	}
	s := ""
	if length > 0 {
		s, err = clob.ReadString(1, int(length))
		if err != nil {
			return errors.New(fmt.Sprint("Failed to read dm clob string:", value))
		}
	}
	bytes := []byte(s)

	result := json.RawMessage{}
	err = json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	result := json.RawMessage{}
	err := result.UnmarshalJSON(b)
	*j = JSON(result)
	return err
}

func (j JSON) String() string {
	return string(j)
}

func (JSON) GormDataType() string {
	return "text"
}
