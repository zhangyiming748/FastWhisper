# FastWhisper

FastWhisper 是一个基于 OpenAI Whisper 的快速字幕生成工具，支持 CUDA 加速。

## 功能特性

- 自动检测 CUDA 可用性并选择相应设备执行
- 支持多种 Whisper 模型
- 支持多种语言识别
- 生成 SRT 格式字幕文件
- 实时显示处理进度

## 安装

### 环境要求

- 能正常运行whisper命令的环境
- Go 1.25+
- [CUDA](https://developer.nvidia.com/cuda-downloads) (可选，用于 GPU 加速)
