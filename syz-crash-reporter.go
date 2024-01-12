package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/syzkaller/pkg/log"
	"github.com/google/syzkaller/pkg/mgrconfig"
	"github.com/google/syzkaller/pkg/report"
)

var (
	flagConfig = flag.String("config", "", "configuration file")
)

func main() {
	flag.Parse()
	cfg, err := mgrconfig.LoadFile(*flagConfig)
	if err != nil {
		log.Fatalf("%v", err)
	}

	myreporter, err := report.NewReporter(cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// 指定要读取的文件路径
	filePath := "./input_case"

	// 使用 ioutil.ReadFile 读取文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件时出错:", err)
		return
	}

	// 打印文件内容
	fmt.Println("文件内容:")
	fmt.Println(string(fileContent))

	rep := myreporter.Parse(fileContent)

	if err := myreporter.Symbolize(rep); err != nil {
		log.Logf(0, "failed to symbolize report: %v", err)
	}

	fmt.Println("Parse内容:")
	fmt.Printf("Type: %s , Frame: %s, Report: \n%s", string(rep.Type), string(rep.Frame), string(rep.Report))

}
