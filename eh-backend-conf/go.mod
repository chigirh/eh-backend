module conf

go 1.18

require (
	adapter v0.0.0-00010101000000-000000000000
	app v0.0.0-00010101000000-000000000000
	github.com/google/wire v0.5.0
	github.com/labstack/echo v3.3.10+incompatible
	gopkg.in/ini.v1 v1.66.4
)

require (
	domain v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.6 // indirect
)

replace adapter => ./../eh-backend-adapter

replace app => ./../eh-backend-app

replace domain => ./../eh-backend-domain
