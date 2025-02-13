---

# Notification Microservice

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/your-username/notification_microservice/ci.yml?branch=main) ![Docker Image Size](https://img.shields.io/docker/image-size/jeffdanurss/notification_microservice/latest) ![License](https://img.shields.io/github/license/your-username/notification_microservice)

> A microservice for managing notifications via email and RabbitMQ.

This microservice is designed to handle notifications by consuming messages from a RabbitMQ queue, processing them, and sending emails using MongoDB as the database.

---

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Docker Deployment](#docker-deployment)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **RabbitMQ Integration**: Consumes messages from a RabbitMQ queue.
- **Email Notifications**: Sends emails using an SMTP server.
- **MongoDB Storage**: Persists notification data in a MongoDB database.
- **REST API**: Provides endpoints to interact with the service.
- **Docker Support**: Easily deployable using Docker and Docker Compose.

---

## Prerequisites

Before you begin, ensure you have the following installed:

- [Node.js](https://nodejs.org/) (v18 or higher)
- [MongoDB](https://www.mongodb.com/)
- [RabbitMQ](https://www.rabbitmq.com/)
- [Docker](https://www.docker.com/) (optional, for containerized deployment)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/notification_microservice.git
   cd notification_microservice
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

---

## Configuration

Create a `.env` file in the root directory and add the following environment variables:

```env
PORT=3002
MONGO_URI=mongodb://root:example@mongodb:27017/notifications?authSource=admin
RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672
EMAIL_USER=your-email@gmail.com
EMAIL_PASS=your-email-password
```

- `PORT`: The port on which the service will run.
- `MONGO_URI`: Connection string for MongoDB.
- `RABBITMQ_URI`: Connection string for RabbitMQ.
- `EMAIL_USER`: Email address for sending notifications.
- `EMAIL_PASS`: Password for the email account.

---

## Usage

### Start the Service Locally

1. Start MongoDB and RabbitMQ (if not already running).
2. Run the service:
   ```bash
   npm start
   ```

The service will start listening on the specified port (default: `3002`).

---

## API Endpoints

### Send a Notification

- **POST** `/notifications`
  - **Request Body**:
    ```json
    {
      "message": "Hello, this is a test notification!",
      "email": "recipient@example.com"
    }
    ```
  - **Response**:
    ```json
    {
      "status": "success",
      "message": "Notification sent successfully."
    }
    ```

---

## Docker Deployment

### Build and Run with Docker Compose

1. Build and start the containers:
   ```bash
   docker-compose up --build
   ```

2. Access the service at `http://localhost:3002`.

### Push Docker Image to Docker Hub

1. Build the Docker image:
   ```bash
   docker build -t jeffdanurss/notification_microservice:latest .
   ```

2. Push the image to Docker Hub:
   ```bash
   docker push jeffdanurss/notification_microservice:latest
   ```

---

## Contributing

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please follow the [code of conduct](CODE_OF_CONDUCT.md).

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Node.js](https://nodejs.org/)
- [Express.js](https://expressjs.com/)
- [MongoDB](https://www.mongodb.com/)
- [RabbitMQ](https://www.rabbitmq.com/)
- [Nodemailer](https://nodemailer.com/)

---
