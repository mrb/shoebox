## Shoebox

Shoebox is a proof of concept Go project that ties together `pat`, a route handler, `noeq`, a meaningful UUID generator, and `riakpbc`, a Riak Protocol Buffers client. It currently supports three routes:

```go
func GetId(w http.ResponseWriter, req *http.Request)
  // Get a new ID
  // [GET] /id/new
  // Success response: {'id':'209130139663990784'}
  // Failure response: 500 Internal Server Error

func PostData(w http.ResponseWriter, req *http.Request)
  // Post data to Riak and get a Unique ID
  // [POST] /data
  // Body: JSON data to store
  // Success response: {'id': '209130139663990784'}
  // Failure response: 400 Bad Request

func GetData(w http.ResponseWriter, req *http.Request)
  // Get the data you've posted
  // [GET] /data/:id
  // Success response: *Your JSON Data*
  // Failure response: 404 Not Found
```

This is a trivial application that doesn't have much practical use, but it shows that `riakpbc` actually works. It also lets me experiment with `noeq` and `pat`, and get a little more Go experience.

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
