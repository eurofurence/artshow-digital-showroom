package server

func (s *Server) sendCommandsToMpv(cmds ...[]any) error {
	for _, cmd := range cmds {
		command := map[string]any{"command": cmd}
		if err := s.mpvEncoder.Encode(command); err != nil {
			return err
		}
	}
	return nil
}
