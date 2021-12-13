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
2. Get the code

    ```shell
    git clone https://github.com/krispykalsi/cueme.git
    cd cueme
    ```
    
3. Create a file named `.env.secret` and copy this data in it. Replace the values with your keys

    ```shell
    TWILIO_PHONE_AUTH_USERNAME=BCa1f7845a336f6492a8a90bbe5b763602
    TWILIO_PHONE_AUTH_PASSWORD=c9757b0fc7fef40774fff78d5f67ce53
    TWILIO_MSG_SERVICE_SID=MGd438c79dd721034c44c249edfa4ce96b
    TWILIO_WA_PHONE_NO=+14155238886
    TWILIO_MSG_API_ENDPOINT=https://api.twilio.com/2010-04-01/Accounts/ACf1f7845a336f6492a8a90bbe5b763602/Messages.json
    TWILIO_EMAIL_AUTH_KEY=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
    TWILIO_VERIFIED_EMAIL=joe@mama.com
    ```

4. Download dependencies

    ```shell
    go mod download
    ```
5. Build the server

    ```shell
    go build ./cmd/cueme-server
    ```

6. Run and access through `http://localhost:8080`

    - Windows
    
    ```shell
    cueme-server.exe
    ```
    - Unix
    
    ```shell
    ./cueme-server
    ```
