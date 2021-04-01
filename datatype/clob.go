package datatype

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/Insua/dm"
)

type Clob string

func (c *Clob) Scan(value interface{})  error {
	clob, ok := value.(*dm.DmClob)
	if !ok {
		return errors.New(fmt.Sprint("Failed to parse dm clob value:", value))
	}
	length, err := clob.GetLength()
	if err != nil {
		return errors.New(fmt.Sprint("Failed to get dm clob length:", value))
	}
	s, err := clob.ReadString(1, int(length))
	if err != nil {
		return errors.New(fmt.Sprint("Failed to read dm clob string:", value))
	}
	*c = Clob(s)
	return nil
}

func (c Clob) Value() (driver.Value, error)  {
	return string(c), nil
}

func (Clob) GormDataType() string {
	return "text"
}
