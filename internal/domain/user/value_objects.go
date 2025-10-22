package user

import (
	"database/sql/driver"
	"errors"
	"regexp"
)

type Email struct {
	Address string
}

func NewEmail(address string) (string, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,}$`)
	if !re.MatchString(address) {
		return "", errors.New("invalid email format")
	}
	return address, nil
}

// GORM compatibility
func (e Email) Value() (driver.Value, error) {
	return e.Address, nil
}

func (e *Email) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid email type")
	}
	e.Address = str
	return nil
}
