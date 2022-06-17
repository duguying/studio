package storage

type S3Oss struct {
}

func (s *S3Oss) List(remotePrefix string) (list []*FileInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *S3Oss) GetFileInfo(remotePath string) (info *FileInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *S3Oss) AddFile(localPath string, remotePath string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *S3Oss) RenameFile(remotePath string, newRemotePath string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *S3Oss) RemoveFile(remotePath string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *S3Oss) FetchFile(remotePath string, localPath string) (err error) {
	//TODO implement me
	panic("implement me")
}

