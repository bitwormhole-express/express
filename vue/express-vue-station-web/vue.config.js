module.exports = {
    devServer: {
        port: 18080,
        host: 'localhost',
        proxy: {
            "/api": {
                target: "http://localhost:28080/"
            }
        }
    }
}