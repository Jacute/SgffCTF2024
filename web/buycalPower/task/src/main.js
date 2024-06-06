const express = require('express');
const bodyParser = require('body-parser');
const routes = require('./routes');
const session = require('express-session');
const path = require('path');

const app = express();

app.use(bodyParser.json());
app.set('view engine', 'ejs');
app.use("/static", express.static(path.join(__dirname, '../static')));
app.use(session({
    secret: process.env["SESSION_SECRET"],
    resave: true,
    saveUninitialized: true
}));

// init first session state
app.use((req, res, next) => {
    if (!req.session.initialized) {
        req.session.initialized = true;
        req.session.data = {};
        req.session.data["lakes"] = {
            buycal: {},
            onezhskoe: {},
            ladozhskoe: {}
        };
    }
    next();
});

app.use('/', routes.router);

const PORT = 3000;
const HOST = '0.0.0.0';

app.listen(PORT, HOST, () => {
    console.log(`Server is running on http://${HOST}:${PORT}`);
});
