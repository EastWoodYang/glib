package glib

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

/* ================================================================================
 * 文件
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
 * ================================================================================ */
type (
	FileInfoList []*FileInfo
	FileInfo     struct {
		Filename string `form:"filename" json:"filename"` //原始文件名（test.jpg）
		Data     []byte `form:"data" json:"data"`         //文件字节切片
		Size     int64  `form:"size" json:"size"`         //大小（单位：字节）
		Duration int64  `form:"duration" json:"duration"` //时长（单位：秒）
		Path     string `form:"path" json:"path"`         //全路径（本地磁盘或第三方文件系统）
	}

	IFileSize interface {
		Size() int64
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前文件执行的全路径
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	paths := strings.Split(path, string(os.PathSeparator))
	return strings.Join(paths[:len(paths)-1], string(os.PathSeparator))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取全文件路径的相对路径
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetRelativePath(fullpath string) string {
	currentFullPath := GetCurrentPath()
	//path := strings.Replace(fullpath, currentFullPath, "", -1)
	//splitstring := strings.Split(path, string(os.PathSeparator))
	//return strings.Join(splitstring[1:], string(os.PathSeparator))
	relPath, _ := filepath.Rel(currentFullPath, fullpath)
	return relPath
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建多级目录
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func CreateDir(perm os.FileMode, args ...string) (string, error) {
	dirs := strings.Join(args, string(os.PathSeparator))
	err := os.MkdirAll(dirs, perm)
	if err != nil {
		return "", err
	}

	return dirs, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据指定的日期在指定的目录下创建多级目录
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func CreateDateDir(rootPath string, datetime time.Time, perm os.FileMode) (string, error) {
	year, month, day := datetime.Date()
	sYear := fmt.Sprintf("%d", year)
	sMonth := fmt.Sprintf("%02d", month)
	sDay := fmt.Sprintf("%02d", day)

	return CreateDir(perm, rootPath, sYear, sMonth, sDay)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据当前日期在指定根目录下创建多级目录
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func CreateCurrentDateDir(rootPath string, perm os.FileMode) (string, error) {
	nowDate := time.Now()
	return CreateDateDir(rootPath, nowDate, perm)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取路径里的路径和文件名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetFilePath(filePath string) (string, string) {
	path := ""
	filename := ""
	paths := strings.Split(filePath, string(os.PathSeparator))
	length := len(paths)
	if length > 0 {
		path = strings.Join(paths[0:length-1], string(os.PathSeparator))
		filename = paths[length-1 : length][0]
	}

	return path, filename
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取路径里的文件名，不带扩展名的文件名，扩展名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetFilename(filePath string) (string, string, string) {
	paths := strings.Split(filePath, string(os.PathSeparator))
	filename := ""
	filenameWithoutExtname := ""
	extname := ""

	if len(paths) > 0 {
		filename = paths[len(paths)-1]
		filenames := strings.Split(filename, ".")
		filenameWithoutExtname = filenames[0]
		extname = filenames[1]
	}

	return filename, filenameWithoutExtname, extname
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取Http请求里的文件数据
 * maxSize: 文件大小限制，0表示不限制
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetHttpRequestFile(req *http.Request, args ...int64) (*FileInfo, error) {
	//获取请求文件
	inputFile, fileHeader, err := req.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	var maxSize int64
	if len(args) > 0 {
		maxSize = args[0]
	}

	//判断大小是否超出限制
	size := inputFile.(IFileSize).Size()
	if maxSize != 0 && size > maxSize {
		return nil, errors.New("file is too large")
	}

	//读取文件数据到字节切片
	dataBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	//返回数据结果
	fileInfo := &FileInfo{
		Filename: fileHeader.Filename,
		Data:     dataBytes,
		Size:     size,
	}

	return fileInfo, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取图片文件
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetImageFile(filename string, args ...string) (image.Image, error) {
	fileType := "png"
	if len(args) == 1 {
		fileType = args[0]
	}

	filename, _ = filepath.Abs(filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decodes := make(map[string]func(io.Reader) (image.Image, error), 0)
	decodes["png"] = png.Decode
	decodes["jpg"] = jpeg.Decode
	return decodes[fileType](file)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 从视频文件截取图片
 * ffmpeg -i ./test.mp4 -ss 00:00:01 -s 120*120 -r 1 -q:v 2 -f image2 -vframes 1 image-%2d.jpg
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func CutVideoImage(sourceFile, newFilename string, width, height uint64, second, q int) (string, error) {
	if !filepath.IsAbs(sourceFile) {
		sourceFile, _ = filepath.Abs(sourceFile)
	}

	if !filepath.IsAbs(newFilename) {
		newFilename, _ = filepath.Abs(newFilename)
	}

	if second >= 60 {
		second = 0
	}

	startTime := fmt.Sprintf("00:00:%d", second)

	params := make([]string, 0)
	params = append(params, "-i")
	params = append(params, sourceFile)
	params = append(params, "-ss")
	params = append(params, startTime)
	params = append(params, "-s")
	params = append(params, fmt.Sprintf("%d*%d", width, height))
	params = append(params, "-r")
	params = append(params, "1")
	params = append(params, "-q:v")
	params = append(params, fmt.Sprintf("%d", q))
	params = append(params, "-f")
	params = append(params, "image2")
	params = append(params, "-vframes")
	params = append(params, "1")
	params = append(params, newFilename)

	log.Printf("sourceFile: " + sourceFile + " ,newFilename: " + newFilename)

	cmd := exec.Command("ffmpeg", params...)
	_, err := cmd.Output()
	if err != nil {
		return strings.Join(params, " "), err
	}
	return strings.Join(params, " "), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 保存Http 上传的文件到磁盘指定目录（返回客户端原文件名，大小，全文件路径，错误）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func SaveHttpFile(req *http.Request, filename, basePath string, maxSize int64, args ...string) (*FileInfo, error) {
	fileInfo, err := GetHttpRequestFile(req, maxSize)
	if err != nil {
		return nil, err
	}

	fullFilename, err := SaveFile(fileInfo.Data, filename, basePath, args...)
	if err != nil {
		return nil, err
	}

	fileInfo.Path = fullFilename
	log.Printf("SaveHttpFile Filename: %s, fullFilename: %s", fileInfo.Filename, fileInfo.Path)

	return fileInfo, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 保存文件
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func SaveFile(data []byte, filename, basePath string, args ...string) (string, error) {
	rootPath, _ := filepath.Abs(basePath)

	if len(args) == 0 {
		//默认保存在上传目录下，且生成当前日期目录
		rootPath, _ = CreateDateDir(rootPath, time.Now(), 0755)
	} else if len(args) == 1 {
		//保存到指定目录
		basePath = args[0]
		rootPath, _ = filepath.Abs(basePath)
		if _, err := CreateDir(0755, rootPath); err != nil {
			return "", err
		}
	}

	fullFilename := rootPath + string(os.PathSeparator) + filename
	log.Printf("SaveFile basePath: %s, filename: %s, fullFilename: %s", basePath, filename, fullFilename)

	//写入文件数据
	outputFile, err := os.OpenFile(fullFilename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	if _, err = io.Copy(outputFile, bytes.NewReader(data)); err != nil {
		return "", err
	}

	return fullFilename, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 移动文件（全路径源文件，目的路径，日期，会根据日期自动创建路径然后连接到目的路径后）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func MoveFile(srcFilename, dstPath string, creationDate time.Time) (string, error) {
	var err error

	srcFullFilename, err := filepath.Abs(srcFilename)
	if err != nil {
		return "", err
	}

	srcFile, err := os.OpenFile(srcFullFilename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	rootPath, err := filepath.Abs(dstPath)
	if err != nil {
		return "", err
	}

	if !creationDate.IsZero() {
		//根据日期创建yyyy/mm/dd目录
		rootPath, _ = CreateDateDir(rootPath, creationDate, 0755)
	}

	dstFullFilename := rootPath + string(os.PathSeparator) + filepath.Base(srcFullFilename)
	dstFile, err := os.OpenFile(dstFullFilename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", err
	} else {
		os.Remove(srcFullFilename)
	}

	return dstFullFilename, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除文件
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DeleteFile(filename string, args ...string) error {
	log.Printf("0 DeleteFile filename: %s", filename)

	fullFilename := filename
	if !filepath.IsAbs(filename) {
		if len(args) > 0 {
			filename = filepath.Join(args[0], filename)
			log.Printf("1 DeleteFile filename: %s", filename)

			fullFilename, _ = filepath.Abs(filename)

			log.Printf("2 DeleteFile fullFilename: %s", fullFilename)
		} else {
			fullFilename, _ = filepath.Abs(filename)
		}
	}

	//删除文件
	return os.Remove(fullFilename)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断文件是否存在
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FileIsExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
