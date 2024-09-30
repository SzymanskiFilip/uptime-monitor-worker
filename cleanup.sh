echo 'cleaning up database...'

docker stop $(docker ps -a)
docker rm $(docker ps -a)
docker compose up -d

echo 'finished'