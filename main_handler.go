package main

import("fmt"
       "net/http"
       "log"
      "strings)
       
        type  cocktails struct {}

      func drinks(w http.ResponseWriter, req * http.Request) {  
        id := strings.TrimPrefix(req.URL.Path, "/provisions/")
        //will be for name cocktails
        name := strings.TrimPrefix(req.URL.Path, "/provisions/")
	      fmt.Fprintf(w, "hello friend!\n")
       }
       
       func main(){
          var c cocktails
           mux := http.NewServeMux()
           mux.handle(“/drinks/”, c)
         
           http.ListenAndServe(“:80”, mux)
       }
