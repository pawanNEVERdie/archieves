package web

import (
	"encoding/json"
	"log"
	"net/http"
	"my-blog/db" "encoding/json"
    "log"
    "net/http"
    "my-blog/db"
    "my-blog/model"
    "github.com/gorilla/mux"
    "strconv"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
	router *mux.Router 
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		router mux.NewRouter()
		handlers: make(map[string]http.HandlerFunc),
	}
	techHandler := app.GetTechnologies
	if !cors {
		techHandler = disableCors(techHandler)
	}
	app.handlers["/api/technologies"] = techHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	app.router.HandleFunc("/api/blogs", app.handleGetBlogs).Methods("GET")
    app.router.HandleFunc("/api/blog/{id:[0-9]+}", app.handleGetBlog).Methods("GET")
    app.router.HandleFunc("/api/blog/create", app.handleCreateBlog).Methods("POST")
    app.router.HandleFunc("/api/blog/update/{id:[0-9]+}", app.handleUpdateBlog).Methods("PUT")
    app.router.HandleFunc("/api/blog/delete/{id:[0-9]+}", app.handleDeleteBlog).Methods("DELETE")
    app.router.PathPrefix("/").Handler(http.FileServer(http.Dir("/webapp")))
    return app
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
func (a *App) Serve() error {
    log.Println("Web server is available on port 3001")
    return http.ListenAndServe(":3001", disableCors(a.router.ServeHTTP))
}func (a *App) handleGetBlogs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    blogs, err := a.d.GetBlogs()
    if err != nil {
        sendErr(w, http.StatusInternalServerError, err.Error())
        return
    }
    err = json.NewEncoder(w).Encode(blogs)
    if err != nil {
        sendErr(w, http.StatusInternalServerError, err.Error())
    }
}
func (a *App) handleGetBlog(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        sendErr(w, http.StatusBadRequest, "Invalid blog ID")
        return
    }
    blog, err := a.d.GetBlog(id)
    if err != nil {
        sendErr(w, http.StatusInternalServerError, err.Error())
        return
    }
    json.NewEncoder(w).Encode(blog)
}
func (a *App) handleCreateBlog(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var b model.Blog
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&b); err != nil {
        sendErr(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()
    if err := a.d.CreateBlog(&b); err != nil {
        sendErr(w, http.StatusInternalServerError, "Error creating the blog")
        return
    }
    json.NewEncoder(w).Encode(b)
}
func (a *App) handleUpdateBlog(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        sendErr(w, http.StatusBadRequest, "Invalid blog ID")
        return
    }
    w.Header().Set("Content-Type", "application/json")
    var b model.Blog
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&b); err != nil {
        sendErr(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()
    if err := a.d.UpdateBlog(id, &b); err != nil {
        sendErr(w, http.StatusInternalServerError, "Error updating the blog")
        return
    }
    json.NewEncoder(w).Encode(b)
}
func (a *App) handleDeleteBlog(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        sendErr(w, http.StatusBadRequest, "Invalid blog ID")
        return
    }
    if err := a.d.DeleteBlog(id); err != nil {
        sendErr(w, http.StatusInternalServerError, "Error deleting the blog")
        return
    }
    w.WriteHeader(http.StatusOK)
}