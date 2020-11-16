package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/amsokol/http-rest-benchmarking/go/rest-std/data"
)

const addr = ":50051"

const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var _rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var _muxRand sync.Mutex

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Printf("invalid HTTP method: %s", r.Method)
			http.Error(w, "404 not found.", http.StatusNotFound)

			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("failed to read body: %v", err)
			http.Error(w, "failed read body", http.StatusBadRequest)

			return
		}
		defer r.Body.Close()

		var in data.In
		if err = in.UnmarshalJSON(body); err != nil {
			log.Printf("failed to unmarshal JSON body: %v", err)
			http.Error(w, "failed to unmarshal JSON body", http.StatusBadRequest)

			return
		}

		l := len(in.Name)
		_muxRand.Lock()
		cap := _rand.Intn(801) + 200
		_muxRand.Unlock()

		var b strings.Builder

		b.Grow(l+cap)
		b.WriteString(in.Name)

		for i := 0; i < cap; i++ {
			_muxRand.Lock()
			n := _rand.Intn(len(charset))
			_muxRand.Unlock()
			b.WriteByte(charset[n])
		}

		out := data.Out{Result: b.String()}
		d, err := out.MarshalJSON()
		if err != nil {
			log.Printf("failed to marshal to JSON: %v", err)
			http.Error(w, "failed to marshal to JSON", http.StatusBadRequest)

			return
		}

		if _, err = w.Write(d); err != nil {
			log.Printf("failed to write response: %v", err)
			http.Error(w, "failed to write response", http.StatusInternalServerError)

			return
		}
	})

	log.Printf("starting HTTP server at %s", addr)
	log.Printf("HTTP server stopped with error: %v",http.ListenAndServe(addr, nil))
}
