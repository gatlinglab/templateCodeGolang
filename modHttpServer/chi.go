package modHttpServer

import (
	"fmt"
	"net/http"

	"wjtemplate1/modUtility"

	"github.com/go-chi/chi/v5"
)

type CChi struct {
	router *chi.Mux
}

var g_singleChi *CChi = &CChi{}

func getSingleChi() *CChi {
	return g_singleChi
}

func (pInst *CChi) Initialize() error {
	r := chi.NewRouter()
	r.Get("/", handlerHomepage)

	// static log
	// if modUtility.G_LocalLogPath != "" {
	// 	fs := http.FileServer(http.Dir(modUtility.G_LocalLogPath))
	// 	r.Handle("/log/*", http.StripPrefix("/log/", fs))
	// }
	pInst.router = r

	return nil
}

func (pInst *CChi) start() error {
	listenStr := fmt.Sprintf(":%d", modUtility.G_HttpPort)
	fmt.Println("start listen: ", listenStr)
	err := http.ListenAndServe(listenStr, pInst.router)
	return err
}

func handlerHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome message"))
}
