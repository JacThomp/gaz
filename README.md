# Green Aluminum Zebrule (because names are hard)
Open Source golang log streaming

Beta v1.0.0

you make a zebrule, then feed it aluminum, becase you are cruel.

The Aluminum struct is both sent to the firehose, and fmt.Printlnd for immediate notification.

Supports: AWS firehose,

Usage
```go
//or z = zebrule.Must(zebrule.NewZebrule(...))
//config := some config structure, reflection figures out what to do
//id := name of the firehose/whatever
//the default destinations created are fatal, error, and warning
z, err := zebrule.NewZebrule(zebrule.NewDestination(id, config), zebrule.Destination{}, zebrule.Destination{})

//To redefine destinations/define undefined destinations
z, err = z.Set[TYPE](id, config)

aluminum = zebrule.Data{
    Type: "" //[FATAL, ERROR, WARNING, DEBUG, INFO, NOTICE]
    Aluminum: struct //struct of data, needs a bytes() method that returns a []byte (json.Marshal), or you can use aluminum.Aluminum(pre-defined struct)
}

z.Feed(aluminum)
```
