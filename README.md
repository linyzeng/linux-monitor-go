
# UNSTABLE VERY ALPHA

## nagios-plugins-go : Nagios plugins written in GO

### Background
I been writting nagios plugins for years (who remember netsaint?), must in
bash, perl and python and now I planning to re-write all these in Go and
make them public

The reason of my choice for Go, is simple, I wanted a single binary


### How to build

create a work directory then set GOPATH : export GOPATH=full-path-work-directory

```
go get github.com/my10c/nagios-plugins-go/check_xxxx
with xxxx the name of the check
```

that's it


Enjoy, Momo
