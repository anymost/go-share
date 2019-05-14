const http = require('http');

const server = http.createServer(function(req, res) {
    res.end(JSON.stringify({
        name: "hello world"
    }))
});

server.listen("4000", function (err) {
    console.log('listen');
});