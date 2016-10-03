package config

type Config struct {
	Server   ServerConf `toml:"server"`
	Database DBConf     `toml:"database"`
	VK       VKConf     `toml:"vk"`
}

type ServerConf struct {
	Address string `toml:"address"`
}

type VKConf struct {
	ClientID         int    `toml:"client_id"`
	ClientSecret     string `toml:"client_secret"`
	Version          string `toml:"version"`
	OAuthRedirectURI string `toml:"oauth_redirect_uri"`
}

type DBConf struct {
	Name     string `toml:"name"`
	Address  string `toml:"address"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}
