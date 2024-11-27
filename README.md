docker build -t scalesync
xhost +local:docker
docker run -e DISPLAY=$DISPLAY \
           -v /tmp/.X11-unix:/tmp/.X11-unix \
           scalesync

prometheus.exe --config.file=config/prometheus.yml