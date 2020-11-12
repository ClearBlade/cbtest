package types

type System struct {
	PlatformURL  string `json:"platformUrl" mapstructure:"platformUrl"`
	MessagingURL string `json:"messagingUrl" mapstructure:"messagingUrl"`
	SystemKey    string `json:"systemKey" mapstructure:"systemKey"`
	SystemSecret string `json:"systemSecret" mapstructure:"systemSecret"`
}

type Developer struct {
	Email    string `json:"email" mapstructure:"email"`
	Password string `json:"password" mapstructure:"password"`
}

type User struct {
	Email    string `json:"email" mapstructure:"email"`
	Password string `json:"password" mapstructure:"password"`
}
