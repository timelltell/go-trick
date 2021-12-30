package bizloader

import (
	"GolangTrick/Engine/dao"
	_struct "GolangTrick/Engine/struct"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type psgFetcher struct {
	config         *_struct.Config
	logDir         string
	logName        string
	currentLoadTag string
	db             *dao.DB
	logFileHandler *os.File
}

func newPsgFetcher(config *_struct.Config) *psgFetcher {
	return &psgFetcher{
		logDir:         config.PopeEditor.LogDir,
		logName:        config.PopeEditor.LogName,
		currentLoadTag: "",
		db:             dao.GetInstance(),
		config:         config,
	}
}

func (f *psgFetcher) fetch(ctx context.Context) error {

	response, err := f.load(ctx, f.config.PopeEditor.IndexerMode)
	var jsonStr []byte
	if err != nil {
		jsonStr, _ = json.Marshal(response)
		//logutil.AddErrorLog(ctx, logutil.MDU_BIZ_LOADER, logutil.IDX_PSG_BIZ_FAILED, err.Error(), string(jsonStr))
		return err
	}
	err = f.db.Fill(response)
	if err != nil {
		jsonStr, _ = json.Marshal(response)
		//logutil.AddErrorLog(ctx, logutil.MDU_BIZ_LOADER, logutil.IDX_PSG_BIZ_FAILED, err.Error(), string(jsonStr))
		return err
	}
	jsonStr, _ = json.Marshal(response)
	currentLoadTag := genHashTag(jsonStr)
	//if currentLoadTag != psgLatestLoadTag {
	//	psgLatestLoadTag = currentLoadTag
	//logutil.AddInfoLog(ctx, logutil.MDU_BIZ_LOADER, logutil.IDX_PSG_BIZ_FETCHER, "load success", string(jsonStr), currentLoadTag)
	//} else {
	//logutil.AddInfoLog(ctx, logutil.MDU_BIZ_LOADER, logutil.IDX_PSG_BIZ_FETCHER, "load success", "load done", currentLoadTag)
	//}
	fmt.Println(currentLoadTag)
	return nil
}

func genHashTag(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (f *psgFetcher) load(ctx context.Context, mode string) (*_struct.PopeIndexerResponse, error) {
	switch mode {
	case "file": // 文件模式
		data, err := f.loadFromFile()
		if err != nil {
			return nil, err
		}
		response := _struct.NewPopeIndexerResponse()
		if err := json.Unmarshal(data, response); err != nil {
			return nil, err
		}
		return response, nil
	default: // 其他模式，目前主要是uri模式
		//client, err := _struct.NewClient(global.GetPopeEditorDisfName())
		//if err != nil {
		//	return nil, err
		//}
		//return client.GetCanvasIndexer(ctx)
	}
	return nil, errors.New("err_indexer_mode_no_match")
}

func (f *psgFetcher) loadFromFile() ([]byte, error) {
	fh, err := os.Open(f.config.PopeEditor.IndexerFile)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	rep, err := ioutil.ReadAll(fh)
	return rep, err
}
