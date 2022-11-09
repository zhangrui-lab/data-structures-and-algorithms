package to

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"testing"
)

var store Store

var (
	rpcEnabled     = flag.Bool("rpc", false, "use rpc callable")
	masterAddr     = flag.String("master", "", "master node address")
	port           = flag.Int64("port", 8888, "http listen port")
	persistentFile = flag.String("file", "", "data store file name. if not set, using memory only")
)

func TestUrlStore(t *testing.T) {
	flag.Parse()
	if *masterAddr == "" {
		fmt.Printf("Master Node Run on: localhost:%d. Persistent File: %s...\n", *port, *persistentFile)
		if *persistentFile == "" {
			log.Fatalln("master nod need persistent file!")
		}
		store = NewUrlStore(*persistentFile)
		if *rpcEnabled {
			err := rpc.RegisterName("Store", store)
			if err != nil {
				log.Fatalln(err)
			}
			rpc.HandleHTTP()
		}
	} else {
		fmt.Printf("Slave Node Run on: localhost:%d. Master Node: %s...\n", *port, *masterAddr)
		store = NewUrlStoreProxy(*masterAddr)
	}
	startServer()
}

func startServer() {
	http.Handle("/", http.HandlerFunc(redirect))
	http.Handle("/add", http.HandlerFunc(add))
	err := http.ListenAndServe(":"+strconv.FormatInt(*port, 10), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var url string
	key := r.URL.Path[1:]
	err := store.Get(&key, &url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func add(w http.ResponseWriter, r *http.Request) {
	const form = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
	if r.Method == http.MethodGet {
		_, _ = fmt.Fprint(w, form)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		_, _ = w.Write(bytes.NewBufferString("params url not found").Bytes())
	}
	var key string
	err := store.Put(&url, &key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte(key))
}
