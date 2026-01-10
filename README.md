# redirect2https
The program redirects HTTP requests to an HTTPS server. For requests with a specified path, it can act as a file server. This can be useful, for example, for obtaining Let's Encrypt certificates. Typically, automated tools like certbot use an HTTP request to verify website ownership using a .well-known directory.
The program runs on various operating systems. It has been tested on Windows, Linux, and FreeBSD.

