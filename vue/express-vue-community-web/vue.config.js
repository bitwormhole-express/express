module.exports = {
    devServer: {
        port: 18080,
        host: '0.0.0.0',
        proxy: {
            "/api": {
                // target: "http://localhost:28080/"
                 target: "http://192.168.2.217:28080/"
                
            }
        }
    }
}