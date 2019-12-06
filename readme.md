# intro
simple file upload and download tool written in go
# install
```bash
go get github.com/IanSmith123/fileupload
```

# upload
```bash
curl -F file=@filename.png 127.0.0.1:2019
```

# download
```bash
wget 127.0.0.1:2019/file
```