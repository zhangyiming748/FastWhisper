package FastWhisper

type WhisperConfig struct {
	ModelType string // 使用的模型名称 如large-v3
	ModelDir  string // 模型存放路径
	Language  string // 语言
	VideoRoot string // 视频存放的绝对路径
}

func (wc *WhisperConfig) SetModelType(s string) {
	wc.ModelType = s
}
func (wc *WhisperConfig) SetModelDir(s string) {
	wc.ModelDir = s
}
func (wc *WhisperConfig) SetLanguage(s string) {
	wc.Language = s
}
func (wc *WhisperConfig) SetVideoRoot(s string) {
	wc.VideoRoot = s
}
