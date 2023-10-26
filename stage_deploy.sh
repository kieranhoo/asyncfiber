docker network rm scsystem_api_network

docker network create \
  --driver=bridge \
  --subnet=172.28.0.0/16 \
  --ip-range=172.28.5.0/24 \
  --gateway=172.28.5.254 \
  scsystem_api_network

docker compose up -d --build

