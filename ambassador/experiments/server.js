express = require("express")
app = express()

app.get('/', (req, res) => {
    res.send(process.env.LABEL)
})
app.listen(3000, () => {
    console.log(`Server is running on port 3000`)
})