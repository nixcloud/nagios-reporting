package main

import (
  "fmt"
  "net/http"
  "time"
  "io/ioutil"
  "encoding/json"
)

type mydata struct {
  Servername   string
  Service      string
  Exitcode     int
  Status       string
  Perfdata     string
};

func handler(rw http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
      fmt.Println("error ReadAll")
  }
  //fmt.Println("----", string(body), "----")
  var t mydata
  err = json.Unmarshal(body, &t)
  if err != nil {
      fmt.Println("error Unmarshal")
      fmt.Fprint(rw, "error Unmarshaling the json string, maybe wrong types?\n\n", string(body))
      rw.WriteHeader(http.StatusNotFound)
      return
  }
  //fmt.Println(t.Servername)

  ti := time.Now();

// http://www.admin-magazine.com/Archive/2014/22/Nagios-Passive-Checks
// https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/configmain.html#command_file
// echo "[1470746488] PROCESS_SERVICE_CHECK_RESULT;nixcloud.io;Powerful_backup;0;neu2|" > /var/lib/nagios/nagios.cmd 
// curl -H "Content-Type: application/json" -X POST -d '{"servername":"nixcloud.io", "service":"Powerful_backup", "exitcode":0, "perfdata":""}' https://nixcloud.io/nagios-reporting/

  a := fmt.Sprintf("[ %v ] PROCESS_SERVICE_CHECK_RESULT;%s;%s;%v;%s|%s\n",  ti.UTC().Unix(), t.Servername, t.Service, t.Exitcode, t.Status, t.Perfdata)
  buff := []byte(a)
  fmt.Println(a)

  fmt.Fprint(rw, a)
  e := ioutil.WriteFile("/var/lib/nagios/nagios.cmd", buff, 0644)
  if e != nil {
    fmt.Println(e)
  }
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9999", nil)
}
