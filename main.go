package main

import (
  "net/http"
  "fmt"
  "strings"
  "flag"
  "image/png"
  "github.com/qpliu/qrencode-go/qrencode"
  "github.com/c9s/go-bitly/bitly"
)

func qrcodeHandler(w http.ResponseWriter, r *http.Request) {
  text := r.FormValue("text")
  if text == "" {
    fmt.Fprintf(w,"form value text is required")
    return
  }

  if strings.Contains(text,"http://") {
    var err error
    bitly.SetUser("o_52nge5mh7c")
    bitly.SetKey("R_eb4b4b532889a43023bb99fc2e81b6f0")
    text, err = bitly.Shorten(text)
    if err != nil {
      fmt.Fprintf(w,"bitly error: %s", err)
      return
    }
    // to test decode: http://zxing.org/w/decode.jspx
    // fmt.Println(text)
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
