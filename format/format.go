package format

import (
	"github.com/zebemce/vdk/av/avutil"
	"github.com/zebemce/vdk/format/aac"
	"github.com/zebemce/vdk/format/flv"
	"github.com/zebemce/vdk/format/mp4"
	"github.com/zebemce/vdk/format/rtmp"
	"github.com/zebemce/vdk/format/rtsp"
	"github.com/zebemce/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
