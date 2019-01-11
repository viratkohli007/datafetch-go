package main 
import ( 
         "html/template"
         "net/http"
          "log"
          "os"
           "encoding/json"
          "io/ioutil"
          "fmt"
          "io"
        )

type welcome struct{

    Title string
}

type form struct{

}

type Data struct{
    Name string       `json:"name"`
    Age string        `json:"age"`
    Phno string       `json: "phno"`
    Bloodgroup string `json:"bloodgroup"`
}

type display struct{}

func main() {
    
    http.HandleFunc("/", Home)
    http.HandleFunc("/form", Form)
    http.HandleFunc("/display", Display)
    http.HandleFunc("/list", List)
    err := http.ListenAndServe(":8080", nil)
    if err != nil{
      log.Fatal(err)
    }

}

func Home(w http.ResponseWriter, r *http.Request) {
    
     title := welcome{Title : "welcome"}
     t, _ := template.ParseFiles("welcome.html")
     t.Execute(w, title)
}

func Form(w http.ResponseWriter, r *http.Request) {
     
     var f form
     t, _ := template.ParseFiles("form.html")
     t.Execute(w, f)
}

func Display(w http.ResponseWriter, r *http.Request) {
    var d Data
    d.Name = r.FormValue("name")
    d.Age = r.FormValue("age")
    d.Phno = r.FormValue("phno")
    d.Bloodgroup = r.FormValue("bloodgroup")
    t, _ := template.ParseFiles("display.html")
    t.Execute(w, d)

    f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil{
        log.Fatal(err)
    }
    defer f.Close()
     info := new(Data)
    info.Name = r.FormValue("name")
    info.Age = r.FormValue("age")
    info.Phno = r.FormValue("phno")
    info.Bloodgroup = r.FormValue("bloodgroup")

    b, err := json.Marshal(info)
    if err != nil{
        log.Fatal(err)
    }
     _, err4 := io.WriteString(f, string(b))
     if err4 != nil {
          log.Fatal(err)
     }
    // r.ParseForm()
     // w.Write(d)
    //fmt.Fprintf(w,d.Name)
    //w.Write(r.Form)
    //fmt.Fprintf(w, r.Form)
    // if r.URL.name != "" {
    //  data := &data{Name: ""|| r.URL.name, Age:""|| r.URL.age, Phno:""|| r.URL.phno, Bloodgroup: ""||r.URL.bloodgroup }
    // }
    // t, _ = template.ParseFiles("data.json")
    // t.Execute(w, info)
    
    
}

func List(w http.ResponseWriter, r *http.Request) {
    
    fdata, err := os.Open("data.json")
    if err != nil {
        log.Fatal(err)
    }
    byteval, _ := ioutil.ReadAll(fdata)
    w.Write(byteval)
    fmt.Fprintf(w, "/n")
    fdata.Close()
}