package config

type Logger struct {
	Level          *string `mapstructure:"level"           json:"level"          validate:"required,oneof=PANIC FATAL ERROR WARN INFO DEBUG"` //nolint:lll
	Encoding       *string `mapstructure:"encoding"        json:"encoding"       validate:"required,oneof=json plain"`
	StdoutEnabled  *bool   `mapstructure:"stdout_enabled"  json:"stdoutEnabled"  validate:"required"`
	SyslogEnabled  *bool   `mapstructure:"syslog_enabled"  json:"syslogEnabled"  validate:"required"`
	SyslogFacility *string `mapstructure:"syslog_facility" json:"syslogFacility" validate:"required,oneof=KERN USER MAIL DAEMON AUTH SYSLOG LPR NEWS UUCP CRON AUTHPRIV FTP LOCAL0 LOCAL1 LOCAL2 LOCAL3 LOCAL4 LOCAL5 LOCAL6 LOCAL7"` //nolint:lll
	SyslogTag      *string `mapstructure:"syslog_tag"      json:"syslogTag"      validate:"required,notblank"`
}
