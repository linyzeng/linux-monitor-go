
# See check for their release status

## nagios-plugins-go : Nagios plugins written in GO

### Background
I been writting nagios plugins for years (who remember netsaint?), most in
bash, perl and python and now I'm planning to re-write all these in Go and
make them public

The reason of my choice for Go, is simple, I wanted a single binary

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
go get github.com/my10c/nagios-plugins-go/check_xxxx
with xxxx the name of the check
```

that's it


Enjoy, Momo
