# Deploy 

## Deploying Applications with Docker

### Docker Run Commands

Docker run commands are used to create and start containers from Docker images.

```bash
docker run -p 8080:8080 ghcr.io/nicnl/borderlands4-serials:latest
```

### Docker Compose

Docker Compose is a tool for defining and running multi-container Docker applications. You can use a `docker-compose.yml` file to configure your application's services.

```yaml
services:
  app:
    image: ghcr.io/nicnl/borderlands4-serials:latest
    container_name: borderlands-serials
    ports:
      - "8080:8080"
    restart: unless-stopped
```
