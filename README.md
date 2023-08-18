# RabbitMQ Order Processing with Go

<!-- ![RabbitMQ Logo](rabbitmq_logo.png) -->

This project showcases an order processing system implemented using RabbitMQ as the message broker and the Go programming language. The system consists of a producer that sends orders to a RabbitMQ queue and a consumer that processes those orders and updates their status.

## Features

- **Order Submission**: The producer application sends example orders to the RabbitMQ queue, simulating incoming orders for processing.
- **Order Processing**: The consumer application listens to the RabbitMQ queue and processes orders as they arrive. It updates the order status and performs necessary actions based on the order content.
- **Real-time Communication**: RabbitMQ ensures efficient communication and coordination between the producer and consumer, allowing orders to be processed in near real-time.
- **Scalability**: The RabbitMQ architecture supports scalability, allowing you to add more consumers as needed to handle increased order volumes.
- **Asynchronous Processing**: Orders are processed asynchronously, reducing the likelihood of delays and bottlenecks in the system.
- **Fault Tolerance**: RabbitMQ provides features like message acknowledgments and durable queues, ensuring that messages are not lost even in the event of system failures.
- **Simple Setup**: The project includes a straightforward setup process, enabling you to quickly get started with RabbitMQ and Go.
- **Easy Customization**: You can easily modify the producer and consumer applications to suit your specific business requirements and integrate them into your existing systems.

## Requirements

- Go (1.13+ recommended)
- RabbitMQ (running locally or on a reachable server)

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/BaseMax/RabbitMQOrderGo.git
cd RabbitMQOrderGo
```

2. Install the streadway/amqp library:
```bash
go get github.com/streadway/amqp
```

3. Start the consumer:
```bash
go run consumer.go
```

4. In a separate terminal, start the producer:
```bash
go run producer.go
```

The producer will send an example order to the RabbitMQ queue, and the consumer will process it and update its status.

## Project Structure

- `producer.go`: The producer application that sends orders to the RabbitMQ queue.
- `consumer.go`: The consumer application that processes orders from the RabbitMQ queue.

## API

### Health Check

**Endpoint:** `/health`

**Method:** `GET`

**Description:** Health check endpoint to verify the server's status.

---

### Submit Order

**Endpoint:** `/orders`

**Method:** `POST`

**Description:** Submit an order for processing.

---

### Retrieve Order Status

**Endpoint:** `/orders/:orderID`

**Method:** `GET`

**Description:** Retrieve the status of a specific order.

---

### List All Orders

**Endpoint:** `/orders`

**Method:** `GET`

**Description:** Get a list of all orders in the system.

---

### Update Order Status

**Endpoint:** `/orders/:orderID/status`

**Method:** `PUT`

**Description:** Update the status of a specific order.

---

### Cancel Order

**Endpoint:** `/orders/:orderID/cancel`

**Method:** `POST`

**Description:** Cancel a specific order.

---

### Get Customer Orders

**Endpoint:** `/customers/:customerID/orders`

**Method:** `GET`

**Description:** Get a list of orders for a specific customer.

---

### Calculate Order Total

**Endpoint:** `/orders/:orderID/total`

**Method:** `GET`

**Description:** Calculate the total cost of a specific order.

---

### Process Payment

**Endpoint:** `/orders/:orderID/payment`

**Method:** `POST`

**Description:** Process the payment for a specific order.

---

### Assign Order to User

**Endpoint:** `/orders/:orderID/assign`

**Method:** `PUT`

**Description:** Assign a specific order to a user or employee.

---

### Get User's Assigned Orders

**Endpoint:** `/users/:userID/orders`

**Method:** `GET`

**Description:** Get a list of orders assigned to a specific user.

---

### Get Order History

**Endpoint:** `/orders/:orderID/history`

**Method:** `GET`

**Description:** Retrieve the history of status changes for a specific order.

---

### Export Order Data

**Endpoint:** `/orders/export`

**Method:** `GET`

**Description:** Export order data in a specified format (CSV, JSON, etc.).

---

### Import Order Data

**Endpoint:** `/orders/import`

**Method:** `POST`

**Description:** Import order data from an external source.

---

### Generate Order Report

**Endpoint:** `/reports/orders`

**Method:** `GET`

**Description:** Generate a report containing order statistics and details.

---

### Request Order Refund

**Endpoint:** `/orders/:orderID/refund`

**Method:** `POST`

**Description:** Request a refund for a specific order.

---

### Check Refund Status

**Endpoint:** `/orders/:orderID/refund/status`

**Method:** `GET`

**Description:** Check the status of a refund request for a specific order.

---

### Get Product Details

**Endpoint:** `/products/:productID`

**Method:** `GET`

**Description:** Retrieve details about a specific product.

---

### List Available Products

**Endpoint:** `/products`

**Method:** `GET`

**Description:** Get a list of all available products in the system.

---

### Update Product Information

**Endpoint:** `/products/:productID`

**Method:** `PUT`

**Description:** Update information for a specific product.

## Database Schema

```
Table: orders
-------------------------------------
| order_id | customer_id | status   |
-------------------------------------
| 1        | 101         | processing|
| 2        | 102         | completed |
| 3        | 103         | cancelled |
-------------------------------------

Table: customers
-------------------------------------
| customer_id | name      | email    |
-------------------------------------
| 101         | John Doe  | john@example.com |
| 102         | Jane Smith| jane@example.com |
| 103         | Alice Lee | alice@example.com|
-------------------------------------

Table: products
-------------------------------------
| product_id | name      | price    |
-------------------------------------
| 201        | Product A | 19.99    |
| 202        | Product B | 29.99    |
| 203        | Product C | 14.99    |
-------------------------------------

Table: users
-------------------------------
| user_id | name      | role |
-------------------------------
| 301     | Admin     | admin|
| 302     | Employee  | user |
-------------------------------

Table: refunds
--------------------------------
| refund_id | order_id | status |
--------------------------------
| 401       | 1        | pending|
| 402       | 2        | approved|
| 403       | 3        | declined|
--------------------------------
```

## Acknowledgements

This project is built using the `streadway/amqp` Go library for RabbitMQ communication. Special thanks to the library authors and the Go community for their contributions.

Feel free to customize and expand upon this project to meet your specific use case and business requirements. If you have any questions or suggestions, please open an issue or pull request in this repository.

Enjoy processing orders efficiently with RabbitMQ and Go!

Copyright 2023, Max Base
