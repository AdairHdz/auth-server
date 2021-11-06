package entity

type User struct {
	ID string `field:"id" json:"id"`
	UserID string `field:"user_id" json:"userId"`
	Names string `field:"names" json:"names"`
	LastName string `field:"lastName" json:"lastName"`
	StateID string `field:"state_id" json:"stateId"`
	EmailAddress string `field:"email_address" json:"emailAddress"`
	Password string `field:"password" json:"-"`
	UserType int `field:"user_type" json:"userType"`
	Verified bool `field:"verified" json:"verified"`
	Token string
}