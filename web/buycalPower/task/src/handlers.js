const getIndex = async (req, res) => {
    if (req.session.data["lakes"]["buycal"].isDrinked) {
        Object.prototype.isDrinked = false;
        return res.send(process.env["FLAG"]);
    }
    return res.render('index');
}

const getDrink = async (req, res) => {
    const lake = req.query.lake;

    if (!lake) {
        return res.status(400).json({"status": "BAD", "error": "Некорректный запрос"});
    }

    if (lake == "buycal") {
        return res.status(400).json({"status": "BAD", "error": "Тупой ящер! Ты не можешь выпить воды из Байкала!"});
    } else {
        if (!req.session.data.lakes[lake]) {
            return res.status(400).json({"status": "BAD", "error": "Такого озера нет!"})
        }
        req.session.data["lakes"][lake].isDrinked = true;
        return res.json({"status": "OK"});
    }
};

const getLakes = async (req, res) => {
    return res.json(req.session.data.lakes);
};

module.exports = { getIndex, getDrink, getLakes };