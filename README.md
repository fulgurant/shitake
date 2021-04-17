# Shitake

Shitake is a user management microservice.

It provides sign-up, sign-in and account removal functionality.

## Setup

This application relies on Redis compatible backing servers, their addresses are specified in the configuration.

## Security

All passwords are client-side hashed before coming to the server, and are re-hashed on the server. The username itself forms part of the salt for the hash. The client-side hashing algorithm used is SHA256, and the server-side is SHA512.

An email address is also stored with the username, but this is also client and server-side hashed in the same way as the password - to prevent privacy breaches.

## Protocol

The protocol to communicate with the server is an HTTP API as follows.

### Sign-Up

POST /sign-up/username

{
    "password": "salted-sha256-password",
    "email": "salted-sha256-email"
}