# Coupon Management System

This repository contains a Coupon Management System implemented with Gorm, Echo, and following the Clean Architecture principles. Users are required to register and log in to upload and view their coupons, while public access is allowed to view all coupons.

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [API](#api)

## 🚀 Getting Started

### Prerequisites

- Go installed on your machine
- GORM
- Echo
- Database (e.g., Mysql) installed and accessible

### 💻 Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/fauzibrillian/CouponSystem.git
2. Change into the project directory:

   ```bash
    cd CouponSystem
3. Install dependencies:

   ```bash
    go mod tidy

### Tech Stack

![TectStack](https://github.com/fauzibrillian/CouponSystem/assets/73748420/c998ad7d-a0a2-48f9-a1ac-d3a9834bb0ba)


### API

### Overview

The Coupon Management System provides a RESTful API for managing and retrieving coupons. Below are the available endpoints:

### Authentication

- **Endpoint**: `/register`
  - **Method**: `POST`
  - **Description**: Register a new user.
  - **Request Body**:
    ```json
    {
      "name": "your_name",
      "password": "your_password"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "User registered successfully"
    }
    ```

- **Endpoint**: `/login`
  - **Method**: `POST`
  - **Description**: Log in an existing user.
  - **Request Body**:
    ```json
    {
      "name": "your_name",
      "password": "your_password"
    }
    ```
  - **Response**:
    ```json
    {
      "token": "your_access_token"
    }
    ```

### Coupon Management

- **Endpoint**: `/coupon`
  - **Method**: `POST`
  - **Description**: Upload coupons for the authenticated user.
  - **Authorization**: Bearer Token
  - **Request Body**:
    ```json
    {
      "id": "1",
      "nama_program": "your_nama_program",
      "link": "your_link",
      "gambar":"your_gambar",
      "user_id":"your_user_id"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Coupons uploaded successfully"
    }
    ```

- **Endpoint**: `/coupon/user`
  - **Method**: `GET`
  - **Description**: View coupons uploaded by the authenticated user.
  - **Authorization**: Bearer Token
  - **Response**:
    ```json
    {
      "id": "1",
      "nama_program": "your_nama_program",
      "link": "your_link",
      "gambar":"your_gambar",
      "user_id":"your_user_id"
    }
    ```

- **Endpoint**: `/coupon/all`
  - **Method**: `GET`
  - **Description**: View a list of all coupons (public access).
  - **Response**:
    ```json
    {
      "id": "1",
      "nama_program": "your_nama_program",
      "link": "your_link",
      "gambar":"your_gambar",
      "user_id":"your_user_id"
    }
    ```

**Note**: Ensure that you include details specific to your API, such as data structures, expected responses, and any other relevant information.





