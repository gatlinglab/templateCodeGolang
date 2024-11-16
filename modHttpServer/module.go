package modHttpServer

func Http_Initialize() error {
	return getSingleChi().Initialize()
}
func Http_Start() error {
	return g_singleChi.start()
}