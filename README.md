# websocketvideostream
use jsmpeg.js  to decode mpeg1 video, use broadway  js to decode h264 video, the video stream transported by websocket, and the server wrote by golang.

jsmpeg.js is from  https://github.com/phoboslab/jsmpeg

broadway js is from https://github.com/mbebenita/Broadway

just run  websocketvideostream.exe

then,u can see video here:

http://127.0.0.1:8080/h264/h264.html

http://127.0.0.1:8080/mpeg1/mpeg1.html





H.264-AVC-ISO_IEC_14496-10-2012.pdf
For MPEG-4 H.264 transcoders that deliver I-frame, P-frame, and B-frame NALUs inside an MPEG-2 transport, the resulting packetized elementary streams (PES) are timestamped with presentation time stamps (PTS) and decoder timestamps (DTS) in time units of 1/90000 of a second.

The NALUs come in DTS timestamp order in a repeating pattern like

I P B B B P B B B ...  
where the intended playback rendering is

I B B B P B B B P ... 
(This transport strategy ensures that both frames that the B-frame bridges are in the decoder before the B-frame is processed.)

For FLV, the Timestamp (FLV spec p.69) tells when the frame should be fed to the decoder in milliseconds, which is

timestamp = DTS / 90.0
The CompositionTime (FLV spec p.72) tells the renderer when to perform ("compose") the video frame on the display device in milliseconds after it enters the decoder; thus it is

compositionTime = (PTS - DTS) / 90.0 
(Because the PTS >= DTS, this delta is never negative.)
