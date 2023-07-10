package common

import "database/sql/driver"

type Price []byte

func (p *Price) Value() (driver.Value, error) {
	return []byte{1, 2}, nil
}

func (p *Price) Scan(val interface{}) error {
	v, _ := val.([]byte)
	*p = v
	return nil
}


//func (p Price) Value() (driver.Value, error) {
//	if p == "" {
//		return nil, nil
//	}
//	buf, err := strconv.Atoi(string(p))
//	return int64(buf), err
//}
//
//func (p *Price) Scan(val interface{}) error {
//	v, _ := val.(int64)
//	buf := strconv.Itoa(int(v))
//	*p = Price(buf)
//	return nil
//}
