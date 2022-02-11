User = {
    type: "object",
    properties: {
        username: { type: "string" },
        name: { type: "string" },
        phone: { type: "string" },
        role: { type: "string" }
    },
    required: ["username", "name", "phone", "role"],
    additionalProperties: false,
}

module.exports = User;