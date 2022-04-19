package controllers

import (
	"log"
	"net/http"
	"strconv"

	"reakgo/utility"
)

func AddForm(w http.ResponseWriter, r *http.Request){
    if (r.Method == "POST") {
        err := r.ParseForm()
        if (err != nil){
            log.Println(err)
        } else {
            // Logic
            name := r.FormValue("name")
            address := r.FormValue("address")

            log.Println(name, address)
            Db.formAddView.Add(name, address)
        }
    }
    utility.RenderTemplate(w, r, "addForm", nil) 
}


func ViewForm(w http.ResponseWriter, r *http.Request){
    if(r.FormValue("action") == "delete"){
        id := r.FormValue("id")
        intId, _ := strconv.Atoi(id)
        Db.formAddView.Delete(intId)
    }
    result, err := Db.formAddView.View()
    if (err != nil){
        log.Println(err)
    }
    utility.RenderTemplate(w, r, "viewForm", result)
}
