package web

func (s *Server) externalize(path string) string {
	//todo handle scheme by params
	return "http://" + s.externalHost + ":" + s.externalPort + "/" + path
}
