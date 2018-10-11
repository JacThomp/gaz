# Green Aluminum Zebrule (because names are hard)
Open Source golang log streaming

you make a zebrule, then feed it aluminum
becase you are cruel.

Supports: AWS firehose, 

Usage
```go
//or z = zebrule.Must(zebrule.NewZebrule(...))
//config := some config structure
//typeOf := supported stream string ("AWS", )
//destination := destination string (/example/log/fatal)
//the default destinations created are fatal, warning, and error
z, err := zebrule.NewZebrule(config, zebrule.NewDestination(typeOf, destination), zebrule.NewDestination(), zebrule.NewDestination())


aluminum = zebrule.Aluminum{
    Type: "" //[FATAL, ERROR, WARNING, DEBUG, INFO, NOTICE]
    Data: "" //string that gets jsonified and streamed
}

z.Feed(aluminum)
```