package util

import (
	"co-studio-e-commerce/conf"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/admin/search"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UploadAFile(c *gin.Context) {
	// Create our instance
	cfg, _ := conf.LoadConfig(".")
	cld, _ := cloudinary.NewFromURL(cfg.Url_HTTPS)

	// Get the prefrered name of the file if its not supplied
	fileName := c.PostForm("name")

	// Add tags
	fileTags := c.PostForm("tags")
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to upload",
		})
		return
	}

	result, err := cld.Upload.Upload(c, file, uploader.UploadParams{
		PublicID: fileName,

		// Split the tags by comma
		Tags: strings.Split(",", fileTags),
	})

	if err != nil {
		c.String(http.StatusConflict, "Upload to cloudinary failed")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Successfully uploaded the file",
		"secureURL": result.SecureURL,
		"publicURL": result.URL,
	})

	//other code
}

func GetUploadedFile(c *gin.Context) {
	// Create instance
	cfg, _ := conf.LoadConfig(".")
	cld, _ := cloudinary.NewFromURL(cfg.Url_HTTPS)

	fileName := c.Param("accessId")

	//Access the filename using a desired file access id
	result, err := cld.Admin.Asset(c, admin.AssetParams{
		PublicID: fileName,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, "We were unable to find the file requested")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":    "Successfully found the image",
		"secureURL":  result.SecureURL,
		"publicURL":  result.URL,
		"created_at": result.CreatedAt.String(),
	})

	// other code
}

func UploadMultiFile(c *gin.Context) {
	cfg, _ := conf.LoadConfig(".")
	cld, _ := cloudinary.NewFromURL(cfg.Url_HTTPS)

	var urls []string

	searchQ := search.Query{
		Expression: "resource_type:image AND uploaded_at>1d AND bytes<1m",
		SortBy: []search.SortByField{
			{
				"created_at": search.Descending,
			},
		},
		MaxResults: 10,
	}

	results, err := cld.Admin.Search(c, searchQ)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err,
			"message": "Failed to query your files",
		})
		return
	}

	for _, asset := range results.Assets {
		urls = append(urls, asset.SecureURL)
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"data":    urls,
	})

	// other code
}

func UpdateFile(c *gin.Context) {
	cfg, _ := conf.LoadConfig(".")
	cld, _ := cloudinary.NewFromURL(cfg.Url_HTTPS)
	fileId := c.Param("publicId")
	newFileName := c.PostForm("fileName")

	// Access the filename using a desired filename
	result, err := cld.Upload.Rename(c, uploader.RenameParams{
		FromPublicID: fileId,
		ToPublicID:   newFileName,
	})
	if err != nil {
		c.String(http.StatusNotFound, "We were unable to find the file requested")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":    "Successfully found the image",
		"secureURL":  result.SecureURL,
		"publicURL":  result.URL,
		"created_at": result.CreatedAt.String(),
	})
}

func DeleteFile(c *gin.Context) {
	cfg, _ := conf.LoadConfig(".")
	cld, _ := cloudinary.NewFromURL(cfg.Url_HTTPS)
	fileId := c.Param("assetId")
	result, err := cld.Upload.Destroy(c, uploader.DestroyParams{
		PublicID: fileId,
	})

	if err != nil {
		c.String(http.StatusBadRequest, "File could not be deleted")
		return
	}

	c.JSON(http.StatusContinue, result.Result)
}
