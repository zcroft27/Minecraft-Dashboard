# Minecraft Server Dashboard

## Overview  
This is a fullstack project to create a dashboard for
managing a minecraft server (or really any server) in a virtual machine sending/executing commands over SSH.
I will host this project on a very lightweight (probably B1ls) virtual machine, pointed to by a subdomain of my zachlearns.com domain.
My Minecraft server is already hosted and available to read about on [my website](https://zachlearns.com/mc.html).

## The Backend
I have implemented authentication and authorization using pgx and raw HTTP requests. The current flow is:
- POST to http://localhost:<port>/sign-up with the body:
  {
  "email" : "myemail@gmail.com",
  "password" : "mysecretpassword"
  }
- POST to http://localhost:<port>/login with the body:
  {
  "email" : "myemail@gmail.com",
  "password" : "mysecretpassword"
  }

The request to the 'login' route will store an access token and userID in the session cookies. These cookies are essential. Any route
within the route /vm/* will use a middleware that validates the access token with the JWTSecret token in the .env and validate
that the role set for the user is 'admin.' This role has to be manually set with a migration script over Supabase CLI, the web interface
disallows mutation.


I have endpoints that allow starting and stopping of the server, executing select commands on the server (/list, /ban <player>, etc.) remotely,
and running vm diagnostics (free -h to see RAM usage, etc.).   
These endpoints are in progress, but will be protected using JSON web tokens(? need to research). They use the [crypto SSH](golang.org/x/crypto/ssh) library
for connecting over SSH and sending commands.

I use a three layer Controller-Service-Repository architecture for this project.

The repository layer currently contains a sanity check 'TestConnection' to validate that the connection is healthy.

## The Frontend
The frontend dashboard is built with React + TypeScript.

## General Design Choices
I intentionally do not want to be able to start/stop the **VM** from this dashboard, and do not want to parameterize the execution
of SSH commands. The VM should only be turned on from the Azure dashboard. If my endpoints become vulnerable, I want an extra layer of security,
ensuring the VM is off when I want it to be. While allowing full console control could be nice, to do so would require allowing any command to be entered,
which I don't feel comfortable securing.


## The Virtual Machine + Server
**My [website](http://zachlearns.com/mc.html) highlights this more extensively.**

My VM is hosted in Azure with a reverse proxy configured for the **mc.zachlearns.com** subdomain. The server is PurPur instead of Mojang's default server.
