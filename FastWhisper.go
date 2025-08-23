package fastwhisper

import (
	"filepath"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func init() {
	os.Setenv("PYTHONWARNINGS", "ignore::FutureWarning")
	os.Setenv("PYTHONIOENCODING", "utf-8")
	// 设置默认时区为 Asia/Shanghai
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err) // 如果加载时区失败，则直接 panic
	}
	time.Local = loc
}

type WhisperConfig struct {
	ModelType string // 使用的模型名称 如large-v3
	ModelDir  string // 模型存放路径
	Language  string // 语言
	VideoRoot string // 视频存放路径
}

/*
生成字幕后返回字幕的绝对路径
*/
func GetSubtitle(fp string, wc WhisperConfig) string {
	var cmd *exec.Cmd
	if isCUDAAvailable() {
		cmd = exec.Command("whisper", fp, "--model", wc.ModelType, "--device", "cuda", "--model_dir", wc.ModelDir, "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", wc.Language, "--output_dir", wc.VideoRoot, "--verbose", "True")
	} else {
		cmd = exec.Command("whisper", fp, "--model", wc.ModelType, "--model_dir", wc.ModelDir, "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", wc.Language, "--output_dir", wc.VideoRoot, "--verbose", "True")
	}
	log.Printf("命令: %s\n", cmd.String())
	startTime := time.Now()

	// 修改开始：引入定期提示信息机制
	type result struct {
		output []byte
		err    error
	}

	done := make(chan result, 1)
	go func() {
		out, err := cmd.CombinedOutput()
		done <- result{out, err}
	}()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	var out []byte
	var err error

loop:
	for {
		select {
		case r := <-done:
			out, err = r.output, r.err
			break loop
		case <-ticker.C:
			fmt.Println("字幕生成中，请稍候...")
		}
	}
	// 修改结束

	if err != nil {
		log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
	} else {
		log.Printf("当前字幕生成成功\t命令原文:%v\t输出:%v\n", cmd.String(), string(out))
	}
	fp = strings.Replace(fp, filepath.Ext(fp), ".srt", 1)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	totalMinutes := duration.Seconds() / 60
	log.Printf("文件%v\n总共用时: %.2f 分钟\n", fp, totalMinutes)
	return fp
}

func isCUDAAvailable() bool {
	// 使用 nvidia-smi 命令检查 CUDA 是否可用
	cmd := exec.Command("nvidia-smi")
	output, err := cmd.CombinedOutput()
	// 如果命令执行出错，或者输出中不包含 "NVIDIA-SMI"，则认为 CUDA 不可用
	if err != nil || !strings.Contains(string(output), "NVIDIA-SMI") {
		return false
	}
	return true
}
