package message

const (
	// KindConnected is sent when user connects
	KindConnected = iota + 1
	// KindUserJoined is sent when someone else joins
	KindUserJoined
	// KindUserLeft is sent when someone leaves
	KindUserLeft
	// KindStroke message specifies a drawn stroke by a user
	KindStroke
	// KindClear message is sent when a user clears the screen
	KindClear
)

// Point ...
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// User ...
type User struct {
	ID    string `json:"id"`
	Color string `json:"color"`
}

// Connected ...
type Connected struct {
	Kind  int    `json:"kind"`
	Color string `json:"color"`
	Users []User `json:"users"`
}

// UserJoined ...
type UserJoined struct {
	Kind int  `json:"kind"`
	User User `json:"user"`
}

// UserLeft ...
type UserLeft struct {
	Kind   int    `json:"kind"`
	UserID string `json:"userId"`
}

// Stroke ...
type Stroke struct {
	Kind   int     `json:"kind"`
	UserID string  `json:"userId"`
	Points []Point `json:"points"`
	Finish bool    `json:"finish"`
}

// Clear ...
type Clear struct {
	Kind   int    `json:"kind"`
	UserID string `json:"userId"`
}

// NewConnected ...
func NewConnected(color string, users []User) *Connected {
	return &Connected{
		Kind:  KindConnected,
		Color: color,
		Users: users,
	}
}

// NewUserJoined ...
func NewUserJoined(userID string, color string) *UserJoined {
	return &UserJoined{
		Kind: KindUserJoined,
		User: User{ID: userID, Color: color},
	}
}

// NewUserLeft ...
func NewUserLeft(userID string) *UserLeft {
	return &UserLeft{
		Kind:   KindUserLeft,
		UserID: userID,
	}
}
