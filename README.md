# SSensor

The sensor's server side

# Installation

```Shell
go get github.com/zjwdmlmx/ssensor
```

# Run

Run in GNU/Linux

```Shell
# redis-server
# $GOPATH/bin/cleaner
# $GOPATH/bin/ssensor
```

or just

```Shell
supervisord -c ./superisord.conf
```

Run with Docker

```Shell
$ docker built -t ssensor --rm 
$ docker run -d ssensor 
``` 
