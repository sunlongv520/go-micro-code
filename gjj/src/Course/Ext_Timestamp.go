package Course

import (
	"database/sql/driver"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"time"
)
//db  数据 =>go类型
func (this *Timestamp) Scan(value interface{}) error {
	switch t := value.(type) {
	case time.Time:
		var err error
		this.Timestamp, err = ptypes.TimestampProto(t)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("timestamp error")
	}
	return nil
}
// insert update   把go数据类型 =>db类型
func (this Timestamp) Value() (driver.Value, error) {
	return ptypes.Timestamp(this.Timestamp)
}

