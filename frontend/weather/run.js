const express = require('express');
const path = require('path');
const app = express();
const port = 4000;
app.use(express.static(path.join(__dirname, 'build')));
app.get('*', function(req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});
app.listen(port);
console.log(`Servidor subiu com sucesso em http://localhost:${port}`);
console.log('Para derrubar o servidor: ctrl + c');
/*const express = require('express');
const app = express();
const baseDir = `${__dirname}/build/`;
const port = 4000;

app.use(express.static(`${baseDir}`));
app.get('*', (req,res) => res.sendFile('index.html' , { root : baseDir }));
app.listen(port, () => console.log(`Servidor subiu com sucesso em http://localhost:${port}`));
*/






