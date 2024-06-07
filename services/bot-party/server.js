const { spawn } = require('child_process');
const express = require('express');
const app = express();
const port = 3000;

/* app.get('/stream', (req, res) => {
  // Set the response headers for an m3u8 playlist (HLS format)
  res.setHeader('Content-Type', 'application/vnd.apple.mpegurl');
  res.setHeader('Access-Control-Allow-Origin', '*');

  // Spawn an FFmpeg process to convert RTSP to HLS
  const ffmpegProcess = spawn('ffmpeg', [
    '-i', 'rtsp://your_rtsp_stream_url', // Your RTSP stream URL
    '-c:v', 'libx264',                   // Video codec: H.264
    '-hls_time', '2',                    // Segment duration
    '-hls_list_size', '3',               // Number of segments in playlist
    '-hls_flags', 'delete_segments+append_list', // HLS flags
    '-f', 'hls',                         // Output format
    'pipe:1'                             // Output to stdout
  ]);

  ffmpegProcess.stdout.pipe(res); // Pipe FFmpeg's stdout to the response

  ffmpegProcess.on('close', (code) => {
    console.log(`FFmpeg process exited with code ${code}`);
  });
}); */


app.use(express.static('public'));

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
