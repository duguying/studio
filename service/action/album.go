package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/service/models"
)

func ListAlbumFiles(c *CustomContext) (interface{}, error) {
	userID := c.UserID()
	files, err := db.ListAllMediaFile(g.Db, userID)
	if err != nil {
		return nil, err
	}
	apiFiles := []*models.MediaFile{}
	for _, file := range files {
		apiFiles = append(apiFiles, file.ToMediaFile())
	}
	return models.ListMediaFileResponse{
		Ok:   true,
		List: apiFiles,
	}, nil
}
