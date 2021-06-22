package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("/bin/sh", "-c", "systemctl restart v2ray", "")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		fmt.Printf("combined out:\n%s\n", string(out))
		fmt.Fprintf(w, "restart success, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
