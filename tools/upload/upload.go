package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"tool_update_service/models"
	"tool_update_service/tools"
	"tool_update_service/tools/fileInfo"
	"tool_update_service/tools/settings"
)

func File(c *gin.Context, ufs *models.ParamUpload) (err error) {

	f, err := c.FormFile("fileName")
	if err != nil {
		return
	}

	var dst string
	if ufs.Classification == "0" {
		dst = path.Join("./files/img", f.Filename)
	} else {
		//dst = path.Join("./files/source/"+strings.Split(f.Filename, ".")[0], f.Filename)
		dstb := path.Join("./files/source/", ufs.Version)
		tools.Mkdir(dstb)
		dst = path.Join(dstb, f.Filename)
	}
	c.SaveUploadedFile(f, dst)
	//url := fmt.Sprintf("http://%s:%v/%s/%s", tools.GetOutboundIP(), settings.Config.Port, tools.GetCurrentPath(), dst)
	url := fmt.Sprintf("http://%s:%v/download/%s", tools.GetOutboundIP(), settings.Config.Port, dst)

	ufs.Url = url
	ufs.Size = fileInfo.CalcSize(f)
	return
}
