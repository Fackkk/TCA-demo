package scan

import (
	"TCA-Plugins-test/logger"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Scan() {
	//SOURCE_DIR：要扫描的代码目录路径
	//DIFF_FILES: 值为一个json文件路径，文件内容为增量扫描的文件列表(增量扫描时可用)
	//SCAN_FILES: 值为一个json文件路径，文件内容为需要扫描的文件列表(增量或全量扫描均可用)
	//TASK_REQUEST: 值为一个json文件路径，文件内容为当前扫描任务参数
	//RESULT_DIR: 结果result.json输出的结果目录路径

	logger.Info("SOURCE_DIR: ", os.Getenv("SOURCE_DIR"))
	logger.Info("DIFF_FILES: ", os.Getenv("DIFF_FILES"))
	logger.Info("SCAN_FILES: ", os.Getenv("SCAN_FILES"))
	logger.Info("TASK_REQUEST: ", os.Getenv("TASK_REQUEST"))
	logger.Info("RESULT_DIR: ", os.Getenv("RESULT_DIR"))

	logger.Infof("=====scan start=====")

	logger.Infof("=====scan end=====")

	srcPath := filepath.Join(os.Getenv("SOURCE_DIR"), "scan", "scan.go")
	result := `[{
	               "path": ` + srcPath + `,
	               'line': 17,
	               'column': 0,
	               'msg': "This is a testcase.",
	               'rule': "DemoRule",
	               "refs": [
	                   {
	                       "line": 19,
	                       "msg": "first ref msg",
	                       "tag": "first_tag",
	                       "path": ` + srcPath + `
	                   },
	                   {
	                       "line": 20,
	                       "msg": "second ref msg",
	                       "tag": "second_tag",
	                       "path": ` + srcPath + `
	                   }
	               ]
	           },{
	               "path": ` + srcPath + `,
	               'line': 18,
	               'column': 12,
	               'msg': "This is a testcase.",
	               'rule': "DemoRule",
	               "refs": [
	                   {
	                       "line": 20,
	                       "msg": "first ref msg",
	                       "tag": "first_tag",
	                       "path": ` + srcPath + `
	                   },
	                   {
	                       "line": 21,
	                       "msg": "second ref msg",
	                       "tag": "second_tag",
	                       "path": ` + srcPath + `
	                   }
	               ]
	           }]`
	err := ioutil.WriteFile(filepath.Join(os.Getenv("RESULT_DIR"), "result.json"), []byte(result), 0644)
	if err != nil {
		logger.Error(err)
	}

}
