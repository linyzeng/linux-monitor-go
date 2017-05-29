
# UNSTABLE VERY ALPHA

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
For each check create a section `check-name:` and under the configs, such as threshold.
example

```
check-momo:
  user: momo
  password: momo
  warning: 10%
  critical: 5%
```



### How to build

create a work directory then set GOPATH : export GOPATH=full-path-work-directory

```
go get github.com/my10c/nagios-plugins-go/check_xxxx
with xxxx the name of the check
```

that's it


Enjoy, Momo
