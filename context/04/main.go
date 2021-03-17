package main

import "sync"

type Server struct {
	serverStopChan chan struct{}
	stopWg         sync.WaitGroup
}

func (s *Server) Stop() {
	if s.serverStopChan == nil {
		panic("gorpc.Server: server must be started before stopping it")
	}
	close(s.serverStopChan)
	s.stopWg.Wait()
	s.serverStopChan = nil
}

func serverHandler(s *Server) {
	for {
		select {
		case <-s.serverStopChan:
			return
		default:
			// .. do something
		}
	}
}

func main() {

}
