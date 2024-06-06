const express = require('express');
const handlers = require('./handlers');

const router = express.Router();

router.get('/', handlers.getIndex);
router.get('/drink', handlers.getDrink);
router.get('/lakes', handlers.getLakes);

module.exports = { router };