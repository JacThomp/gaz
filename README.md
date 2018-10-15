# Green Aluminum Zebrule (because names are hard)
Open Source golang log streaming

Super alpha v0.0.0

you make a zebrule, then feed it aluminum, becase you are cruel.

Supports: AWS firehose, 

Usage
```go
//or z = zebrule.Must(zebrule.NewZebrule(...))
//config := some config structure
//typeOf := supported stream string ("AWS", )
//destination := destination string (/example/log/fatal) unsupported by AWS, the hose itself determines where things go
//id := name of the firehose/whatever
//the default destinations created are fatal, error, and warning
z, err := zebrule.NewZebrule(config, zebrule.NewDestination(typeOf, destination, id), nil, nil)

//To redefine destinations/define undefined destinations
z.Set[TYPE](typeOf, destination, id)

aluminum = zebrule.Aluminum{
    Type: "" //[FATAL, ERROR, WARNING, DEBUG, INFO, NOTICE]
    Data: "" //string that gets streamed as a byte array
}

z.Feed(aluminum)
```