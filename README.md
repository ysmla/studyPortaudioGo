## Golang中portaudio库的尝试
### 使用的三方库
```Golang
errorlog "code.breezecode.tech/Breezecode/errorlog"
port "github.com/gordonklaus/portaudio"
```
### 代码说明

####  var selectDevice  port.StreamDeviceParameters
录制使用的设备参数
##### Device
指向设备的指针
##### Channels
通道数，int类型，单声道为1
##### Latency
延迟数，默认为DefaultLowInputLatency

#### var audioStream port.StreamParameters
录制使用的流参数
##### SampleRate
采样率 float64类型
##### FramesPerBuffer
缓冲空间大小

#### func audioProcesss(in []float32){}
回调函数，用于处理收到的音频数据

#### func savePcmData(filename string)error{}
转换并保存pcm数据

###附件说明
#### tomp3.py
将pcm数据文件转化为mp3文件，方便试听测试音频质量