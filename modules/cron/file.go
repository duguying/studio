package cron

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/imgtools"
	"duguying/studio/utils"
	"log"
	"time"
)

func scanFile() error {
	files, err := db.ListAllMediaFile(g.Db, 0)
	if err != nil {
		return err
	}

	for _, file := range files {
		time.Sleep(time.Second)

		localPath := utils.GetFileLocalPath(file.Path)
		if file.MediaWidth <= 0 || file.MediaHeight <= 0 {
			width, height, err := imgtools.GetImgSize(localPath)
			if err != nil {
				log.Println("获取文件尺寸失败, err:", err.Error(), "localPath:", localPath)
				continue
			}
			err = db.UpdateFileMediaSize(g.Db, file.ID, int(width), int(height))
			if err != nil {
				log.Println("更新文件尺寸失败, err:", err.Error())
				continue
			}
		}

		if file.Thumbnail == "" {
			thumbKey, err := imgtools.MakeThumbnail(localPath, 220)
			if err != nil {
				log.Println("制作缩略图失败, err:", err.Error())
				continue
			}
			err = db.UpdateFileThumbneil(g.Db, file.ID, thumbKey)
			if err != nil {
				log.Println("更新缩略图失败, err:", err.Error())
				continue
			}
		}
	}

	return nil
}
