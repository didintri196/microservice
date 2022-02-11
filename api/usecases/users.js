const jwt = require("jsonwebtoken");
const ajv = require('ajv');
const validator = new ajv();
const usersModel = require("../models/users");
const loginModel = require("../models/login");
const randomString = require("../utils/randomString");
const usersRepository = require("../repositories/users");
const jwt_library = require("../library/jwt_library");

exports.user_signup = (req, res, next) => {
  const valid = validator.validate(usersModel, req.body);
  if (valid) {

    const password = randomString.build(4);
    req.body.password = password;
    //cek username
    if (usersRepository.cekUsername(req.body.username)) {
      // register user
      if (usersRepository.addUser(req.body)) {
        res.status(201).json({
          message: "User created",
          data: req.body,
        });
      } else {
        res.status(406).json({
          message: "Error insert DB",
          data: req.body,
        });
      }
    } else {
      res.status(406).json({
        message: "Username alredy taken",
      });
    }

  } else {
    res.status(400).json({
      message: validator.errors[0].message
    });
  }
};

exports.user_login = (req, res, next) => {
  const valid = validator.validate(loginModel, req.body);
  if (valid) {
    const userLogin = usersRepository.getUserByPhonePassword(req.body.phone, req.body.password);
    if(userLogin.length>0){
      res.status(201).json({
        message: "Success Login",
        data: userLogin[0],
        token :jwt_library.generateToken(userLogin[0])
      });
    }else{
      res.status(401).json({
        message: "Wrong number phone or password",
      });
    }
  } else {
    res.status(400).json({
      message: validator.errors[0].message
    });
  }
};

exports.user_profile = (req, res, next) => {
  res.status(201).json({
    message: "User Authenticated",
    data: req.userData,
  });
};