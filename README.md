# Ascii Art Web

This small project takes the main [Ascii Art Project](https://zone01normandie.org/git/faoudia/ascii-art) and runs it on the web. The server is fully written in go and uses HTML templates to display content on the front-end.

## Usage

To launch the server, use this command.

```bash
go run .
```

if the server runs correctly, you should see this:

```bash
ascii-web-server:00:58:12  Server started at http://localhost:8080                                                          
```

Now, you can visit [localhost:8080](localhost:8080) to see the website.

## Code

Two main handlers are used: the homehandler and the posthandler. The home handler displays the home page which contains an HTML form to submit the input and the output of the Ascii Art.

The post handler takes the input typed on the home page and turns it into Ascii Art before sending it back to the front-end.

The server handles error codes 404, 400 and 500 with a specific template for each case.