package main

import (
    "fmt"
    "database/sql"
    "log"
    "encoding/json"
	"io/ioutil"
    "net/http"
	"github.com/gorilla/mux" 
	"text/template"
    _ "github.com/go-sql-driver/mysql"
)
func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "sreekanth_Go_API"
    dbPass := "Mu$@021324"
    dbName := "MyGo"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
type Product struct {
    Id    int `json:"id"`
    Title  string `json:"title"`
    Description string `json:"description"`
    
}

var tmpl = template.Must(template.ParseGlob("forms/*"))

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/",Home)
	r.HandleFunc("/p/{id}",singleProduct)
	r.HandleFunc("/new", createProduct)
	r.HandleFunc("/update/{id}", updateProduct)
	r.HandleFunc("/delete/{id}", deleteProduct)
	http.ListenAndServe(":8080",r)
	
}



// Get all products
func Home(w http.ResponseWriter, r *http.Request){
	var device Product

	log.Println("welcome",device)
	db :=dbConn()
	rows,err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}
	
	defer rows.Close()
    var deviceSlice []Product
	log.Println("deviceslice",deviceSlice)
    for rows.Next(){
        rows.Scan(&device.Id, &device.Title, &device.Description)
                  
        deviceSlice = append(deviceSlice, device)
    }
    fmt.Println(deviceSlice)
	tmpl.ExecuteTemplate(w, "Home", deviceSlice)
	json.NewEncoder(w).Encode(deviceSlice)

}

// Get by products id
func singleProduct(w http.ResponseWriter, r *http.Request){
	// var tmpl = template.Must(template.ParseGlob("forms/byid.tmpl"))
	var device Product
	db :=dbConn()
	params := mux.Vars(r)
	result,err := db.Query("SELECT id, title,description FROM products  WHERE id = ?", params["id"])
	defer result.Close()
	
    if err != nil {
      panic(err.Error())
    }
	var deviceSlice []Product
	for result.Next() {
		result.Scan(&device.Id, &device.Title, &device.Description)
		
		fmt.Println(device)
		deviceSlice = append(deviceSlice, device)
	}
	
	tmpl.ExecuteTemplate(w, "singleProduct", deviceSlice)
	json.NewEncoder(w).Encode(deviceSlice)
}

// Create products by id
func createProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create")
	db :=dbConn()
	stmt, err := db.Prepare("INSERT INTO products(title,description) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	_,created_data := stmt.Exec("React","UI framework")
	json.Unmarshal(body, &created_data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, `
	<h1>New post was created</h1>
	<p>Note: Feed the create data in code</p`)
}

// Update products by id
func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := dbConn()
	query:=`UPDATE products SET title=?, description=? WHERE id = ?`
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body)
	_,updated_data:= stmt.Exec("Golang","It is a Programming language", params["id"])
	json.Unmarshal(body, &updated_data)
	
	fmt.Fprintf(w, `
	<h1>Post with ID = %s was updated</h1>
	<p>Note: Feed the update data in code</p`, params["id"])
	
}
// Delete products by id
func deleteProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("delete")
	db:= dbConn()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, `
	<h1>Post with ID = %s was deleted</h1>
	<p>The id with the given number is deleted and check the original data </p`, params["id"])
	
}









