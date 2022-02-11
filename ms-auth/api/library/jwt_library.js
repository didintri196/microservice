const jwt = require("jsonwebtoken");

class JWTLibrary {
    constructor() {
        this.secret = process.env.JWT_SECRET;
    }

    generateToken(user) {
        const payload = {
            username: user.username,
            phone: user.phone,
            role: user.role,
            timestamp: new Date().getTime()
        };
        return jwt.sign(payload, this.secret, {
            expiresIn: '1h'
        });
    }

    verifyToken(token) {
        return jwt.verify(token, this.secret);
    }
}

module.exports = new JWTLibrary();