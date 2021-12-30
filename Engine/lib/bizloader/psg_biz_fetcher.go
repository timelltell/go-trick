package bizloader

import (
	"GolangTrick/Engine/dao"
	_struct "GolangTrick/Engine/struct"
	"context"
	"encoding/json"
	"errors"
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
		// f.addLog(jsonStr, "ERROR", "", err.Error())
		return err
	}
	err = f.db.Fill(response)
	if err != nil {
		jsonStr, _ = json.Marshal(response)
		//logutil.AddErrorLog(ctx, logutil.MDU_BIZ_LOADER, logutil.IDX_PSG_BIZ_FAILED, err.Error(), string(jsonStr))
		// f.addLog(jsonStr, "ERROR", "", err.Error())
		return err
	}
	jsonStr, _ = json.Marshal(response)
	currentLoadTag := genHashTag(jsonStr)
	if currentLoadTag != psgLatestLoadTag {
		psgLatestLoadTag = currentLoadTag
		//logutil.AddInfoLog(ctx, logutil.MDU_BIZ_LOADER, logutil.IDX_PSG_BIZ_FETCHER, "load success", string(jsonStr), currentLoadTag)
	} else {
		//logutil.AddInfoLog(ctx, logutil.MDU_BIZ_LOADER, logutil.IDX_PSG_BIZ_FETCHER, "load success", "load done", currentLoadTag)
	}
	return nil
}

func (f *psgFetcher) load(ctx context.Context, mode string) (*pope_engine_editor.PopeIndexerResponse, error) {
	switch mode {
	case "file": // 文件模式
		data, err := f.loadFromFile()
		if err != nil {
			return nil, err
		}
		response := pope_engine_editor.NewPopeIndexerResponse()
		if err := json.Unmarshal(data, response); err != nil {
			return nil, err
		}
		return response, nil
	default: // 其他模式，目前主要是uri模式
		client, err := pope_engine_editor.NewClient(global.GetPopeEditorDisfName())
		if err != nil {
			return nil, err
		}
		return client.GetCanvasIndexer(ctx)
	}
	return nil, errors.New(ERR_INDEXER_MODE_NO_MATCH)
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
