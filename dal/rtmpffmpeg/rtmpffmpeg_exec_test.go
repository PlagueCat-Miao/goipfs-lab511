package rtmpffmpeg

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)



func TestRmptFfmpeg(t *testing.T) {

	Convey("[TestRmptFfmpeg] ffmpeg环境测试", t, func() {
		rmptCtrl:=NewRmptFfmpeg()
		err:=rmptCtrl.FfmpegExecCheck()
		So(err,ShouldBeNil)
	})
}

