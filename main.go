package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// The person Type (more like an object)


// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func renderTemplate(w http.ResponseWriter, r *http.Request){
	// gather the data to pass to the template
	data := "Test"
	// load the template
	tmpl := GetTemplate("templates/index.gohtml")
	tmpl.ExecuteTemplate(w,"index.gohtml", data	)
	writeFile("html/index.html", data)
}
func GetTemplate(tpath string)*template.Template{
	cmd := exec.Command("touch", tpath)
	cmd.Run()
var defaultTPath = "templates/index.gohtml"
if len(tpath) < 1 {
	tpath = defaultTPath
}
	tmpl, err := template.ParseFiles(tpath)
	if err != nil{
		log.Fatal(err)
	}
return	tmpl

}

func writeFile(location string, data string){
	locations := strings.Split(location,"/")
	var i = 0
	var d string
	for i < len(locations){
		if i == len(locations) - 1{
			var re = regexp.MustCompile(`(\.[a-z]+)$`)
			t, err := template.ParseFiles(string(re.ReplaceAll([]byte("templates/"+locations[i]),[]byte(".gohtml"))))
			if err != nil {
				log.Print(err)
				return
			}
			f, _ := os.Create(d+"/"+locations[i])
			t.Execute(f, data)
			defer f.Close()
			println(locations[i]+" created")
		}else{
			d = string(d) + string(locations[i])
			os.Mkdir(locations[i], 755)
			println(locations[i]+" created")
		}
	i++
	}
	
}
