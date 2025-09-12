express = require("express")
var app = express()

app.get("/", (req, res) => {
    console.log("received request")
    res.send("Hello")
})

app.listen(8099, () => {
    console.log("now listening")
    console.log("listening at port 8099")
})