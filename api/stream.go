package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/dreamvo/gilfoyle/api/util"
	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/schema"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

// @ID getMediaPlaylistFile
// @Tags Stream
// @Summary Get HLS playlist file of a media
// @Description Get HLS playlist file of a media
// @Produce  octet-stream
// @Success 200 {string} string
// @Failure 404 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Header 200 {string} Content-Type "application/octet-stream"
// @Param media_id path string true "Media identifier" validate(required)
// @Param filename path string true "HLS filename" validate(required)
// @Router /medias/{media_id}/stream/{filename} [get]
func (s *Server) getMediaPlaylistFile(ctx *gin.Context) {
	mediaUUID := ctx.Param("id")
	filename := ctx.Param("filename")

	parsedUUID, err := util.ValidateUUID(mediaUUID)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, ErrInvalidUUID)
		return
	}

	m, _ := s.db.Media.Query().Where(media.ID(parsedUUID)).Only(context.Background())
	if m == nil {
		util.NewError(ctx, http.StatusNotFound, ErrResourceNotFound)
		return
	}

	if m.Status == schema.MediaStatusErrored {
		util.NewError(ctx, http.StatusBadRequest, errors.New("media contains errors"))
		return
	}

	stat, err := s.storage.Stat(context.Background(), path.Join(m.ID.String(), filename))
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}

	// Open the file
	f, err := s.storage.Open(context.Background(), path.Join(m.ID.String(), filename))
	if err != nil {
		util.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	filenameParts := strings.Split(filename, "/")
	file := filenameParts[len(filenameParts)-1]

	ctx.DataFromReader(http.StatusOK, stat.Size, mime.TypeByExtension(filepath.Ext(file)), f, map[string]string{
		"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, file),
	})
}
