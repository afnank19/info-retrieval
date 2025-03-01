## Term Freq - Inverse Document Freq

Under development!

In order to run this project locally, you must ensure the following is present:

1. Crow (micro web framework for cpp), It is included in the repository or available at:
   <br></br>
   `https://github.com/CrowCpp/Crow.git`

2. Asio development headers (NOT INCLUDED, look into Crow's getting started docs for the exact version). Asio is available here:
   `https://github.com/chriskohlhoff/asio.git`

#### How to include:

1. Clone the repo to a place on your system.
   `git clone --depth 1 https://github.com/chriskohlhoff/asio.git`

2. Head over to the include folder. (/asio/include)
3. Copy all the content from the folder and paste it in your clone of this repo.

4. Run the Makefile

```make run // compiles and runs at the same time
    make build // compiles only
    make run-build // runs the build
    make clean // removes the binary
```

#### How to use:

Must have a collection folder containing .txt files (May add other files like .md or .html)
An endpoint has been provided:
<br></br>
`GET /search?q=your-query`

A request to this endpoint when you run the binary will return JSON containg the path to the file along with the relevancy to the query.
