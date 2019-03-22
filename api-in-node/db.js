const { Pool } = require('pg');

const localDbOptions = {
  host: 'localhost',
  user: 'sammendelson'
}

const pool = new Pool(localDbOptions);

function query(text, params, callback) {
  return pool.query(text, params, callback);
}

function getRow(id) {
  return query(
    'SELECT * FROM already_messaged_users WHERE id=$1',
    [
      id
    ]
  ).then((dbResult) => {
    if(!dbResult.rows[0]) return JSON.stringify({
      error: true,
      msg: "No user for that ID."
    });
    const row = dbResult.rows[0];
    return JSON.stringify({
      id: row.id,
      user_id: row.user_id,
      timestamp: row.timestamp
    })
  });
}
module.exports = getRow;
