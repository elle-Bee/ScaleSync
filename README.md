docker build -t scalesync
xhost +local:docker
docker run -e DISPLAY=$DISPLAY \
           -v /tmp/.X11-unix:/tmp/.X11-unix \
           scalesync

prometheus.exe --config.file=config/prometheus.yml

docker run --net=host -it --rm -e POSTGRES_PASSWORD=amity postgres

# Start an example database
docker run --net=host -it --rm -e POSTGRES_PASSWORD=password postgres
# Connect to it
docker run \
  --net=host \
  -e DATA_SOURCE_URI="localhost:5432/postgres?sslmode=disable" \
  -e DATA_SOURCE_USER=postgres \
  -e DATA_SOURCE_PASS=amity \
  quay.io/prometheuscommunity/postgres-exporter
