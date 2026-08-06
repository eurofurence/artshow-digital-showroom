package server

func (s *Server) sendCommandToMpv(commands ...[]any) error {
	cmd := map[string]any{"command": commands}
	return s.mpvEncoder.Encode(cmd)
}
func (s *Server) sendCommandsToMpv(cmds ...[]any) error {
    for _, cmd := range cmds {
		command := map[string]any{"command": cmd}
		if err := s.mpvEncoder.Encode(command); err != nil {
            return err
        }
    }
    return nil
}
