package packageOne

type User struct {
	ID       string
	Name     string
	Username string
	Password string
}

type Order struct {
	ID     string
	UserID string
}
