const express = require('express');
const getRow = require('./db');
const app = express();

app.get('/', (req, res) => {
  res.send('Test result...');
})

app.get('/db', (req, res) => {
  if(!req.query.id) return res.send({error: true, msg: 'No ID.'});
  getRow(req.query.id).then((dbResult) => {
    res.send(dbResult)
  })
})
app.listen(8001);
