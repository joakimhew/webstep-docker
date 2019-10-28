const express = require("express");
const app = express();
const port = 8080;

app.use("/", (req, res) => {
  res.send("Hello, Webstep!");
});

app.listen(port, () => console.log(`Hello listening on port ${port}!`));
