# Green Aluminum Zebrule (because names are hard)
Open Source golang log streaming

Alpha v0.1.0

you make a zebrule, then feed it aluminum, becase you are cruel.

The Aluminum struct is both sent to the firehose, and fmt.Printlnd for immediate notification.

Supports: AWS firehose, 

Usage
```go
//or z = zebrule.Must(zebrule.NewZebrule(...))
//config := some config structure, reflection figures out what to do
//id := name of the firehose/whatever
//the default destinations created are fatal, error, and warning
z, err := zebrule.NewZebrule(config, zebrule.NewDestination(id), zebrule.Destination{}, zebrule.Destination{})

//To redefine destinations/define undefined destinations
z.Set[TYPE](id)

aluminum = zebrule.Data{
    Type: "" //[FATAL, ERROR, WARNING, DEBUG, INFO, NOTICE]
    Aluminum: struct //struct of data, needs a bytes() method that returns a []byte (json.Marshal), or you can use aluminum.Aluminum(pre-defined struct)
}

z.Feed(aluminum)
```