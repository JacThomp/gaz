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
//Seperator := seperation string (EX. "====="), seperates entries in the stream (for human readability)
//id := name of the firehose/whatever
//the default destinations created are fatal, error, and warning
z, err := zebrule.NewZebrule(config, zebrule.NewDestination(typeOf, id, seperator), zebrule.Destination{}, zebrule.Destination{})

//To redefine destinations/define undefined destinations
z.Set[TYPE](typeOf, destination, id)

aluminum = zebrule.Aluminum{
    Type: "" //[FATAL, ERROR, WARNING, DEBUG, INFO, NOTICE]
    Data: struct //struct of data, needs a String() method that returns a []byte (json.Marshal)
}

z.Feed(aluminum)
```