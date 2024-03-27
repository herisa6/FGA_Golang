package models

type User struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Age            int    `json:"age"`
	Password       string `json:"password"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Photo struct {
	ID          int    `json:"id"`
	Caption     string `json:"caption"`
	Title       string `json:"title"`
	PhotoURL    string `json:"photo_url"`
	UserID      int    `json:"user_id"`
	User        *User  `json:"user,omitempty"`
}

type Comment struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	PhotoID  int    `json:"photo_id"`
	UserID   int    `json:"user_id"`
	User     *User  `json:"user,omitempty"`
	Photo    *Photo `json:"photo,omitempty"`
}

type SocialMedia struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
	User           *User  `json:"user,omitempty"`
}
