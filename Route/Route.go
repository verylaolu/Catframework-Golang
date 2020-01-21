package Route
import (
	"sso/Middleware"
	"sso/App/Controller"
	"github.com/gorilla/mux"
	"net/http"
)

type route struct{}

func Init()  {
	var routeObj route
	routeObj.RegisterRoutes()
}
func (routeObj route) RegisterRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", Middleware.Auth(Controller.Home))
	r.HandleFunc("/jwt/encode",Middleware.Auth(Controller.Encodetoken))
	r.HandleFunc("/jwt/decode", Middleware.Auth(Controller.Decodetoken))
	http.Handle("/", r)
}
