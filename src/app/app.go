package app 

import (
    "fmt"
    "log"
	"net/http"
    "github.com/gorilla/mux"
    
)
type  App struct  {
    Router  *mux.Router

}

func (a *App) Init() {
    a.Router = mux.NewRouter()
    a.SetRouters()
}

func (b *App) SetRouters() {
    b.Get("/", GetAllTasks)
    b.Get("/GetTask/{taskId}", GetTask)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
    a.Router.HandleFunc(path, f).Methods("GET")
}
 
func (a *App) Run(host string) {
    log.Fatal(http.ListenAndServe(host, a.Router))
}
func GetTask(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    taskId := vars["taskId"]
    fmt.Fprintln(w, "Task id is", taskId)
}

func  GetAllTasks(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to Go!!!")
}
