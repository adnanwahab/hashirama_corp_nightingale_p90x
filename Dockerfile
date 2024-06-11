# Use the Stereolabs ZED base image with NVIDIA runtime support
FROM stereolabs/zed:4.1-tools-devel-l4t-r36.3

# Set environment variables to avoid interactive prompts during package installation
ENV DEBIAN_FRONTEND=noninteractive

# Update and install essential packages
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    git \
    && apt-get remove -y libopencv-dev \
    && apt-get install -y \
    libopencv-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-contrib-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-calib3d-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-core-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-dnn-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-features2d-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-flann-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-highgui-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-imgcodecs-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-imgproc-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-ml-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-objdetect-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-photo-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-stitching-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-video-dev=4.5.4+dfsg-9ubuntu4 \
    libopencv-videoio-dev=4.5.4+dfsg-9ubuntu4 \
    libgstreamer1.0-0 \
    gstreamer1.0-libav \
    libgstrtspserver-1.0-0 \
    gstreamer1.0-tools \
    gstreamer1.0-x \
    gstreamer1.0-alsa \
    gstreamer1.0-gl \
    gstreamer1.0-gtk3 \
    gstreamer1.0-qt5 \
    gstreamer1.0-pulseaudio \
    libgstreamer1.0-dev \
    libgstrtspserver-1.0-dev \
    libgstreamer-plugins-base1.0-0 \
    libgstreamer-plugins-base1.0-dev \
    libgstreamer-plugins-good1.0-0 \
    libgstreamer-plugins-good1.0-dev \
    libgstreamer-plugins-bad1.0-0 \
    libgstreamer-plugins-bad1.0-dev \
    && rm -rf /var/lib/apt/lists/*

# Clone the ZED GStreamer plugin repository
RUN git clone https://github.com/stereolabs/zed-gstreamer.git /usr/local/zed-gstreamer

# Build and install the ZED GStreamer plugin
WORKDIR /usr/local/zed-gstreamer
RUN mkdir build && cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release .. && \
    make && \
    make install

# Expose the RTSP port
EXPOSE 8554

# Command to start the RTSP server
CMD ["gst-zed-rtsp-launch", "zedsrc", "!", "videoconvert", "!", "video/x-raw, format=(string)I420", "!", "x264enc", "!", "rtph264pay", "pt=96", "name=pay0"]
