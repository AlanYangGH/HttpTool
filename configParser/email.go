package configParser

type Config struct {
	HTTP struct {
		Port int `toml:"port"`
	} `toml:"http"`
	MailMailGun struct {
		ApiKey string `toml:"mm_api_key"`
		Domain string `toml:"mm_domain"`
		From   string `toml:"mm_from"`
	} `toml:"mail_mailgun"`
	Sms struct {
		Sign      string `toml:"sms_sign"`
		ContentCn string `toml:"sms_content_cn"`
		ContentEn string `toml:"sms_content_en"`
	}
	SmsTwilio struct {
		ApiUrl string `toml:"st_api_url"`
		Sid    string `toml:"st_sid"`
		Token  string `toml:"st_token"`
		From   string `toml:"st_from"`
	} `toml:"sms_twilio"`
	SmsQiXun struct {
		ApiUrl string `toml:"sq_api_url"`
		Sid    string `toml:"sq_sid"`
		Pwd    string `toml:"sq_pwd"`
	} `toml:"sms_qixun"`
}
