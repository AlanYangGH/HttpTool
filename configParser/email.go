package configParser

type Config struct {
	HTTP struct {
		Port int `toml:"port"`
	} `toml:"http"`
	MailGun struct {
		ApiKey string `toml:"api_key"`
		Domain string `toml:"domain"`
		From   string `toml:"from"`
	} `toml:"mail_gun"`
}
