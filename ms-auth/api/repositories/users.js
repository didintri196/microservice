const db = require('diskdb');

class UsersRepository {
    
    getUserByPhonePassword(phone, password){
        return db.users.find({phone: phone, password: password});
    }

    addUser(user) {
        return db.users.save(user);
    }
    
    cekUsername(username) {
        const cekUsername = db.users.find({ username: username });
        if (cekUsername.length > 0) {
            return false;
        }
        return true;
    }

}

module.exports = new UsersRepository();