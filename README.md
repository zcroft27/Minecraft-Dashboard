# Minecraft Server Dashboard

## Overview  
This is a full-stack project to create a dashboard for
managing a minecraft server--or really any server--in a virtual machine by sending commands over SSH.

The frontend + backend are hosted on a Standard B1ms (1 vcpu, 2 GiB memory), pointed to by a subdomain of my zachlearns.com domain through a Cloudflared
tunnel to hide the VM ip.
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

A request to the 'login' route will store an access token and userID in the session cookies. These cookies are essential. Any route
within the route /vm/* will use a middleware that validates the access token with the JWTSecret token in the .env and validate
that the role set for the user is 'admin.' This role has to be manually set with a migration script over Supabase CLI as the web interface disallows mutation within the auth schema.


I have endpoints that allow starting and stopping of the server, executing select commands on the server (/list, /ban <player>, etc.) remotely,
and running vm diagnostics (free -h to see RAM usage, etc.).   
These endpoints use the [crypto SSH](golang.org/x/crypto/ssh) library
for connecting over SSH and sending commands.

I use a three layer Controller-Service-Repository architecture for this project.

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

## Things That Caught Me Up
I had a difficult time debugging an issue with 'prepared statements' already existing. [This thread](https://forum.bubble.io/t/sql-connector-issue-prepared-statement-supabase-integration/303849/3) helped, highlighting the differences between using ports 6543 and 5432. I need to learn more about PgBouncer, pooling, and prepared statements.   
Identifiers such as table names in PostgreSQL are converted to lowercase by default--nomenclature is important.  
When the backend and frontend are on different origins (hostname/port), [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is necessary to be considered. A critical piece of code was:
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
