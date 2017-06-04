
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
An other thing was configure the flags in a single file, my choice is yaml. The reason
is very simple, is easy to read and to create. Given the flag `--setup`, the check will 
show you what available configuration such as threshlold name that are available.
Other flags option can be query with `help`, example `--mode help` to show valid modes
options.

For each check create a section `check-name:` and under the configs, such as threshold.
example

```
check-momo:
  user: momo
  password: momo
  warning: 10%
  critical: 5%
```

Here are the shared configurarion:
seems like a lot, but you should only need to configure these once, or disable some of the
feature such as syslog, email, pagerdurty and slack. 

```
	# Optional add these values in the common section.
	# Values shown are the default values. If either emailfrom or emailto is empty then no email will be sent.
	# tagfile and tagkeyname are use to get the tag info by looking for the key tagkeyname in the
	# configured file tagfile, the format need to be just 'keyname value' nothing fancy!
common:
  debug: false
  emailhostport: 25
  emailhost: localhost
  emailto:
  emailtoname:
  emailfrom:
  emailfromname:
  emailuser:
  emailpass:
  tagfile:
  tagkeyname:
  noalert: false
  nolog: false
  logmaxsize: 128
  logmaxbackups: 3
  logmaxage: 10
  logdir: /var/log/nagios-plugins-go
  logfile: /var/log/nagios-plugins-go/{check name}.log

	# Syslog support, to disable set tag value to off or leave it empty, syslogtag default to check name.
syslog:
  syslogtag:
  syslogfacility: LOG_SYSLOG
  syslogpriority: LOG_INFO

	# Optional for pagerduty support, if pdservicekey and/or pdservicename is keys are empty then pagerduty is not used.
pagerduty:
  pdservicekey:
  pdservicename:
  pdservicename: MONITOR
  pdvalidunit: hour
  pdevent: [MONITOR]

	# Optional for slack support, if slackservicekey and/or slackchannel is empty then slack is not used.
slack:
  slackservicekey:
  slackchannel:
  iconemoji: ":bangbang::doge:"

NOTE
	* Any key that has any of these charaters: '#[]()*' in their value must be double quoted!
	* Syslog Valid Priority: LOG_ALERT LOG_CRIT LOG_ERR LOG_NOTICE LOG_INFO LOG_EMERG LOG_WARNING LOG_DEBUG
	* Syslog Valid Facility: LOG_DAEMON LOG_LPR LOG_CRON LOG_LOCAL1 LOG_LOCAL2 LOG_LOCAL7
		LOG_SYSLOG LOG_NEWS LOG_LOCAL0 LOG_UUCP LOG_FTP LOG_LOCAL3 LOG_MAIL
		LOG_AUTH LOG_AUTHPRIV LOG_LOCAL4 LOG_LOCAL5 LOG_LOCAL6
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
