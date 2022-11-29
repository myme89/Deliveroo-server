package model

type GroupTypeInfo struct {
	ID          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
}

type MenuRestaurantInfo struct {
	ID          int
	Name        string
	Description string
	Price       int
	Image       string
}

type RestaurantInfo struct {
	ID               int                  `json:"id"`
	IDGroupType      int                  `json:"id_group_type"`
	Tittle           string               `json:"tittle"`
	Rating           int                  `json:"rating"`
	Genre            string               `json:"genre"`
	Address          string               `json:"address"`
	ShortDescription string               `json:"short_escription"`
	Dishes           []MenuRestaurantInfo `json:"dishes"`
	Long             float32              `json:"long"`
	Lat              float32              `json:"lat"`
}

type UserLoginInfo struct {
	UserName  string `json:"user_name"`
	GroupName string `json:"group_name"`
	Token     string `json:"token"`
}
