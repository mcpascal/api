package responses

type Login struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
