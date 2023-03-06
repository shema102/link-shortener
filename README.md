_Start Redis:_

`docker-compose -f docker-compose.yml up -d`

_Stop Redis:_

`docker-compose -f docker-compose.yml down`

_Build Docker Image:_

`docker build -t link-shortener .`

_Run Docker Image:_

`docker run -p 8080:8080 -e ENV=prod -e REDIS_PASSWORD=redis_password -e REDIS_HOST=host.docker.internal -e REDIS_PORT=6379 -e HOST=localhost:8080 link-shortener`
