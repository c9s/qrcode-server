package main

import (
  "net/http"
  "fmt"
  "flag"
  "image/png"
  "github.com/qpliu/qrencode-go/qrencode"
)

func qrcodeHandler(w http.ResponseWriter, r *http.Request) {
  text := r.FormValue("text")
  if text == "" {
    fmt.Fprintf(w,"form value text is required")
    return
  }
  w.Header().Add("Content-Type", "image/png")
  grid, err := qrencode.Encode( text , qrencode.ECLevelQ)
  if err != nil {
    fmt.Fprintf(w,"QRCode error:%s",err)
    return
  }
  png.Encode(w, grid.Image(8))
}

func main() {
  var port *string = flag.String("port","8888","Set port")
  flag.Parse()
  fmt.Printf("Starting qrcode server at http://0.0.0.0:%s ...\n",*port)
  http.HandleFunc("/", qrcodeHandler)
  http.ListenAndServe(":" + *port, nil)
}
