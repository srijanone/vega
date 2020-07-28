const express = require('express')
const app = express()
const port =  env = process.env.NODE_ENV || 3000

app.get('/', (req, res) => res.json({ success: true, message: "Hello World!" }))

app.listen(port, () => console.log(`Example app listening at http://localhost:${port}`))