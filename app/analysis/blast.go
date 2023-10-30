package analysis

import (
	"bytes"
	"log"
	"os/exec"
)

// BlastAnalyze Blast 分析
func BlastAnalyze() {
	// 指定要执行的Shell脚本和参数
	cmd := exec.Command("sh", "app/analysis/blast.sh")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error(), stderr.String())
	} else {
		log.Printf(out.String())
	}
}
