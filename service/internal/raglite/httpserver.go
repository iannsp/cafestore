package raglite

import(
    "net/http"
)

type HttpServer struct{
    service *http.Server
    handler *http.ServeMux
}

type HttpServerHandlerFunc func(http.ResponseWriter, *http.Request)

func (hs *HttpServer) ListenAndServe() error{
    return hs.service.ListenAndServe()
}

func (hs *HttpServer) AttachRoutes(route string, handler HttpServerHandlerFunc){
    hs.handler.HandleFunc(route, handler)
}

func NewHttpServer(port string) HttpServer{
    hs := HttpServer{}
    hs.handler = http.NewServeMux()
    hs.service = &http.Server{
            Addr: ":"+ port,
            Handler: hs.handler,
    }
    return hs
}
/*

*/
