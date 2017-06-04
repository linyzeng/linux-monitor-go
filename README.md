
# See check for their release status
# CURRENT : working on logging
# NEXT    : stats

## nagios-plugins-go : Nagios plugins written in GO

### Background
I been writting nagios plugins for years (who remember netsaint?), most in
bash, perl and python and now I'm planning to re-write all these in Go and
make them public

The reason of my choice for Go, is simple, I wanted a single binary and able
to use a single configuration file. A side effect is you could use this code
as a nagios-plugin framework without to have re-investing the wheel :)

### Single configuration
An other thing was the check's flags, sometime there are a lot of them, so now instead
of having these given on the command-line, they are now defined in the configuration file,
Example `warning` instead of `-w value`, just set `warning: value` in configuration file.

My choice for the file format is yaml. The reason is very simple, is easy to read
and to create. Given the flag `-setup`, the check will show you what are the available
configurations such as threshlold name. Some flags can be query with the keyword `help`,
example `-mode help` to show all the valid modes and the required configuration keys name.

For each check create a section `check-name:` and under it add the configuration value,
such warning- and critical-thresholds.
example

```
	check-momo:
	  user: momo
	  password: momo
	  warning: 10%
	  critical: 5%
```

#### Trick
With the flag `-config` you can use the same check for different needs. An other neat trick
is that could copy the same binary and now you can have single configuration for the same check.
Example:
	ln -s check-a check-b
	ln -s check-a check-c
the configuration would then look like this
```
check-a:
  username: momo
check-b:
  username: mimi
check-c:
  username: mumu
```

Here are the shared configurarion:
seems like a lot, but you should only need to configure these once, or disable the one
you do not care about or use the default.

`Values shown are the default values. Any section can be ommited, it will then use the default values.`
```html
<style="color:red;">
if the key's value to disabled a section is shown empty, then the section is disable by default
```
```
	common:
	  nolog: false
	  debug: false
	  noalert: false
	# to disable set an empty `logfile`.
	log:
	  logdir: /var/log/nagios-plugins-go
	  logfile: check-mysql.log
	  logmaxsize: 128
	  logmaxbackups: 3
	  logmaxage: 10
	# to disable set an empty `statsfile`.
	stats:
	  statsdir: /var/log/nagios-plugins-go-stats
	  statsfile: check-mysql.stats
	# to disable set an empty `emailto`.
	email:
	  emailfrom:
	  emailto:
	  emailpass:
	  emailhost: localhost
	  emailfromname:
	  emailtoname:
	  emailsubjecttag: [MONITOR]
	  emailuser:
	  emailhostport: 25
	# to disable set an empty `tagfile`.
	tag:
	  tagfile:
	  tagkeyname:
	# to disable set `syslogtag: off`.
	syslog:
	  syslogtag: [{name-of-check}]
	  syslogpriority: LOG_INFO
	  syslogfacility: LOG_SYSLOG
	# to disable set an empty `pdservicekey`.
	pagerduty:
	  pdservicekey:
	  pdservicename:
	  pdvalidunit: hour
	  pdevent: MONITOR ALERT
	# to disable set an empty `slackservicekey`.
	slack:
	  slackservicekey:
	  slackchannel:
	  slackuser: MONITOR
	  iconemoji: :bangbang:

NOTE
	* The key most be all lowercase!
	* Any key that has any of these charaters: ':#[]()*' in their value must be double quoted!
	* tagfile and tagkeyname are use to get the tag info by looking for the key `tagkeyname` in the
	  configured file `tagfile`, the format need to be just 'keyname value' nothing fancy!
	* pagerduty `pdvalidunit` is the unit used to create an event-id so no duplicate is created.
	  valid choices are hour or minute, so a even create at hour X (or minute X) will result that
	  pagerduty will not create a new event, it sees it as an update to the previous event,
	  but do realize there always the possiblity that it could overlaps.
	  If the `pdvalidunit` is invalid then it defaults to hour, valid options are `hour` and `minute`.
	* `emailsubjecttag` is use for email filtering.
	* Syslog Valid `syslogpriority`: LOG_EMERG LOG_CRIT LOG_ERR LOG_WARNING LOG_NOTICE LOG_INFO LOG_DEBUG LOG_ALERT
	* Syslog Valid `syslogfacility`: LOG_LOCAL3 LOG_LOCAL5 LOG_AUTH LOG_SYSLOG LOG_NEWS LOG_CRON
		LOG_AUTHPRIV LOG_LOCAL0 LOG_DAEMON LOG_FTP LOG_LOCAL1 LOG_LOCAL2 LOG_LOCAL4
		LOG_MAIL LOG_LPR LOG_UUCP LOG_LOCAL6 LOG_LOCAL7
```


### Checks
This is the list of check I plan to build:

check-cert status `not started yet` 	: check cert expiration

check-fd status `not started yet` 		: check file descriptors

check-http status `not started yet`		: check http port reply

check-disk status `not started yet`		: check disk space

check-load status `not started yet`		: check system load

check-memory status `not started yet`	: check available memory

check-mysql status `first release`		: check mysql health include slave/replication

check-network status `not started yet`	: check network status such as TX, RX and error

check-nginx status `not started yet`	: check nginx status

check-process status `not started yet`	: check if a given process is running, /prod basesd

check-psql status `not started yet`		: check mysql health include slave/replication

Any other that you would like to see? shoot me an email

### How to build

create a work directory then set GOPATH : export GOPATH=full-path-work-directory

```
go get github.com/my10c/nagios-plugins-go/check-xxxx
with xxxx the name of the check
```

that's it


### Feedback
Feedback and bug report welcome...

Enjoy, Momo
