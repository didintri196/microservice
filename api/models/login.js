Login = {
    type: "object",
    properties: {
        phone: { type: "string" },
        password: { type: "string" }
    },
    required: ["phone", "password"],
    additionalProperties: false,
}

module.exports = Login;
