package server

import (
	"fmt"
	audit "github.com/GOLANG-NINJA/crud-audit-log/pkg/domain"
	"net"
)

type Server struct {
	grpcSrv *grpc.Server
	audit   audit.AuditServiceServer
}

func New(auditServer audit.AuditServiceServer) *Server {
	return &Server{
		grpcSrv:     grpc.NewServer(),
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	audit.RegisterAuditServiceServer(s.grpcSrv, s.auditServer)

	if err := s.gprcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}