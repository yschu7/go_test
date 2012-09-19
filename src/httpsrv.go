package main

import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type",
        "text/html",
    )
    io.WriteString(
        res,
        `<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
    )
}
func main() {
    http.HandleFunc("/hello", hello)
    http.Handle(
        "/example/",
        http.StripPrefix(
            "/example/",
            http.FileServer(http.Dir("example")),
        ),
    )
    http.Handle(
        "/chksum/",
        http.StripPrefix(
            "/chksum/",
            http.FileServer(http.Dir("chksum")),
        ),
    )
    http.Handle(
        "/all/",
        http.StripPrefix(
            "/all/",
            http.FileServer(http.Dir(".")),
        ),
    )
    http.ListenAndServe(":9000", nil)
}

