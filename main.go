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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
)

const (
	AppName = "log-audit"
)

func main() {
	fmt.Println("Server Started...")
	routine := 0
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
		routine += 1
		go OpenLevelDB(eventList, routine)
		if err != nil {
			fmt.Println(err)
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

func OpenLevelDB(list *v1beta1.EventList, routine int) error {
	fmt.Printf("Routine Number %d\n", routine)
	fmt.Printf("Hello %s\n", list.Kind)
	if list == nil {
		 fmt.Println("Nil")
		 return fmt.Errorf("%s", "Empty event list")
	}
	for _, val := range list.Items {
		fmt.Println("-----------------------")
		fmt.Println(routine)
		_, err := json.MarshalIndent(val, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Println(string(mar))
		fmt.Println(val.ObjectRef)
		// fmt.Println(val.ObjectRef.Name)
		// fmt.Println(val.ObjectRef.Namespace)
		if val.ResponseObject != nil {
			fmt.Println("********************")
			var ro runtime.TypeMeta
			if err := json.Unmarshal(val.ResponseObject.Raw, &ro); err != nil {
				return err
			}
			kind := ro.GetObjectKind().GroupVersionKind()
			versionedObject, err := scheme.Scheme.New(kind)
			err = json.Unmarshal(val.ResponseObject.Raw, versionedObject)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("---", versionedObject)
			// obj := versionedObject.DeepCopyObject()
			// fmt.Println(obj)

			fmt.Println(string(val.ResponseObject.Raw))
		}
	}

	path := filepath.Join(os.TempDir(), AppName)

	if _, err := os.Stat(path); err != nil {
		err := os.Mkdir(path, 0755)
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
