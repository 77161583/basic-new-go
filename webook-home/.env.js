module.exports = {
    // 开发环境
    dev: {
        NODE_ENV: 'development',
        BASE_API: 'https://cloud.chinahrt.com', // 本地服务器地址
        WEB_API: 'https://web.chinahrt.com', // 本地服务器地址
        STUDENT_URL: 'https://gp.chinahrt.com',
        KEYWORDS: 'AES!@#S!@S)(SAS%GD<>?:S))!#$DKEY', //AES的key
        platformId: '151',
        trainplanId: 'c2be72c587cf43e3921846b40c81fec2',
        NEWS_ID:'10660', //政策法规
        NOTICE_ID:'10661', //通知公告
        CAOZUO_ID:'10662', //操作说明
        LINK_ID:'10675', //友情链接
        ALERT_ID:'12342', //弹出框
        GONGXU_ID:'baea7cab29e24a559e183b0b5ccfc23a',//公需课
        ZHUANYE_ID:'0cf65c75f69b460697831ab84535ccbf', //专业课
    },
    // 测试环境
    test: {
        NODE_ENV: 'test',
        BASE_API: 'https://cloud.chinahrt.com', // 本地服务器地址
        STUDENT_URL: 'https://gp.chinahrt.com',
        KEYWORDS: 'AES!@#S!@S)(SAS%GD<>?:S))!#$DKEY', //AES的key
        platformId: '151',
        trainplanId: 'c2be72c587cf43e3921846b40c81fec2',
        NEWS_ID:'10660', //政策法规
        NOTICE_ID:'10661', //通知公告
        CAOZUO_ID:'10662', //操作说明
        LINK_ID:'10675', //友情链接
        ALERT_ID:'12342', //弹出框
        GONGXU_ID:'baea7cab29e24a559e183b0b5ccfc23a',//公需课
        ZHUANYE_ID:'0cf65c75f69b460697831ab84535ccbf', //专业课
    },
    // 生产环境
    product: {
        NODE_ENV: 'production',
        BASE_API: 'https://cloud.chinahrt.com', // 本地服务器地址
        WEB_API: 'https://web.chinahrt.com', // 本地服务器地址
        STUDENT_URL: 'https://gp.chinahrt.com',
        KEYWORDS: 'AES!@#S!@S)(SAS%GD<>?:S))!#$DKEY', //AES的key
        NODE_PORT: '3042',
        platformId: '151',
        trainplanId: 'c2be72c587cf43e3921846b40c81fec2',
        NEWS_ID:'10660', //政策法规
        NOTICE_ID:'10661', //通知公告
        CAOZUO_ID:'10662', //操作说明
        LINK_ID:'10675', //友情链接
        ALERT_ID:'12342', //弹出框
        GONGXU_ID:'baea7cab29e24a559e183b0b5ccfc23a',//公需课
        ZHUANYE_ID:'0cf65c75f69b460697831ab84535ccbf', //专业课
        
    }
}
