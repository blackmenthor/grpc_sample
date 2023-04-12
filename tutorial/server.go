package tutorial

import (
	context "context"
	"fmt"
	"strconv"
)

type Server struct {
	UnimplementedAlbumServiceServer
}

func (s *Server) GetAlbum(ctx context.Context, in *AlbumRequest) (*AlbumResponse, error) {
	var listOfData []*Album
	var maximumDataSize = 1000000
	for i := 1; i < maximumDataSize; i++ {
		var newAlbum = &Album{
			Id:     strconv.Itoa(i),
			Title:  fmt.Sprintf("Album %d", i),
			Artist: fmt.Sprintf("Artist %d", i),
			Price:  56.99,
		}
		listOfData = append(
			listOfData,
			newAlbum,
		)
	}

	return &AlbumResponse{
		Albums: listOfData,
	}, nil
}
