

package main

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id    int
    Task  string
    Assignee string
    Dateline string
    Mark  string
    
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "user"
    dbPass := "password"
    dbName := "tugas"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT id,task,assignee,dateline,IF(STRCMP(mark,'N'),'Done','Mark As Done') as mark FROM Employee1 ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    res := []Employee{}
    for selDB.Next() {
        var id int
        var task, assignee, dateline, mark string
        err = selDB.Scan(&id, &task, &assignee, &dateline,&mark)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Task = task
        emp.Assignee = assignee
        emp.Dateline = dateline
        emp.Mark = mark

        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT id,task,assignee,dateline, mark FROM Employee1 WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var task, assignee, dateline, mark string
        err = selDB.Scan(&id, &task, &assignee, &dateline,&mark)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Task = task
        emp.Assignee = assignee
        emp.Dateline = dateline
        emp.Mark = mark

    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT id,task,assignee,dateline, mark FROM Employee1 WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var task, assignee, dateline, mark string
        err = selDB.Scan(&id, &task, &assignee, &dateline,&mark)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Task = task
        emp.Assignee = assignee
        emp.Dateline = dateline
        emp.Mark = mark

    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        task := r.FormValue("task")
        assignee := r.FormValue("assignee")
        dateline := r.FormValue("dateline")
        mark := r.FormValue("mark")
        insForm, err := db.Prepare("INSERT INTO Employee1(task, assignee,dateline,mark) VALUES(?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(task, assignee, dateline,mark)
        log.Println("INSERT: Task: " + task + " | Assignee: " + assignee)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        task := r.FormValue("task")
        assignee := r.FormValue("assignee")
        dateline := r.FormValue("dateline")
        mark := r.FormValue("mark")

        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE Employee1 SET task=?, assignee=?, dateline=?, mark=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(task, assignee, dateline,mark, id)
        log.Println("UPDATE: Task: " + task + " | Assignee: " + assignee)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Mark(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("UPDATE Employee1 SET mark='Y' WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("update")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Employee1 WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.HandleFunc("/mark", Mark)
    http.ListenAndServe(":8080", nil)
}

