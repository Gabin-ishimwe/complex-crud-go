package auth

func HashPassword(password string) (string, error) {
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// if err != nil {
	// 	return "", err
	// }
	return password, nil
}
