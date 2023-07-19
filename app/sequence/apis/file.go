package apis

import (
	"bufio"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/app/sequence/service"
	"go-admin/app/sequence/service/dto"
	"os"
	"strings"
)

type ReqParam struct {
	Path string `json:"path"`
}

func (req *ReqParam) GetId() interface{} {
	return 0
}

// UploadFile 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单图，2：多图, 3：base64图片)
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/public/uploadFile [post]
// @Security Bearer
func (e Sequence) UploadFile(c *gin.Context) {
	reqParam := ReqParam{}
	err := c.BindJSON(&reqParam)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&reqParam, binding.JSON).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if reqParam.Path == "" {
		e.Error(500, errors.New("path is empty"), "path is empty")
		return
	}
	// 解析文件，写入数据库
	err = e.insertIntoDB(reqParam.Path)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK("", "success")
}

func (e Sequence) insertIntoDB(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("failed to open file: %s", err)
		return err
	}
	defer file.Close()

	var sequences []dto.SequenceInsertReq
	var seq dto.SequenceInsertReq
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if seq.SequenceId != "" {
				sequences = append(sequences, seq)
			}
			line = line[1:] // trimming ">" symbol
			parts := strings.SplitN(line, " ", 2)
			seq = dto.SequenceInsertReq{}
			seq.SequenceId = parts[0]
			if len(parts) > 1 {
				seq.SequenceDescription = parts[1]
			}
		} else {
			seq.Sequence += line
		}
	}
	// Append last read sequence
	if seq.SequenceId != "" {
		sequences = append(sequences, seq)
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s := service.Sequence{}
	err = e.MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		// e.Error(500, err, err.Error())
		return err
	}
	for _, seq := range sequences {
		err = s.Insert(&seq)
		if err != nil {
			// e.Error(500, err, fmt.Sprintf("创建Sequence失败，\r\n失败信息 %s", err.Error()))
			return err
		}
	}

	// Example of usage
	for _, seq := range sequences {
		log.Debugf("ID: %s, Description: %s, Sequence: %s", seq.SequenceId, seq.SequenceDescription, seq.Sequence)
	}
	return nil
}
