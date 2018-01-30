package main

import (
	"fmt"
	"net/http"
	"html"
	"io/ioutil"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/appscode/go/log"
	"os"
	"path/filepath"
	"encoding/json"
	"k8s.io/apiserver/pkg/apis/audit/v1beta1"
)

const (
	AppName = "log-audit"
)

func main() {
	fmt.Println("Server Started...")

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/events" {
			http.NotFound(w, r)
			return
		}
		fmt.Println("hello request")
		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(resp))
		eventList := &v1beta1.EventList{}
		err = json.Unmarshal(resp, eventList)
		if err != nil {
			log.Fatalln(err)
		}
		// err = OpenLevelDB(eventList)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.HandleFunc("/get-logs", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/get-logs" {
			return
		}
		fmt.Println("hello request222")
		fmt.Fprintf(writer, "Hello %q", html.EscapeString(request.URL.Path))
	})


	log.Fatal(http.ListenAndServe(":8080", nil))
}

func OpenLevelDB(list *v1beta1.EventList) error {
	fmt.Printf("Hello %s\n", list.Kind)
	if list == nil {
		 fmt.Println("Nil")
		 return fmt.Errorf("%s", "Empty event list")
	}
	for i, val := range list.Items {
		fmt.Println(i)
		fmt.Println(val)
	}

	path := filepath.Join(os.TempDir(), AppName)

	if _, err := os.Stat(path); err != nil {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return err
	}
	defer db.Close()


	return nil
}
