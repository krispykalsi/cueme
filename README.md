<img src="https://i.imgur.com/gUw7ANz.jpg" alt="Cueme logo" align="right" width="200"/>

# Cueme
![Go version](https://img.shields.io/github/go-mod/go-version/krispykalsi/cueme)

An API to send a text message via email/WhatsApp/sms on a specific time

## 🧭 Usage

### Endpoints
- **POST** `/api/wa`
- **POST** `/api/sms`
- **POST** `/api/email`

### Request body
It can have four parameters -
```json
{
    "message": "hello cueme :3",
    "phone":   "+919876543210",
    "email":   "joe@mama.com",
    "time":    "2021-07-25T15:59:59.000Z",
}
```
- `phone` or `email` can be omitted if it's not required by the medium. (Eg: `api/sms` doesn't require the `email` field)
- `time` is in RFC3339 format and is an **optional** field.

## 🚧 Run Locally
1. Make sure you have [Go](https://go.dev/dl/) installed
2. Download the code

    ```shell
    git clone https://github.com/krispykalsi/cueme.git
    cd cueme
    ```
2. Download dependencies

    ```shell
    go mod download
    ```
3. Build the server

    ```shell
    go build ./cmd/cueme-server
    ```

4. Run locally and access through `http://localhost:8080`

    - Windows
    
    ```shell
    cueme-server.exe
    ```
    - Unix
    
    ```shell
    ./cueme-server
    ```
