const express = require("express");
const router = express.Router();

const UserController = require('../usecases/users');
const checkAuth = require('../middlewares/jwt_middleware');

router.post("/signup", UserController.user_signup);

router.post("/login", UserController.user_login);

router.get("/profile", checkAuth, UserController.user_profile);

module.exports = router;
