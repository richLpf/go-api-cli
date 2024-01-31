package gen_frame

import (
	"fmt"
	"go-api-cli/utils"
	"os"
	"path"
	"strings"
)

type GenFrameConfig struct {
	OutPath  string
	PrjName  string
	JsonCase string
}
type GenFrameService struct {
	cfg        *GenFrameConfig
	genPrjName string
	genPrjPath string
	outPrjPath string
}

func NewGenFrameConfig(outPath, prjName string) *GenFrameConfig {
	cfg := &GenFrameConfig{
		OutPath: outPath,
		PrjName: prjName,
	}
	if outPath == "" {
		cfg.OutPath = "./"
	}
	if prjName == "" {
		cfg.PrjName = "prj_aiee"
	}
	return cfg
}

func NewGenFrameService(cfg *GenFrameConfig) *GenFrameService {
	s := &GenFrameService{
		cfg:        cfg,
		genPrjName: "go-api-cli-prj",
		genPrjPath: "./prj",
		outPrjPath: path.Join(cfg.OutPath, cfg.PrjName),
	}
	return s
}

func (s *GenFrameService) GenFrame() (err error) {
	if err = utils.CreateDir(s.outPrjPath); err != nil {
		return err
	}
	if err = s.genWithPrjFile("main.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("go.mod", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("config/config.go", nil); err != nil {
		return err
	}
	if err = s.genWithPrjFile("config/config.yml", nil); err != nil {
		return err
	}
	if err = s.genWithPrjFile("router/router.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/logger/logger.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/mysql/mysql.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/apis/apis.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/apis/common/common.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/apis/common/const.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/apis/common/reply.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	return nil
}

func (s *GenFrameService) replaceContent(content string, oldStr, newStr string) string {
	content = strings.ReplaceAll(content, oldStr, newStr)
	return content
}

func (s *GenFrameService) genWithPrjFile(fileName string, replace map[string]string) error {
	filePath := path.Join(s.outPrjPath, fileName)
	contentPath := path.Join(s.genPrjPath, fileName)
	datab, err := os.ReadFile(contentPath)
	if err != nil {
		return err
	}
	content := string(datab)
	for oldStr, newStr := range replace {
		content = s.replaceContent(content, oldStr, newStr)
	}
	if s.cfg.JsonCase == "camel" {
		content = utils.JsonToCamel(content)
	}
	err = utils.SaveFile(filePath, []byte(content))
	if err != nil {
		fmt.Println("filePath err", err.Error())
		return err
	}

	return err
}
