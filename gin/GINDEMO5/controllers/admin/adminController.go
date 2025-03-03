package admin

import (
	"mime/multipart"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// For decoupling
type IAdminController interface {
	Add(c *gin.Context)
	DoUpload(c *gin.Context)
	Edit(c *gin.Context)
	EditUpload(c *gin.Context)
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

	//get file1
	file1, err := c.FormFile("file1")
	file1_dst := setDstAndSave(c, file1, err)

	//get file2
	file2, err := c.FormFile("file2")
	file2_dst := setDstAndSave(c, file2, err)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"file": gin.H{
			"file1": file1,
			"file2": file2,
		},
		"dst": gin.H{
			"file1_dst": file1_dst,
			"file2_dst": file2_dst,
		},
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
