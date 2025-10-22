package user

type User struct {
	ID    uint
	Email Email
	// Password PasswordHash
}

func NewUser(email, password string) (*User, error) {
	emailVO, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	// hash, err := HashPassword(password)
	// if err != nil {
	//  return nil, err
	// }

	return &User{
		Email: emailVO,
		// Password: hash,
	}, nil
}