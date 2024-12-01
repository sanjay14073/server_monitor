# Web Service Monitor

A Go application that monitors the availability of a web service by checking a provided URL every 30 seconds(You can change the CRON expression to modify timings). If the service is down, it sends an email notification to a specified maintenance user.

## Requirements

- Go 1.x
- SMTP credentials for sending emails
- `.env` file for environment variables

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/web-service-monitor.git
   ```

2. Install dependencies:
   ```bash
   go get -u github.com/robfig/cron/v3
   go get -u gopkg.in/gomail.v2
   go get -u github.com/joho/godotenv
   ```

3. Create a `.env` file in the root of the project with the following variables:
   ```env
   URL=<your-web-service-url>
   SMTP_USER=<your-smtp-email>
   SMTP_PASSWORD=<your-smtp-password>
   MAINTAIN_USER=<maintenance-user-email>
   ```

## Usage

1. Run the application:
   ```bash
   go run main.go
   ```

2. The app will check the provided URL every 30 seconds(You can change it). If the service is down, an email will be sent to the maintenance user.
