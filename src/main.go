package main

// https://ashishb.net/tech/docker-101-a-basic-web-server-displaying-hello-world/
// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
// https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/
// https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker


import (
    "fmt"
    "log"
    "net/http"
    "os"
    "sync"
)

// default noun
var messageTo = "World"

// built into binary using ldflags
var Version string
var BuildTime string

// request count for this container
var counter int
var mutex = &sync.Mutex{}
func incrementCounter() {
    mutex.Lock()
    counter++
    mutex.Unlock()
}


func StartWebServer() {

    // handlers
    http.HandleFunc("/healthz", handleHealth)
    http.HandleFunc("/shutdown", handleShutdown)

    // APP_CONTEXT defaults to root
    appContext := getenv("APP_CONTEXT","/")
    log.Printf("app context: %s", appContext)
    http.HandleFunc(appContext, handleApp)

    port := getenv("PORT","8080")
    log.Printf("Starting web server on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        panic(err)
    }

}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type","application/json")
    fmt.Fprintf(w, "{\"health\":\"ok\", \"Version\":\"%s\", \"BuildTime\":\"%s\"}", Version, BuildTime )
}

func handleApp(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type","text/plain")

    // print main hello message
    fmt.Fprintf(w, "Hello, %s\n", messageTo)

    // writes count and path
    mainMsgFormat := "request %d %s %s\n"
    log.Printf(mainMsgFormat, counter, r.Method, r.URL.Path)
    fmt.Fprintf(w, mainMsgFormat, counter, r.Method, r.URL.Path)

    // 'Host' header is promoted to Request.Host field and removed from Header map
    fmt.Fprintf(w, "Host: %s\n", provideDefault(r.Host,"empty"))

    incrementCounter()
}

// provide default for value
func provideDefault(value,defaultVal string) string {
  if len(value)==0 { 
    return defaultVal
  }
  return value
}
// pull from OS environment variable, provide default
func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
// non-graceful and abrupt exit
func handleShutdown(w http.ResponseWriter, r *http.Request) {
    log.Printf("About to abruptly exit")
    os.Exit(0)
}

func main() {
    StartWebServer()
}
