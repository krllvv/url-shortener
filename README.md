# URL-shortener app

### Project setup

1. #### Clone repository
```bash
  git clone https://github.com/krllvv/url-shortener
  cd url-shortener
```

2. ####  Create file `.env`. Copy values from `env.example` to `.env` and update them with your PostgreSQL credentials and other settings
### Running Locally

1. #### Install dependencies
```bash
  go mod download     
```
2. #### Build the application
```bash
  go build -o app.exe ./cmd/main.go
```
3. #### Run the application
```bash
  ./app.exe    # using memory storage
```
OR
```bash
  ./app.exe -d  # using postgres storage
```
The application will run on `http://localhost:8080`

### Running using Docker
 #### Build and start containers
```bash
  docker-compose up --build
```
 By default, the application starts with `-d` flag to use the PostgreSQL as the storage. To run application with memory storage, you need to modify `Dockerfile` and remove `-d` flag. 

## Endpoints

### `POST /`
#### Request Body: URL to shorten
```text
 http://cjdr17afeihmk.biz/123/kdni9/z9d112423421
```

#### Response: Shortened URL 
```text
 http://localhost:8080/bvnPf
```

### `GET /alias`
#### Request URL: URL with the shortened alias 
```text
 http://localhost:8080/bvnPf
```

#### Response: Original URL
```text
 http://cjdr17afeihmk.biz/123/kdni9/z9d112423421
```