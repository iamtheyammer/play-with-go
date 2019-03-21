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
    if(!dbResult.rows) return "No rows.";
    const row = dbResult.rows[0];
    return "ID: " + row.id + ", User ID: " + row.user_id + ", Timestamp: " + row.timestamp
  });
}
module.exports = getRow;
