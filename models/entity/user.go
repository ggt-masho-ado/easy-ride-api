package entity

type User struct {
	ID               int    `json:"id"`
	Full_name        string `json:"full_name"`
	Email            string `json:"email"`
	Home_coordinates string `json:"home_coordinates"`
	Home_address     string `json:"home_address"`
	Is_active        int16  `json:"is_active"`
	Created_at       string `json:"created_at"`
}
