package constdef

//检查文件夹
const PushLogPath = "./output/log/rtmppush"
const PushLogPathFormat = "./output/log/rtmppush/push-%v-%v.log"

const (
	FfmpegCmd = "ffmpeg" // 该常量应为目前环境下 ffmpeg的启动函数 ，请检查你的liunx系统PATH下是否有ffmpeg
	FfplayCmd = "ffplay" // 该常量应为目前环境下 ffplay的启动函数 ，请检查你的liunx系统PATH下是否有ffplay

	FfmpegArgFormat = "-r 30 -i /dev/video0 -vcodec h264 -max_delay 100 -f flv -g 5 -b 700000 %s -map 0:0 -map 0:2"

	FfplayArgFormat = "-fflags nobuffer -analyzeduration 500000 -i %s"
	RtmpFormat      = "rtmp://%s:%v/%s"
)
const (
	LiveRoom = "live"
	HlsRoom  = "hls"
	VodRoom  = "vod"
)

const RtmpDefaultPort = 1935
