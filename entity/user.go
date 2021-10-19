package entity

type User struct {
	ID string `field:"id"`
	UserID string `field:"user_id"`
	Names string `field:"names"`
	LastName string `field:"lastName"`
	StateID string `field:"state_id"`
	EmailAddress string `field:"email_address"`
	Password string `field:"password" json:"-"`
	UserType int `field:"user_type"`
	Verified bool `field:"verified"`
	Token string
}