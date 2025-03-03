package admin

import (
	"fmt"
	"gin_tutorial/GINDEMO5/Utiles"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

// For decoupling
type IAdminController interface {
	Add(c *gin.Context)
	DoUpload(c *gin.Context)
	Edit(c *gin.Context)
	EditUpload(c *gin.Context)
	Multipart(c *gin.Context)
	MultipartUpload(c *gin.Context)
}

type AdminController struct{}

// factory pattern
func NewAdminController() IAdminController {
	return &AdminController{}
}

// show add format html
func (admin *AdminController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/admin.html", gin.H{})
}

// do upload files
func (admin *AdminController) DoUpload(c *gin.Context) {

	//Todo:
	user := c.PostForm("user")

	image, err := c.FormFile("image")
	if err != nil {
		return
	}

	dst := path.Join("./static/upload", image.Filename)
	c.SaveUploadedFile(image, dst)

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"image": image,
		"dst":   dst,
	})
	// c.String(http.StatusOK, "do upload")
}

func (admin *AdminController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/adminedit.html", nil)
}

func (admin *AdminController) EditUpload(c *gin.Context) {
	//GET username
	user := c.PostForm("user")
	//get file
	file, err := c.FormFile("file1")
	if err != nil {
		return
	}

	//get ext
	ext := path.Ext(file.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[ext]; !ok {
		log.Printf("文件格式不正确")
		c.Redirect(http.StatusTemporaryRedirect, "/admin/edit")
	}

	//get current day
	day := Utiles.GetDay()
	dir := "./static/upload/" + day

	err = os.MkdirAll(dir, 0666)
	if err != nil {
		fmt.Println(err)
		c.String(200, "mkdirALL failed")
		return
	}

	//生成文件名称和文件保存的目录
	fileName := strconv.FormatInt(Utiles.GetUnix(), 10) + ext

	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)

	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"user": user,
	// 	"file": gin.H{
	// 		"file1": file1,
	// 		"file2": file2,
	// 	},
	// 	"dst": gin.H{
	// 		"file1_dst": file1_dst,
	// 		"file2_dst": file2_dst,
	// 	},
	// })

}

func (admin *AdminController) Multipart(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/adminMultipart.html", nil)
}

func (admin *AdminController) MultipartUpload(c *gin.Context) {
	user := c.PostForm("user")

	form, err := c.MultipartForm()
	if err != nil {
		return
	}

	files := form.File["file"]

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"success": true,
	})
}

func setDstAndSave(c *gin.Context, file *multipart.FileHeader, err error) string {
	if err != nil {
		return err.Error()
	}
	dst := path.Join("./static/upload", file.Filename)
	c.SaveUploadedFile(file, dst)

	return dst
}
