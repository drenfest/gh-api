package main

import (
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
	data := "Test"
	writeFile("html/test.html", data)
	// router := mux.NewRouter()
	// router.HandleFunc("/", renderTemplate).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8000", router))
}
func renderTemplate(w http.ResponseWriter, r *http.Request){
	// gather the data to pass to the template
	data := "Test"
	// load the template
	tmpl := GetTemplate("templates/index.gohtml")
	tmpl.ExecuteTemplate(w,"index.gohtml", data	)
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
			tByte := re.ReplaceAll([]byte("templates/"+locations[i]),[]byte(".gohtml"))
			if Exists(string(tByte)) {
				t, err := template.ParseFiles(string(tByte))
				if err != nil {
					log.Print(err)
				}
				f, _ := os.Create(d + "/" + locations[i])
				t.Execute(f, data)
				defer f.Close()
				println(locations[i] + " created")
			}else {
				t, err := template.ParseFiles(string("templates/index.gohtml"))
					if err != nil {
						log.Print(err)
					}
					f, _ := os.Create(d+"/"+locations[i])
					t.Execute(f, data)
					defer f.Close()
					println(locations[i]+" created")
			}
		}else{
			d = string(d) + string(locations[i])
			os.Mkdir(locations[i], 755)
			println(locations[i]+" created")
		}
	i++
	}
}



func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}