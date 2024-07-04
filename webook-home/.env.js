module.exports = {
    // 开发环境
    dev: {
        NODE_ENV: 'development',
        BASE_API: 'http://localhost:8080', // 本地服务器地址
        WEB_API: 'https://web.chinahrt.com', // 本地服务器地址
    },
    // 测试环境
    test: {
        NODE_ENV: 'test',
        BASE_API: 'https://cloud.chinahrt.com', // 本地服务器地址
        STUDENT_URL: 'https://gp.chinahrt.com',
       
    },
    // 生产环境
    product: {
        NODE_ENV: 'production',
        BASE_API: 'https://cloud.chinahrt.com', // 本地服务器地址
        WEB_API: 'https://web.chinahrt.com', // 本地服务器地址     
    }
}
