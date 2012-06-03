## Shoebox

Shoebox is a proof of concept Go project that ties together `pat`, a route handler, `noeq`, a meaningful UUID generator, and `riakpbc`, a Riak Protocol Buffers client. It currently supports two routes:

```go
func GetId(w http.ResponseWriter, req *http.Request)
  // Get a new ID
  // [GET] /id/new

func PostData(w http.ResponseWriter, req *http.Request)
  // Post data to Riak and get a Unique ID
  // [POST] /data
  // Body: JSON data to store
```

### Requirements

#### A running Riak cluster

With a protocol buffer port listening on port 8087.

#### A running noeqd instance

Listening on port 4444.

### Running

`go build && ./shoebox`

### Credits

riakpbc is (c) Michael R. Bernstein, 2012

`pat` and `noeqd` are awesome Go libraries by Blake Mizerany. Check his work out here: https://github.com/bmizerany

### License

shoebox is distributed under the MIT License, see `LICENSE` file for details.
