const br = require('bahtrext');
const express = require('express');
const app = express();

const PORT = 3000;

app.get("/br/:money", (req, res) => {
  res.json(br.OB(req.params.money));
});

app.listen(PORT, () => console.log(`Listening on port ${PORT}`));