module github.com/brieefly

require (
	github.com/go-chi/chi v3.3.3+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/uuid v1.1.0
)

replace github.com/brieefly/config => /config

replace github.com/brieefly/util => /util

replace github.com/brieefly/log => /log

replace github.com/brieefly/net => /net

replace github.com/brieefly/model => /model

replace github.com/brieefly/model/agency => /model/agency

replace github.com/brieefly/model/market => /model/market

replace github.com/brieefly/model/panel => /model/panel

replace github.com/brieefly/model/project => /model/project

replace github.com/brieefly/ctrl/agency => /ctrl/agency

replace github.com/brieefly/ctrl/market => /ctrl/market

replace github.com/brieefly/ctrl/panel => /ctrl/panel

replace github.com/brieefly/ctrl/project => /ctrl/project

replace github.com/brieefly/ctrl/process => /ctrl/process

replace github.com/brieefly/ctrl/user => /ctrl/user

replace github.com/brieefly/db/agency => /db/agency

replace github.com/brieefly/db/panel => /db/panel

replace github.com/brieefly/db/project => /db/project

replace github.com/brieefly/db/user => /db/user
