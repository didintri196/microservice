const db = require('diskdb');

class UsersRepository {
    
    // async getAllUsers(){
    //     return await this.model.find({});
    // }

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

    // async updateUser(id, user){
    //     return await this.model.findByIdAndUpdate(id, user);
    // }

    // async deleteUser(id){
    //     return await this.model.findByIdAndDelete(id);
    // }
}

module.exports = new UsersRepository();