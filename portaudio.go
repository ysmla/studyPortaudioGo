package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	errorlog "code.breezecode.tech/Breezecode/errorlog"
	port "github.com/gordonklaus/portaudio"
)

var pcmData []float32

func main() {
	//初始化portaudio
	_err := port.Initialize()
	//最后释放资源
	defer port.Terminate()
	if _err != nil {
		errorlog.Write("Unable to initialize portaudio library", "y")
		return
	}

	v := port.VersionText() //获取portaudio版本号
	fmt.Println("Version:", v)

	devs, _err := port.Devices()
	if _err != nil {
		errorlog.Write(_err.Error(), "y")
		return
	}

	var inputDevices []*port.DeviceInfo
	for _, dev := range devs {
		// 直接使用 dev，而不是尝试间接引用 *dev
		if dev.MaxInputChannels > 0 {
			inputDevices = append(inputDevices, dev)
		}
	}

	//定义设备参数
	selectDevice := port.StreamDeviceParameters{
		Device:   inputDevices[0],
		Channels: 1,                                      //通道数
		Latency:  inputDevices[0].DefaultLowInputLatency, //延时，此处使用默认最低延时
	}

	//定义流参数
	audioStream := port.StreamParameters{
		Input:           selectDevice,
		SampleRate:      float64(16000), //采样率
		FramesPerBuffer: 0,
	}
	//打开音频流
	stream, err := port.OpenStream(audioStream, audioProcess)
	if err != nil {
		errorlog.Write(err.Error(), "y")
		return
	}
	defer stream.Close() //最后释放资源

	//开始音频流
	err = stream.Start()
	if err != nil {
		errorlog.Write(err.Error(), "y")
		return
	}

	//录制10秒
	fmt.Println("Start recording.....")
	time.Sleep(10 * time.Second)

	//停止录制
	err = stream.Stop()
	if err != nil {
		errorlog.Write(err.Error(), "y")
		return
	}
	fmt.Println("Stop...")

	//保存pcm数据
	err = savePcmData("output.pcm")
	if err != nil {
		errorlog.Write(err.Error(), "y")
		return
	}
}

// 录音回调函数
func audioProcess(in []float32) {
	pcmData = append(pcmData, in...)
}

// 保存pcm数据
func savePcmData(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		errorlog.Write(err.Error(), "y")
		return err
	}
	defer file.Close()

	//将浮点数转换为int16并写入文件
	for _, sample := range pcmData {
		intSample := int16(sample * 32767)
		err = binary.Write(file, binary.LittleEndian, intSample)
		if err != nil {
			errorlog.Write(err.Error(), "y")
			return err
		}
	}
	return nil
}
