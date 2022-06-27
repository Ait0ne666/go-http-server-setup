package config

type Config struct {
	HTTP        *HTTP         `yaml:"http"`
	PostgresDsn string        `yaml:"postgres_dsn"`
	JWTKey      string        `yaml:"jwtkey"`
	Telegram    *Telegram     `yaml:"telegram"`
	SMS         *Sms          `yaml:"sms"`
	Email       *ForSendEmail `yaml:"server_email"`
	Check       *Check        `yaml:"check"`
}

type HTTP struct {
	Port string `yaml:"server_port"`
}

type Telegram struct {
	TelegramToken string `yaml:"telegram_token"`
	ChatID        string `yaml:"chat_id"`
}

type Sms struct {
	SMSURL    string `yaml:"sms_url"`
	SmsPass   string `yaml:"sms_pass"`
	SmsLog    string `yaml:"sms_login"`
	SmsSender string `yaml:"sms_sender"`
}

type ForSendEmail struct {
	EmailHost        string `yaml:"host"`
	EmailPort        int    `yaml:"port"`
	EmailLogin       string `yaml:"login"`
	EmailFrom        string `yaml:"from"`
	EmailPass        string `yaml:"pass"`
	EmailUnsubscribe string `yaml:"email_unsubscribe"`
	NameSender       string `yaml:"email_name_sender"`
	EmailSender      string `yaml:"email_sender"`
}

type Check struct {
	Url   string `yaml:"url"`
	Token string `yaml:"token"`
}
